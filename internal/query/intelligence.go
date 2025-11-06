package query

import (
	"net/url"
)

// NewsSentiment builds query parameters for the NEWS_SENTIMENT API.
// Returns live and historical market news & sentiment data from premier news outlets worldwide.
type NewsSentiment url.Values

// NewNewsSentiment creates a new NEWS_SENTIMENT query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewNewsSentiment(apikey string) NewsSentiment {
	return NewsSentiment{
		KeyFunction: []string{FunctionNewsSentiment},
		KeyAPIKey:   []string{apikey},
	}
}

// Tickers sets the stock/crypto/forex symbols to filter articles.
// Example: "IBM", "COIN,CRYPTO:BTC,FOREX:USD"
func (q NewsSentiment) Tickers(tickers string) NewsSentiment {
	q[KeyTickers] = []string{tickers}
	return q
}

// Topics sets the news topics to filter by.
// Example: "technology", "technology,ipo"
func (q NewsSentiment) Topics(topics string) NewsSentiment {
	q[KeyTopics] = []string{topics}
	return q
}

// TimeFrom sets the start time range in YYYYMMDDTHHMM format.
// Example: "20220410T0130"
func (q NewsSentiment) TimeFrom(timeFrom string) NewsSentiment {
	q[KeyTimeFrom] = []string{timeFrom}
	return q
}

// TimeTo sets the end time range in YYYYMMDDTHHMM format.
// Defaults to current time if time_from is specified.
func (q NewsSentiment) TimeTo(timeTo string) NewsSentiment {
	q[KeyTimeTo] = []string{timeTo}
	return q
}

// Sort sets the sort order.
// Valid values: "LATEST", "EARLIEST", "RELEVANCE"
func (q NewsSentiment) Sort(sort SortOption) NewsSentiment {
	q[KeySort] = sort.values()
	return q
}

// SortLatest sets the sort order to latest (most recent first).
func (q NewsSentiment) SortLatest() NewsSentiment {
	return q.Sort(SortOptionLatest)
}

// SortEarliest sets the sort order to earliest (oldest first).
func (q NewsSentiment) SortEarliest() NewsSentiment {
	return q.Sort(SortOptionEarliest)
}

// SortRelevance sets the sort order to relevance.
func (q NewsSentiment) SortRelevance() NewsSentiment {
	return q.Sort(SortOptionRelevance)
}

// Limit sets the number of results to return.
// Default: 50, Max: 1000
func (q NewsSentiment) Limit(limit string) NewsSentiment {
	q[KeyLimit] = []string{limit}
	return q
}

// Validate checks if all required parameters are present and enum values are valid.
func (q NewsSentiment) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	// Validate optional fields if present
	if sort, ok := q[KeySort]; ok {
		if err := validateSort(sort); err != nil {
			return err
		}
	}

	// Note: tickers, topics, time_from, time_to, limit are free-form strings

	return nil
}

// Encode returns the URL-encoded query string.
func (q NewsSentiment) Encode() string { return encode(q) }

// EarningsCallTranscript builds query parameters for the EARNINGS_CALL_TRANSCRIPT API.
// Returns earnings call transcript for a company in a specific quarter.
type EarningsCallTranscript url.Values

// NewEarningsCallTranscript creates a new EARNINGS_CALL_TRANSCRIPT query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: Ticker symbol (e.g., "IBM")
//	quarter: Fiscal quarter in YYYYQM format (e.g., "2024Q1"). Supports quarters since 2010Q1.
func NewEarningsCallTranscript(apikey, symbol, quarter string) EarningsCallTranscript {
	return EarningsCallTranscript{
		KeyFunction: []string{FunctionEarningsCallTranscript},
		KeySymbol:   []string{symbol},
		KeyQuarter:  []string{quarter},
		KeyAPIKey:   []string{apikey},
	}
}

// Validate checks if all required parameters are present.
func (q EarningsCallTranscript) Validate() error {
	return validateRequired(q, KeyAPIKey, KeySymbol, KeyQuarter)
}

// Encode returns the URL-encoded query string.
func (q EarningsCallTranscript) Encode() string { return encode(q) }

// TopGainersLosers builds query parameters for the TOP_GAINERS_LOSERS API.
// Returns top 20 gainers, losers, and most active traded tickers in the US market.
type TopGainersLosers url.Values

// NewTopGainersLosers creates a new TOP_GAINERS_LOSERS query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewTopGainersLosers(apikey string) TopGainersLosers {
	return TopGainersLosers{
		KeyFunction: []string{FunctionTopGainersLosers},
		KeyAPIKey:   []string{apikey},
	}
}

// Validate checks if all required parameters are present.
func (q TopGainersLosers) Validate() error {
	return validateRequired(q, KeyAPIKey)
}

// Encode returns the URL-encoded query string.
func (q TopGainersLosers) Encode() string { return encode(q) }

// InsiderTransactions builds query parameters for the INSIDER_TRANSACTIONS API.
// Returns latest and historical insider transactions by key stakeholders.
type InsiderTransactions url.Values

// NewInsiderTransactions creates a new INSIDER_TRANSACTIONS query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: Ticker symbol (e.g., "IBM")
func NewInsiderTransactions(apikey, symbol string) InsiderTransactions {
	return InsiderTransactions{
		KeyFunction: []string{FunctionInsiderTransactions},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

// Validate checks if all required parameters are present.
func (q InsiderTransactions) Validate() error {
	return validateRequired(q, KeyAPIKey, KeySymbol)
}

// Encode returns the URL-encoded query string.
func (q InsiderTransactions) Encode() string { return encode(q) }

// AnalyticsFixedWindow builds query parameters for the ANALYTICS_FIXED_WINDOW API.
// Returns advanced analytics metrics for time series over a fixed temporal window.
type AnalyticsFixedWindow url.Values

// NewAnalyticsFixedWindow creates a new ANALYTICS_FIXED_WINDOW query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbols: Comma-separated list of symbols (e.g., "AAPL,MSFT,IBM")
//	rangeParam: Date range for the series (e.g., "2023-07-01", "2020-12")
//	interval: Time interval (e.g., "DAILY", "WEEKLY", "MONTHLY")
//	calculations: Comma-separated analytics metrics to calculate
func NewAnalyticsFixedWindow(apikey, symbols, rangeParam, interval, calculations string) AnalyticsFixedWindow {
	return AnalyticsFixedWindow{
		KeyFunction:          []string{FunctionAnalyticsFixedWindow},
		KeySymbols:           []string{symbols},
		KeyRange:             []string{rangeParam},
		KeyIntervalAnalytics: []string{interval},
		KeyCalculations:      []string{calculations},
		KeyAPIKey:            []string{apikey},
	}
}

// OHLC sets the OHLC field for calculation.
// Valid values: "open", "high", "low", "close" (default: "close")
func (q AnalyticsFixedWindow) OHLC(ohlc SeriesTypeOption) AnalyticsFixedWindow {
	q[KeyOHLC] = ohlc.values()
	return q
}

// Validate checks if all required parameters are present and enum values are valid.
func (q AnalyticsFixedWindow) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeySymbols, KeyRange, KeyIntervalAnalytics, KeyCalculations); err != nil {
		return err
	}

	// Validate optional fields if present
	if ohlc, ok := q[KeyOHLC]; ok {
		if err := validateOHLC(ohlc); err != nil {
			return err
		}
	}

	return nil
}

// Encode returns the URL-encoded query string.
func (q AnalyticsFixedWindow) Encode() string { return encode(q) }

// AnalyticsSlidingWindow builds query parameters for the ANALYTICS_SLIDING_WINDOW API.
// Returns advanced analytics metrics for time series over sliding time windows.
type AnalyticsSlidingWindow url.Values

// NewAnalyticsSlidingWindow creates a new ANALYTICS_SLIDING_WINDOW query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbols: Comma-separated list of symbols
//	rangeParam: Date range for the series
//	interval: Time interval
//	windowSize: Size of moving window (minimum 10)
//	calculations: Comma-separated analytics metrics to calculate
func NewAnalyticsSlidingWindow(apikey, symbols, rangeParam, interval, windowSize, calculations string) AnalyticsSlidingWindow {
	return AnalyticsSlidingWindow{
		KeyFunction:          []string{FunctionAnalyticsSlidingWindow},
		KeySymbols:           []string{symbols},
		KeyRange:             []string{rangeParam},
		KeyIntervalAnalytics: []string{interval},
		KeyWindowSize:        []string{windowSize},
		KeyCalculations:      []string{calculations},
		KeyAPIKey:            []string{apikey},
	}
}

// OHLC sets the OHLC field for calculation.
// Valid values: "open", "high", "low", "close" (default: "close")
func (q AnalyticsSlidingWindow) OHLC(ohlc SeriesTypeOption) AnalyticsSlidingWindow {
	q[KeyOHLC] = ohlc.values()
	return q
}

// Validate checks if all required parameters are present and enum values are valid.
func (q AnalyticsSlidingWindow) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeySymbols, KeyRange, KeyIntervalAnalytics, KeyWindowSize, KeyCalculations); err != nil {
		return err
	}

	// Validate optional fields if present
	if ohlc, ok := q[KeyOHLC]; ok {
		if err := validateOHLC(ohlc); err != nil {
			return err
		}
	}

	return nil
}

// Encode returns the URL-encoded query string.
func (q AnalyticsSlidingWindow) Encode() string { return encode(q) }
