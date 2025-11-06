package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/portfoliotree/alphavantage/query"
)

func TestOverview(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewOverview(apiKeyTestValue, "IBM")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionOverview, v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewOverview(apiKeyTestValue, "IBM")
		s := q.Encode()

		assert.Contains(t, s, "function=OVERVIEW")
		assert.Contains(t, s, "symbol=IBM")
	})
}

func TestETFProfile(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewETFProfile(apiKeyTestValue, "QQQ")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionETFProfile, v.Get(query.KeyFunction))
		assert.Equal(t, "QQQ", v.Get(query.KeySymbol))
	})
}

func TestDividends(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewDividends(apiKeyTestValue, "IBM")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionDividends, v.Get(query.KeyFunction))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewDividends(apiKeyTestValue, "IBM").DataTypeCSV()
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("invalid dataType", func(t *testing.T) {
		q := query.NewDividends(apiKeyTestValue, "IBM")
		q[query.KeyDataType] = []string{"cake"}
		err := q.Validate()
		assert.Error(t, err)
	})
}

func TestSplits(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewSplits(apiKeyTestValue, "IBM")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionSplits, v.Get(query.KeyFunction))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewSplits(apiKeyTestValue, "IBM").DataTypeJSON()
		err := q.Validate()
		assert.NoError(t, err)
	})
}

func TestIncomeStatement(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewIncomeStatement(apiKeyTestValue, "IBM")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionIncomeStatement, v.Get(query.KeyFunction))
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
	})
}

func TestBalanceSheet(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewBalanceSheet(apiKeyTestValue, "IBM")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionBalanceSheet, v.Get(query.KeyFunction))
	})
}

func TestCashFlow(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewCashFlow(apiKeyTestValue, "IBM")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionCashFlow, v.Get(query.KeyFunction))
	})
}

func TestSharesOutstanding(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewSharesOutstanding(apiKeyTestValue, "MSFT")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionSharesOutstanding, v.Get(query.KeyFunction))
		assert.Equal(t, "MSFT", v.Get(query.KeySymbol))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewSharesOutstanding(apiKeyTestValue, "MSFT").DataTypeCSV()
		err := q.Validate()
		assert.NoError(t, err)
	})
}

func TestEarnings(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewEarnings(apiKeyTestValue, "IBM")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionEarnings, v.Get(query.KeyFunction))
	})
}

func TestEarningsEstimates(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewEarningsEstimates(apiKeyTestValue, "IBM")
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionEarningsEstimates, v.Get(query.KeyFunction))
	})
}

func TestListingStatus(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewListingStatus(apiKeyTestValue)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionListingStatus, v.Get(query.KeyFunction))
	})

	t.Run("with date and state", func(t *testing.T) {
		q := query.NewListingStatus(apiKeyTestValue).
			Date("2014-07-10").
			State("delisted")

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "2014-07-10", v.Get(query.KeyDate))
		assert.Equal(t, "delisted", v.Get(query.KeyState))
	})

	t.Run("invalid state", func(t *testing.T) {
		q := query.NewListingStatus(apiKeyTestValue)
		q[query.KeyState] = []string{"pending"} // Directly inject invalid value to test validator
		err := q.Validate()
		assert.Error(t, err)
	})
}

func TestEarningsCalendar(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewEarningsCalendar(apiKeyTestValue)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "EARNINGS_CALENDAR", v.Get(query.KeyFunction))
	})

	t.Run("with symbol and horizon", func(t *testing.T) {
		q := query.NewEarningsCalendar(apiKeyTestValue).
			Symbol("IBM").
			Horizon("12month")

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "IBM", v.Get(query.KeySymbol))
		assert.Equal(t, "12month", v.Get(query.KeyHorizon))
	})

	t.Run("invalid horizon", func(t *testing.T) {
		q := query.NewEarningsCalendar(apiKeyTestValue)
		q[query.KeyHorizon] = []string{"24month"} // Directly inject invalid value to test validator
		err := q.Validate()
		assert.Error(t, err)
	})
}

func TestIPOCalendar(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewIPOCalendar(apiKeyTestValue)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, query.FunctionIPOCalendar, v.Get(query.KeyFunction))
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewIPOCalendar(apiKeyTestValue)
		s := q.Encode()

		assert.Contains(t, s, "function=IPO_CALENDAR")
		assert.Contains(t, s, "apikey=demo")
	})
}
