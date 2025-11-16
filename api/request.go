package api

import (
	"context"
	"net/http"
	"net/url"
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type Encoder interface {
	Encode() string
}

func DoQuery(ctx context.Context, client Doer, query Encoder) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, (&url.URL{
		Scheme:   DefaultScheme,
		Host:     DefaultHost,
		Path:     DefaultPath,
		RawQuery: query.Encode(),
	}).String(), nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
