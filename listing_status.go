package alphavantage

import (
	"cmp"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ListingStatus represents the listing status information for a security.
type ListingStatus struct {
	Symbol        string    `column-name:"symbol"`        // The security symbol
	Name          string    `column-name:"name"`          // The company or security name
	Exchange      string    `column-name:"exchange"`      // The exchange where it's listed
	AssetType     string    `column-name:"assetType"`     // Type of asset (Stock, ETF, etc.)
	IPODate       time.Time `column-name:"ipoDate"`       // Initial public offering date
	DeListingDate time.Time `column-name:"delistingDate"` // Date when delisted (if applicable)
	Status        string    `column-name:"status"`        // Current status (Active, Delisted)
}

// Listing status constants.
const (
	ListingStatusActive   = "Active"   // Security is actively listed
	ListingStatusDelisted = "Delisted" // Security has been delisted
)

// Asset type constants.
const (
	AssetTypeStock = "Stock" // Stock security type
	AssetTypeETF   = "ETF"   // Exchange-traded fund type
)

func NewListingStatusURL(host string, isListed bool) (string, error) {
	state := ListingStatusActive
	if !isListed {
		state = ListingStatusDelisted
	}
	state = strings.ToLower(state)

	u := url.URL{
		Scheme: "https",
		Host:   cmp.Or(host, DefaultHost),
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

// DoListingStatusRequest fetches listing or delisting status data.
// If isListed is true, it returns currently active listings.
// If isListed is false, it returns delisted securities.
// The response is returned as CSV data in an io.ReadCloser that must be closed by the caller.
func (client *Client) DoListingStatusRequest(ctx context.Context, isListed bool) (io.ReadCloser, error) {
	requestURL, err := NewListingStatusURL(client.host(), isListed)
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
