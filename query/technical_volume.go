package query

import (
	"net/url"
	"strconv"
	"time"
)

// MFI builds query parameters for the MFI API.
// Returns the money flow index (MFI) values.
type MFI url.Values

// NewMFI creates a new MFI query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
func NewMFI(apikey, symbol, interval string, timePeriod int) MFI {
	return MFI{
		KeyFunction:   []string{FunctionMFI},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q MFI) MonthString(month string) MFI {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q MFI) Month(year int, month time.Month) MFI {
	return encodeMonth(q, year, month)
}

func (q MFI) DataType(o DataTypeOption) MFI { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q MFI) DataTypeCSV() MFI { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q MFI) DataTypeJSON() MFI { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q MFI) Validate() error {
	if err := validateRequired(q, KeyAPIKey, KeySymbol, KeyInterval, KeyTimePeriod); err != nil {
		return err
	}
	if err := validateEnum(q[KeyInterval], KeyInterval, IDWMIntervalOptions()); err != nil {
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
func (q MFI) Encode() string { return encode(q) }

// TRIX builds query parameters for the TRIX API.
// Returns the 1-day rate of change of a triple smooth exponential moving average (TRIX) values.
type TRIX url.Values

// NewTRIX creates a new TRIX query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
//	seriesType: The desired price type (close, open, high, low)
func NewTRIX(apikey, symbol, interval string, timePeriod int, seriesType SeriesTypeOption) TRIX {
	return TRIX{
		KeyFunction:   []string{FunctionTRIX},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q TRIX) MonthString(month string) TRIX {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q TRIX) Month(year int, month time.Month) TRIX {
	return encodeMonth(q, year, month)
}

func (q TRIX) DataType(o DataTypeOption) TRIX { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q TRIX) DataTypeCSV() TRIX { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q TRIX) DataTypeJSON() TRIX { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q TRIX) Validate() error {
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
func (q TRIX) Encode() string { return encode(q) }

// AD builds query parameters for the AD API.
// Returns the Chaikin A/D line (AD) values.
type AD url.Values

// NewAD creates a new AD query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
func NewAD(apikey, symbol, interval string) AD {
	return AD{
		KeyFunction: []string{FunctionAD},
		KeySymbol:   []string{symbol},
		KeyInterval: []string{interval},
		KeyAPIKey:   []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q AD) MonthString(month string) AD {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q AD) Month(year int, month time.Month) AD {
	return encodeMonth(q, year, month)
}

func (q AD) DataType(o DataTypeOption) AD { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q AD) DataTypeCSV() AD { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q AD) DataTypeJSON() AD { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q AD) Validate() error {
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
func (q AD) Encode() string { return encode(q) }

// ADOSC builds query parameters for the ADOSC API.
// Returns the Chaikin A/D oscillator (ADOSC) values.
type ADOSC url.Values

// NewADOSC creates a new ADOSC query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
func NewADOSC(apikey, symbol, interval string) ADOSC {
	return ADOSC{
		KeyFunction: []string{FunctionADOSC},
		KeySymbol:   []string{symbol},
		KeyInterval: []string{interval},
		KeyAPIKey:   []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q ADOSC) MonthString(month string) ADOSC {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q ADOSC) Month(year int, month time.Month) ADOSC {
	return encodeMonth(q, year, month)
}

// FastPeriod sets the time period of the fast EMA.
func (q ADOSC) FastPeriod(p int) ADOSC {
	q[KeyFastPeriod] = []string{strconv.Itoa(p)}
	return q
}

// SlowPeriod sets the time period of the slow EMA.
func (q ADOSC) SlowPeriod(p int) ADOSC {
	q[KeySlowPeriod] = []string{strconv.Itoa(p)}
	return q
}

func (q ADOSC) DataType(o DataTypeOption) ADOSC { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q ADOSC) DataTypeCSV() ADOSC { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q ADOSC) DataTypeJSON() ADOSC { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q ADOSC) Validate() error {
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
func (q ADOSC) Encode() string { return encode(q) }

// OBV builds query parameters for the OBV API.
// Returns the on balance volume (OBV) values.
type OBV url.Values

// NewOBV creates a new OBV query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
func NewOBV(apikey, symbol, interval string) OBV {
	return OBV{
		KeyFunction: []string{FunctionOBV},
		KeySymbol:   []string{symbol},
		KeyInterval: []string{interval},
		KeyAPIKey:   []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q OBV) MonthString(month string) OBV {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q OBV) Month(year int, month time.Month) OBV {
	return encodeMonth(q, year, month)
}

func (q OBV) DataType(o DataTypeOption) OBV { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q OBV) DataTypeCSV() OBV { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q OBV) DataTypeJSON() OBV { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q OBV) Validate() error {
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
func (q OBV) Encode() string { return encode(q) }
