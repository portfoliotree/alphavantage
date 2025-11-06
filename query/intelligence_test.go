package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/portfoliotree/alphavantage/query"
)

func TestNewsSentiment(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewNewsSentiment(apiKeyTestValue)
		err := q.Validate()
		assert.NoError(t, err)
	})

	t.Run("with all options", func(t *testing.T) {
		q := query.NewNewsSentiment(apiKeyTestValue).
			Tickers("AAPL").
			Topics("technology").
			TimeFrom("20220410T0130").
			Sort("LATEST").
			Limit("50")

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "AAPL", v.Get("tickers"))
	})

	t.Run("invalid sort", func(t *testing.T) {
		q := query.NewNewsSentiment(apiKeyTestValue)
		q[query.KeySort] = []string{"INVALID"} // Directly inject invalid value to test validator
		err := q.Validate()
		assert.Error(t, err)
	})
}

func TestEarningsCallTranscript(t *testing.T) {
	t.Run("valid query", func(t *testing.T) {
		q := query.NewEarningsCallTranscript(apiKeyTestValue, "IBM", "2024Q1")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "EARNINGS_CALL_TRANSCRIPT", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "2024Q1", v.Get(query.KeyQuarter))
	})
}

func TestTopGainersLosers(t *testing.T) {
	t.Run("valid query", func(t *testing.T) {
		q := query.NewTopGainersLosers(apiKeyTestValue)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "TOP_GAINERS_LOSERS", v.Get(query.KeyFunction))
	})
}

func TestInsiderTransactions(t *testing.T) {
	t.Run("valid query", func(t *testing.T) {
		q := query.NewInsiderTransactions(apiKeyTestValue, "IBM")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "INSIDER_TRANSACTIONS", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})
}

func TestAnalyticsFixedWindow(t *testing.T) {
	t.Run("valid query", func(t *testing.T) {
		q := query.NewAnalyticsFixedWindow(apiKeyTestValue, "AAPL,MSFT", "2023-07-01", "DAILY", "MEAN,STDDEV")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "ANALYTICS_FIXED_WINDOW", v.Get(query.KeyFunction))
	})

	t.Run("with OHLC", func(t *testing.T) {
		q := query.NewAnalyticsFixedWindow(apiKeyTestValue, "AAPL", "2023-07-01", "DAILY", "MEAN").OHLC(query.SeriesTypeOptionClose)
		err := q.Validate()
		assert.NoError(t, err)
	})

	t.Run("invalid OHLC", func(t *testing.T) {
		q := query.NewAnalyticsFixedWindow(apiKeyTestValue, "AAPL", "2023-07-01", "DAILY", "MEAN").OHLC("volume")
		err := q.Validate()
		assert.Error(t, err)
	})
}

func TestAnalyticsSlidingWindow(t *testing.T) {
	t.Run("valid query", func(t *testing.T) {
		q := query.NewAnalyticsSlidingWindow(apiKeyTestValue, "AAPL,IBM", "2month", "DAILY", "20", "MEAN")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "ANALYTICS_SLIDING_WINDOW", v.Get(query.KeyFunction))
		assert.Equal(t, "20", v.Get(query.KeyWindowSize))
	})

	t.Run("with OHLC", func(t *testing.T) {
		q := query.NewAnalyticsSlidingWindow(apiKeyTestValue, "AAPL", "2month", "DAILY", "20", "MEAN").OHLC(query.SeriesTypeOptionHigh)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "high", v.Get("OHLC"))
	})
}
