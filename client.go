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
	"log"
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

func (client *Client) ETFProfile(ctx context.Context, symbol string) (ETFProfile, error) {
	u := url.URL{
		Scheme: "https",
		Host:   "www.alphavantage.co",
		Path:   "/query",
		RawQuery: url.Values{
			"function": []string{"ETF_PROFILE"},
			"symbol":   []string{symbol},
		}.Encode(),
	}

	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		u.String(),
		nil,
	)
	if err != nil {
		return ETFProfile{}, fmt.Errorf("failed to create ETF profile request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return ETFProfile{}, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return ETFProfile{}, err
	}

	var result ETFProfile
	err = json.Unmarshal(buf, &result)
	return result, err
}

type ETFProfile struct {
	Symbol            string       `json:"symbol,omitempty"`
	NetAssets         string       `json:"net_assets,omitempty"`
	NetExpenseRatio   string       `json:"net_expense_ratio,omitempty"`
	PortfolioTurnover string       `json:"portfolio_turnover,omitempty"`
	DividendYield     string       `json:"dividend_yield,omitempty"`
	InceptionDate     string       `json:"inception_date,omitempty"`
	Leveraged         string       `json:"leveraged,omitempty"`
	Sectors           []ETFSector  `json:"sectors,omitempty"`
	Holdings          []ETFHolding `json:"holdings,omitempty"`
}

type ETFSector struct {
	Sector string `json:"sector,omitempty"`
	Weight string `json:"weight,omitempty"`
}

type ETFHolding struct {
	Symbol      string `json:"symbol,omitempty"`
	Description string `json:"description,omitempty"`
	Weight      string `json:"weight,omitempty"`
}

// QuoteFunction represents the different time series functions available
// from the AlphaVantage API.
type QuoteFunction string

// Time series function constants for different data intervals and types.
const (
	TimeSeriesIntraday        QuoteFunction = "TIME_SERIES_INTRADAY"         // Intraday time series data
	TimeSeriesDaily           QuoteFunction = "TIME_SERIES_DAILY"            // Daily time series data
	TimeSeriesDailyAdjusted   QuoteFunction = "TIME_SERIES_DAILY_ADJUSTED"   // Daily adjusted time series data
	TimeSeriesWeekly          QuoteFunction = "TIME_SERIES_WEEKLY"           // Weekly time series data
	TimeSeriesWeeklyAdjusted  QuoteFunction = "TIME_SERIES_WEEKLY_ADJUSTED"  // Weekly adjusted time series data
	TimeSeriesMonthly         QuoteFunction = "TIME_SERIES_MONTHLY"          // Monthly time series data
	TimeSeriesMonthlyAdjusted QuoteFunction = "TIME_SERIES_MONTHLY_ADJUSTED" // Monthly adjusted time series data
)

// NewQuotesURL constructs a URL for time series data requests.
// It validates the function parameter and returns a properly formatted API URL.
func NewQuotesURL(symbol string, function QuoteFunction) (string, error) {
	err := function.Validate()
	if err != nil {
		return "", err
	}

	u := url.URL{
		Scheme: "https",
		Host:   "www.alphavantage.co",
		Path:   "/query",
		RawQuery: url.Values{
			"datatype":   []string{"csv"},
			"outputsize": []string{"full"},
			"function":   []string{string(function)},
			"symbol":     []string{symbol},
		}.Encode(),
	}

	return u.String(), nil
}

// Validate checks if the QuoteFunction is one of the supported time series functions.
// It returns an error if the function is not recognized.
func (fn QuoteFunction) Validate() error {
	switch fn {
	case TimeSeriesIntraday,
		TimeSeriesDaily,
		TimeSeriesDailyAdjusted,
		TimeSeriesWeekly,
		TimeSeriesWeeklyAdjusted,
		TimeSeriesMonthly,
		TimeSeriesMonthlyAdjusted:
		return nil
	default:
		return errors.New("unknown time series function")
	}
}

// Quotes fetches time series data for the specified symbol and function.
// It parses the CSV response into a slice of Quote structs with dates in the given location.
// The location parameter is used for parsing timestamps; use time.UTC for UTC times.
func (client *Client) Quotes(ctx context.Context, symbol string, function QuoteFunction, location *time.Location) ([]Quote, error) {
	rc, err := client.DoQuotesRequest(ctx, symbol, function)
	if err != nil {
		return nil, err
	}
	defer closeAndIgnoreError(rc)

	switch function {
	case TimeSeriesIntraday:
		list, err := ParseIntraDayQuotes(rc, location)
		if err != nil {
			return nil, err
		}
		return convertQuoteElements(list, func(q IntraDayQuote) Quote { return Quote(q) }), nil
	default:
		quotes, err := ParseQuotes(rc, location)
		if err != nil {
			return nil, err
		}
		return quotes, nil
	}
}

// Deprecated: use DoQuotesRequest instead. This method will be removed before 2023.
func (client *Client) QuotesRequest(ctx context.Context, symbol string, function QuoteFunction, fn func(r io.Reader) error) error {
	rc, err := client.DoQuotesRequest(ctx, symbol, function)
	if err != nil {
		return err
	}
	defer closeAndIgnoreError(rc)
	return fn(rc)
}

// DoQuotesRequest fetches time series data for the specified symbol and function.
// It returns the raw CSV response as an io.ReadCloser that must be closed by the caller.
// This method provides direct access to the CSV data without parsing.
func (client *Client) DoQuotesRequest(ctx context.Context, symbol string, function QuoteFunction) (io.ReadCloser, error) {
	requestURL, err := NewQuotesURL(symbol, function)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		requestURL,
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return checkError(res.Body)
}

type Quote struct {
	Time             time.Time `column-name:"timestamp"`
	Open             float64   `column-name:"open"`
	High             float64   `column-name:"high"`
	Low              float64   `column-name:"low"`
	Close            float64   `column-name:"close"`
	Volume           float64   `column-name:"volume"`
	DividendAmount   float64   `column-name:"dividend_amount"`
	SplitCoefficient float64   `column-name:"split_coefficient"`
}

// ParseQuotes handles parsing the following "Stock Time Series" functions
// - TIME_SERIES_DAILY
// - TIME_SERIES_DAILY_ADJUSTED
// - TIME_SERIES_MONTHLY
// - TIME_SERIES_MONTHLY_ADJUSTED
func ParseQuotes(r io.Reader, location *time.Location) ([]Quote, error) {
	var list []Quote
	return list, ParseCSV(r, &list, location)
}

func NewIntraDayQuotesURL(symbol string) string {
	u := url.URL{
		Scheme: "https",
		Host:   "www.alphavantage.co",
		Path:   "/query",
		RawQuery: url.Values{
			"datatype":       []string{"csv"},
			"outputsize":     []string{"compact"},
			"function":       []string{"TIME_SERIES_INTRADAY"},
			"symbol":         []string{symbol},
			"interval":       []string{"15min"},
			"extended_hours": []string{"true"},
		}.Encode(),
	}
	return u.String()
}

func (client *Client) TimeSeriesIntraday(ctx context.Context, symbol string) ([]IntraDayQuote, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, NewIntraDayQuotesURL(symbol), nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer closeAndIgnoreError(res.Body)
	quotes, err := ParseIntraDayQuotes(res.Body, time.UTC)
	if err != nil {
		return nil, err
	}
	return quotes, nil
}

// IntraDayQuote is convertable to Quote. The only difference is the time-layout includes additional time information.
type IntraDayQuote struct {
	Time             time.Time `column-name:"timestamp" time-layout:"2006-01-02 15:04:05"`
	Open             float64   `column-name:"open"`
	High             float64   `column-name:"high"`
	Low              float64   `column-name:"low"`
	Close            float64   `column-name:"close"`
	Volume           float64   `column-name:"volume"`
	DividendAmount   float64   `column-name:"dividend_amount"`
	SplitCoefficient float64   `column-name:"split_coefficient"`
}

var _ = Quote(IntraDayQuote{})

// ParseIntraDayQuotes handles parsing the following "Stock Time Series" functions
// - TIME_SERIES_INTRADAY
func ParseIntraDayQuotes(r io.Reader, location *time.Location) ([]IntraDayQuote, error) {
	var list []IntraDayQuote
	return list, ParseCSV(r, &list, location)
}

//TODO: use this instead after bumping to 1.18
//func convertElements[T1, T2 any](list []T1, convert func(T1) T2) []T2 {
//	result := make([]T2, len(list))
//	for i := range list {
//		result[i] = convert(list[i])
//	}
//	return result
//}

func convertQuoteElements(list []IntraDayQuote, convert func(quote IntraDayQuote) Quote) []Quote {
	result := make([]Quote, len(list))
	for i := range list {
		result[i] = convert(list[i])
	}
	return result
}

// ListingStatus represents the listing status information for a security.
type ListingStatus struct {
	Symbol        string    `column-name:"symbol"`        // The security symbol
	Name          string    `column-name:"name"`          // The company or security name
	Exchange      string    `column-name:"exchange"`      // The exchange where it's listed
	AssetType     string    `column-name:"assetType"`     // Type of asset (Stock, ETF, etc.)
	IPODate       time.Time `column-name:"ipoDate"`       // Initial public offering date
	DeListingDate time.Time `column-name:"delistingDate"` // Date when delisted (if applicable)
	Status        string    `column-name:"status"`        // Current status (Active, Delisted)
}

// Listing status constants.
const (
	ListingStatusActive   = "Active"   // Security is actively listed
	ListingStatusDelisted = "Delisted" // Security has been delisted
)

// Asset type constants.
const (
	AssetTypeStock = "Stock" // Stock security type
	AssetTypeETF   = "ETF"   // Exchange-traded fund type
)

func NewListingStatusURL(isListed bool) (string, error) {
	state := ListingStatusActive
	if !isListed {
		state = ListingStatusDelisted
	}
	state = strings.ToLower(state)

	u := url.URL{
		Scheme: "https",
		Host:   "www.alphavantage.co",
		Path:   "/query",
		RawQuery: url.Values{
			"datatype": []string{"csv"},
			"function": []string{"LISTING_STATUS"},
			"state":    []string{state},
		}.Encode(),
	}

	return u.String(), nil
}

func (client *Client) ListingStatus(ctx context.Context, isListed bool) ([]ListingStatus, error) {
	rc, err := client.DoListingStatusRequest(ctx, isListed)
	if err != nil {
		return nil, err
	}
	defer closeAndIgnoreError(rc)
	var result []ListingStatus
	return result, ParseCSV(rc, &result, nil)
}

// DoListingStatusRequest fetches listing or delisting status data.
// If isListed is true, it returns currently active listings.
// If isListed is false, it returns delisted securities.
// The response is returned as CSV data in an io.ReadCloser that must be closed by the caller.
func (client *Client) DoListingStatusRequest(ctx context.Context, isListed bool) (io.ReadCloser, error) {
	requestURL, err := NewListingStatusURL(isListed)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		requestURL,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create listing status request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return checkError(res.Body)
}

// Deprecated: use DoListingStatusRequest instead. This method will be removed before 2023.
func (client *Client) ListingStatusRequest(ctx context.Context, isListed bool, fn func(io.Reader) error) error {
	rc, err := client.DoListingStatusRequest(ctx, isListed)
	if err != nil {
		return err
	}
	defer closeAndIgnoreError(rc)
	return fn(rc)
}

// CompanyOverview fetches comprehensive company information for the specified symbol.
// It returns detailed company data including financial metrics, sector information,
// and key statistics as a CompanyOverview struct.
func (client *Client) CompanyOverview(ctx context.Context, symbol string) (CompanyOverview, error) {
	u := url.URL{
		Scheme: "https",
		Host:   "www.alphavantage.co",
		Path:   "/query",
		RawQuery: url.Values{
			"function": []string{"OVERVIEW"},
			"symbol":   []string{symbol},
		}.Encode(),
	}

	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		u.String(),
		nil,
	)
	if err != nil {
		return CompanyOverview{}, fmt.Errorf("failed to create listing status request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return CompanyOverview{}, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return CompanyOverview{}, err
	}

	var result CompanyOverview
	err = json.Unmarshal(buf, &result)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

// CompanyOverview contains comprehensive company information and financial metrics
// returned by the AlphaVantage OVERVIEW function.
type CompanyOverview struct {
	CIK                        string    `av-json:"CIK"`
	Symbol                     string    `av-json:"Symbol"`
	AssetType                  string    `av-json:"AssetType"`
	Name                       string    `av-json:"Name"`
	Description                string    `av-json:"Description"`
	Exchange                   string    `av-json:"Exchange"`
	Currency                   string    `av-json:"Currency"`
	Country                    string    `av-json:"Country"`
	Sector                     string    `av-json:"Sector"`
	Industry                   string    `av-json:"Industry"`
	Address                    string    `av-json:"Address"`
	FiscalYearEnd              string    `av-json:"FiscalYearEnd"`
	LatestQuarter              time.Time `av-json:"LatestQuarter"`
	MarketCapitalization       int       `av-json:"MarketCapitalization"`
	EBITDA                     int       `av-json:"EBITDA"`
	PERatio                    float64   `av-json:"PERatio"`
	PEGRatio                   float64   `av-json:"PEGRatio"`
	BookValue                  float64   `av-json:"BookValue"`
	DividendPerShare           float64   `av-json:"DividendPerShare"`
	DividendYield              float64   `av-json:"DividendYield"`
	EPS                        float64   `av-json:"EPS"`
	RevenuePerShareTTM         float64   `av-json:"RevenuePerShareTTM"`
	ProfitMargin               float64   `av-json:"ProfitMargin"`
	OperatingMarginTTM         float64   `av-json:"OperatingMarginTTM"`
	ReturnOnAssetsTTM          float64   `av-json:"ReturnOnAssetsTTM"`
	ReturnOnEquityTTM          float64   `av-json:"ReturnOnEquityTTM"`
	RevenueTTM                 int       `av-json:"RevenueTTM"`
	GrossProfitTTM             int       `av-json:"GrossProfitTTM"`
	DilutedEPSTTM              float64   `av-json:"DilutedEPSTTM"`
	QuarterlyEarningsGrowthYOY float64   `av-json:"QuarterlyEarningsGrowthYOY"`
	QuarterlyRevenueGrowthYOY  float64   `av-json:"QuarterlyRevenueGrowthYOY"`
	AnalystTargetPrice         float64   `av-json:"AnalystTargetPrice"`
	TrailingPE                 float64   `av-json:"TrailingPE"`
	ForwardPE                  float64   `av-json:"ForwardPE"`
	PriceToSalesRatioTTM       float64   `av-json:"PriceToSalesRatioTTM"`
	PriceToBookRatio           float64   `av-json:"PriceToBookRatio"`
	EVToRevenue                float64   `av-json:"EVToRevenue"`
	EVToEBITDA                 float64   `av-json:"EVToEBITDA"`
	Beta                       float64   `av-json:"Beta"`
	FiftyTwoWeekHigh           float64   `av-json:"52WeekHigh"`
	FiftyTwoWeekLow            float64   `av-json:"52WeekLow"`
	FiftyDayMovingAverage      float64   `av-json:"50DayMovingAverage"`
	TwoHundredDayMovingAverage float64   `av-json:"200DayMovingAverage"`
	SharesOutstanding          int       `av-json:"SharesOutstanding"`
	SharesFloat                int       `av-json:"SharesFloat"`
	SharesShort                int       `av-json:"SharesShort"`
	SharesShortPriorMonth      int       `av-json:"SharesShortPriorMonth"`
	ShortRatio                 float64   `av-json:"ShortRatio"`
	ShortPercentOutstanding    float64   `av-json:"ShortPercentOutstanding"`
	ShortPercentFloat          float64   `av-json:"ShortPercentFloat"`
	PercentInsiders            float64   `av-json:"PercentInsiders"`
	PercentInstitutions        float64   `av-json:"PercentInstitutions"`
	ForwardAnnualDividendRate  float64   `av-json:"ForwardAnnualDividendRate"`
	ForwardAnnualDividendYield float64   `av-json:"ForwardAnnualDividendYield"`
	PayoutRatio                float64   `av-json:"PayoutRatio"`
	DividendDate               time.Time `av-json:"DividendDate"`
	ExDividendDate             time.Time `av-json:"ExDividendDate"`
	LastSplitFactor            string    `av-json:"LastSplitFactor"`
	LastSplitDate              time.Time `av-json:"LastSplitDate"`
}

func (c *CompanyOverview) UnmarshalJSON(in []byte) error {
	var data map[string]string

	err := json.Unmarshal(in, &data)
	if err != nil {
		return err
	}

	rv := reflect.ValueOf(c)

	numFields := rv.Type().Elem().NumField()
	for i := 0; i < numFields; i++ {
		ft := rv.Elem().Type().Field(i)
		fv := rv.Elem().Field(i)
		jsonKey := ft.Tag.Get("av-json")

		v, ok := data[jsonKey]
		if !ok || v == "" || v == "None" {
			continue
		}

		switch fv.Interface().(type) {
		case string:
			fv.SetString(v)
		case int:
			in, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse %s: %w", jsonKey, err)
			}
			fv.SetInt(in)
		case float64:
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return fmt.Errorf("failed to parse %s: %w", jsonKey, err)
			}
			fv.SetFloat(f)
		case time.Time:
			if v == "0000-00-00" {
				continue
			}
			t, err := time.ParseInLocation(DefaultDateFormat, v, time.UTC)
			if err != nil {
				return err
			}
			fv.Set(reflect.ValueOf(t))
		default:
			return fmt.Errorf("unsupported type %T", fv.Interface())
		}
	}

	return nil
}

// GlobalQuote fetches the latest price and volume information for the specified equity symbol.
// It returns the data in CSV format as an io.ReadCloser that must be closed by the caller.
//
// The CSV response includes columns for symbol, open, high, low, price, volume, latestDay,
// previousClose, change, and changePercent.
func (client *Client) GlobalQuote(ctx context.Context, symbol string) (io.ReadCloser, error) {
	makeURL := func(scheme, host string) string {
		u := url.URL{
			Scheme: scheme,
			Host:   host,
			Path:   "/query",
			RawQuery: url.Values{
				"function": []string{"GLOBAL_QUOTE"},
				"symbol":   []string{symbol},
				"datatype": []string{"csv"},
			}.Encode(),
		}
		return u.String()
	}

	return client.doWithFallback(ctx, makeURL)
}

// SymbolSearchResult represents a single result from the symbol search API.
type SymbolSearchResult struct {
	Symbol      string  `column-name:"symbol"`      // The security symbol
	Name        string  `column-name:"name"`        // Company or security name
	Type        string  `column-name:"type"`        // Security type (Equity, ETF, etc.)
	Region      string  `column-name:"region"`      // Geographic region
	MarketOpen  string  `column-name:"marketOpen"`  // Market opening time
	MarketClose string  `column-name:"marketClose"` // Market closing time
	TimeZone    string  `column-name:"timezone"`    // Market timezone
	Currency    string  `column-name:"currency"`    // Trading currency
	MatchScore  float64 `column-name:"matchScore"`  // Relevance score (0.0 to 1.0)
}

func NewSymbolSearchURL(keywords string) (string, error) {
	u := url.URL{
		Scheme: "https",
		Host:   "www.alphavantage.co",
		Path:   "/query",
		RawQuery: url.Values{
			"datatype": []string{"csv"},
			"function": []string{"SYMBOL_SEARCH"},
			"keywords": []string{keywords},
		}.Encode(),
	}
	return u.String(), nil
}

func (client *Client) SymbolSearch(ctx context.Context, keywords string) ([]SymbolSearchResult, error) {
	rc, err := client.DoSymbolSearchRequest(ctx, keywords)
	if err != nil {
		return nil, err
	}
	defer closeAndIgnoreError(rc)
	return ParseSymbolSearchQuery(rc)
}

// DoSymbolSearchRequest searches for securities matching the given keywords.
// It returns CSV data containing symbol search results as an io.ReadCloser that must be closed by the caller.
// The results include symbol, name, type, region, market times, timezone, currency, and match score.
func (client *Client) DoSymbolSearchRequest(ctx context.Context, keywords string) (io.ReadCloser, error) {
	requestURL, err := NewSymbolSearchURL(keywords)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		requestURL,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create quotes request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	rc, err := checkError(res.Body)
	if err != nil {
		closeAndIgnoreError(res.Body)
		return nil, err
	}

	return rc, nil
}

func ParseSymbolSearchQuery(r io.Reader) ([]SymbolSearchResult, error) {
	var list []SymbolSearchResult
	return list, ParseCSV(r, &list, nil)
}

func (r *SymbolSearchResult) ParseTimezone() (*time.Location, error) {
	return time.LoadLocation(r.TimeZone)
}

func closeAndIgnoreError(c io.Closer) {
	_ = c.Close()
}

type multiReadCloser struct {
	io.Reader
	close func() error
}

func (mrc multiReadCloser) Close() error {
	return mrc.close()
}
