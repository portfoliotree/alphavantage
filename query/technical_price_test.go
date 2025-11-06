package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/portfoliotree/alphavantage/query"
)

func TestMIDPOINT(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewMidpoint(apiKeyTestValue, "IBM", "weekly", 14, query.SeriesTypeOptionClose)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "Midpoint", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
		assert.Equal(t, "14", v.Get(query.KeyTimePeriod))
		assert.Equal(t, "close", v.Get(query.KeySeriesType))
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewMidpoint(apiKeyTestValue, "IBM", "weekly", 14, query.SeriesTypeOptionClose)
		s := q.Encode()

		assert.Contains(t, s, "function=Midpoint")
		assert.Contains(t, s, "symbol=IBM")
	})
}

func TestMIDPRICE(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewMIDPRICE(apiKeyTestValue, "IBM", "weekly", 14)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "MIDPRICE", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
		assert.Equal(t, "14", v.Get(query.KeyTimePeriod))
	})
}
