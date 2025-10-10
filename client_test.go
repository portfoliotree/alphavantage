package alphavantage_test

import (
	"bytes"
	"context"
	"crypto/tls"
	_ "embed"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/portfoliotree/alphavantage"
)

type doerFunc func(*http.Request) (*http.Response, error)

func (fn doerFunc) Do(req *http.Request) (*http.Response, error) { return fn(req) }

type waitFunc func(ctx context.Context) error

func (wf waitFunc) Wait(ctx context.Context) error {
	return wf(ctx)
}

func TestParse(t *testing.T) {
	t.Run("nil data", func(t *testing.T) {
		assert.Panics(t, func() {
			_ = alphavantage.ParseCSV(bytes.NewReader(nil), (*[]alphavantage.Quote)(nil), nil)
		})
	})

	t.Run("real data", func(t *testing.T) {
		var someFolks []struct {
			ID           int       `column-name:"id"`
			FirstInitial string    `column-name:"first_initial"`
			BirthDate    time.Time `column-name:"birth_date" time-layout:"2006/01/02"`
			Mass         float64   `column-name:"mass"`
		}

		err := alphavantage.ParseCSV(strings.NewReader(panthersCSV), &someFolks, nil)
		require.NoError(t, err)
		assert.Len(t, someFolks, 3)

		assert.Equal(t, 1, someFolks[0].ID)
		assert.Equal(t, "N", someFolks[0].FirstInitial)
		assert.Equal(t, mustParseDate(t, "2020-02-17"), someFolks[0].BirthDate)
		assert.Equal(t, 70.0, someFolks[0].Mass)

		assert.Equal(t, 2, someFolks[1].ID)
		assert.Equal(t, "S", someFolks[1].FirstInitial)
		assert.Equal(t, mustParseDate(t, "2020-10-22"), someFolks[1].BirthDate)
		assert.Equal(t, 68.2, someFolks[1].Mass)

		assert.Equal(t, 3, someFolks[2].ID)
		assert.Equal(t, "C", someFolks[2].FirstInitial)
		assert.Equal(t, mustParseDate(t, "2021-08-31"), someFolks[2].BirthDate)
		assert.Equal(t, 72.9, someFolks[2].Mass)
	})
}

const panthersCSV = `id,first_initial,birth_date,mass
1, N, 2020/02/17, 70
2, S, 2020/10/22, 68.2
3, C, 2021/08/31, 72.9
`

func mustParseDate(t *testing.T, date string) time.Time {
	tm, err := time.ParseInLocation(alphavantage.DefaultDateFormat, date, time.UTC)
	if err != nil {
		t.Fatal(err)
	}
	return tm
}

// TestClientHostConfiguration verifies that the Client uses a configured host
// when making API requests. This test creates a fake server and ensures the
// client targets it instead of the default AlphaVantage host.
func TestClientHostConfiguration(t *testing.T) {
	// Create a TLS test server that returns mock CSV data
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request has the API key query parameter
		assert.Equal(t, "test-api-key", r.URL.Query().Get("apikey"))

		// Return mock CSV data
		w.Header().Set("Content-Type", "text/csv")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("timestamp,open,high,low,close,volume,dividend_amount,split_coefficient\n2024-01-01,100,110,90,105,1000000,0,1\n"))
	}))
	defer server.Close()

	// Create a client with the test server's host and a custom HTTP client that accepts self-signed certs
	client := alphavantage.NewClient("test-api-key")
	client.Host = server.URL[len("https://"):] // Remove the scheme
	client.Limiter = nil                       // Disable rate limiting for tests
	client.Client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Make a request
	ctx := context.Background()
	_, err := client.DoQuotesRequest(ctx, "TEST", alphavantage.TimeSeriesDaily)

	require.NoError(t, err)
}

// TestClientHostFallback verifies that when Host is not set, the client
// falls back to the DefaultHost.
func TestClientHostFallback(t *testing.T) {
	client := alphavantage.NewClientWithHost("", "test-api-key")

	// Host should be empty by default
	assert.Equal(t, "", client.Host)

	// When creating a URL, it should use the default host
	url, err := alphavantage.NewQuotesURL("", "TEST", alphavantage.TimeSeriesDaily)
	require.NoError(t, err)

	assert.Contains(t, url, alphavantage.DefaultHost)
}

// TestClientHostEnvironmentVariable verifies that the host can be configured
// via the Client.Host field, which can be set from an environment variable.
func TestClientHostEnvironmentVariable(t *testing.T) {
	// Create a TLS test server
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/csv")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("symbol,name,exchange,assetType,ipoDate,delistingDate,status\nTEST,Test Company,NYSE,Stock,2020-01-01,null,Active\n"))
	}))
	defer server.Close()

	// Simulate setting host from environment variable
	customHost := server.URL[len("https://"):]

	client := alphavantage.NewClient("test-api-key")
	client.Host = customHost
	client.Limiter = nil
	client.Client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Test with ListingStatus endpoint
	ctx := context.Background()
	results, err := client.ListingStatus(ctx, true)

	require.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, "TEST", results[0].Symbol)
}
