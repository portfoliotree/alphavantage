package alphavantage

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

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
