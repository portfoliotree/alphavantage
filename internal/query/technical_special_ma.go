package query

import (
	"net/url"
	"time"
)

// MAMA builds query parameters for the MAMA API.
// Returns the MESA adaptive moving average (MAMA) values.
type MAMA url.Values

// NewMAMA creates a new MAMA query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	seriesType: The desired price type (close, open, high, low)
func NewMAMA(apikey, symbol, interval string, seriesType SeriesTypeOption) MAMA {
	return MAMA{
		KeyFunction:   []string{FunctionMAMA},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q MAMA) MonthString(month string) MAMA {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q MAMA) Month(year int, month time.Month) MAMA {
	return encodeMonth(q, year, month)
}

// FastLimit sets the fast limit parameter.
// Positive floats are accepted. By default, fastlimit=0.01.
func (q MAMA) FastLimit(v string) MAMA {
	q[KeyFastLimit] = []string{v}
	return q
}

// SlowLimit sets the slow limit parameter.
// Positive floats are accepted. By default, slowlimit=0.01.
func (q MAMA) SlowLimit(v string) MAMA {
	q[KeySlowLimit] = []string{v}
	return q
}

func (q MAMA) DataType(o DataTypeOption) MAMA { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q MAMA) DataTypeCSV() MAMA { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q MAMA) DataTypeJSON() MAMA { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q MAMA) Validate() error {
	if err := validateRequired(q, KeyAPIKey, KeySymbol, KeyInterval, KeySeriesType); err != nil {
		return err
	}
	if seriesType, ok := q[KeySeriesType]; ok {
		if err := validateSeriesType(seriesType); err != nil {
			return err
		}
	}
	if dt, ok := q[KeyDataType]; ok {
		if err := validateDatatype(dt); err != nil {
			return err
		}
	}
	return nil
}

// Encode returns the URL-encoded query string.
func (q MAMA) Encode() string { return encode(q) }

// VWAP builds query parameters for the VWAP API.
// Returns the volume weighted average price (VWAP) for intraday time series.
type VWAP url.Values

// NewVWAP creates a new VWAP query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
func NewVWAP(apikey, symbol, interval string) VWAP {
	return VWAP{
		KeyFunction: []string{FunctionVWAP},
		KeySymbol:   []string{symbol},
		KeyInterval: []string{interval},
		KeyAPIKey:   []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q VWAP) MonthString(month string) VWAP {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q VWAP) Month(year int, month time.Month) VWAP {
	return encodeMonth(q, year, month)
}

func (q VWAP) DataType(o DataTypeOption) VWAP { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q VWAP) DataTypeCSV() VWAP { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q VWAP) DataTypeJSON() VWAP { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q VWAP) Validate() error {
	if err := validateRequired(q, KeyAPIKey, KeySymbol, KeyInterval); err != nil {
		return err
	}
	if dt, ok := q[KeyDataType]; ok {
		if err := validateDatatype(dt); err != nil {
			return err
		}
	}
	return nil
}

// Encode returns the URL-encoded query string.
func (q VWAP) Encode() string { return encode(q) }
