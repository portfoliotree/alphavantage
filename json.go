package alphavantage

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"

	"github.com/portfoliotree/alphavantage/query/fundamental"
)

func (client *Client) ETFProfile(ctx context.Context, q fundamental.ETFProfileQuery) (fundamental.ETFProfile, error) {
	res, err := client.Query(ctx, url.Values(q))
	if err != nil {
		return fundamental.ETFProfile{}, fmt.Errorf("failed to create ETF profile request: %w", err)
	}
	defer closeAndIgnoreError(res.Body)

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return fundamental.ETFProfile{}, err
	}

	var result fundamental.ETFProfile
	err = json.Unmarshal(buf, &result)
	return result, err
}

// CompanyOverview fetches comprehensive company information for the specified symbol.
// It returns detailed company data including financial metrics, sector information,
// and key statistics as a CompanyOverview struct.
func (client *Client) CompanyOverview(ctx context.Context, q fundamental.OverviewQuery) (fundamental.CompanyOverview, error) {
	res, err := client.Query(ctx, url.Values(q))
	if err != nil {
		return fundamental.CompanyOverview{}, fmt.Errorf("failed to create listing status request: %w", err)
	}
	defer closeAndIgnoreError(res.Body)

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return fundamental.CompanyOverview{}, err
	}

	var result fundamental.CompanyOverview
	err = json.Unmarshal(buf, &result)
	if err != nil {
		log.Println(err)
	}
	return result, err
}
