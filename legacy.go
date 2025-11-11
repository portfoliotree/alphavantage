package alphavantage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"time"
)

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
	q := QueryListingStatus(client.APIKey)
	if isListed {
		q = q.StateActive()
	} else {
		q = q.StateDelisted()
	}
	req, err := client.newRequest(ctx, url.Values(q))
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

// DoSymbolSearchRequest searches for securities matching the given keywords.
// It returns CSV data containing symbol search results as an io.ReadCloser that must be closed by the caller.
// The results include symbol, name, type, region, market times, timezone, currency, and match score.
func (client *Client) DoSymbolSearchRequest(ctx context.Context, keywords string) (io.ReadCloser, error) {
	req, err := client.newRequest(ctx, url.Values(QuerySymbolSearch(client.APIKey, keywords).DataTypeCSV()))
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
		return convertElements(list, func(q IntraDayQuote) Quote { return Quote(q) }), nil
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
	err := function.Validate()
	if err != nil {
		return nil, err
	}
	req, err := client.newRequest(ctx, url.Values{
		"datatype":   []string{"csv"},
		"outputsize": []string{"full"},
		"function":   []string{string(function)},
		"symbol":     []string{symbol},
		"apikey":     []string{client.APIKey},
	})
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return checkError(res.Body)
}

func ParseSymbolSearchQuery(r io.Reader) ([]SymbolSearchResult, error) {
	var list []SymbolSearchResult
	return list, ParseCSV(r, &list, nil)
}

func (r *SymbolSearchResult) ParseTimezone() (*time.Location, error) {
	return time.LoadLocation(r.TimeZone)
}

// QuoteFunction represents the different time series functions available
// from the AlphaVantage API.
type QuoteFunction string

// Time series function constants for different data intervals and types.
const (
	TimeSeriesIntraday        QuoteFunction = "TIME_SERIES_INTRADAY"
	TimeSeriesDaily           QuoteFunction = "TIME_SERIES_DAILY"
	TimeSeriesDailyAdjusted   QuoteFunction = "TIME_SERIES_DAILY_ADJUSTED"
	TimeSeriesWeekly          QuoteFunction = "TIME_SERIES_WEEKLY"
	TimeSeriesWeeklyAdjusted  QuoteFunction = "TIME_SERIES_WEEKLY_ADJUSTED"
	TimeSeriesMonthly         QuoteFunction = "TIME_SERIES_MONTHLY"
	TimeSeriesMonthlyAdjusted QuoteFunction = "TIME_SERIES_MONTHLY_ADJUSTED"
)

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

func convertElements[T1, T2 any](list []T1, convert func(T1) T2) []T2 {
	result := make([]T2, len(list))
	for i := range list {
		result[i] = convert(list[i])
	}
	return result
}
