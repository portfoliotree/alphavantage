package alphavantage

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

// GlobalQuote fetches the latest price and volume information for the specified equity symbol.
// It returns the data in CSV format as an io.ReadCloser that must be closed by the caller.
//
// The CSV response includes columns for symbol, open, high, low, price, volume, latestDay,
// previousClose, change, and changePercent.
func (client *Client) GlobalQuote(ctx context.Context, symbol string) (io.ReadCloser, error) {
	u := url.URL{
		Scheme: "https",
		Host:   "www.alphavantage.co",
		Path:   "/query",
		RawQuery: url.Values{
			"function": []string{"GLOBAL_QUOTE"},
			"symbol":   []string{symbol},
			"datatype": []string{"csv"},
		}.Encode(),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return checkError(res.Body)
}
