package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/portfoliotree/alphavantage/query"
)

func TestTimeSeriesIntraday(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewTimeSeriesIntraday(apiKeyTestValue, "IBM", "5min")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "TIME_SERIES_INTRADAY", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "5min", v.Get(query.KeyInterval))
		assert.Equal(t, apiKeyTestValue, v.Get(query.KeyAPIKey))
	})

	t.Run("with all options", func(t *testing.T) {
		q := query.NewTimeSeriesIntraday(apiKeyTestValue, "IBM", "15min").
			Adjusted(false).
			ExtendedHours(false).
			MonthString("2009-01").
			OutputSize(query.OutputSizeOptionFull).
			DataTypeCSV()

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "false", v.Get("adjusted"))
		assert.Equal(t, "false", v.Get("extended_hours"))
		assert.Equal(t, "2009-01", v.Get(query.KeyMonth))
		assert.Equal(t, "full", v.Get(query.KeyOutputSize))
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("invalid interval", func(t *testing.T) {
		q := query.NewTimeSeriesIntraday(apiKeyTestValue, "IBM", "10min")
		err := q.Validate()
		assert.Error(t, err)
	})

	t.Run("invalid adjusted boolean", func(t *testing.T) {
		q := query.NewTimeSeriesIntraday(apiKeyTestValue, "IBM", "5min")
		q[query.KeyAdjusted] = []string{"banana"}

		err := q.Validate()
		assert.Error(t, err)
	})

	t.Run("invalid extended_hours boolean", func(t *testing.T) {
		q := query.NewTimeSeriesIntraday(apiKeyTestValue, "IBM", "5min")
		q[query.KeyExtendedHours] = []string{"banana"}
		err := q.Validate()
		assert.Error(t, err)
	})
}

func TestTimeSeriesDaily(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewTimeSeriesDaily(apiKeyTestValue, "IBM")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "TIME_SERIES_DAILY", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})

	t.Run("with outputSize and dataType", func(t *testing.T) {
		q := query.NewTimeSeriesDaily(apiKeyTestValue, "IBM").
			OutputSize("full").
			DataTypeCSV()

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "full", v.Get(query.KeyOutputSize))
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("invalid outputsize", func(t *testing.T) {
		q := query.NewTimeSeriesDaily(apiKeyTestValue, "IBM")
		q[query.KeyOutputSize] = []string{"all"} // Directly inject invalid value to test validator
		err := q.Validate()
		assert.Error(t, err)
	})
}

func TestTimeSeriesDailyAdjusted(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewTimeSeriesDailyAdjusted(apiKeyTestValue, "IBM")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "TIME_SERIES_DAILY_ADJUSTED", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})

	t.Run("with options", func(t *testing.T) {
		q := query.NewTimeSeriesDailyAdjusted(apiKeyTestValue, "IBM").
			OutputSize("compact").
			DataTypeJSON()

		err := q.Validate()
		assert.NoError(t, err)
	})
}

func TestTimeSeriesWeekly(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewTimeSeriesWeekly(apiKeyTestValue, "IBM")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "TIME_SERIES_WEEKLY", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewTimeSeriesWeekly(apiKeyTestValue, "IBM").DataTypeCSV()
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("invalid dataType", func(t *testing.T) {
		q := query.NewTimeSeriesWeekly(apiKeyTestValue, "IBM")
		q[query.KeyDataType] = []string{"cake"}
		err := q.Validate()
		assert.Error(t, err)
	})
}

func TestTimeSeriesWeeklyAdjusted(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewTimeSeriesWeeklyAdjusted(apiKeyTestValue, "IBM")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "TIME_SERIES_WEEKLY_ADJUSTED", v.Get(query.KeyFunction))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewTimeSeriesWeeklyAdjusted(apiKeyTestValue, "IBM").DataTypeJSON()
		err := q.Validate()
		assert.NoError(t, err)
	})
}

func TestTimeSeriesMonthly(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewTimeSeriesMonthly(apiKeyTestValue, "IBM")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "TIME_SERIES_MONTHLY", v.Get(query.KeyFunction))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewTimeSeriesMonthly(apiKeyTestValue, "TSCO.LON").DataTypeCSV()
		err := q.Validate()
		assert.NoError(t, err)
	})
}

func TestTimeSeriesMonthlyAdjusted(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewTimeSeriesMonthlyAdjusted(apiKeyTestValue, "IBM")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "TIME_SERIES_MONTHLY_ADJUSTED", v.Get(query.KeyFunction))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewTimeSeriesMonthlyAdjusted(apiKeyTestValue, "IBM").DataTypeCSV()
		err := q.Validate()
		assert.NoError(t, err)
	})
}

func TestRealtimeBulkQuotes(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewRealtimeBulkQuotes(apiKeyTestValue, "MSFT,AAPL,IBM")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "REALTIME_BULK_QUOTES", v.Get(query.KeyFunction))
		assert.Equal(t, "MSFT,AAPL,IBM", v.Get(query.KeySymbol))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewRealtimeBulkQuotes(apiKeyTestValue, "MSFT,AAPL").DataTypeCSV()
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewRealtimeBulkQuotes(apiKeyTestValue, "MSFT,AAPL")
		s := q.Encode()

		assert.Contains(t, s, "function=REALTIME_BULK_QUOTES")
		assert.Contains(t, s, "symbol=MSFT")
	})
}

func TestSymbolSearch(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewSymbolSearch(apiKeyTestValue, "microsoft")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "SYMBOL_SEARCH", v.Get(query.KeyFunction))
		assert.Equal(t, "microsoft", v.Get(query.KeyKeywords))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewSymbolSearch(apiKeyTestValue, "tesco").DataTypeJSON()
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "json", v.Get(query.KeyDataType))
	})

	t.Run("missing required field", func(t *testing.T) {
		q := query.SymbolSearch{
			query.KeyFunction: []string{query.FunctionSymbolSearch},
			query.KeyAPIKey:   []string{apiKeyTestValue},
		}
		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeyKeywords)
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewSymbolSearch(apiKeyTestValue, "BA")
		s := q.Encode()

		assert.Contains(t, s, "function=SYMBOL_SEARCH")
		assert.Contains(t, s, "keywords=BA")
	})
}

func TestMarketStatus(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewMarketStatus(apiKeyTestValue)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "MARKET_STATUS", v.Get(query.KeyFunction))
		assert.Equal(t, apiKeyTestValue, v.Get(query.KeyAPIKey))
	})

	t.Run("missing required field", func(t *testing.T) {
		q := query.MarketStatus{
			query.KeyFunction: []string{query.FunctionMarketStatus},
		}
		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeyAPIKey)
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewMarketStatus(apiKeyTestValue)
		s := q.Encode()

		assert.Contains(t, s, "function=MARKET_STATUS")
		assert.Contains(t, s, "apikey=demo")
	})
}

func TestGlobalQuote(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewGlobalQuote(apiKeyTestValue, "IBM")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "GLOBAL_QUOTE", v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, apiKeyTestValue, v.Get(query.KeyAPIKey))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewGlobalQuote(apiKeyTestValue, "IBM").DataTypeCSV()
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("invalid dataType", func(t *testing.T) {
		q := query.NewGlobalQuote(apiKeyTestValue, "IBM")
		q[query.KeyDataType] = []string{"cake"}
		err := q.Validate()
		assert.Error(t, err)
	})

	t.Run("missing required field", func(t *testing.T) {
		q := query.GlobalQuote{
			query.KeyFunction: []string{query.FunctionGlobalQuote},
			query.KeyAPIKey:   []string{apiKeyTestValue},
		}
		err := q.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), query.KeySymbol)
	})

	t.Run("Encode() method", func(t *testing.T) {
		q := query.NewGlobalQuote(apiKeyTestValue, "IBM")
		s := q.Encode()

		assert.Contains(t, s, "function=GLOBAL_QUOTE")
		assert.Contains(t, s, "symbol=IBM")
		assert.Contains(t, s, "apikey=demo")
	})
}
