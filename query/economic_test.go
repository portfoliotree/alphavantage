package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/portfoliotree/alphavantage/query"
)

func TestRealGDP(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewRealGDP(apiKeyTestValue)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "REAL_GDP", v.Get(query.KeyFunction))
	})

	t.Run("with interval and dataType", func(t *testing.T) {
		q := query.NewRealGDP(apiKeyTestValue).Interval(query.QAIntervalOptionQuarterly).DataTypeCSV()
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "quarterly", v.Get(query.KeyInterval))
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("invalid interval", func(t *testing.T) {
		q := query.NewRealGDP(apiKeyTestValue).Interval("monthly")
		err := q.Validate()
		assert.Error(t, err)
	})
}

func TestRealGDPPerCapita(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewRealGDPPerCapita(apiKeyTestValue)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "REAL_GDP_PER_CAPITA", v.Get(query.KeyFunction))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewRealGDPPerCapita(apiKeyTestValue).DataTypeJSON()
		err := q.Validate()
		assert.NoError(t, err)
	})
}

func TestTreasuryYield(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewTreasuryYield(apiKeyTestValue)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "TREASURY_YIELD", v.Get(query.KeyFunction))
	})

	t.Run("with all options", func(t *testing.T) {
		q := query.NewTreasuryYield(apiKeyTestValue).
			Interval(query.DWMIntervalOptionMonthly).
			Maturity(query.MaturityOption10Year).
			DataTypeCSV()

		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "monthly", v.Get(query.KeyInterval))
		assert.Equal(t, "10year", v.Get(query.KeyMaturity))
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("invalid maturity", func(t *testing.T) {
		q := query.NewTreasuryYield(apiKeyTestValue)
		q[query.KeyMaturity] = []string{"1year"} // Directly inject invalid value to test validator
		err := q.Validate()
		assert.Error(t, err)
	})
}

func TestFederalFundsRate(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewFederalFundsRate(apiKeyTestValue)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "FEDERAL_FUNDS_RATE", v.Get(query.KeyFunction))
	})

	t.Run("with interval", func(t *testing.T) {
		q := query.NewFederalFundsRate(apiKeyTestValue).Interval(query.DWMIntervalOptionWeekly)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "weekly", v.Get(query.KeyInterval))
	})

	t.Run("invalid interval", func(t *testing.T) {
		q := query.NewFederalFundsRate(apiKeyTestValue).Interval("annual")
		err := q.Validate()
		assert.Error(t, err)
	})
}

func TestCPI(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewCPI(apiKeyTestValue)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "CPI", v.Get(query.KeyFunction))
	})

	t.Run("with monthly interval", func(t *testing.T) {
		q := query.NewCPI(apiKeyTestValue).Interval("monthly")
		err := q.Validate()
		assert.NoError(t, err)
	})

	t.Run("with semiannual interval", func(t *testing.T) {
		q := query.NewCPI(apiKeyTestValue).Interval("semiannual")
		err := q.Validate()
		assert.NoError(t, err)
	})

	t.Run("invalid interval", func(t *testing.T) {
		q := query.NewCPI(apiKeyTestValue).Interval("quarterly")
		err := q.Validate()
		assert.Error(t, err)
	})
}

func TestInflation(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewInflation(apiKeyTestValue)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "INFLATION", v.Get(query.KeyFunction))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewInflation(apiKeyTestValue).DataTypeCSV()
		err := q.Validate()
		assert.NoError(t, err)
	})
}

func TestRetailSales(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewRetailSales(apiKeyTestValue)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "RETAIL_SALES", v.Get(query.KeyFunction))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewRetailSales(apiKeyTestValue).DataTypeJSON()
		err := q.Validate()
		assert.NoError(t, err)
	})
}

func TestDurables(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewDurables(apiKeyTestValue)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "DURABLES", v.Get(query.KeyFunction))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewDurables(apiKeyTestValue).DataTypeCSV()
		err := q.Validate()
		assert.NoError(t, err)
	})
}

func TestUnemployment(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewUnemployment(apiKeyTestValue)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "UNEMPLOYMENT", v.Get(query.KeyFunction))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewUnemployment(apiKeyTestValue).DataTypeJSON()
		err := q.Validate()
		assert.NoError(t, err)
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewUnemployment(apiKeyTestValue)
		s := q.Encode()

		assert.Contains(t, s, "function=UNEMPLOYMENT")
		assert.Contains(t, s, "apikey=demo")
	})
}

func TestNonfarmPayroll(t *testing.T) {
	t.Run("valid basic query", func(t *testing.T) {
		q := query.NewNonfarmPayroll(apiKeyTestValue)
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "NONFARM_PAYROLL", v.Get(query.KeyFunction))
	})

	t.Run("with dataType", func(t *testing.T) {
		q := query.NewNonfarmPayroll(apiKeyTestValue).DataTypeCSV()
		err := q.Validate()
		assert.NoError(t, err)

		v := q.Values()
		assert.Equal(t, "csv", v.Get(query.KeyDataType))
	})

	t.Run("String() method", func(t *testing.T) {
		q := query.NewNonfarmPayroll(apiKeyTestValue)
		s := q.Encode()

		assert.Contains(t, s, "function=NONFARM_PAYROLL")
		assert.Contains(t, s, "apikey=demo")
	})
}
