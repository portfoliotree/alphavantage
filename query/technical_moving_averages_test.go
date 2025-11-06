package query_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/portfoliotree/alphavantage/query"
)

func TestSMA(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewSMA(apiKeyTestValue, "IBM", "weekly", 10, query.SeriesTypeOptionOpen)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "SMA", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
		assert.Equal(t, "10", v.Get(query.KeyTimePeriod))
		assert.Equal(t, "open", v.Get(query.KeySeriesType))
	})

	t.Run("with optional parameters", func(t *testing.T) {
		q := query.NewSMA(apiKeyTestValue, "IBM", "weekly", 10, query.SeriesTypeOptionOpen).
			Month(2009, time.January).
			DataTypeCSV()
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "2009-01", v.Get(query.KeyMonth))
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("invalid series_type", func(t *testing.T) {
		q := query.NewSMA(apiKeyTestValue, "IBM", "weekly", 10, query.SeriesTypeOptionClose)
		q[query.KeySeriesType] = []string{"volume"} // Directly inject invalid value to test validator
		err := q.Validate()
		assert.Error(t, err)
	})
}

func TestEMA(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewEMA(apiKeyTestValue, "IBM", "weekly", 10, query.SeriesTypeOptionOpen)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "EMA", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})
}

func TestWMA(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewWMA(apiKeyTestValue, "IBM", "weekly", 10, query.SeriesTypeOptionOpen)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "WMA", v.Get(query.KeyFunction))
	})
}

func TestDEMA(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewDEMA(apiKeyTestValue, "IBM", "weekly", 10, query.SeriesTypeOptionOpen)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "DEMA", v.Get(query.KeyFunction))
	})
}

func TestTEMA(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewTEMA(apiKeyTestValue, "IBM", "weekly", 10, query.SeriesTypeOptionOpen)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "TEMA", v.Get(query.KeyFunction))
	})
}

func TestTRIMA(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewTRIMA(apiKeyTestValue, "IBM", "weekly", 10, query.SeriesTypeOptionOpen)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "TRIMA", v.Get(query.KeyFunction))
	})
}

func TestKAMA(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewKAMA(apiKeyTestValue, "IBM", "weekly", 10, query.SeriesTypeOptionOpen)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "KAMA", v.Get(query.KeyFunction))
	})
}

func TestT3(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewT3(apiKeyTestValue, "IBM", "weekly", 10, query.SeriesTypeOptionOpen)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "T3", v.Get(query.KeyFunction))
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewT3(apiKeyTestValue, "IBM", "weekly", 10, query.SeriesTypeOptionOpen)
		s := q.Encode()

		assert.Contains(t, s, "function=T3")
		assert.Contains(t, s, "symbol=IBM")
	})
}
