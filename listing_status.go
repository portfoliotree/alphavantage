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

func NewListingStatusURL(isListed bool) (string, error) {
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

	return u.String(), nil
}

func (client *Client) ListingStatus(ctx context.Context, isListed bool) ([]ListingStatus, error) {
	rc, err := client.DoListingStatusRequest(ctx, isListed)
	if err != nil {
		return nil, err
	}
	defer closeAndIgnoreError(rc)
	var result []ListingStatus
	return result, ParseCSV(rc, &result, nil)
}

func (client *Client) DoListingStatusRequest(ctx context.Context, isListed bool) (io.ReadCloser, error) {
	requestURL, err := NewListingStatusURL(isListed)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		requestURL,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create listing status request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return checkError(res.Body)
}

// Deprecated: use DoListingStatusRequest instead. This method will be removed before 2023.
func (client *Client) ListingStatusRequest(ctx context.Context, isListed bool, fn func(io.Reader) error) error {
	rc, err := client.DoListingStatusRequest(ctx, isListed)
	if err != nil {
		return err
	}
	defer closeAndIgnoreError(rc)
	return fn(rc)
}
