package alphavantage

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var timezone *time.Location

func init() {
	var err error
	timezone, err = time.LoadLocation("US/Eastern")
	if err != nil {
		panic(err)
	}
}

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type Service struct {
	Client Doer

	APIKey string
}

func (service Service) Do(req *http.Request) (*http.Response, error) {
	if service.Client == nil {
		service.Client = http.DefaultClient
	}
	u, _ := url.Parse("https://www.alphavantage.co")
	u.Path = req.URL.Path
	if req.URL.Query().Get("apiKey") == "" {
		req.URL.Query().Set("apiKey", service.APIKey)
	}
	req.URL.Query().Set("datatype", "csv")
	u.RawQuery = req.URL.Query().Encode()
	req.URL = u
	return service.Client.Do(req)
}

type Quote struct {
	Time                           time.Time
	Open, High, Low, Close, Volume float64
}

var expectedColumns = []string{"timestamp", "open", "high", "low", "close", "volume"}

// ParseStockQuery handles parsing the following "Stock Time Series" functions
// - TIME_SERIES_INTRADAY
// - TIME_SERIES_DAILY
// - TIME_SERIES_DAILY_ADJUSTED
// - TIME_SERIES_MONTHLY
// - TIME_SERIES_MONTHLY_ADJUSTED
func ParseStockQuery(r io.Reader) ([]Quote, error) {
	reader := csv.NewReader(r)
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = len(expectedColumns)

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for i := range expectedColumns {
		if rows[0][i] != expectedColumns[i] {
			return nil, fmt.Errorf("header %d %s not match expected header %s", i, rows[0][i], expectedColumns[i])
		}
	}

	rows = rows[1:]

	timestampFormat := "2006-01-02"
	if len(rows) > 0 {
		_, err := time.Parse(timestampFormat, rows[0][0])
		if err != nil {
			intradayFormat := "2006-01-02 15:04:05"
			_, err := time.Parse(intradayFormat, rows[0][0])
			if err != nil {
				return nil, fmt.Errorf("could not parse timestamp %q: %s", rows[0][0], err)
			}
			timestampFormat = intradayFormat
		}
	}

	quotes := make([]Quote, len(rows))

	for i, row := range rows {
		var err error

		quotes[i].Time, err = time.ParseInLocation(timestampFormat, row[0], timezone)
		if err != nil {
			return nil, fmt.Errorf("failed to parse row %d: %s", i, err)
		}

		quotes[i].Open, err = strconv.ParseFloat(row[1], 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse row %d: %s", i, err)
		}

		quotes[i].High, err = strconv.ParseFloat(row[2], 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse row %d: %s", i, err)
		}

		quotes[i].Low, err = strconv.ParseFloat(row[3], 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse row %d: %s", i, err)
		}

		quotes[i].Close, err = strconv.ParseFloat(row[4], 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse row %d: %s", i, err)
		}

		quotes[i].Volume, err = strconv.ParseFloat(row[5], 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse row %d: %s", i, err)
		}
	}

	return quotes, nil
}
