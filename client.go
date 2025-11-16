// Package alphavantage provides a Go client for the AlphaVantage financial data API.
//
// It supports fetching stock quotes, time series data, company fundamentals,
// and symbol search functionality from https://www.alphavantage.co.
//
// See the package examples for usage patterns:
package alphavantage

import (
	"cmp"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"golang.org/x/time/rate"

	"github.com/portfoliotree/alphavantage/api"
)

//go:generate go run ./cmd/av-generate

const (
	envVarPrefix = "ALPHA_VANTAGE_"

	// APIKeyEnvironmentVariableName is the standard environment variable
	// name for storing the AlphaVantage API key.
	APIKeyEnvironmentVariableName = envVarPrefix + "API_KEY"

	// APIURLEnvironmentVariableName is the standard environment variable
	// name for overriding the AlphaVantage API URL.
	APIURLEnvironmentVariableName = envVarPrefix + "URL"

	// RequestsPerMinuteEnvironmentVariableName is the number of requests per minute
	// the API rate limiter should be configured to permit.
	RequestsPerMinuteEnvironmentVariableName = envVarPrefix + "REQUEST_PER_MINUTE"
)

// Client represents an AlphaVantage API client with configurable rate limiting
// and HTTP client behavior.
type Client struct {
	// Limiter controls the rate at which API requests are made.
	// The default limiter allows 5 requests per minute to comply with
	// free tier limits.
	Limiter Waiter

	// Client is the HTTP client used for making requests.
	// Defaults to http.DefaultClient.
	Client interface {
		Do(*http.Request) (*http.Response, error)
	}

	// APIKey is the AlphaVantage API key used for authentication.
	APIKey string

	BaseURL url.URL
}

type Waiter interface {
	Wait(ctx context.Context) error
}

type Doer = api.Doer

// NewClient creates a new AlphaVantage client with the specified API key.
// The client will use environment variable ALPHA_VANTAGE_URL if set, otherwise defaults
// to https://www.alphavantage.co.
func NewClient() *Client {
	var limit Waiter = nil
	if val, ok := os.LookupEnv(RequestsPerMinuteEnvironmentVariableName); ok {
		n, err := strconv.Atoi(val)
		if err != nil {
			slog.Error("failed to parse requests per minute environment variable while setting up alphavantage client",
				slog.String("message", ""),
				slog.String("error", err.Error()),
			)
		} else {
			limit = rate.NewLimiter(RequestsPerMinute(n).Limit(), n)
		}
	}
	var baseURL url.URL
	if bu, err := url.Parse(cmp.Or(os.Getenv(APIURLEnvironmentVariableName), "https://www.alphavantage.co")); err != nil {
		bu, _ = url.Parse("https://www.alphavantage.co")
		baseURL = *bu
		slog.Error("failed to parse api URL for alphavantage client",
			slog.String("error", err.Error()),
		)
	} else {
		baseURL = *bu
	}

	return &Client{
		Client:  http.DefaultClient,
		Limiter: limit,
		APIKey:  cmp.Or(os.Getenv(APIKeyEnvironmentVariableName), os.Getenv("ALPHA_VANTAGE_TOKEN"), "demo"),
		BaseURL: baseURL,
	}
}

func (client *Client) QueryRequest(ctx context.Context, values url.Values) (*http.Request, error) {
	return http.NewRequestWithContext(ctx,
		http.MethodGet,
		(&url.URL{
			Scheme:   api.DefaultScheme,
			Host:     api.DefaultHost,
			Path:     api.DefaultPath,
			RawQuery: values.Encode(),
		}).String(),
		nil,
	)
}

func (client *Client) Do(req *http.Request) (*http.Response, error) {
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	if client.Limiter != nil {
		err := client.Limiter.Wait(req.Context())
		if err != nil {
			return &http.Response{}, err
		}
	}

	res, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 300 || res.StatusCode < 200 {
		buf, err := io.ReadAll(io.LimitReader(res.Body, 1<<10))
		if err != nil {
			buf = []byte(err.Error())
		}
		return res, fmt.Errorf("request failed with status %d %s: %s",
			res.StatusCode, http.StatusText(res.StatusCode), string(buf))
	}

	return res, nil
}

func closeAndIgnoreError(c io.Closer) {
	_ = c.Close()
}
