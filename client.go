package alphavantage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	metaDataKey = "Meta Data"

	timeSeriesDaily          = "Time Series (Daily)"
	timeSeriesWeekly         = "Weekly Time Series"
	timeSeriesWeeklyAdjusted = "Weekly Adjusted Time Series"
	timeSeriesIntraday1min   = "Time Series (1min)"
	timeSeriesIntraday5min   = "Time Series (5min)"
	timeSeriesIntraday15min  = "Time Series (15min)"
	timeSeriesIntraday30min  = "Time Series (30min)"
	timeSeriesIntraday60min  = "Time Series (60min)"
)

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
	req.URL.Host = u.Host
	req.URL.Scheme = u.Scheme

	q := req.URL.Query()
	if q.Get("apikey") == "" {
		q.Set("apikey", service.APIKey)
	}
	req.URL.RawQuery = q.Encode()
	return service.Client.Do(req)
}

type Quote struct {
	Time                           time.Time
	Open, High, Low, Close, Volume float64
}

func ParseQuotesResponse(in []byte) ([]Quote, error) {
	body := make(map[string]json.RawMessage)

	if err := json.Unmarshal(in, &body); err != nil {
		return nil, err
	}

	var rawMeta map[string]string

	if err := json.Unmarshal(body[metaDataKey], &rawMeta); err != nil {
		return nil, fmt.Errorf("could not decode meta: %w", err)
	}

	var (
		timeZone *time.Location
	)

	for k, val := range rawMeta {
		switch {
		// case strings.HasSuffix(k, "Last Refreshed"):

		case strings.HasSuffix(k, "Time Zone"):
			timeZone, _ = time.LoadLocation(val)
		}
	}

	if timeZone == nil {
		return nil, fmt.Errorf("could not parse timezone from meta")
	}

	delete(body, metaDataKey)

	var (
		rawQuotes map[string]json.RawMessage
		format    = "2006-01-02"
	)

	for key := range body {
		switch key {
		default:
			return nil, fmt.Errorf("unknown key %q", key)

		case timeSeriesDaily,
			timeSeriesWeekly,
			timeSeriesWeeklyAdjusted:

		case timeSeriesIntraday1min,
			timeSeriesIntraday5min,
			timeSeriesIntraday15min,
			timeSeriesIntraday30min,
			timeSeriesIntraday60min:
			format += " 15:04:05"
		}

		if err := json.Unmarshal(body[key], &rawQuotes); err != nil {
			return nil, fmt.Errorf("could not decode quote: %w", err)
		}

		break
	}

	var quotes []Quote

	for rawTime, rawQuote := range rawQuotes {
		tm, err := time.ParseInLocation(format, rawTime, timeZone)
		if err != nil {
			continue
		}

		quote := Quote{
			Time: tm,
		}

		var valuesMap map[string]string

		if err := json.Unmarshal(rawQuote, &valuesMap); err != nil {
			return nil, err
		}

		for k, val := range valuesMap {
			var n float64

			if val != "" {
				n, err = strconv.ParseFloat(val, 64)
				if err != nil {
					return nil, err
				}
			}

			switch {
			case strings.HasSuffix(k, "open"):
				quote.Open = n
			case strings.HasSuffix(k, "high"):
				quote.High = n
			case strings.HasSuffix(k, "low"):
				quote.Low = n
			case strings.HasSuffix(k, "close"):
				quote.Close = n
			case strings.HasSuffix(k, "volume"):
				quote.Volume = n
			}
		}

		quotes = append(quotes, quote)
	}

	sort.Sort(QuotesInDecreasingDateOrder(quotes))

	return quotes, nil
}

type QuotesInIncreasingDateOrder []Quote

func (q QuotesInIncreasingDateOrder) Len() int {
	return len(q)
}

func (q QuotesInIncreasingDateOrder) Less(i, j int) bool {
	return q[i].Time.Before(q[j].Time)
}

func (q QuotesInIncreasingDateOrder) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

type QuotesInDecreasingDateOrder []Quote

func (q QuotesInDecreasingDateOrder) Len() int {
	return len(q)
}

func (q QuotesInDecreasingDateOrder) Less(i, j int) bool {
	return q[j].Time.Before(q[i].Time)
}

func (q QuotesInDecreasingDateOrder) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
