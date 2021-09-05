package alphavantage

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

var timezone *time.Location

func init() {
	var err error
	timezone, err = time.LoadLocation("US/Eastern")
	if err != nil {
		panic(err)
	}
}

type QuoteFunction string

const (
	TimeSeriesIntraday        QuoteFunction = "TIME_SERIES_INTRADAY"
	TimeSeriesDaily           QuoteFunction = "TIME_SERIES_DAILY"
	TimeSeriesDailyAdjusted   QuoteFunction = "TIME_SERIES_DAILY_ADJUSTED"
	TimeSeriesMonthly         QuoteFunction = "TIME_SERIES_MONTHLY"
	TimeSeriesMonthlyAdjusted QuoteFunction = "TIME_SERIES_MONTHLY_ADJUSTED"
)

type Client struct {
	Limiter interface {
		Wait(ctx context.Context) error
	}
	Client interface {
		Do(*http.Request) (*http.Response, error)
	}
	APIKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		Client:  http.DefaultClient,
		Limiter: rate.NewLimiter(rate.Every(time.Minute/5), 2),
		APIKey:  apiKey,
	}
}

func (client *Client) Do(req *http.Request) (*http.Response, error) {
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	if client.Limiter != nil {
		err := client.Limiter.Wait(req.Context())
		if err != nil {
			return &http.Response{}, err
		}
	}

	q := req.URL.Query()
	q.Set("apikey", client.APIKey)
	req.URL.RawQuery = q.Encode()

	return client.Client.Do(req)
}

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

	if res.StatusCode >= 300 || res.StatusCode < 200 {
		buf, err := ioutil.ReadAll(io.LimitReader(res.Body, 1<<10))
		if err != nil {
			buf = []byte(err.Error())
		}
		return nil, fmt.Errorf("failed to request got status %d %s: %s",
			res.StatusCode, http.StatusText(res.StatusCode), string(buf))
	}

	var buf [1]byte
	n, err := res.Body.Read(buf[:])
	if err != nil {
		return nil, fmt.Errorf("could not read request response: %w", err)
	}

	mr := io.MultiReader(bytes.NewReader(buf[:]), res.Body)
	if n > 0 && buf[0] == '{' {
		var message struct {
			Note        string `json:"Note"`
			Information string `json:"Information"`
		}
		err = json.NewDecoder(mr).Decode(&message)
		if err != nil {
			return nil, fmt.Errorf("could not read response for: %w", err)
		}
		if strings.Contains(message.Note, " higher API call frequency") {
			return nil, fmt.Errorf("reached alphavantage rate limit")
		}

		return nil, fmt.Errorf("alphavantage request did not return csv; got notice: %w", errors.New(strings.Join([]string{message.Note, message.Information}, " ")))
	}

	return ParseStockQuery(mr)
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

	if res.StatusCode >= 300 || res.StatusCode < 200 {
		buf, err := ioutil.ReadAll(io.LimitReader(res.Body, 1<<10))
		if err != nil {
			buf = []byte(err.Error())
		}
		return nil, fmt.Errorf("failed to request got status %d %s: %s",
			res.StatusCode, http.StatusText(res.StatusCode), string(buf))
	}

	var buf [1]byte
	n, err := res.Body.Read(buf[:])
	if err != nil {
		return nil, fmt.Errorf("could not read request response: %w", err)
	}

	mr := io.MultiReader(bytes.NewReader(buf[:]), res.Body)
	if n > 0 && buf[0] == '{' {
		var message struct {
			Note        string `json:"Note"`
			Information string `json:"Information"`
		}
		err = json.NewDecoder(mr).Decode(&message)
		if err != nil {
			return nil, fmt.Errorf("could not read response for: %w", err)
		}
		if strings.Contains(message.Note, " higher API call frequency") {
			return nil, fmt.Errorf("reached alphavantage rate limit")
		}

		return nil, fmt.Errorf("alphavantage request did not return csv; got notice: %w", errors.New(strings.Join([]string{message.Note, message.Information}, " ")))
	}

	return ParseSearchQuery(mr)
}
