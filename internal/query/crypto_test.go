package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/portfoliotree/alphavantage/internal/query"
)

func TestCryptoIntraday(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewCryptoIntraday(apiKeyTestValue, "ETH", "USD", "5min")

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionCryptoIntraday, v.Get(query.KeyFunction))
		assert.Equal(t, "ETH", v.Get(query.KeySymbol))
		assert.Equal(t, "USD", v.Get(query.KeyMarket))
		assert.Equal(t, "5min", v.Get(query.KeyInterval))
		assert.Equal(t, apiKeyTestValue, v.Get(query.KeyAPIKey))
	})

	t.Run("with outputsize", func(t *testing.T) {
		q := query.NewCryptoIntraday(apiKeyTestValue, "ETH", "USD", query.IIntervalOption5min).OutputSize(query.OutputSizeOptionFull)

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "full", v.Get(query.KeyOutputSize))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewCryptoIntraday(apiKeyTestValue, "ETH", "USD", query.IIntervalOption5min).DataTypeCSV()

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("invalid interval", func(t *testing.T) {
		q := query.NewCryptoIntraday(apiKeyTestValue, "ETH", "USD", "10min")

		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeyInterval)
	})

	t.Run("missing required field", func(t *testing.T) {
		q := query.CryptoIntraday{
			query.KeyFunction: []string{query.FunctionCryptoIntraday},
			query.KeySymbol:   []string{"ETH"},
			query.KeyMarket:   []string{"USD"},
		}

		err := q.Validate()
		assert.Error(t, err)
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewCryptoIntraday(apiKeyTestValue, "ETH", "USD", query.IIntervalOption5min).OutputSize(query.OutputSizeOptionFull)
		s := q.Encode()

		assert.Contains(t, s, "function=CRYPTO_INTRADAY")
		assert.Contains(t, s, "symbol=ETH")
		assert.Contains(t, s, "market=USD")
		assert.Contains(t, s, "interval=5min")
		assert.Contains(t, s, "outputsize=full")
	})
}

func TestDigitalCurrencyDaily(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewDigitalCurrencyDaily(apiKeyTestValue, "BTC", "EUR")

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionDigitalCurrencyDaily, v.Get(query.KeyFunction))
		assert.Equal(t, "BTC", v.Get(query.KeySymbol))
		assert.Equal(t, "EUR", v.Get(query.KeyMarket))
		assert.Equal(t, apiKeyTestValue, v.Get(query.KeyAPIKey))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewDigitalCurrencyDaily(apiKeyTestValue, "BTC", "EUR").DataTypeCSV()

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("invalid dataType", func(t *testing.T) {
		q := query.NewDigitalCurrencyDaily(apiKeyTestValue, "BTC", "EUR")
		q[query.KeyDataType] = []string{"cake"}

		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeyDataType)
	})

	t.Run("missing required field", func(t *testing.T) {
		q := query.DigitalCurrencyDaily{
			query.KeyFunction: []string{query.FunctionDigitalCurrencyDaily},
			query.KeySymbol:   []string{"BTC"},
		}

		err := q.Validate()
		assert.Error(t, err)
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewDigitalCurrencyDaily(apiKeyTestValue, "BTC", "EUR")
		s := q.Encode()

		assert.Contains(t, s, "function=DIGITAL_CURRENCY_DAILY")
		assert.Contains(t, s, "symbol=BTC")
		assert.Contains(t, s, "market=EUR")
	})
}

func TestDigitalCurrencyWeekly(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewDigitalCurrencyWeekly(apiKeyTestValue, "BTC", "EUR")

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionDigitalCurrencyWeekly, v.Get(query.KeyFunction))
		assert.Equal(t, "BTC", v.Get(query.KeySymbol))
		assert.Equal(t, "EUR", v.Get(query.KeyMarket))
		assert.Equal(t, apiKeyTestValue, v.Get(query.KeyAPIKey))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewDigitalCurrencyWeekly(apiKeyTestValue, "BTC", "EUR").DataTypeCSV()

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("invalid dataType", func(t *testing.T) {
		q := query.NewDigitalCurrencyWeekly(apiKeyTestValue, "BTC", "EUR")
		q[query.KeyDataType] = []string{"cake"}

		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeyDataType)
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewDigitalCurrencyWeekly(apiKeyTestValue, "BTC", "EUR").DataTypeJSON()
		s := q.Encode()

		assert.Contains(t, s, "function=DIGITAL_CURRENCY_WEEKLY")
		assert.Contains(t, s, "dataType=json")
	})
}

func TestDigitalCurrencyMonthly(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewDigitalCurrencyMonthly(apiKeyTestValue, "BTC", "EUR")

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionDigitalCurrencyMonthly, v.Get(query.KeyFunction))
		assert.Equal(t, "BTC", v.Get(query.KeySymbol))
		assert.Equal(t, "EUR", v.Get(query.KeyMarket))
		assert.Equal(t, apiKeyTestValue, v.Get(query.KeyAPIKey))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewDigitalCurrencyMonthly(apiKeyTestValue, "BTC", "EUR").DataTypeCSV()

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("missing required field", func(t *testing.T) {
		q := query.DigitalCurrencyMonthly{
			query.KeyFunction: []string{query.FunctionDigitalCurrencyMonthly},
			query.KeyMarket:   []string{"EUR"},
			query.KeyAPIKey:   []string{apiKeyTestValue},
		}

		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeySymbol)
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewDigitalCurrencyMonthly(apiKeyTestValue, "BTC", "EUR").DataTypeCSV()
		s := q.Encode()

		assert.Contains(t, s, "function=DIGITAL_CURRENCY_MONTHLY")
		assert.Contains(t, s, "symbol=BTC")
		assert.Contains(t, s, "market=EUR")
		assert.Contains(t, s, "dataType=csv")
	})
}
