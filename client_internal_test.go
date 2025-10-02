package alphavantage

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

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

func TestClient_FallbackURL(t *testing.T) {
	t.Run("uses primary URL when successful", func(t *testing.T) {
		var primaryCalls, fallbackCalls atomic.Int32

		csvData := "symbol,open,high,low,price,volume,latestDay,previousClose,change,changePercent\nTEST,100.0,105.0,99.0,102.0,1000,2024-01-01,101.0,1.0,0.99%\n"

		primary := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			primaryCalls.Add(1)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(csvData))
		}))
		defer primary.Close()

		fallback := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fallbackCalls.Add(1)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(csvData))
		}))
		defer fallback.Close()

		primaryScheme, primaryHost := parseURLSchemeHost(primary.URL)
		fallbackScheme, fallbackHost := parseURLSchemeHost(fallback.URL)

		client := &Client{
			Client:         http.DefaultClient,
			Limiter:        nil,
			APIKey:         "test",
			PrimaryHost:    primaryHost,
			PrimaryScheme:  primaryScheme,
			FallbackHost:   fallbackHost,
			FallbackScheme: fallbackScheme,
		}

		_, err := client.GlobalQuote(context.Background(), "TEST")
		require.NoError(t, err)
		require.Equal(t, int32(1), primaryCalls.Load())
		require.Equal(t, int32(0), fallbackCalls.Load())
	})

	t.Run("uses fallback URL when primary fails", func(t *testing.T) {
		var primaryCalls, fallbackCalls atomic.Int32

		csvData := "symbol,open,high,low,price,volume,latestDay,previousClose,change,changePercent\nTEST,100.0,105.0,99.0,102.0,1000,2024-01-01,101.0,1.0,0.99%\n"

		primary := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			primaryCalls.Add(1)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("primary server error"))
		}))
		defer primary.Close()

		fallback := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fallbackCalls.Add(1)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(csvData))
		}))
		defer fallback.Close()

		primaryScheme, primaryHost := parseURLSchemeHost(primary.URL)
		fallbackScheme, fallbackHost := parseURLSchemeHost(fallback.URL)

		client := &Client{
			Client:         http.DefaultClient,
			Limiter:        nil,
			APIKey:         "test",
			PrimaryHost:    primaryHost,
			PrimaryScheme:  primaryScheme,
			FallbackHost:   fallbackHost,
			FallbackScheme: fallbackScheme,
		}

		_, err := client.GlobalQuote(context.Background(), "TEST")
		require.NoError(t, err)
		require.Equal(t, int32(1), primaryCalls.Load())
		require.Equal(t, int32(1), fallbackCalls.Load())
	})

	t.Run("returns error when both URLs fail", func(t *testing.T) {
		var primaryCalls, fallbackCalls atomic.Int32

		primary := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			primaryCalls.Add(1)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("primary server error"))
		}))
		defer primary.Close()

		fallback := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fallbackCalls.Add(1)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("fallback server error"))
		}))
		defer fallback.Close()

		primaryScheme, primaryHost := parseURLSchemeHost(primary.URL)
		fallbackScheme, fallbackHost := parseURLSchemeHost(fallback.URL)

		client := &Client{
			Client:         http.DefaultClient,
			Limiter:        nil,
			APIKey:         "test",
			PrimaryHost:    primaryHost,
			PrimaryScheme:  primaryScheme,
			FallbackHost:   fallbackHost,
			FallbackScheme: fallbackScheme,
		}

		_, err := client.GlobalQuote(context.Background(), "TEST")
		require.Error(t, err)
		require.Equal(t, int32(1), primaryCalls.Load())
		require.Equal(t, int32(1), fallbackCalls.Load())
	})
}
