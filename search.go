package alphavantage

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"
)

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

	quotes := make([]SearchResult, len(rows))

	for i, row := range rows {
		if err := quotes[i].ParseRow(header, row); err != nil {
			return nil, err
		}
	}

	return quotes, nil
}

func (r *SearchResult) ParseTimezone() (*time.Location, error) {
	return time.LoadLocation(r.TimeZone)
}
