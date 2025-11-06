package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/portfoliotree/alphavantage/query"
)

func TestMAMA(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewMAMA(apiKeyTestValue, "IBM", "daily", query.SeriesTypeOptionClose)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "MAMA", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "daily", v.Get(query.KeyInterval))
		assert.Equal(t, "close", v.Get(query.KeySeriesType))
	})

	t.Run("with optional parameters", func(t *testing.T) {
		q := query.NewMAMA(apiKeyTestValue, "IBM", "daily", query.SeriesTypeOptionClose).
			FastLimit("0.02").
			SlowLimit("0.01")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "0.02", v.Get("fastlimit"))
		assert.Equal(t, "0.01", v.Get("slowlimit"))
	})
}

func TestVWAP(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewVWAP(apiKeyTestValue, "IBM", "15min")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "VWAP", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "15min", v.Get(query.KeyInterval))
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewVWAP(apiKeyTestValue, "IBM", "15min")
		s := q.Encode()

		assert.Contains(t, s, "function=VWAP")
		assert.Contains(t, s, "symbol=IBM")
	})
}
