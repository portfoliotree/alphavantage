package alphavantage

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (client *Client) ETFProfile(ctx context.Context, symbol string) (ETFProfile, error) {
	u := url.URL{
		Scheme: "https",
		Host:   "www.alphavantage.co",
		Path:   "/query",
		RawQuery: url.Values{
			"function": []string{"ETF_PROFILE"},
			"symbol":   []string{symbol},
		}.Encode(),
	}

	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		u.String(),
		nil,
	)
	if err != nil {
		return ETFProfile{}, fmt.Errorf("failed to create ETF profile request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return ETFProfile{}, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return ETFProfile{}, err
	}

	var result ETFProfile
	err = json.Unmarshal(buf, &result)
	return result, err
}

type ETFProfile struct {
	Symbol            string        `json:"symbol,omitempty"`
	NetAssets         string        `json:"net_assets,omitempty"`
	NetExpenseRatio   string        `json:"net_expense_ratio,omitempty"`
	PortfolioTurnover string        `json:"portfolio_turnover,omitempty"`
	DividendYield     string        `json:"dividend_yield,omitempty"`
	InceptionDate     string        `json:"inception_date,omitempty"`
	Leveraged         string        `json:"leveraged,omitempty"`
	Sectors           []ETFSector   `json:"sectors,omitempty"`
	Holdings          []ETFHolding  `json:"holdings,omitempty"`
}

type ETFSector struct {
	Sector string `json:"sector,omitempty"`
	Weight string `json:"weight,omitempty"`
}

type ETFHolding struct {
	Symbol      string `json:"symbol,omitempty"`
	Description string `json:"description,omitempty"`
	Weight      string `json:"weight,omitempty"`
}