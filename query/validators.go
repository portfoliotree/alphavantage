package query

import (
	"fmt"
	"net/url"
)

type ErrorMissing string

func (e ErrorMissing) Error() string {
	return fmt.Sprintf("required field %q is missing or empty", string(e))
}

func missingKeyError(field string) error {
	return ErrorMissing(field)
}

// validateRequired checks that all required fields are present and non-empty.
func validateRequired[Values ~urlValues](v Values, fields ...string) error {
	for _, field := range fields {
		if !url.Values(v).Has(field) || len(v[field]) == 0 || v[field][0] == "" {
			return missingKeyError(field)
		}
	}
	return nil
}

// validateEnum checks if value is one of the allowed options.
// Returns nil if the field is empty (for optional fields).
func validateEnum[T ~string](values []string, fieldName string, allowed []T) error {
	if len(values) == 0 || values[0] == "" {
		return nil
	}
	val := values[0]
	for _, a := range allowed {
		if val == string(a) {
			return nil
		}
	}
	return fmt.Errorf("%s must be one of %v, got '%s'", fieldName, allowed, val)
}

// validateBoolean checks if value is "true" or "false".
func validateBoolean(values []string, fieldName string) error {
	return validateEnum(values, fieldName, []string{booleanValueTrue, booleanValueFalse})
}

// validateDatatype checks if dataType is "json" or "csv".
func validateDatatype(values []string) error {
	return validateEnum(values, KeyDataType, DataTypeOptions())
}

// validateOutputSize checks if outputSize is "compact" or "full".
func validateOutputSize(values []string) error {
	return validateEnum(values, KeyOutputSize, OutputSizeOptions())
}

// validateSeriesType checks if series_type is valid.
func validateSeriesType(values []string) error {
	return validateEnum(values, KeySeriesType, SeriesTypeOptions())
}

// validateOHLC checks if OHLC is valid.
func validateOHLC(values []string) error {
	return validateEnum(values, KeyOHLC, SeriesTypeOptions())
}

// validateState checks if state is valid for listing status.
func validateState(values []string) error {
	return validateEnum(values, KeyState, StateOptions())
}

// validateHorizon checks if horizon is valid for earnings calendar.
func validateHorizon(values []string) error {
	return validateEnum(values, KeyHorizon, HorizonOptions())
}

// validateSort checks if sort is valid.
func validateSort(values []string) error {
	return validateEnum(values, KeySort, SortOptions())
}

// validateMaturity checks if maturity is valid for treasury yield.
func validateMaturity(values []string) error {
	return validateEnum(values, KeyMaturity, MaturityOptions())
}

// validateMAType checks if moving average type is valid (0-8).
func validateMAType(values []string, fieldName string) error {
	return validateEnum(values, fieldName, []string{"0", "1", "2", "3", "4", "5", "6", "7", "8"})
}
