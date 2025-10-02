// Package alphavantage provides a Go client for the AlphaVantage financial data API.
//
// It supports fetching stock quotes, time series data, company fundamentals,
// and symbol search functionality from https://www.alphavantage.co.
//
// See the package examples for usage patterns:
package alphavantage

import (
	"bufio"
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"iter"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

const (
	// StandardTokenEnvironmentVariableName is the standard environment variable
	// name for storing the AlphaVantage API key.
	StandardTokenEnvironmentVariableName = "ALPHA_VANTAGE_TOKEN"

	// StandardURLEnvironmentVariableName is the standard environment variable
	// name for storing the AlphaVantage API base URL.
	StandardURLEnvironmentVariableName = "ALPHA_VANTAGE_URL"
)

// DefaultDateFormat is the RFC 3339 date format used for parsing dates.
const DefaultDateFormat = "2006-01-02"

// Client represents an AlphaVantage API client with configurable rate limiting
// and HTTP client behavior.
type Client struct {
	// Limiter controls the rate at which API requests are made.
	// The default limiter allows 5 requests per minute to comply with
	// free tier limits.
	Limiter interface {
		Wait(ctx context.Context) error
	}

	// Client is the HTTP client used for making requests.
	// Defaults to http.DefaultClient.
	Client interface {
		Do(*http.Request) (*http.Response, error)
	}

	// APIKey is the AlphaVantage API key used for authentication.
	APIKey string

	// PrimaryScheme is the URL scheme for the primary API endpoint (http or https).
	// Defaults to "https" if not specified.
	PrimaryScheme string

	// PrimaryHost is the hostname for the primary API endpoint.
	// Defaults to "www.alphavantage.co" if not specified.
	PrimaryHost string

	// FallbackScheme is the URL scheme for the fallback API endpoint (http or https).
	// If specified along with FallbackHost, requests will retry using this URL if the primary fails.
	FallbackScheme string

	// FallbackHost is the hostname for the fallback API endpoint.
	// If specified along with FallbackScheme, requests will retry using this URL if the primary fails.
	FallbackHost string
}

// NewClient creates a new AlphaVantage client with the specified API key.
// It uses default rate limiting (5 requests per minute) and the default HTTP client.
// The client will use environment variable ALPHA_VANTAGE_URL if set, otherwise defaults
// to https://www.alphavantage.co.
func NewClient(apiKey string) *Client {
	return &Client{
		Client:  http.DefaultClient,
		Limiter: rate.NewLimiter(rate.Every(time.Minute/5), 5),
		APIKey:  apiKey,
	}
}

// getSchemeAndHost returns the scheme and host for API requests.
// It checks environment variable ALPHA_VANTAGE_URL first, then falls back to
// client configuration, and finally defaults to https://www.alphavantage.co.
func (client *Client) getSchemeAndHost() (scheme, host string) {
	// Try environment variable first
	if envURL := getEnvURL(); envURL != "" {
		return parseURLSchemeHost(envURL)
	}

	// Use client configuration or defaults
	scheme = client.PrimaryScheme
	if scheme == "" {
		scheme = "https"
	}

	host = client.PrimaryHost
	if host == "" {
		host = "www.alphavantage.co"
	}

	return scheme, host
}

// hasFallback returns true if the client has a fallback URL configured.
func (client *Client) hasFallback() bool {
	return client.FallbackHost != "" && client.FallbackScheme != ""
}

// getEnvURL returns the ALPHA_VANTAGE_URL environment variable value if set.
func getEnvURL() string {
	return os.Getenv(StandardURLEnvironmentVariableName)
}

// parseURLSchemeHost parses a URL string and extracts scheme and host.
// Returns "https" and "www.alphavantage.co" as defaults if parsing fails.
func parseURLSchemeHost(urlStr string) (scheme, host string) {
	u, err := url.Parse(urlStr)
	if err != nil || u.Host == "" {
		return "https", "www.alphavantage.co"
	}

	scheme = u.Scheme
	if scheme == "" {
		scheme = "https"
	}

	return scheme, u.Host
}

// doWithFallback executes a request with automatic fallback support.
// It first tries the primary URL, and if that fails and a fallback is configured,
// it retries with the fallback URL.
func (client *Client) doWithFallback(ctx context.Context, makeURL func(scheme, host string) string) (io.ReadCloser, error) {
	// Try primary URL first
	scheme, host := client.getSchemeAndHost()
	requestURL := makeURL(scheme, host)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	res, primaryErr := client.Do(req)
	if primaryErr == nil {
		return checkError(res.Body)
	}

	// If primary failed and we have a fallback, try it
	if client.hasFallback() {
		fallbackURL := makeURL(client.FallbackScheme, client.FallbackHost)

		req, err = http.NewRequestWithContext(ctx, http.MethodGet, fallbackURL, nil)
		if err != nil {
			return nil, fmt.Errorf("primary request failed: %w; fallback request creation failed: %v", primaryErr, err)
		}

		res, fallbackErr := client.Do(req)
		if fallbackErr == nil {
			return checkError(res.Body)
		}

		return nil, fmt.Errorf("primary request failed: %w; fallback request failed: %v", primaryErr, fallbackErr)
	}

	return nil, primaryErr
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

	q := req.URL.Query()
	q.Set("apikey", client.APIKey)
	req.URL.RawQuery = q.Encode()

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

func checkError(rc io.ReadCloser) (io.ReadCloser, error) {
	var buf [1]byte
	n, err := rc.Read(buf[:])
	if err != nil {
		return nil, fmt.Errorf("could not read request response: %w", err)
	}

	mr := io.MultiReader(bytes.NewReader(buf[:]), rc)
	if n > 0 && buf[0] == '{' {
		var message struct {
			Note         string `json:"Note,omitempty"`
			Information  string `json:"Information,omitempty"`
			ErrorMessage string `json:"Error Message,omitempty"`
			Detail       string `json:"detail,omitempty"`
		}
		err = json.NewDecoder(mr).Decode(&message)
		if err != nil {
			return nil, fmt.Errorf("could not read response for: %w", err)
		}
		if strings.Contains(message.Note, " higher API call frequency") {
			return nil, fmt.Errorf("reached alphavantage rate limit")
		}

		if message.ErrorMessage != "" {
			return nil, fmt.Errorf("alphavantage request did not return csv; got notice: %w", errors.New(message.ErrorMessage))
		}
		if message.Detail != "" {
			return nil, fmt.Errorf("alphavantage request did not return csv; got notice: %w", errors.New(message.Detail))
		}
		if message.Note != "" || message.Information != "" {
			return nil, fmt.Errorf("alphavantage request did not return csv; got notice: %w", errors.New(strings.Join([]string{message.Note, message.Information}, " ")))
		}

		return nil, fmt.Errorf("alphavantage request did not return csv")
	}

	return multiReadCloser{
		Reader: mr,
		close:  rc.Close,
	}, nil
}

var typeType = reflect.TypeOf(time.Time{})

// ParseCSV parses CSV data into a slice of structs using reflection.
//
// Supported field types:
//   - string: Direct mapping from CSV column value
//   - int: Parsed using strconv.ParseInt with base 10
//   - float64: Parsed using strconv.ParseFloat
//   - time.Time: Parsed using time.ParseInLocation (see time-layout tag)
//
// Struct field tags:
//   - `column-name:"header"`: Maps field to CSV column header (required)
//   - `time-layout:"layout"`: Custom time format for time.Time fields (optional, defaults to "2006-01-02")
//
// Example struct:
//
//	type StockPrice struct {
//	    Date   time.Time `column-name:"timestamp"`
//	    Open   float64   `column-name:"open"`
//	    High   float64   `column-name:"high"`
//	    Volume int       `column-name:"volume"`
//	}
//
// Unmapped columns are ignored. Fields without matching columns keep their zero value.
// Time fields with "null" values remain as zero time.Time.
func ParseCSV[T any](r io.Reader, data *[]T, location *time.Location) error {
	if data == nil {
		panic(fmt.Errorf("data must not be nil"))
	}
	var err error
	for row := range ParseCSVRows[T](r, location, func(e error) bool {
		err = e
		return false
	}) {
		*data = append(*data, row)
	}
	return err
}

// ParseCSVRows returns an iterator that parses CSV data row by row into structs.
// This is memory-efficient for large datasets as it processes one row at a time.
//
// Uses the same struct field tagging system as ParseCSV:
//   - `column-name:"header"`: Maps field to CSV column header (required)
//   - `time-layout:"layout"`: Custom time format for time.Time fields (optional)
//
// The handleErr function is called when parsing errors occur. Return true to continue
// processing, false to stop. Location defaults to UTC if nil.
//
// Example usage:
//
//	for price := range ParseCSVRows[StockPrice](reader, time.UTC, func(err error) bool {
//	    log.Printf("Parse error: %v", err)
//	    return true // continue on errors
//	}) {
//	    fmt.Printf("Price: %+v\n", price)
//	}
func ParseCSVRows[T any](r io.Reader, location *time.Location, handleErr func(error) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		if location == nil {
			location = time.UTC
		}

		rowType := reflect.TypeFor[T]()

		reader := csv.NewReader(bufio.NewReader(r))
		reader.TrimLeadingSpace = true
		header, err := reader.Read()
		if err != nil {
			handleErr(err)
			return
		}
		reader.FieldsPerRecord = len(header)

		if rowType.Kind() != reflect.Struct {
			panic(fmt.Errorf("expected a struct kind: got %s", rowType.Kind()))
			return
		}

		structType := rowType

		columnToField := make(map[int]int, len(header))
		for columnHeaderIndex, columnHeaderName := range header {
			for fieldIndex := 0; fieldIndex < structType.NumField(); fieldIndex++ {
				fieldType := structType.Field(fieldIndex)

				csvTag := fieldType.Tag.Get("column-name")
				if csvTag != columnHeaderName {
					continue
				}

				columnToField[columnHeaderIndex] = fieldIndex
			}
		}

		for rowIndex := 1; ; rowIndex++ {
			row, err := reader.Read()
			if err != nil {
				if err == io.EOF {
					return
				}
				handleErr(err)
				return
			}

			structValue := reflect.New(structType)

			for columnIndex, value := range row {
				fieldIndex, ok := columnToField[columnIndex]
				if !ok {
					continue
				}

				structFieldType := structType.Field(fieldIndex)

				switch structFieldType.Type.Kind() {
				case reflect.String:
					structValue.Elem().Field(fieldIndex).SetString(value)
				case reflect.Float64:
					fl, err := strconv.ParseFloat(value, 64)
					if err != nil {
						if handleErr(fmt.Errorf("failed to parse float64 value %q on row %d column %d (%s): %w", value, rowIndex, columnIndex, header[columnIndex], err)) {
							continue
						}
						return
					}
					structValue.Elem().Field(fieldIndex).SetFloat(fl)
				case reflect.Int:
					in, err := strconv.ParseInt(value, 10, 64)
					if err != nil {
						if !handleErr(fmt.Errorf("failed to parse float64 value %q on row %d column %d (%s): %w", value, rowIndex, columnIndex, header[columnIndex], err)) {
							continue
						}
						return
					}
					structValue.Elem().Field(fieldIndex).SetInt(in)
				default:
					if structFieldType.Type != typeType {
						if handleErr(fmt.Errorf("unsupported type %T for field %s", structFieldType.Type, structFieldType.Name)) {
							continue
						}
						return
					}

					layout := DefaultDateFormat
					tagLayout := structFieldType.Tag.Get("time-layout")
					if tagLayout != "" {
						layout = tagLayout
					}
					if value == "null" {
						continue
					}
					tm, err := time.ParseInLocation(layout, value, location)
					if err != nil {
						if handleErr(fmt.Errorf("failed to parse time value on row %d column %d (%s): %w", rowIndex, columnIndex, header[columnIndex], err)) {
							continue
						}
						return
					}
					structValue.Elem().Field(fieldIndex).Set(reflect.ValueOf(tm))
				}
			}

			if !yield(structValue.Elem().Interface().(T)) {
				return
			}
		}
	}
}
