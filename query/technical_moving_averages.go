package query

import (
	"net/url"
	"strconv"
	"time"
)

// SMA builds query parameters for the SMA API.
// Returns the simple moving average (SMA) values.
type SMA url.Values

// NewSMA creates a new SMA query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
//	seriesType: The desired price type (close, open, high, low)
func NewSMA(apikey, symbol, interval string, timePeriod int, seriesType SeriesTypeOption) SMA {
	return SMA{
		KeyFunction:   []string{FunctionSMA},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q SMA) MonthString(month string) SMA {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q SMA) Month(year int, month time.Month) SMA {
	return encodeMonth(q, year, month)
}

func (q SMA) DataType(o DataTypeOption) SMA { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q SMA) DataTypeCSV() SMA { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q SMA) DataTypeJSON() SMA { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q SMA) Validate() error {
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
func (q SMA) Encode() string { return encode(q) }

// EMA builds query parameters for the EMA API.
// Returns the exponential moving average (EMA) values.
type EMA url.Values

// NewEMA creates a new EMA query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
//	seriesType: The desired price type (close, open, high, low)
func NewEMA(apikey, symbol, interval string, timePeriod int, seriesType SeriesTypeOption) EMA {
	return EMA{
		KeyFunction:   []string{FunctionEMA},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q EMA) MonthString(month string) EMA {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q EMA) Month(year int, month time.Month) EMA {
	return encodeMonth(q, year, month)
}

func (q EMA) DataType(o DataTypeOption) EMA { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q EMA) DataTypeCSV() EMA { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q EMA) DataTypeJSON() EMA { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q EMA) Validate() error {
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
func (q EMA) Encode() string { return encode(q) }

// WMA builds query parameters for the WMA API.
// Returns the weighted moving average (WMA) values.
type WMA url.Values

// NewWMA creates a new WMA query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
//	seriesType: The desired price type (close, open, high, low)
func NewWMA(apikey, symbol, interval string, timePeriod int, seriesType SeriesTypeOption) WMA {
	return WMA{
		KeyFunction:   []string{FunctionWMA},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q WMA) MonthString(month string) WMA {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q WMA) Month(year int, month time.Month) WMA {
	return encodeMonth(q, year, month)
}

func (q WMA) DataType(o DataTypeOption) WMA { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q WMA) DataTypeCSV() WMA { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q WMA) DataTypeJSON() WMA { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q WMA) Validate() error {
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
func (q WMA) Encode() string { return encode(q) }

// DEMA builds query parameters for the DEMA API.
// Returns the double exponential moving average (DEMA) values.
type DEMA url.Values

// NewDEMA creates a new DEMA query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
//	seriesType: The desired price type (close, open, high, low)
func NewDEMA(apikey, symbol, interval string, timePeriod int, seriesType SeriesTypeOption) DEMA {
	return DEMA{
		KeyFunction:   []string{FunctionDEMA},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q DEMA) MonthString(month string) DEMA {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q DEMA) Month(year int, month time.Month) DEMA {
	return encodeMonth(q, year, month)
}

func (q DEMA) DataType(o DataTypeOption) DEMA { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q DEMA) DataTypeCSV() DEMA { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q DEMA) DataTypeJSON() DEMA { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q DEMA) Validate() error {
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
func (q DEMA) Encode() string { return encode(q) }

// TEMA builds query parameters for the TEMA API.
// Returns the triple exponential moving average (TEMA) values.
type TEMA url.Values

// NewTEMA creates a new TEMA query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
//	seriesType: The desired price type (close, open, high, low)
func NewTEMA(apikey, symbol, interval string, timePeriod int, seriesType SeriesTypeOption) TEMA {
	return TEMA{
		KeyFunction:   []string{FunctionTEMA},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q TEMA) MonthString(month string) TEMA {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q TEMA) Month(year int, month time.Month) TEMA {
	return encodeMonth(q, year, month)
}

func (q TEMA) DataType(o DataTypeOption) TEMA { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q TEMA) DataTypeCSV() TEMA { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q TEMA) DataTypeJSON() TEMA { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q TEMA) Validate() error {
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
func (q TEMA) Encode() string { return encode(q) }

// TRIMA builds query parameters for the TRIMA API.
// Returns the triangular moving average (TRIMA) values.
type TRIMA url.Values

// NewTRIMA creates a new TRIMA query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
//	seriesType: The desired price type (close, open, high, low)
func NewTRIMA(apikey, symbol, interval string, timePeriod int, seriesType SeriesTypeOption) TRIMA {
	return TRIMA{
		KeyFunction:   []string{FunctionTRIMA},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q TRIMA) MonthString(month string) TRIMA {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q TRIMA) Month(year int, month time.Month) TRIMA {
	return encodeMonth(q, year, month)
}

func (q TRIMA) DataType(o DataTypeOption) TRIMA { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q TRIMA) DataTypeCSV() TRIMA { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q TRIMA) DataTypeJSON() TRIMA { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q TRIMA) Validate() error {
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
func (q TRIMA) Encode() string { return encode(q) }

// KAMA builds query parameters for the KAMA API.
// Returns the Kaufman adaptive moving average (KAMA) values.
type KAMA url.Values

// NewKAMA creates a new KAMA query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
//	seriesType: The desired price type (close, open, high, low)
func NewKAMA(apikey, symbol, interval string, timePeriod int, seriesType SeriesTypeOption) KAMA {
	return KAMA{
		KeyFunction:   []string{FunctionKAMA},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q KAMA) MonthString(month string) KAMA {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q KAMA) Month(year int, month time.Month) KAMA {
	return encodeMonth(q, year, month)
}

func (q KAMA) DataType(o DataTypeOption) KAMA { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q KAMA) DataTypeCSV() KAMA { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q KAMA) DataTypeJSON() KAMA { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q KAMA) Validate() error {
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
func (q KAMA) Encode() string { return encode(q) }

// T3 builds query parameters for the T3 API.
// Returns the triple exponential moving average (T3) values.
type T3 url.Values

// NewT3 creates a new T3 query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
//	seriesType: The desired price type (close, open, high, low)
func NewT3(apikey, symbol, interval string, timePeriod int, seriesType SeriesTypeOption) T3 {
	return T3{
		KeyFunction:   []string{FunctionT3},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q T3) MonthString(month string) T3 {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q T3) Month(year int, month time.Month) T3 {
	return encodeMonth(q, year, month)
}

func (q T3) DataType(o DataTypeOption) T3 { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q T3) DataTypeCSV() T3 { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q T3) DataTypeJSON() T3 { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q T3) Validate() error {
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
func (q T3) Encode() string { return encode(q) }
