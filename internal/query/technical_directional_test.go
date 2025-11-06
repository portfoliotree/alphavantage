package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/portfoliotree/alphavantage/internal/query"
)

func TestDX(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewDX(apiKeyTestValue, "IBM", "weekly", 14)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "DX", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
		assert.Equal(t, "14", v.Get(query.KeyTimePeriod))
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewDX(apiKeyTestValue, "IBM", "weekly", 14)
		s := q.Encode()

		assert.Contains(t, s, "function=DX")
		assert.Contains(t, s, "symbol=IBM")
	})
}

func TestADX(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewADX(apiKeyTestValue, "IBM", "weekly", 14)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "ADX", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "14", v.Get(query.KeyTimePeriod))
	})
}

func TestADXR(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewADXR(apiKeyTestValue, "IBM", "weekly", 14)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "ADXR", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})
}

func TestMINUSDI(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewMINUSDI(apiKeyTestValue, "IBM", "weekly", 14)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "MINUS_DI", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})
}

func TestPLUSDI(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewPLUSDI(apiKeyTestValue, "IBM", "weekly", 14)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "PLUS_DI", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})
}

func TestMINUSDM(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewMINUSDM(apiKeyTestValue, "IBM", "weekly", 14)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "MINUS_DM", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})
}

func TestPLUSDM(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewPLUSDM(apiKeyTestValue, "IBM", "weekly", 14)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "PLUS_DM", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})
}

func TestAROON(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewAROON(apiKeyTestValue, "IBM", "weekly", 25)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "AROON", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "25", v.Get(query.KeyTimePeriod))
	})
}

func TestAROONOSC(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewAROONOSC(apiKeyTestValue, "IBM", "weekly", 25)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "AROONOSC", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "25", v.Get(query.KeyTimePeriod))
	})
}
