package alphavantage

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

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
