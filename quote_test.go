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
	result, err := mockClient.GlobalQuote(ctx, "IBM")

	require.NoError(t, err)
	defer result.Close()

	// Verify we can read the CSV response
	content, err := io.ReadAll(result)
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

	client := NewClient(token)
	ctx := context.Background()

	result, err := client.GlobalQuote(ctx, "IBM")
	require.NoError(t, err)
	defer result.Close()

	// Verify we get valid CSV data
	content, err := io.ReadAll(result)
	require.NoError(t, err)

	responseText := string(content)
	// Should contain the CSV header
	assert.Contains(t, responseText, "symbol")
	// And should contain IBM data
	assert.Contains(t, responseText, "IBM")
}

func TestTimeSeriesWeekly_Constants(t *testing.T) {
	// Test that new weekly constants exist and validate
	require.NoError(t, TimeSeriesWeekly.Validate())
	require.NoError(t, TimeSeriesWeeklyAdjusted.Validate())

	// Test constant values
	assert.Equal(t, "TIME_SERIES_WEEKLY", string(TimeSeriesWeekly))
	assert.Equal(t, "TIME_SERIES_WEEKLY_ADJUSTED", string(TimeSeriesWeeklyAdjusted))
}

func TestClient_WeeklyQuotes(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping unit test in short mode")
	}

	// Mock client that intercepts HTTP requests
	mockClient := &Client{
		Client: doerFunc(func(req *http.Request) (*http.Response, error) {
			// Verify the request
			assert.Equal(t, "/query", req.URL.Path)
			assert.Equal(t, "TIME_SERIES_WEEKLY", req.URL.Query().Get("function"))
			assert.Equal(t, "IBM", req.URL.Query().Get("symbol"))
			assert.Equal(t, "csv", req.URL.Query().Get("datatype"))
			assert.Equal(t, "test-key", req.URL.Query().Get("apikey"))

			// Return mock CSV response
			mockResponse := `timestamp,open,high,low,close,volume
2023-12-01,129.00,130.50,128.50,129.75,1234567
2023-11-24,128.25,129.80,127.90,129.00,1098765`

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
	result, err := mockClient.DoQuotesRequest(ctx, "IBM", TimeSeriesWeekly)

	require.NoError(t, err)
	defer result.Close()

	// Verify we can read the CSV response
	content, err := io.ReadAll(result)
	require.NoError(t, err)

	responseText := string(content)
	assert.Contains(t, responseText, "timestamp,open,high,low,close,volume")
	assert.Contains(t, responseText, "2023-12-01,129.00,130.50")
	assert.Contains(t, responseText, "2023-11-24,128.25,129.80")
}

func TestClient_WeeklyAdjustedQuotes(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping unit test in short mode")
	}

	// Mock client that intercepts HTTP requests
	mockClient := &Client{
		Client: doerFunc(func(req *http.Request) (*http.Response, error) {
			// Verify the request for adjusted weekly data
			assert.Equal(t, "/query", req.URL.Path)
			assert.Equal(t, "TIME_SERIES_WEEKLY_ADJUSTED", req.URL.Query().Get("function"))
			assert.Equal(t, "IBM", req.URL.Query().Get("symbol"))
			assert.Equal(t, "csv", req.URL.Query().Get("datatype"))
			assert.Equal(t, "test-key", req.URL.Query().Get("apikey"))

			// Return mock CSV response with adjusted fields
			mockResponse := `timestamp,open,high,low,close,adjusted_close,volume,dividend_amount
2023-12-01,129.00,130.50,128.50,129.75,129.75,1234567,0.0000
2023-11-24,128.25,129.80,127.90,129.00,129.00,1098765,0.0000`

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
	result, err := mockClient.DoQuotesRequest(ctx, "IBM", TimeSeriesWeeklyAdjusted)

	require.NoError(t, err)
	defer result.Close()

	// Verify we can read the CSV response
	content, err := io.ReadAll(result)
	require.NoError(t, err)

	responseText := string(content)
	assert.Contains(t, responseText, "timestamp,open,high,low,close,adjusted_close,volume,dividend_amount")
	assert.Contains(t, responseText, "2023-12-01,129.00,130.50")
}

func TestClient_WeeklyQuotes_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	token := os.Getenv("ALPHA_VANTAGE_TOKEN")
	if token == "" {
		t.Skip("ALPHA_VANTAGE_TOKEN not set, skipping integration test")
	}

	client := NewClient(token)
	ctx := context.Background()

	// Test weekly data
	result, err := client.DoQuotesRequest(ctx, "IBM", TimeSeriesWeekly)
	require.NoError(t, err)
	defer result.Close()

	// Verify we get valid CSV data
	content, err := io.ReadAll(result)
	require.NoError(t, err)

	responseText := string(content)
	// Should contain the CSV header
	assert.Contains(t, responseText, "timestamp")
	// And should contain IBM data with weekly intervals
	assert.Contains(t, responseText, ",") // CSV format check
}

// Helper types for mocking
type doerFunc func(*http.Request) (*http.Response, error)

func (fn doerFunc) Do(req *http.Request) (*http.Response, error) { return fn(req) }

type waitFunc func(ctx context.Context) error

func (wf waitFunc) Wait(ctx context.Context) error {
	return wf(ctx)
}
