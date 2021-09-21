package alphavantage

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type QuoteFunction string

const (
	TimeSeriesIntraday        QuoteFunction = "TIME_SERIES_INTRADAY"
	TimeSeriesDaily           QuoteFunction = "TIME_SERIES_DAILY"
	TimeSeriesDailyAdjusted   QuoteFunction = "TIME_SERIES_DAILY_ADJUSTED"
	TimeSeriesMonthly         QuoteFunction = "TIME_SERIES_MONTHLY"
	TimeSeriesMonthlyAdjusted QuoteFunction = "TIME_SERIES_MONTHLY_ADJUSTED"
)

func (client *Client) Quotes(ctx context.Context, symbol string, function QuoteFunction) ([]Quote, error) {
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

	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		u.String(),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create quotes request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	r, err := checkError(res.Body)
	if err != nil {
		return nil, err
	}

	switch function {
	case TimeSeriesIntraday:
		list, err := ParseIntraDayQuotes(r)
		if err != nil {
			return nil, err
		}
		quotes := make([]Quote, len(list))
		for i := range list {
			quotes[i] = Quote(list[i])
		}
		return quotes, nil
	default:
		return ParseQuotes(r)
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

// ParseQuotes handles parsing the following "Stock Time Series" functions
// - TIME_SERIES_DAILY
// - TIME_SERIES_DAILY_ADJUSTED
// - TIME_SERIES_MONTHLY
// - TIME_SERIES_MONTHLY_ADJUSTED
func ParseQuotes(r io.Reader) ([]Quote, error) {
	var list []Quote
	return list, ParseCSV(r, &list, nil)
}

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
func ParseIntraDayQuotes(r io.Reader) ([]IntraDayQuote, error) {
	var list []IntraDayQuote
	return list, ParseCSV(r, &list, nil)
}
