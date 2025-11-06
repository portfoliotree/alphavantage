package query_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/portfoliotree/alphavantage/internal/query"
)

func TestHT_TRENDLINE(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewHilbertTransformTrendLine(apiKeyTestValue, "IBM", "weekly", query.SeriesTypeOptionClose)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "HT_TRENDLINE", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
		assert.Equal(t, "close", v.Get(query.KeySeriesType))
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewHilbertTransformTrendLine(apiKeyTestValue, "IBM", "weekly", query.SeriesTypeOptionClose)
		s := q.Encode()

		assert.Contains(t, s, "function=HT_TRENDLINE")
		assert.Contains(t, s, "symbol=IBM")
	})
}

func TestHT_SINE(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewHilbertTransformSine(apiKeyTestValue, "IBM", "weekly", query.SeriesTypeOptionClose).Month(2004, time.May)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "HT_SINE", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "2004-05", v.Get(query.KeyMonth))
		assert.Equal(t, "close", v.Get(query.KeySeriesType))
	})
}

func TestHT_TRENDMODE(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewHilbertTransformTrendMode(apiKeyTestValue, "IBM", "weekly", query.SeriesTypeOptionClose)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "HT_TRENDMODE", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})
}

func TestHT_DCPERIOD(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewHilbertTransformDCPeriod(apiKeyTestValue, "IBM", "weekly", query.SeriesTypeOptionClose)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "HT_DCPERIOD", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})
}

func TestHT_DCPHASE(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewHilbertTransformDCPhase(apiKeyTestValue, "IBM", "weekly", query.SeriesTypeOptionClose)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "HT_DCPHASE", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})
}

func TestHT_PHASOR(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewHilbertTransformPhasor(apiKeyTestValue, "IBM", "weekly", query.SeriesTypeOptionClose)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "HT_PHASOR", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})
}
