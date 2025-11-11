package alphavantage

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_checkError(t *testing.T) {
	t.Run("error message", func(t *testing.T) {
		rc := io.NopCloser(bytes.NewBufferString(`{"Error Message": "the parameter apikey is invalid or missing. Please claim your free API key on (https://www.alphavantage.co/support/#api-key). It should take less than 20 seconds."}`))

		_, err := checkError(rc)
		require.ErrorContains(t, err, "the parameter apikey")
	})

	t.Run("detail", func(t *testing.T) {
		rc := io.NopCloser(bytes.NewBufferString(`{"detail": "Could not satisfy the request Accept header."}`))
		_, err := checkError(rc)
		require.ErrorContains(t, err, "Could not satisfy")
	})
}

func TestClient_GlobalQuote(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping unit test in short mode")
	}

	// Mock client that intercepts HTTP requests
	mockClient := &Client{
		Client: doerFunc(func(req *http.Request) (*http.Response, error) {
			// Verify the request
			assert.Equal(t, "/query", req.URL.Path)
			assert.Equal(t, "GLOBAL_QUOTE", req.URL.Query().Get("function"))
			assert.Equal(t, "IBM", req.URL.Query().Get("symbol"))
			assert.Equal(t, "csv", req.URL.Query().Get("datatype"))
			assert.Equal(t, "test-key", req.URL.Query().Get("apikey"))

			// Return mock CSV response
			mockResponse := `symbol,open,high,low,price,volume,latestDay,previousClose,change,changePercent
IBM,129.00,130.50,128.50,129.75,1234567,2023-12-01,129.25,0.50,0.3867%`

			return &http.Response{
				StatusCode: 200,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader([]byte(mockResponse))),
			}, nil
		}),
		Limiter: waitFunc(func(ctx context.Context) error { return nil }),
		APIKey:  "test-key",
	}

	ctx := context.Background()
	res, err := mockClient.GetGlobalQuote(ctx, QueryGlobalQuote(mockClient.APIKey, "IBM").DataTypeCSV())

	require.NoError(t, err)
	defer res.Body.Close()

	// Verify we can read the CSV response
	content, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	responseText := string(content)
	assert.Contains(t, responseText, "symbol,open,high,low,price,volume")
	assert.Contains(t, responseText, "IBM,129.00,130.50")
}

func TestClient_GlobalQuote_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	token := os.Getenv("ALPHA_VANTAGE_TOKEN")
	if token == "" {
		t.Skip("ALPHA_VANTAGE_TOKEN not set, skipping integration test")
	}

	client := NewClient()
	ctx := context.Background()

	res, err := client.GetGlobalQuote(ctx, QueryGlobalQuote(token, "IBM"))
	require.NoError(t, err)
	defer res.Body.Close()

	// Verify we get valid CSV data
	content, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	responseText := string(content)
	// Should contain the CSV header
	assert.Contains(t, responseText, "symbol")
	// And should contain IBM data
	assert.Contains(t, responseText, "IBM")
}

// Helper types for mocking
type doerFunc func(*http.Request) (*http.Response, error)

func (fn doerFunc) Do(req *http.Request) (*http.Response, error) { return fn(req) }

type waitFunc func(ctx context.Context) error

func (wf waitFunc) Wait(ctx context.Context) error {
	return wf(ctx)
}
