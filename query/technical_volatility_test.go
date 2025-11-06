package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/portfoliotree/alphavantage/query"
)

func TestBBANDS(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewBBANDS(apiKeyTestValue, "IBM", "weekly", 20, query.SeriesTypeOptionClose)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "BBANDS", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
		assert.Equal(t, "20", v.Get(query.KeyTimePeriod))
		assert.Equal(t, "close", v.Get(query.KeySeriesType))
	})

	t.Run("with optional parameters", func(t *testing.T) {
		q := query.NewBBANDS(apiKeyTestValue, "IBM", "weekly", 20, query.SeriesTypeOptionClose).
			NBDEVUP(2).
			NBDEVDN(2).
			MAType(0)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "2", v.Get("nbdevup"))
		assert.Equal(t, "2", v.Get("nbdevdn"))
		assert.Equal(t, "0", v.Get("matype"))
	})

	t.Run("invalid matype", func(t *testing.T) {
		q := query.NewBBANDS(apiKeyTestValue, "IBM", "weekly", 20, query.SeriesTypeOptionClose).
			MAType(9)
		err := q.Validate()
		assert.Error(t, err)
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewBBANDS(apiKeyTestValue, "IBM", "weekly", 20, query.SeriesTypeOptionClose)
		s := q.Encode()

		assert.Contains(t, s, "function=BBANDS")
		assert.Contains(t, s, "symbol=IBM")
	})
}

func TestTRANGE(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewTRANGE(apiKeyTestValue, "IBM", "weekly")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "TRANGE", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
	})
}

func TestATR(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewATR(apiKeyTestValue, "IBM", "weekly", 14)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "ATR", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "14", v.Get(query.KeyTimePeriod))
	})
}

func TestNATR(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewNATR(apiKeyTestValue, "IBM", "weekly", 14)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "NATR", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "14", v.Get(query.KeyTimePeriod))
	})
}

func TestSAR(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewSAR(apiKeyTestValue, "IBM", "weekly")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "SAR", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
	})

	t.Run("with optional parameters", func(t *testing.T) {
		q := query.NewSAR(apiKeyTestValue, "IBM", "weekly").
			Acceleration(0.01).
			Maximum(0.20)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "0.01", v.Get("acceleration"))
		assert.Equal(t, "0.2", v.Get("maximum"))
	})
}

func TestULTOSC(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewULTOSC(apiKeyTestValue, "IBM", "weekly")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "ULTOSC", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
	})

	t.Run("with optional parameters", func(t *testing.T) {
		q := query.NewULTOSC(apiKeyTestValue, "IBM", "weekly").
			TimePeriod1(7).
			TimePeriod2(14).
			TimePeriod3(28)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "7", v.Get("timeperiod1"))
		assert.Equal(t, "14", v.Get("timeperiod2"))
		assert.Equal(t, "28", v.Get("timeperiod3"))
	})
}
