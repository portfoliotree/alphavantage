package alphavantage

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
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

type Quote struct {
	Time                           time.Time
	Open, High, Low, Close, Volume float64
}

func (q *Quote) SetTime(str string) error {
	var err error
	q.Time, err = time.ParseInLocation("2006-01-02", str, timezone)
	if err != nil {
		return fmt.Errorf("failed to set Time: %s", err)
	}
	return nil
}

func (q *Quote) SetTimeIntraDay(str string) error {
	var err error
	q.Time, err = time.ParseInLocation("2006-01-02 15:04:05", str, timezone)
	if err != nil {
		return fmt.Errorf("failed to set Time: %s", err)
	}
	return nil
}

func (q *Quote) SetOpen(str string) error {
	var err error
	q.Open, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("failed to set Open: %s", err)
	}
	return nil
}

func (q *Quote) SetHigh(str string) error {
	var err error
	q.High, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("failed to set High: %s", err)
	}
	return nil
}

func (q *Quote) SetLow(str string) error {
	var err error
	q.Low, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("failed to set Low: %s", err)
	}
	return nil
}

func (q *Quote) SetClose(str string) error {
	var err error
	q.Close, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("failed to set Close: %s", err)
	}
	return nil
}

func (q *Quote) SetVolume(str string) error {
	var err error
	q.Volume, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("failed to set Volume: %s", err)
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
				if err := q.SetTimeIntraDay(row[i]); err != nil {
					return nil
				}
			} else {
				if err := q.SetTime(row[i]); err != nil {
					return nil
				}
			}
		case "open":
			if err := q.SetOpen(row[i]); err != nil {
				return nil
			}
		case "high":
			if err := q.SetHigh(row[i]); err != nil {
				return nil
			}
		case "low":
			if err := q.SetLow(row[i]); err != nil {
				return nil
			}
		case "close":
			if err := q.SetClose(row[i]); err != nil {
				return nil
			}
		case "volume":
			if err := q.SetVolume(row[i]); err != nil {
				return nil
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

	quotes := make([]Quote, len(rows))

	for i, row := range rows {
		if err := quotes[i].ParseRow(header, row); err != nil {
			return nil, err
		}
	}

	return quotes, nil
}
