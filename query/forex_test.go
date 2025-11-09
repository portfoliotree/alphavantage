package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/portfoliotree/alphavantage/query"
)

func TestCurrencyExchangeRate(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewCurrencyExchangeRate(apiKeyTestValue, "USD", "JPY")

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionCurrencyExchangeRate, v.Get(query.KeyFunction))
		assert.Equal(t, "USD", v.Get(query.KeyFromCurrency))
		assert.Equal(t, "JPY", v.Get(query.KeyToCurrency))
		assert.Equal(t, apiKeyTestValue, v.Get(query.KeyAPIKey))
	})

	t.Run("crypto to fiat", func(t *testing.T) {
		q := query.NewCurrencyExchangeRate(apiKeyTestValue, "BTC", "EUR")

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "BTC", v.Get(query.KeyFromCurrency))
		assert.Equal(t, "EUR", v.Get(query.KeyToCurrency))
	})

	t.Run("missing apikey", func(t *testing.T) {
		q := query.CurrencyExchangeRate{
			query.KeyFunction:     []string{query.FunctionCurrencyExchangeRate},
			query.KeyFromCurrency: []string{"USD"},
			query.KeyToCurrency:   []string{"JPY"},
		}

		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeyAPIKey)
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewCurrencyExchangeRate(apiKeyTestValue, "USD", "JPY")
		s := q.Encode()

		assert.Contains(t, s, "function=CURRENCY_EXCHANGE_RATE")
		assert.Contains(t, s, "from_currency=USD")
		assert.Contains(t, s, "to_currency=JPY")
		assert.Contains(t, s, "apikey=demo")
	})
}

func TestFXIntraday(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewFXIntraday(apiKeyTestValue, "EUR", "USD", "5min")

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionFXIntraday, v.Get(query.KeyFunction))
		assert.Equal(t, "EUR", v.Get(query.KeyFromSymbol))
		assert.Equal(t, "USD", v.Get(query.KeyToSymbol))
		assert.Equal(t, "5min", v.Get(query.KeyInterval))
		assert.Equal(t, apiKeyTestValue, v.Get(query.KeyAPIKey))
	})

	t.Run("with outputsize", func(t *testing.T) {
		q := query.NewFXIntraday(apiKeyTestValue, "EUR", "USD", "5min").OutputSize(query.OutputSizeOptionFull)

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "full", v.Get(query.KeyOutputSize))
	})

	t.Run("with datatype", func(t *testing.T) {
		q := query.NewFXIntraday(apiKeyTestValue, "EUR", "USD", "5min").DataTypeCSV()

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("invalid interval", func(t *testing.T) {
		q := query.NewFXIntraday(apiKeyTestValue, "EUR", "USD", "10min")

		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeyInterval)
	})

	t.Run("invalid outputsize", func(t *testing.T) {
		q := query.NewFXIntraday(apiKeyTestValue, "EUR", "USD", "5min")
		q[query.KeyOutputSize] = []string{"partial"} // Directly inject invalid value to test validator

		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeyOutputSize)
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewFXIntraday(apiKeyTestValue, "EUR", "USD", "5min").OutputSize(query.OutputSizeOptionFull)
		s := q.Encode()

		assert.Contains(t, s, "function=FX_INTRADAY")
		assert.Contains(t, s, "from_symbol=EUR")
		assert.Contains(t, s, "to_symbol=USD")
		assert.Contains(t, s, "interval=5min")
		assert.Contains(t, s, "outputsize=full")
	})
}

func TestFXDaily(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewFXDaily(apiKeyTestValue, "EUR", "USD")

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionFXDaily, v.Get(query.KeyFunction))
		assert.Equal(t, "EUR", v.Get(query.KeyFromSymbol))
		assert.Equal(t, "USD", v.Get(query.KeyToSymbol))
		assert.Equal(t, apiKeyTestValue, v.Get(query.KeyAPIKey))
	})

	t.Run("with outputsize", func(t *testing.T) {
		q := query.NewFXDaily(apiKeyTestValue, "EUR", "USD").OutputSize(query.OutputSizeOptionFull)

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "full", v.Get(query.KeyOutputSize))
	})

	t.Run("with datatype", func(t *testing.T) {
		q := query.NewFXDaily(apiKeyTestValue, "EUR", "USD").DataTypeCSV()

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("missing required field", func(t *testing.T) {
		q := query.FXDaily{
			query.KeyFunction:   []string{query.FunctionFXDaily},
			query.KeyFromSymbol: []string{"EUR"},
		}

		err := q.Validate()
		assert.Error(t, err)
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewFXDaily(apiKeyTestValue, "EUR", "USD").OutputSize(query.OutputSizeOptionFull).DataTypeCSV()
		s := q.Encode()

		assert.Contains(t, s, "function=FX_DAILY")
		assert.Contains(t, s, "outputsize=full")
		assert.Contains(t, s, "datatype=csv")
	})
}

func TestFXWeekly(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewFXWeekly(apiKeyTestValue, "EUR", "USD")

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionFXWeekly, v.Get(query.KeyFunction))
		assert.Equal(t, "EUR", v.Get(query.KeyFromSymbol))
		assert.Equal(t, "USD", v.Get(query.KeyToSymbol))
		assert.Equal(t, apiKeyTestValue, v.Get(query.KeyAPIKey))
	})

	t.Run("with datatype", func(t *testing.T) {
		q := query.NewFXWeekly(apiKeyTestValue, "EUR", "USD").DataTypeJSON()

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "json", v.Get(query.KeyDataType))
	})

	t.Run("invalid datatype", func(t *testing.T) {
		q := query.NewFXWeekly(apiKeyTestValue, "EUR", "USD")
		q[query.KeyDataType] = []string{"cake"}

		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeyDataType)
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewFXWeekly(apiKeyTestValue, "EUR", "USD")
		s := q.Encode()

		assert.Contains(t, s, "function=FX_WEEKLY")
		assert.Contains(t, s, "from_symbol=EUR")
		assert.Contains(t, s, "to_symbol=USD")
	})
}

func TestFXMonthly(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewFXMonthly(apiKeyTestValue, "EUR", "USD")

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionFXMonthly, v.Get(query.KeyFunction))
		assert.Equal(t, "EUR", v.Get(query.KeyFromSymbol))
		assert.Equal(t, "USD", v.Get(query.KeyToSymbol))
		assert.Equal(t, apiKeyTestValue, v.Get(query.KeyAPIKey))
	})

	t.Run("with datatype", func(t *testing.T) {
		q := query.NewFXMonthly(apiKeyTestValue, "EUR", "USD").DataTypeCSV()

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("missing required field", func(t *testing.T) {
		q := query.FXMonthly{
			query.KeyFunction: []string{query.FunctionFXMonthly},
			query.KeyToSymbol: []string{"USD"},
			query.KeyAPIKey:   []string{apiKeyTestValue},
		}

		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeyFromSymbol)
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewFXMonthly(apiKeyTestValue, "EUR", "USD").DataTypeCSV()
		s := q.Encode()

		assert.Contains(t, s, "function=FX_MONTHLY")
		assert.Contains(t, s, "datatype=csv")
	})
}
