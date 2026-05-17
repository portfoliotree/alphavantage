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
	"golang.org/x/time/rate"

	"github.com/portfoliotree/alphavantage/specification"
)

// fetchLimiter throttles example fetches to stay under AlphaVantage's
// per-second burst limit. Paid tiers cap at 5 req/s; we hold to 4 for safety.
var fetchLimiter = rate.NewLimiter(rate.Every(250*time.Millisecond), 1)

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

	const maxAttempts = 3
	for attempt := 1; ; attempt++ {
		if err := fetchLimiter.Wait(ctx); err != nil {
			return "", fmt.Errorf("rate limiter wait: %w", err)
		}
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
		if err != nil {
			return "", fmt.Errorf("failed to create request: %w", err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", fmt.Errorf("failed to do example request: %w", err)
		}
		if resp.StatusCode != http.StatusOK {
			_ = resp.Body.Close()
			return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
		buf, readErr := io.ReadAll(resp.Body)
		closeErr := resp.Body.Close()
		if readErr != nil {
			return "", fmt.Errorf("failed to read example response: %w", readErr)
		}
		if closeErr != nil {
			return "", fmt.Errorf("failed to close response body: %w", closeErr)
		}

		ct, _, err := mime.ParseMediaType(resp.Header.Get("content-type"))
		if err != nil {
			return "", fmt.Errorf("failed to parse response content type: %w", err)
		}

		// AlphaVantage signals throttling with HTTP 200 + a JSON envelope
		// containing a Note/Information/Error Message field. Detect that and
		// retry rather than caching the error as legitimate data.
		if ct == "application/json" {
			if msg, isErr := alphavantageErrorEnvelope(buf); isErr {
				if attempt >= maxAttempts {
					return "", fmt.Errorf("alphavantage rate-limited after %d attempts: %s", attempt, msg)
				}
				const cooldown = 30 * time.Second
				t.Logf("alphavantage error envelope on attempt %d (%q); sleeping %s before retry", attempt, msg, cooldown)
				select {
				case <-ctx.Done():
					return "", ctx.Err()
				case <-time.After(cooldown):
				}
				continue
			}
		}

		var ext string
		switch ct {
		case "application/json":
			ext = ".json"
		case "application/x-download":
			ext = ".csv"
		default:
			return "", fmt.Errorf("unexpected content type: %s", ct)
		}
		bodyFilename := filepath.Join(dir, filename+ext)
		if err := os.WriteFile(bodyFilename, buf, 0644); err != nil {
			return "", fmt.Errorf("failed to write example response: %w", err)
		}
		return bodyFilename, nil
	}
}

// alphavantageErrorEnvelope reports whether buf is a JSON object whose only
// meaningful content is one of AlphaVantage's notice fields (rate limiting,
// invalid params, etc.). When true, the second return is the notice text.
func alphavantageErrorEnvelope(buf []byte) (string, bool) {
	var msg struct {
		Note         string `json:"Note,omitempty"`
		Information  string `json:"Information,omitempty"`
		ErrorMessage string `json:"Error Message,omitempty"`
	}
	if err := json.Unmarshal(buf, &msg); err != nil {
		return "", false
	}
	switch {
	case msg.Note != "":
		return msg.Note, true
	case msg.Information != "":
		return msg.Information, true
	case msg.ErrorMessage != "":
		return msg.ErrorMessage, true
	}
	return "", false
}
