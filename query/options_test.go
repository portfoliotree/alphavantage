package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/portfoliotree/alphavantage/query"
)

// TestRealtimeOptions tests the REALTIME_OPTIONS query builder
func TestRealtimeOptions(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewRealtimeOptions(apiKeyTestValue, "IBM")

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "REALTIME_OPTIONS", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, apiKeyTestValue, v.Get(query.KeyAPIKey))
	})

	t.Run("with require_greeks", func(t *testing.T) {
		q := query.NewRealtimeOptions(apiKeyTestValue, "IBM").RequireGreeks(true)

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "true", v.Get("require_greeks"))
	})

	t.Run("with contract", func(t *testing.T) {
		q := query.NewRealtimeOptions(apiKeyTestValue, "IBM").
			RequireGreeks(true).
			Contract("IBM270115C00390000")

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "IBM270115C00390000", v.Get("contract"))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewRealtimeOptions(apiKeyTestValue, "IBM").DataTypeCSV()

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("missing apikey", func(t *testing.T) {
		q := query.RealtimeOptions{
			query.KeyFunction: []string{query.FunctionRealtimeOptions},
			query.KeySymbol:   []string{"IBM"},
		}

		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeyAPIKey)
	})

	t.Run("missing symbol", func(t *testing.T) {
		q := query.RealtimeOptions{
			query.KeyFunction: []string{query.FunctionRealtimeOptions},
			query.KeyAPIKey:   []string{apiKeyTestValue},
		}

		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeySymbol)
	})

	t.Run("invalid require_greeks", func(t *testing.T) {
		q := query.NewRealtimeOptions(apiKeyTestValue, "IBM")
		q[query.KeyRequireGreeks] = []string{"nope"}

		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "require_greeks")
	})

	t.Run("invalid dataType", func(t *testing.T) {
		q := query.NewRealtimeOptions(apiKeyTestValue, "IBM")
		q[query.KeyDataType] = []string{"cake"}

		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeyDataType)
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewRealtimeOptions(apiKeyTestValue, "IBM")
		s := q.Encode()

		assert.Contains(t, s, "function=REALTIME_OPTIONS")
		assert.Contains(t, s, "symbol=IBM")
		assert.Contains(t, s, "apikey=demo")
	})

	t.Run("String() with optional params", func(t *testing.T) {
		q := query.NewRealtimeOptions(apiKeyTestValue, "IBM").
			RequireGreeks(true).
			Contract("IBM270115C00390000")
		s := q.Encode()

		assert.Contains(t, s, "require_greeks=true")
		assert.Contains(t, s, "contract=IBM270115C00390000")
	})
}

// TestHistoricalOptions tests the HISTORICAL_OPTIONS query builder
func TestHistoricalOptions(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewHistoricalOptions(apiKeyTestValue, "IBM")

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "HISTORICAL_OPTIONS", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, apiKeyTestValue, v.Get(query.KeyAPIKey))
	})

	t.Run("with date", func(t *testing.T) {
		q := query.NewHistoricalOptions(apiKeyTestValue, "IBM").Date("2017-11-15")

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "2017-11-15", v.Get("date"))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewHistoricalOptions(apiKeyTestValue, "IBM").DataTypeJSON()

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "json", v.Get(query.KeyDataType))
	})

	t.Run("missing apikey", func(t *testing.T) {
		q := query.HistoricalOptions{
			query.KeyFunction: []string{query.FunctionHistoricalOptions},
			query.KeySymbol:   []string{"IBM"},
		}

		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeyAPIKey)
	})

	t.Run("missing symbol", func(t *testing.T) {
		q := query.HistoricalOptions{
			query.KeyFunction: []string{query.FunctionHistoricalOptions},
			query.KeyAPIKey:   []string{apiKeyTestValue},
		}

		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeySymbol)
	})

	t.Run("invalid dataType", func(t *testing.T) {
		q := query.NewHistoricalOptions(apiKeyTestValue, "IBM")
		q[query.KeyDataType] = []string{"cake"}

		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeyDataType)
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewHistoricalOptions(apiKeyTestValue, "IBM")
		s := q.Encode()

		assert.Contains(t, s, "function=HISTORICAL_OPTIONS")
		assert.Contains(t, s, "symbol=IBM")
		assert.Contains(t, s, "apikey=demo")
	})

	t.Run("String() with optional params", func(t *testing.T) {
		q := query.NewHistoricalOptions(apiKeyTestValue, "IBM").
			Date("2017-11-15").
			DataTypeCSV()
		s := q.Encode()

		assert.Contains(t, s, "date=2017-11-15")
		assert.Contains(t, s, "dataType=csv")
	})
}
