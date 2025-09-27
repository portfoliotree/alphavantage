package alphavantage_test

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/portfoliotree/alphavantage"
)

func TestClient_ListingStatus_listed(t *testing.T) {
	f, err := os.Open(filepath.FromSlash("testdata/listing_status.csv"))
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
	}).ListingStatus(ctx, true)
	require.NoError(t, err)
	assert.Len(t, results, 8)

	assert.Equal(t, "www.alphavantage.co", avReq.Host)
	assert.Equal(t, "https", avReq.URL.Scheme)
	assert.Equal(t, "/query", avReq.URL.Path)
	assert.Equal(t, "LISTING_STATUS", avReq.URL.Query().Get("function"))
	assert.Equal(t, "active", avReq.URL.Query().Get("state"))
	assert.Equal(t, "demo", avReq.URL.Query().Get("apikey"))
	assert.Equal(t, "csv", avReq.URL.Query().Get("datatype"))
	assert.Equal(t, 1, waitCallCount)
}

func TestClient_ListingStatus_delisted(t *testing.T) {
	f, err := os.Open(filepath.FromSlash("testdata/listing_status.csv"))
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

	_, err = (&alphavantage.Client{
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
	}).ListingStatus(ctx, false)
	require.NoError(t, err)
	assert.Equal(t, "delisted", avReq.URL.Query().Get("state"))
}
