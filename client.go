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
	"reflect"
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

	res, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 300 || res.StatusCode < 200 {
		buf, err := ioutil.ReadAll(io.LimitReader(res.Body, 1<<10))
		if err != nil {
			buf = []byte(err.Error())
		}
		return res, fmt.Errorf("request failed with status %d %s: %s",
			res.StatusCode, http.StatusText(res.StatusCode), string(buf))
	}

	return res, nil
}

func checkError(r io.Reader) (io.Reader, error) {
	var buf [1]byte
	n, err := r.Read(buf[:])
	if err != nil {
		return nil, fmt.Errorf("could not read request response: %w", err)
	}

	mr := io.MultiReader(bytes.NewReader(buf[:]), r)
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

	return mr, nil
}

func ParseCSV(r io.Reader, data interface{}) error {
	rv := reflect.ValueOf(data)
	if rv.Kind() != reflect.Ptr {
		panic("parse must receive pointer to data")
	}
	if rv.IsNil() {
		panic("parse must not receive pointer to nil data")
	}

	return nil
}
