package alphavantage

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

	return ParseStockQuery(r)
}

type Quote struct {
	Time time.Time
	Open, High, Low, Close, Volume,
	DividendAmount, SplitCoefficient float64
}

func (q *Quote) setTime(str string) error {
	var err error
	q.Time, err = time.ParseInLocation("2006-01-02", str, timezone)
	if err != nil {
		return fmt.Errorf("failed to set Time: %s", err)
	}
	return nil
}

func (q *Quote) setTimeIntraDay(str string) error {
	var err error
	q.Time, err = time.ParseInLocation("2006-01-02 15:04:05", str, timezone)
	if err != nil {
		return fmt.Errorf("failed to set Time: %s", err)
	}
	return nil
}

func (q *Quote) setOpen(str string) error {
	var err error
	q.Open, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("failed to set Open: %s", err)
	}
	return nil
}

func (q *Quote) setHigh(str string) error {
	var err error
	q.High, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("failed to set High: %s", err)
	}
	return nil
}

func (q *Quote) setLow(str string) error {
	var err error
	q.Low, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("failed to set Low: %s", err)
	}
	return nil
}

func (q *Quote) setClose(str string) error {
	var err error
	q.Close, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("failed to set Close: %s", err)
	}
	return nil
}

func (q *Quote) setVolume(str string) error {
	var err error
	q.Volume, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("failed to set Volume: %s", err)
	}
	return nil
}

func (q *Quote) setSplitCoefficient(str string) error {
	var err error
	q.SplitCoefficient, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("failed to set SplitCoefficient: %s", err)
	}
	return nil
}

func (q *Quote) setDividendAmount(str string) error {
	var err error
	q.DividendAmount, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("failed to set SplitCoefficient: %s", err)
	}
	return nil
}

func (q *Quote) ParseRow(header, row []string) error {
	if len(header) != len(row) {
		return fmt.Errorf("row has %d fields but %d were expected", len(row), len(header))
	}

	for i, h := range header {
		switch h {
		case "timestamp":
			if strings.Contains(row[i], ":") {
				if err := q.setTimeIntraDay(row[i]); err != nil {
					return err
				}
			} else {
				if err := q.setTime(row[i]); err != nil {
					return err
				}
			}
		case "open":
			if err := q.setOpen(row[i]); err != nil {
				return err
			}
		case "high":
			if err := q.setHigh(row[i]); err != nil {
				return err
			}
		case "low":
			if err := q.setLow(row[i]); err != nil {
				return err
			}
		case "close":
			if err := q.setClose(row[i]); err != nil {
				return err
			}
		case "split_coefficient":
			if err := q.setSplitCoefficient(row[i]); err != nil {
				return err
			}
		case "dividend_amount":
			if err := q.setDividendAmount(row[i]); err != nil {
				return err
			}
		case "volume":
			if err := q.setVolume(row[i]); err != nil {
				return err
			}
		}
	}

	return nil
}

// ParseStockQuery handles parsing the following "Stock Time Series" functions
// - TIME_SERIES_INTRADAY
// - TIME_SERIES_DAILY
// - TIME_SERIES_DAILY_ADJUSTED
// - TIME_SERIES_MONTHLY
// - TIME_SERIES_MONTHLY_ADJUSTED
func ParseStockQuery(r io.Reader) ([]Quote, error) {
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

	list := make([]Quote, len(rows))

	for i, row := range rows {
		if err := list[i].ParseRow(header, row); err != nil {
			return nil, err
		}
	}

	return list, nil
}
