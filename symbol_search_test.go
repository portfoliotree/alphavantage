package alphavantage_test

import (
	"context"
	_ "embed"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/portfoliotree/alphavantage"
)

func TestSearch(t *testing.T) {
	f, err := os.Open("test_data/search_results.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()

	ctx := context.Background()

	var (
		avReq         *http.Request
		waitCallCount = 0
	)

	results, err := (&alphavantage.Client{
		Client: doerFunc(func(request *http.Request) (*http.Response, error) {
			avReq = request
			return &http.Response{
				Body:       io.NopCloser(f),
				StatusCode: http.StatusOK,
			}, nil
		}),
		APIKey: "demo",
		Limiter: waitFunc(func(ctx context.Context) error {
			waitCallCount++
			return nil
		}),
	}).SymbolSearch(ctx, "BA")
	require.NoError(t, err)
	assert.Len(t, results, 10)

	assert.Equal(t, "www.alphavantage.co", avReq.Host)
	assert.Equal(t, "https", avReq.URL.Scheme)
	assert.Equal(t, "/query", avReq.URL.Path)
	assert.Equal(t, "SYMBOL_SEARCH", avReq.URL.Query().Get("function"))
	assert.Equal(t, "BA", avReq.URL.Query().Get("keywords"))
	assert.Equal(t, "demo", avReq.URL.Query().Get("apikey"))
	assert.Equal(t, "csv", avReq.URL.Query().Get("datatype"))
	assert.Equal(t, 1, waitCallCount)
}

func TestParseSearchQuery(t *testing.T) {
	f, err := os.Open("test_data/search_results.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()

	results, err := alphavantage.ParseSymbolSearchQuery(f)
	require.NoError(t, err)
	assert.Len(t, results, 10)

	assert.Equal(t, []alphavantage.SymbolSearchResult{
		{
			Symbol:      "BA",
			Name:        "Boeing Company",
			Type:        "Equity",
			Region:      "United States",
			MarketOpen:  "09:30",
			MarketClose: "16:00",
			TimeZone:    "UTC-04",
			Currency:    "USD",
			MatchScore:  1,
		},
		{
			Symbol:      "BAB",
			Name:        "Invesco Taxable Municipal Bond ETF",
			Type:        "ETF",
			Region:      "United States",
			MarketOpen:  "09:30",
			MarketClose: "16:00",
			TimeZone:    "UTC-04",
			Currency:    "USD",
			MatchScore:  0.8,
		},
	}, results[:2])
}
