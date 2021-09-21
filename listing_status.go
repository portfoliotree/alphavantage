package alphavantage

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
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

func (client *Client) ListingStatus(ctx context.Context, isListed bool) ([]ListingStatus, error) {
	state := "active"
	if !isListed {
		state = "delisted"
	}

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
		return nil, fmt.Errorf("failed to create listing status request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	r, err := checkError(res.Body)
	if err != nil {
		return nil, err
	}

	var result []ListingStatus
	return result, ParseCSV(r, &result, nil)
}
