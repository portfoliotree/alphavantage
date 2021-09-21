package alphavantage

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

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

type SearchResult struct {
	Symbol      string
	Name        string
	Type        string
	Region      string
	MarketOpen  string
	MarketClose string
	TimeZone    string
	Currency    string
	MatchScore  float64
}

func (r *SearchResult) ParseRow(header []string, row []string) error {
	if len(header) != len(row) {
		return fmt.Errorf("row has %d fields but %d were expected", len(row), len(header))
	}

	for i, h := range header {
		switch h {
		case "symbol":
			r.Symbol = row[i]
		case "name":
			r.Name = row[i]
		case "type":
			r.Type = row[i]
		case "region":
			r.Region = row[i]
		case "marketOpen":
			r.MarketOpen = row[i]
		case "marketClose":
			r.MarketClose = row[i]
		case "timezone":
			r.TimeZone = row[i]
		case "currency":
			r.Currency = row[i]
		case "matchScore":
			var err error
			r.MatchScore, err = strconv.ParseFloat(row[i], 64)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func ParseSearchQuery(r io.Reader) ([]SearchResult, error) {
	reader := csv.NewReader(r)
	reader.TrimLeadingSpace = true
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	reader.FieldsPerRecord = len(header)

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	list := make([]SearchResult, len(rows))

	for i, row := range rows {
		if err := list[i].ParseRow(header, row); err != nil {
			return nil, err
		}
	}

	return list, nil
}

func (r *SearchResult) ParseTimezone() (*time.Location, error) {
	return time.LoadLocation(r.TimeZone)
}
