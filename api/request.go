package api

import (
	"cmp"
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

func DoQuery(ctx context.Context, client Doer, u url.URL, query Encoder) (*http.Response, error) {
	u.Scheme = cmp.Or(u.Scheme, DefaultScheme)
	u.Host = cmp.Or(u.Host, DefaultHost)
	u.Path = "/query"
	u.RawQuery = query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
