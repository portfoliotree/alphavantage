package query

import (
	"net/url"
	"strconv"
	"time"
)

// Midpoint builds query parameters for the Midpoint API.
// Returns the midpoint (Midpoint) values. Midpoint = (highest value + lowest value)/2.
type Midpoint url.Values

// NewMidpoint creates a new Midpoint query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
//	seriesType: The desired price type (close, open, high, low)
func NewMidpoint(apikey, symbol, interval string, timePeriod int, seriesType SeriesTypeOption) Midpoint {
	return Midpoint{
		KeyFunction:   []string{FunctionMidpoint},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q Midpoint) MonthString(month string) Midpoint {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q Midpoint) Month(year int, month time.Month) Midpoint {
	return encodeMonth(q, year, month)
}

func (q Midpoint) DataType(o DataTypeOption) Midpoint { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q Midpoint) DataTypeCSV() Midpoint { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q Midpoint) DataTypeJSON() Midpoint { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q Midpoint) Validate() error {
	if err := validateRequired(q, KeyAPIKey, KeySymbol, KeyInterval, KeyTimePeriod, KeySeriesType); err != nil {
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
func (q Midpoint) Encode() string { return encode(q) }

// MIDPRICE builds query parameters for the MIDPRICE API.
// Returns the midpoint price (MIDPRICE) values. MIDPRICE = (highest high + lowest low)/2.
type MIDPRICE url.Values

// NewMIDPRICE creates a new MIDPRICE query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
func NewMIDPRICE(apikey, symbol, interval string, timePeriod int) MIDPRICE {
	return MIDPRICE{
		KeyFunction:   []string{FunctionMIDPRICE},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q MIDPRICE) MonthString(month string) MIDPRICE {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q MIDPRICE) Month(year int, month time.Month) MIDPRICE {
	return encodeMonth(q, year, month)
}

func (q MIDPRICE) DataType(o DataTypeOption) MIDPRICE { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q MIDPRICE) DataTypeCSV() MIDPRICE { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q MIDPRICE) DataTypeJSON() MIDPRICE { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q MIDPRICE) Validate() error {
	if err := validateRequired(q, KeyAPIKey, KeySymbol, KeyInterval, KeyTimePeriod); err != nil {
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
func (q MIDPRICE) Encode() string { return encode(q) }
