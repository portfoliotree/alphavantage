package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/portfoliotree/alphavantage/query"
)

func TestWTI(t *testing.T) {
	q := query.NewWTI(apiKeyTestValue).Interval(query.DWMIntervalOptionMonthly)
	assert.NoError(t, q.Validate())
	assert.Equal(t, query.FunctionWTI, q.Values().Get(query.KeyFunction))
}

func TestBrent(t *testing.T) {
	q := query.NewBrent(apiKeyTestValue).Interval(query.DWMIntervalOptionWeekly)
	assert.NoError(t, q.Validate())
	assert.Equal(t, query.FunctionBrent, q.Values().Get(query.KeyFunction))
}

func TestNaturalGas(t *testing.T) {
	q := query.NewNaturalGas(apiKeyTestValue).Interval(query.DWMIntervalOptionDaily)
	assert.NoError(t, q.Validate())
	assert.Equal(t, query.FunctionNaturalGas, q.Values().Get(query.KeyFunction))
}

func TestCopper(t *testing.T) {
	q := query.NewCopper(apiKeyTestValue).Interval(query.MQAIntervalOptionQuarterly)
	assert.NoError(t, q.Validate())
	assert.Equal(t, query.FunctionCopper, q.Values().Get(query.KeyFunction))
}

func TestAluminum(t *testing.T) {
	q := query.NewAluminum(apiKeyTestValue).Interval(query.MQAIntervalOptionAnnual)
	assert.NoError(t, q.Validate())
	assert.Equal(t, query.FunctionAluminum, q.Values().Get(query.KeyFunction))
}

func TestWheat(t *testing.T) {
	q := query.NewWheat(apiKeyTestValue).Interval(query.MQAIntervalOptionMonthly)
	assert.NoError(t, q.Validate())
	assert.Equal(t, query.FunctionWheat, q.Values().Get(query.KeyFunction))
}

func TestCorn(t *testing.T) {
	q := query.NewCorn(apiKeyTestValue).Interval(query.MQAIntervalOptionQuarterly)
	assert.NoError(t, q.Validate())
	assert.Equal(t, query.FunctionCorn, q.Values().Get(query.KeyFunction))
}

func TestCotton(t *testing.T) {
	q := query.NewCotton(apiKeyTestValue).Interval(query.MQAIntervalOptionAnnual)
	assert.NoError(t, q.Validate())
	assert.Equal(t, query.FunctionCotton, q.Values().Get(query.KeyFunction))
}

func TestSugar(t *testing.T) {
	q := query.NewSugar(apiKeyTestValue).Interval(query.MQAIntervalOptionMonthly)
	assert.NoError(t, q.Validate())
	assert.Equal(t, query.FunctionSugar, q.Values().Get(query.KeyFunction))
}

func TestCoffee(t *testing.T) {
	q := query.NewCoffee(apiKeyTestValue).Interval(query.MQAIntervalOptionQuarterly)
	assert.NoError(t, q.Validate())
	assert.Equal(t, query.FunctionCoffee, q.Values().Get(query.KeyFunction))
}

func TestAllCommodities(t *testing.T) {
	q := query.NewAllCommodities(apiKeyTestValue).Interval(query.MQAIntervalOptionAnnual)
	assert.NoError(t, q.Validate())
	assert.Equal(t, query.FunctionAllCommodities, q.Values().Get(query.KeyFunction))
}

// Test invalid intervals
func TestCommoditiesInvalidInterval(t *testing.T) {
	t.Run("WTI invalid interval", func(t *testing.T) {
		q := query.NewWTI(apiKeyTestValue).Interval("banana")
		assert.Error(t, q.Validate())
	})

	t.Run("Copper invalid interval", func(t *testing.T) {
		q := query.NewCopper(apiKeyTestValue).Interval("daily")
		assert.Error(t, q.Validate())
	})
}
