package specification_test

import (
	"cmp"
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
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/portfoliotree/alphavantage/specification"
)

const testdataExampleIndex = "testdata/examples/index.json"

type TestdataExampleEntree struct {
	ID      string    `json:"ID"`
	Path    string    `json:"path"`
	Fetched time.Time `json:"fetched"`
	URL     string    `json:"url"`
}

func loadTestdataExampleIndex(t *testing.T) []TestdataExampleEntree {
	var data []TestdataExampleEntree
	if buf, err := os.ReadFile(filepath.FromSlash(testdataExampleIndex)); err == nil {
		require.NoError(t, json.Unmarshal(buf, &data))
	}
	return data
}

func saveTestdataExampleIndex(t *testing.T, data []TestdataExampleEntree) {
	slices.SortFunc(data, func(a, b TestdataExampleEntree) int {
		return cmp.Compare(a.ID, b.ID)
	})
	buf, err := json.MarshalIndent(data, "", specification.JSONIndent)
	assert.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.FromSlash(testdataExampleIndex), buf, 0644))
}

func cacheEntreeID(t *testing.T, fn specification.Function, example string) string {
	t.Helper()
	s := sha256.New()
	s.Write([]byte(example))
	exampleHash := hex.EncodeToString(s.Sum(nil))
	exampleHash = exampleHash[:8]
	fileNamePrefix := fn.Name + "_" + exampleHash
	return fileNamePrefix
}

func testdataExampleBody(t *testing.T, dir, apikey string, entrees []TestdataExampleEntree, fn specification.Function, exampleURL string, run func(p string)) []TestdataExampleEntree {
	const maxDaysInCache = 7
	t.Helper()

	id := cacheEntreeID(t, fn, exampleURL)

	idx := slices.IndexFunc(entrees, func(e TestdataExampleEntree) bool {
		return id == e.ID
	})

	now := time.Now()

	if idx < 0 {
		t.Log("cache miss: index entree not found")

		bodyFilepath, err := downloadExample(t, apikey, exampleURL, dir, id)
		require.NoError(t, err)

		entrees = append(entrees, TestdataExampleEntree{
			ID:      id,
			Path:    filepath.ToSlash(bodyFilepath),
			Fetched: now,
			URL:     exampleURL,
		})

		run(bodyFilepath)
	} else if entrees[idx].Fetched.Before(now.AddDate(0, 0, -maxDaysInCache)) {
		t.Logf("cache miss: index entree more than %d day old", maxDaysInCache)

		bodyFilepath, err := downloadExample(t, apikey, exampleURL, dir, id)
		require.NoError(t, err)

		entrees[idx].Fetched = now
		entrees[idx].Path = filepath.ToSlash(bodyFilepath)
		entrees[idx].URL = exampleURL

		run(bodyFilepath)
	} else {
		run(filepath.FromSlash(entrees[idx].Path))
	}
	return entrees
}

func removeMissingExampleFileEntries(t *testing.T, entrees []TestdataExampleEntree) []TestdataExampleEntree {
	t.Helper()

	var cacheFileIDs []string
	exampleBodyFiles, err := filepath.Glob(filepath.FromSlash("testdata/examples/*/*"))
	require.NoError(t, err)
	for _, exampleBodyFile := range exampleBodyFiles {
		cacheFileIDs = append(cacheFileIDs, strings.TrimSuffix(filepath.Base(exampleBodyFile), filepath.Ext(exampleBodyFile)))
	}

	return slices.DeleteFunc(entrees, func(entree TestdataExampleEntree) bool {
		return !slices.Contains(cacheFileIDs, entree.ID)
	})
}

func downloadExample(t *testing.T, apikey, example, dir, filename string) (string, error) {
	t.Helper()
	ctx := t.Context()
	u, err := url.Parse(example)
	if err != nil {
		return "", fmt.Errorf("failed to parse example url: %w", err)
	}
	q := u.Query()
	q.Set("datatype", "csv")
	q.Set("apikey", apikey)
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to do example request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	buf, err := io.ReadAll(resp.Body)
	if err := resp.Body.Close(); err != nil {
		return "", fmt.Errorf("failed to close response body: %w", err)
	}
	if err != nil {
		return "", fmt.Errorf("failed to read example response: %w", err)
	}

	ct, _, err := mime.ParseMediaType(resp.Header.Get("content-type"))
	if err != nil {
		return "", fmt.Errorf("failed to parse response content type: %w", err)
	}

	switch ct {
	case "application/json":
		bodyFilename := filepath.Join(dir, filename+".json")
		if err := os.WriteFile(bodyFilename, buf, 0644); err != nil {
			return "", fmt.Errorf("failed to write example response: %w", err)
		}
		return bodyFilename, nil
	case "application/x-download":
		bodyFilename := filepath.Join(dir, filename+".csv")
		if err := os.WriteFile(bodyFilename, buf, 0644); err != nil {
			return "", fmt.Errorf("failed to write example response: %w", err)
		}
		return bodyFilename, nil
	default:
		return "", fmt.Errorf("unexpected content type: %s", ct)
	}
}
