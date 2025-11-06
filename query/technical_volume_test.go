package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/portfoliotree/alphavantage/query"
)

func TestMFI(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewMFI(apiKeyTestValue, "IBM", "weekly", 14)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "MFI", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
		assert.Equal(t, "14", v.Get(query.KeyTimePeriod))
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewMFI(apiKeyTestValue, "IBM", "weekly", 14)
		s := q.Encode()

		assert.Contains(t, s, "function=MFI")
		assert.Contains(t, s, "symbol=IBM")
	})
}

func TestTRIX(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewTRIX(apiKeyTestValue, "IBM", "weekly", 30, query.SeriesTypeOptionClose)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "TRIX", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "30", v.Get(query.KeyTimePeriod))
		assert.Equal(t, "close", v.Get(query.KeySeriesType))
	})
}

func TestAD(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewAD(apiKeyTestValue, "IBM", "weekly")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "AD", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
	})
}

func TestADOSC(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewADOSC(apiKeyTestValue, "IBM", "weekly")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "ADOSC", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
	})

	t.Run("with optional parameters", func(t *testing.T) {
		q := query.NewADOSC(apiKeyTestValue, "IBM", "weekly").
			FastPeriod(3).
			SlowPeriod(10)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "3", v.Get("fastperiod"))
		assert.Equal(t, "10", v.Get("slowperiod"))
	})
}

func TestOBV(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewOBV(apiKeyTestValue, "IBM", "weekly")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "OBV", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
	})
}
