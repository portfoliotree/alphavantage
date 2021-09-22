package alphavantage

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ListingStatus struct {
	Symbol        string    `column-name:"symbol"`
	Name          string    `column-name:"name"`
	Exchange      string    `column-name:"exchange"`
	AssetType     string    `column-name:"assetType"`
	IPODate       time.Time `column-name:"ipoDate"`
	DeListingDate time.Time `column-name:"delistingDate"`
	Status        string    `column-name:"status"`
}

const (
	ListingStatusActive   = "Active"
	ListingStatusDelisted = "Delisted"

	AssetTypeStock = "Stock"
	AssetTypeETF   = "ETF"
)

func (client *Client) ListingStatus(ctx context.Context, isListed bool) ([]ListingStatus, error) {
	var result []ListingStatus
	return result, client.ListingStatusRequest(ctx, isListed, func(r io.Reader) error {
		return ParseCSV(r, &result, nil)
	})
}

func (client *Client) ListingStatusRequest(ctx context.Context, isListed bool, fn func(io.Reader) error) error {
	state := ListingStatusActive
	if !isListed {
		state = ListingStatusDelisted
	}
	state = strings.ToLower(state)

	u := url.URL{
		Scheme: "https",
		Host:   "www.alphavantage.co",
		Path:   "/query",
		RawQuery: url.Values{
			"datatype": []string{"csv"},
			"function": []string{"LISTING_STATUS"},
			"state":    []string{state},
		}.Encode(),
	}

	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		u.String(),
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to create listing status request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	r, err := checkError(res.Body)
	if err != nil {
		return err
	}

	return fn(r)
}
