package alphavantage

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type SearchResult struct {
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

func (client *Client) Search(ctx context.Context, keywords string) ([]SearchResult, error) {
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

	return ParseSearchQuery(r)
}

func ParseSearchQuery(r io.Reader) ([]SearchResult, error) {
	var list []SearchResult
	return list, ParseCSV(r, &list, nil)
}

func (r *SearchResult) ParseTimezone() (*time.Location, error) {
	return time.LoadLocation(r.TimeZone)
}
