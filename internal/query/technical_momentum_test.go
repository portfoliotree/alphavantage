package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/portfoliotree/alphavantage/internal/query"
)

func TestWILLR(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewWILLR(apiKeyTestValue, "IBM", "weekly", 14)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "WILLR", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
		assert.Equal(t, "14", v.Get(query.KeyTimePeriod))
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewWILLR(apiKeyTestValue, "IBM", "weekly", 14)
		s := q.Encode()

		assert.Contains(t, s, "function=WILLR")
		assert.Contains(t, s, "symbol=IBM")
	})
}

func TestAPO(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewAPO(apiKeyTestValue, "IBM", "weekly", query.SeriesTypeOptionClose)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "APO", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
		assert.Equal(t, "close", v.Get(query.KeySeriesType))
	})

	t.Run("with optional parameters", func(t *testing.T) {
		q := query.NewAPO(apiKeyTestValue, "IBM", "weekly", query.SeriesTypeOptionClose).
			FastPeriod(12).
			SlowPeriod(26).
			MAType(0)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "12", v.Get("fastperiod"))
		assert.Equal(t, "26", v.Get("slowperiod"))
		assert.Equal(t, "0", v.Get("matype"))
	})

	t.Run("invalid matype", func(t *testing.T) {
		q := query.NewAPO(apiKeyTestValue, "IBM", "weekly", query.SeriesTypeOptionClose).
			MAType(9)
		err := q.Validate()
		assert.Error(t, err)
	})
}

func TestPPO(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewPPO(apiKeyTestValue, "IBM", "weekly", query.SeriesTypeOptionClose)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "PPO", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})

	t.Run("invalid matype", func(t *testing.T) {
		q := query.NewPPO(apiKeyTestValue, "IBM", "weekly", query.SeriesTypeOptionClose).
			MAType(10)
		err := q.Validate()
		assert.Error(t, err)
	})
}

func TestMOM(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewMOM(apiKeyTestValue, "IBM", "weekly", 10, query.SeriesTypeOptionClose)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "MOM", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
		assert.Equal(t, "10", v.Get(query.KeyTimePeriod))
		assert.Equal(t, "close", v.Get(query.KeySeriesType))
	})
}

func TestBOP(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewBOP(apiKeyTestValue, "IBM", "weekly")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "BOP", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
	})
}

func TestCCI(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewCCI(apiKeyTestValue, "IBM", "weekly", 20)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "CCI", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "20", v.Get(query.KeyTimePeriod))
	})
}

func TestCMO(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewCMO(apiKeyTestValue, "IBM", "weekly", 14, query.SeriesTypeOptionClose)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "CMO", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "14", v.Get(query.KeyTimePeriod))
		assert.Equal(t, "close", v.Get(query.KeySeriesType))
	})
}

func TestROC(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewROC(apiKeyTestValue, "IBM", "weekly", 10, query.SeriesTypeOptionClose)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "ROC", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "10", v.Get(query.KeyTimePeriod))
		assert.Equal(t, "close", v.Get(query.KeySeriesType))
	})
}

func TestROCR(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewROCR(apiKeyTestValue, "IBM", "weekly", 10, query.SeriesTypeOptionClose)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "ROCR", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "10", v.Get(query.KeyTimePeriod))
		assert.Equal(t, "close", v.Get(query.KeySeriesType))
	})
}
