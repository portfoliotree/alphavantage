package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/portfoliotree/alphavantage"
)

type indexEntry struct {
	ID      string `json:"ID"`
	Path    string `json:"path"`
	Fetched string `json:"fetched"`
	URL     string `json:"url"`
}

var fakeServer *httptest.Server

func TestMain(m *testing.M) {
	// Load index.json for the fake server
	indexPath := filepath.Join(filepath.FromSlash("../../specification/testdata/examples/index.json"))
	indexData, err := os.ReadFile(indexPath)
	if err != nil {
		panic("failed to read index.json: " + err.Error())
	}

	var index []indexEntry
	if err := json.Unmarshal(indexData, &index); err != nil {
		panic("failed to parse index.json: " + err.Error())
	}

	// Create a fake AlphaVantage server
	fakeServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Match request to index entry
		reqParams := r.URL.Query()
		reqParams.Del("apikey")

		for _, entry := range index {
			entryURL, err := url.Parse(entry.URL)
			if err != nil {
				continue
			}

			// Compare query parameters (ignore apikey)
			entryParams := entryURL.Query()
			entryParams.Del("apikey")

			// Compare all parameters including multi-value ones
			if matchQueryParams(reqParams, entryParams) {
				// Found matching entry, serve the cached response
				responsePath := filepath.Join("..", "..", "specification", entry.Path)
				data, err := os.ReadFile(responsePath)
				if err != nil {
					http.Error(w, "failed to read response file: "+err.Error(), http.StatusInternalServerError)
					return
				}

				// Set content type based on file extension
				if strings.HasSuffix(entry.Path, ".json") {
					w.Header().Set("Content-Type", "application/json")
				} else {
					w.Header().Set("Content-Type", "text/csv")
				}

				w.WriteHeader(http.StatusOK)
				w.Write(data)
				return
			}
		}

		// No matching entry found - log for debugging
		// Only log if we have a function parameter to avoid noise
		if fn := reqParams.Get("function"); fn != "" {
			fmt.Fprintf(os.Stderr, "DEBUG: No match for function=%s, params=%v\n", fn, reqParams)
		}
		http.Error(w, "no cached response for this request", http.StatusNotFound)
	}))

	// Set environment variable to use fake server
	os.Setenv(alphavantage.APIURLEnvironmentVariableName, fakeServer.URL)

	// Run tests
	exitCode := m.Run()

	fakeServer.Close()
	os.Exit(exitCode)
}

// matchQueryParams compares two url.Values maps for equality, including multi-value parameters
// It handles comma-separated values by splitting them before comparison
func matchQueryParams(a, b url.Values) bool {
	if len(a) != len(b) {
		return false
	}

	for key, aValues := range a {
		bValues, ok := b[key]
		if !ok {
			return false
		}

		// Expand comma-separated values into multiple values
		// This handles the case where the CLI sends "2023-07-01,2023-08-31"
		// but the index expects ["2023-07-01", "2023-08-31"]
		aExpanded := expandCommaSeparated(aValues)
		bExpanded := expandCommaSeparated(bValues)

		// Sort both slices for comparison
		aSorted := make([]string, len(aExpanded))
		copy(aSorted, aExpanded)
		slices.Sort(aSorted)

		bSorted := make([]string, len(bExpanded))
		copy(bSorted, bExpanded)
		slices.Sort(bSorted)

		if len(aSorted) != len(bSorted) {
			return false
		}

		for i := range aSorted {
			if aSorted[i] != bSorted[i] {
				return false
			}
		}
	}

	return true
}

// expandCommaSeparated splits comma-separated values in a slice
func expandCommaSeparated(values []string) []string {
	var result []string
	for _, value := range values {
		if strings.Contains(value, ",") {
			result = append(result, strings.Split(value, ",")...)
		} else {
			result = append(result, value)
		}
	}
	return result
}

func TestGlobalQuote(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "GLOBAL_QUOTE", "--symbol=IBM")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("command failed: %v\nOutput: %s", err, output)
	}

	if !strings.Contains(string(output), "IBM") {
		t.Errorf("expected output to contain 'IBM', got: %s", output)
	}
}

func TestTimeSeriesDaily(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "TIME_SERIES_DAILY", "--symbol=IBM")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("command failed: %v\nOutput: %s", err, output)
	}

	if !strings.Contains(string(output), "timestamp") {
		t.Errorf("expected output to contain 'timestamp', got: %s", output)
	}
}

func TestWTI(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "WTI", "--interval=monthly")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("command failed: %v\nOutput: %s", err, output)
	}

	if !strings.Contains(string(output), "value") {
		t.Errorf("expected output to contain 'value', got: %s", output)
	}
}

func TestHelp(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "help")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("command failed: %v\nOutput: %s", err, output)
	}

	if !strings.Contains(string(output), "AlphaVantage CLI") {
		t.Errorf("expected output to contain 'AlphaVantage CLI', got: %s", output)
	}
}

func TestUnknownFunction(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "UNKNOWN_FUNCTION")
	output, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatal("expected command to fail, but it succeeded")
	}

	if !strings.Contains(string(output), "unknown function") {
		t.Errorf("expected error message about unknown function, got: %s", output)
	}
}

func TestMissingRequiredParameter(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "GLOBAL_QUOTE")
	output, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatal("expected command to fail, but it succeeded")
	}

	if !strings.Contains(string(output), "required flag") {
		t.Errorf("expected error about required flag, got: %s", output)
	}
}

// TestGeneratedArgs runs all test cases from testdata/args.json
func TestGeneratedArgs(t *testing.T) {
	type testCase struct {
		Name string   `json:"name"`
		Args []string `json:"args"`
	}

	// Load test cases
	testDataPath := filepath.Join("testdata", "args.json")
	data, err := os.ReadFile(testDataPath)
	if err != nil {
		t.Fatalf("failed to read test data: %v", err)
	}

	var testCases []testCase
	if err := json.Unmarshal(data, &testCases); err != nil {
		t.Fatalf("failed to parse test data: %v", err)
	}

	t.Setenv(alphavantage.APIURLEnvironmentVariableName, fakeServer.URL)

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			// Prepend "go run ." to the args
			cmdArgs := append([]string{"run", "."}, tc.Args...)
			cmd := exec.Command("go", cmdArgs...)

			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("command failed: %v\nOutput: %s", err, output)
			}

			// Check that we got some output
			if len(output) == 0 {
				t.Error("expected non-empty output")
			}
		})
	}
}
