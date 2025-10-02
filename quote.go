package alphavantage

import (
	"context"
	"io"
	"net/url"
)

// GlobalQuote fetches the latest price and volume information for the specified equity symbol.
// It returns the data in CSV format as an io.ReadCloser that must be closed by the caller.
//
// The CSV response includes columns for symbol, open, high, low, price, volume, latestDay,
// previousClose, change, and changePercent.
func (client *Client) GlobalQuote(ctx context.Context, symbol string) (io.ReadCloser, error) {
	makeURL := func(scheme, host string) string {
		u := url.URL{
			Scheme: scheme,
			Host:   host,
			Path:   "/query",
			RawQuery: url.Values{
				"function": []string{"GLOBAL_QUOTE"},
				"symbol":   []string{symbol},
				"datatype": []string{"csv"},
			}.Encode(),
		}
		return u.String()
	}

	return client.doWithFallback(ctx, makeURL)
}
