package alphavantage

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
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
	req.URL, _ = url.Parse("https://www.alphavantage.co")
	req.URL.Query().Set("apikey", service.APIKey)
	return service.Client.Do(req)
}

type Quote struct {
	Time                           time.Time
	Open, High, Low, Close, Volume float64
}

func ParseQueryResponse(r io.Reader) ([]Quote, error) {
	body := make(map[string]interface{})

	if err := json.NewDecoder(r).Decode(&body); err != nil {
		return nil, err
	}

	delete(body, "Meta Data")

	var values interface{}
	for _, values = range body {
		break
	}

	rawQuotes := values.(map[string]interface{})

	var quotes []Quote

	for rawTime, values := range rawQuotes {
		tm, err := time.Parse("2006-01-02 15:04:05", rawTime)
		if err != nil {
			continue
		}

		valuesMap := values.(map[string]interface{})

		var quote Quote

		quote.Time = tm

		for k, val := range valuesMap {
			f := strings.Fields(k)
			if len(f) < 2 {
				continue
			}
			switch f[1] {
			case "open":
				quote.Open = val.(float64)
			case "high":
				quote.High = val.(float64)
			case "low":
				quote.Low = val.(float64)
			case "close":
				quote.Close = val.(float64)
			case "volume":
				quote.Volume = val.(float64)
			}
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}
