package query

import (
	"fmt"
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateRequired(t *testing.T) {
	t.Run("all fields present", func(t *testing.T) {
		v := url.Values{
			"field1": []string{"value1"},
			"field2": []string{"value2"},
			"field3": []string{"value3"},
		}
		err := validateRequired(v, "field1", "field2", "field3")
		assert.NoError(t, err)
	})

	t.Run("missing field", func(t *testing.T) {
		v := url.Values{
			"field1": []string{"value1"},
			"field3": []string{"value3"},
		}
		err := validateRequired(v, "field1", "field2", "field3")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field2")
	})

	t.Run("empty field", func(t *testing.T) {
		v := url.Values{
			"field1": []string{"value1"},
			"field2": []string{""},
			"field3": []string{"value3"},
		}
		err := validateRequired(v, "field1", "field2", "field3")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field2")
	})
}

func TestValidateEnum(t *testing.T) {
	allowed := []string{"option1", "option2", "option3"}

	t.Run("valid value", func(t *testing.T) {
		err := validateEnum([]string{"option2"}, "field", allowed)
		assert.NoError(t, err)
	})

	t.Run("invalid value", func(t *testing.T) {
		err := validateEnum([]string{"invalid"}, "field", allowed)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field")
		assert.Contains(t, err.Error(), "invalid")
	})

	t.Run("empty value", func(t *testing.T) {
		err := validateEnum([]string{""}, "field", allowed)
		assert.NoError(t, err) // Empty is allowed for optional fields
	})

	t.Run("nil value", func(t *testing.T) {
		err := validateEnum(nil, "field", allowed)
		assert.NoError(t, err) // Nil is allowed for optional fields
	})
}

func TestValidateBoolean(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		err := validateBoolean([]string{"true"}, "field")
		assert.NoError(t, err)
	})

	t.Run("false", func(t *testing.T) {
		err := validateBoolean([]string{"false"}, "field")
		assert.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		err := validateBoolean([]string{"yes"}, "field")
		assert.Error(t, err)
	})

	t.Run("empty", func(t *testing.T) {
		err := validateBoolean([]string{""}, "field")
		assert.NoError(t, err)
	})
}

func TestValidateDatatype(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		err := validateDatatype([]string{"json"})
		assert.NoError(t, err)
	})

	t.Run("csv", func(t *testing.T) {
		err := validateDatatype([]string{"csv"})
		assert.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		err := validateDatatype([]string{"xml"})
		assert.Error(t, err)
	})
}

func TestValidateOutputSize(t *testing.T) {
	t.Run("compact", func(t *testing.T) {
		err := validateOutputSize([]string{"compact"})
		assert.NoError(t, err)
	})

	t.Run("full", func(t *testing.T) {
		err := validateOutputSize([]string{"full"})
		assert.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		err := validateOutputSize([]string{"partial"})
		assert.Error(t, err)
	})
}

func TestValidateSeriesType(t *testing.T) {
	validTypes := []string{"close", "open", "high", "low"}

	for _, seriesType := range validTypes {
		t.Run(seriesType, func(t *testing.T) {
			err := validateSeriesType([]string{seriesType})
			assert.NoError(t, err)
		})
	}

	t.Run("invalid", func(t *testing.T) {
		err := validateSeriesType([]string{"volume"})
		assert.Error(t, err)
	})
}

func TestValidateState(t *testing.T) {
	t.Run("active", func(t *testing.T) {
		err := validateState([]string{"active"})
		assert.NoError(t, err)
	})

	t.Run("delisted", func(t *testing.T) {
		err := validateState([]string{"delisted"})
		assert.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		err := validateState([]string{"pending"})
		assert.Error(t, err)
	})
}

func TestValidateHorizon(t *testing.T) {
	t.Run("3month", func(t *testing.T) {
		err := validateHorizon([]string{"3month"})
		assert.NoError(t, err)
	})

	t.Run("6month", func(t *testing.T) {
		err := validateHorizon([]string{"6month"})
		assert.NoError(t, err)
	})

	t.Run("12month", func(t *testing.T) {
		err := validateHorizon([]string{"12month"})
		assert.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		err := validateHorizon([]string{"9month"})
		assert.Error(t, err)
	})
}

func TestValidateSort(t *testing.T) {
	t.Run("LATEST", func(t *testing.T) {
		err := validateSort([]string{"LATEST"})
		assert.NoError(t, err)
	})

	t.Run("EARLIEST", func(t *testing.T) {
		err := validateSort([]string{"EARLIEST"})
		assert.NoError(t, err)
	})

	t.Run("RELEVANCE", func(t *testing.T) {
		err := validateSort([]string{"RELEVANCE"})
		assert.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		err := validateSort([]string{"RANDOM"})
		assert.Error(t, err)
	})
}

func TestValidateMaturity(t *testing.T) {
	validMaturities := []string{"3month", "2year", "5year", "7year", "10year", "30year"}

	for _, maturity := range validMaturities {
		t.Run(maturity, func(t *testing.T) {
			err := validateMaturity([]string{maturity})
			assert.NoError(t, err)
		})
	}

	t.Run("invalid", func(t *testing.T) {
		err := validateMaturity([]string{"1year"})
		assert.Error(t, err)
	})
}

func TestValidateMAType(t *testing.T) {
	for i := 0; i <= 8; i++ {
		t.Run(fmt.Sprintf("type_%d", i), func(t *testing.T) {
			err := validateMAType([]string{strconv.Itoa(i)}, "matype")
			assert.NoError(t, err)
		})
	}

	t.Run("invalid", func(t *testing.T) {
		err := validateMAType([]string{"9"}, "matype")
		assert.Error(t, err)
	})

	t.Run("negative", func(t *testing.T) {
		err := validateMAType([]string{"-1"}, "matype")
		assert.Error(t, err)
	})
}
