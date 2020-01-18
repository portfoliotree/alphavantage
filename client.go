package alphavantage

import (
	"net/http"
	"net/url"
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
