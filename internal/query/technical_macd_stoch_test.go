package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/portfoliotree/alphavantage/internal/query"
)

func TestMACD(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewMACD(apiKeyTestValue, "IBM", "daily", query.SeriesTypeOptionOpen)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "MACD", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})

	t.Run("with optional parameters", func(t *testing.T) {
		q := query.NewMACD(apiKeyTestValue, "IBM", "daily", query.SeriesTypeOptionOpen).
			FastPeriod(12).
			SlowPeriod(26).
			SignalPeriod(9)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "12", v.Get("fastperiod"))
		assert.Equal(t, "26", v.Get("slowperiod"))
		assert.Equal(t, "9", v.Get("signalperiod"))
	})
}

func TestMACDEXT(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewMACDEXT(apiKeyTestValue, "IBM", "daily", query.SeriesTypeOptionOpen)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "MACDEXT", v.Get(query.KeyFunction))
	})

	t.Run("with matype parameters", func(t *testing.T) {
		q := query.NewMACDEXT(apiKeyTestValue, "IBM", "daily", query.SeriesTypeOptionOpen).
			FastMAType(1).
			SlowMAType(1).
			SignalMAType(1)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "1", v.Get("fastmatype"))
	})

	t.Run("invalid matype", func(t *testing.T) {
		q := query.NewMACDEXT(apiKeyTestValue, "IBM", "daily", query.SeriesTypeOptionOpen).
			FastMAType(10)
		err := q.Validate()
		assert.Error(t, err)
	})
}

func TestSTOCH(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewSTOCH(apiKeyTestValue, "IBM", "daily")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "STOCH", v.Get(query.KeyFunction))
	})

	t.Run("with optional parameters", func(t *testing.T) {
		q := query.NewSTOCH(apiKeyTestValue, "IBM", "daily").
			FastKPeriod(5).
			SlowKPeriod(3).
			SlowDPeriod(3)
		err := q.Validate()
		assert.NoError(t, err)
	})
}

func TestSTOCHF(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewSTOCHF(apiKeyTestValue, "IBM", "daily")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "STOCHF", v.Get(query.KeyFunction))
	})
}

func TestRSI(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewRSI(apiKeyTestValue, "IBM", "weekly", 10, query.SeriesTypeOptionOpen)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "RSI", v.Get(query.KeyFunction))
		assert.Equal(t, "10", v.Get(query.KeyTimePeriod))
	})
}

func TestSTOCHRSI(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewSTOCHRSI(apiKeyTestValue, "IBM", "weekly", 10, query.SeriesTypeOptionOpen)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "STOCHRSI", v.Get(query.KeyFunction))
	})

	t.Run("with optional parameters", func(t *testing.T) {
		q := query.NewSTOCHRSI(apiKeyTestValue, "IBM", "weekly", 10, query.SeriesTypeOptionOpen).
			FastKPeriod(5).
			FastDPeriod(3)
		err := q.Validate()
		assert.NoError(t, err)
	})
}
