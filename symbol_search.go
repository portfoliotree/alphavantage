package alphavantage

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type SymbolSearchResult struct {
	Symbol      string  `column-name:"symbol"`
	Name        string  `column-name:"name"`
	Type        string  `column-name:"type"`
	Region      string  `column-name:"region"`
	MarketOpen  string  `column-name:"marketOpen"`
	MarketClose string  `column-name:"marketClose"`
	TimeZone    string  `column-name:"timezone"`
	Currency    string  `column-name:"currency"`
	MatchScore  float64 `column-name:"matchScore"`
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
