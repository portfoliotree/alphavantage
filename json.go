package alphavantage

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"

	"github.com/portfoliotree/alphavantage/query/forex"
	"github.com/portfoliotree/alphavantage/query/fundamental"
	"github.com/portfoliotree/alphavantage/query/intelligence"
	"github.com/portfoliotree/alphavantage/query/options"
	"github.com/portfoliotree/alphavantage/query/timeseries"
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

func (f *ForexFunctions) CurrencyExchangeRate(ctx context.Context, query forex.CurrencyExchangeRateQuery) (io.ReadCloser, error) {
	res, err := (*Client)(f).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (f *FundamentalFunctions) BalanceSheet(ctx context.Context, query fundamental.BalanceSheetQuery) (io.ReadCloser, error) {
	res, err := (*Client)(f).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (f *FundamentalFunctions) CashFlow(ctx context.Context, query fundamental.CashFlowQuery) (io.ReadCloser, error) {
	res, err := (*Client)(f).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (f *FundamentalFunctions) ETFProfile(ctx context.Context, query fundamental.ETFProfileQuery) (io.ReadCloser, error) {
	res, err := (*Client)(f).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (f *FundamentalFunctions) Earnings(ctx context.Context, query fundamental.EarningsQuery) (io.ReadCloser, error) {
	res, err := (*Client)(f).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (f *FundamentalFunctions) EarningsEstimates(ctx context.Context, query fundamental.EarningsEstimatesQuery) (io.ReadCloser, error) {
	res, err := (*Client)(f).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (f *FundamentalFunctions) IncomeStatement(ctx context.Context, query fundamental.IncomeStatementQuery) (io.ReadCloser, error) {
	res, err := (*Client)(f).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (f *FundamentalFunctions) Overview(ctx context.Context, query fundamental.OverviewQuery) (io.ReadCloser, error) {
	res, err := (*Client)(f).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (f *IntelligenceFunctions) AnalyticsFixedWindow(ctx context.Context, query intelligence.AnalyticsFixedWindowQuery) (io.ReadCloser, error) {
	res, err := (*Client)(f).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (f *IntelligenceFunctions) AnalyticsSlidingWindow(ctx context.Context, query intelligence.AnalyticsSlidingWindowQuery) (io.ReadCloser, error) {
	res, err := (*Client)(f).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (f *IntelligenceFunctions) EarningsCallTranscript(ctx context.Context, query intelligence.EarningsCallTranscriptQuery) (io.ReadCloser, error) {
	res, err := (*Client)(f).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (f *IntelligenceFunctions) InsiderTransactions(ctx context.Context, query intelligence.InsiderTransactionsQuery) (io.ReadCloser, error) {
	res, err := (*Client)(f).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (f *IntelligenceFunctions) NewsSentiment(ctx context.Context, query intelligence.NewsSentimentQuery) (io.ReadCloser, error) {
	res, err := (*Client)(f).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (f *IntelligenceFunctions) TopGainersLosers(ctx context.Context, query intelligence.TopGainersLosersQuery) (io.ReadCloser, error) {
	res, err := (*Client)(f).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (f *OptionsFunctions) Realtime(ctx context.Context, query options.RealtimeQuery) (io.ReadCloser, error) {
	res, err := (*Client)(f).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (f *TimeSeriesFunctions) MarketStatus(ctx context.Context, query timeseries.MarketStatusQuery) (io.ReadCloser, error) {
	res, err := (*Client)(f).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (f *TimeSeriesFunctions) RealtimeBulkQuotes(ctx context.Context, query timeseries.RealtimeBulkQuotesQuery) (io.ReadCloser, error) {
	res, err := (*Client)(f).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
