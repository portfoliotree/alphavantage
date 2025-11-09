package specification

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"slices"
)

type QueryParameter struct {
	Ident     string                    `json:"ident,omitempty"`
	Name      string                    `json:"name"`
	Type      string                    `json:"type"`
	Values    []json.RawMessage         `json:"values,omitempty"`
	ValuesURL string                    `json:"values_url,omitempty"`
	Validate  []QueryParameterValidator `json:"validate,omitempty"`
	Format    string                    `json:"format,omitempty"`
}

type QueryParameterValidator struct {
	Message    string `json:"message"`
	Identifier string `json:"identifier"`
	Expression string `json:"expression"`
}

type Generate struct {
	Files []File `json:"files"`
}

type File struct {
	File      string     `json:"file"`
	Functions []Function `json:"functions"`
}

type Function struct {
	Name       string              `json:"name"`
	Ident      string              `json:"ident,omitempty"`
	Required   []string            `json:"required"`
	Optional   []string            `json:"optional"`
	EnumSubset map[string][]string `json:"enum_subset,omitempty"`
	Examples   []string            `json:"examples"`
	CSVColumns []string            `json:"csv_columns,omitempty"`
}

func (fn Function) HasDatatypeParameter() bool {
	return slices.Contains(fn.Optional, QueryKeyDataType) || slices.Contains(fn.Required, QueryKeyDataType)
}

func FetchCSVExamples(ctx context.Context, dir, apikey string, functions []Function) error {
	for _, fn := range functions {
		fileNamePrefixes := fn.ExampleFileNamePrefixes()
		if !fn.HasDatatypeParameter() {
			continue
		}
		for _, example := range fn.Examples {
			if err := downloadExample(ctx, apikey, example, dir, fileNamePrefixes[example]); err != nil {
				return fmt.Errorf("download example for %q failed: %w", fn.Name, err)
			}
		}
	}
	return nil
}

func (fn Function) ExampleFileNamePrefixes() map[string]string {
	m := make(map[string]string)
	for _, example := range fn.Examples {
		s := sha256.New()
		s.Write([]byte(example))
		exampleHash := hex.EncodeToString(s.Sum(nil))
		exampleHash = exampleHash[:8]
		fileNamePrefix := fn.Name + "_" + exampleHash
		m[example] = fileNamePrefix
	}
	return m
}

func downloadExample(ctx context.Context, apikey, example, dir, filename string) error {
	u, err := url.Parse(example)
	if err != nil {
		return fmt.Errorf("failed to parse example url: %w", err)
	}
	q := u.Query()
	q.Set("datatype", "csv")
	q.Set("apikey", apikey)
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to do example request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	buf, err := io.ReadAll(resp.Body)
	if err := resp.Body.Close(); err != nil {
		return fmt.Errorf("failed to close response body: %w", err)
	}
	if err != nil {
		return fmt.Errorf("failed to read example response: %w", err)
	}

	ct, _, err := mime.ParseMediaType(resp.Header.Get("content-type"))
	if err != nil {
		return fmt.Errorf("failed to parse response content type: %w", err)
	}

	switch ct {
	case "application/json":
		if err := os.WriteFile(filepath.Join(dir, filename+".json"), buf, 0644); err != nil {
			return fmt.Errorf("failed to write example response: %w", err)
		}
		return nil
	case "application/x-download":
		if err := os.WriteFile(filepath.Join(dir, filename+".csv"), buf, 0644); err != nil {
			return fmt.Errorf("failed to write example response: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unexpected content type: %s", ct)
	}
}
