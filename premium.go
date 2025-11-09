package alphavantage

import (
	"cmp"
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	"golang.org/x/time/rate"
)

// PremiumPlan helps you set up an HTTP request rate limiter.
// See plan details here: https://www.alphavantage.co/premium/
type PremiumPlan int

const (
	PremiumPlan75   PremiumPlan = 75
	PremiumPlan150  PremiumPlan = 150
	PremiumPlan300  PremiumPlan = 300
	PremiumPlan600  PremiumPlan = 600
	PremiumPlan1200 PremiumPlan = 1200
)

func (plan PremiumPlan) String() string {
	return fmt.Sprintf("premium plan with %d requests per minute", plan)
}

func (plan PremiumPlan) Limit() rate.Limit {
	return rate.Every(time.Minute / time.Duration(cmp.Or(plan, 1)))
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
