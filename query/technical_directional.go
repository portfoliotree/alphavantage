package query

import (
	"net/url"
	"strconv"
	"time"
)

// DX builds query parameters for the DX API.
// Returns the directional movement index (DX) values.
type DX url.Values

// NewDX creates a new DX query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
func NewDX(apikey, symbol, interval string, timePeriod int) DX {
	return DX{
		KeyFunction:   []string{FunctionDX},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q DX) MonthString(month string) DX {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q DX) Month(year int, month time.Month) DX {
	return encodeMonth(q, year, month)
}

func (q DX) DataType(o DataTypeOption) DX { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q DX) DataTypeCSV() DX { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q DX) DataTypeJSON() DX { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q DX) Validate() error {
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
func (q DX) Encode() string { return encode(q) }

// ADX builds query parameters for the ADX API.
// Returns the average directional movement index (ADX) values.
type ADX url.Values

// NewADX creates a new ADX query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
func NewADX(apikey, symbol, interval string, timePeriod int) ADX {
	return ADX{
		KeyFunction:   []string{FunctionADX},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q ADX) MonthString(month string) ADX {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q ADX) Month(year int, month time.Month) ADX {
	return encodeMonth(q, year, month)
}

func (q ADX) DataType(o DataTypeOption) ADX { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q ADX) DataTypeCSV() ADX { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q ADX) DataTypeJSON() ADX { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q ADX) Validate() error {
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
func (q ADX) Encode() string { return encode(q) }

// ADXR builds query parameters for the ADXR API.
// Returns the average directional movement index rating (ADXR) values.
type ADXR url.Values

// NewADXR creates a new ADXR query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
func NewADXR(apikey, symbol, interval string, timePeriod int) ADXR {
	return ADXR{
		KeyFunction:   []string{FunctionADXR},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q ADXR) MonthString(month string) ADXR {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q ADXR) Month(year int, month time.Month) ADXR {
	return encodeMonth(q, year, month)
}

func (q ADXR) DataType(o DataTypeOption) ADXR { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q ADXR) DataTypeCSV() ADXR { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q ADXR) DataTypeJSON() ADXR { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q ADXR) Validate() error {
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
func (q ADXR) Encode() string { return encode(q) }

// MINUSDI builds query parameters for the MINUS_DI API.
// Returns the minus directional indicator (MINUS_DI) values.
type MINUSDI url.Values

// NewMINUSDI creates a new MINUS_DI query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
func NewMINUSDI(apikey, symbol, interval string, timePeriod int) MINUSDI {
	return MINUSDI{
		KeyFunction:   []string{FunctionMINUSDI},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q MINUSDI) MonthString(month string) MINUSDI {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q MINUSDI) Month(year int, month time.Month) MINUSDI {
	return encodeMonth(q, year, month)
}

func (q MINUSDI) DataType(o DataTypeOption) MINUSDI { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q MINUSDI) DataTypeCSV() MINUSDI { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q MINUSDI) DataTypeJSON() MINUSDI { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q MINUSDI) Validate() error {
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
func (q MINUSDI) Encode() string { return encode(q) }

// PLUSDI builds query parameters for the PLUS_DI API.
// Returns the plus directional indicator (PLUS_DI) values.
type PLUSDI url.Values

// NewPLUSDI creates a new PLUS_DI query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
func NewPLUSDI(apikey, symbol, interval string, timePeriod int) PLUSDI {
	return PLUSDI{
		KeyFunction:   []string{FunctionPLUSDI},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q PLUSDI) MonthString(month string) PLUSDI {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q PLUSDI) Month(year int, month time.Month) PLUSDI {
	return encodeMonth(q, year, month)
}

func (q PLUSDI) DataType(o DataTypeOption) PLUSDI { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q PLUSDI) DataTypeCSV() PLUSDI { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q PLUSDI) DataTypeJSON() PLUSDI { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q PLUSDI) Validate() error {
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
func (q PLUSDI) Encode() string { return encode(q) }

// MINUSDM builds query parameters for the MINUS_DM API.
// Returns the minus directional movement (MINUS_DM) values.
type MINUSDM url.Values

// NewMINUSDM creates a new MINUS_DM query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
func NewMINUSDM(apikey, symbol, interval string, timePeriod int) MINUSDM {
	return MINUSDM{
		KeyFunction:   []string{FunctionMINUSDM},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q MINUSDM) MonthString(month string) MINUSDM {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q MINUSDM) Month(year int, month time.Month) MINUSDM {
	return encodeMonth(q, year, month)
}

func (q MINUSDM) DataType(o DataTypeOption) MINUSDM { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q MINUSDM) DataTypeCSV() MINUSDM { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q MINUSDM) DataTypeJSON() MINUSDM { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q MINUSDM) Validate() error {
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
func (q MINUSDM) Encode() string { return encode(q) }

// PLUSDM builds query parameters for the PLUS_DM API.
// Returns the plus directional movement (PLUS_DM) values.
type PLUSDM url.Values

// NewPLUSDM creates a new PLUS_DM query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
func NewPLUSDM(apikey, symbol, interval string, timePeriod int) PLUSDM {
	return PLUSDM{
		KeyFunction:   []string{FunctionPLUSDM},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q PLUSDM) MonthString(month string) PLUSDM {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q PLUSDM) Month(year int, month time.Month) PLUSDM {
	return encodeMonth(q, year, month)
}

func (q PLUSDM) DataType(o DataTypeOption) PLUSDM { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q PLUSDM) DataTypeCSV() PLUSDM { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q PLUSDM) DataTypeJSON() PLUSDM { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q PLUSDM) Validate() error {
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
func (q PLUSDM) Encode() string { return encode(q) }

// AROON builds query parameters for the AROON API.
// Returns the Aroon (AROON) values.
type AROON url.Values

// NewAROON creates a new AROON query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
func NewAROON(apikey, symbol, interval string, timePeriod int) AROON {
	return AROON{
		KeyFunction:   []string{FunctionAROON},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q AROON) MonthString(month string) AROON {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q AROON) Month(year int, month time.Month) AROON {
	return encodeMonth(q, year, month)
}

func (q AROON) DataType(o DataTypeOption) AROON { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q AROON) DataTypeCSV() AROON { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q AROON) DataTypeJSON() AROON { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q AROON) Validate() error {
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
func (q AROON) Encode() string { return encode(q) }

// AROONOSC builds query parameters for the AROONOSC API.
// Returns the Aroon oscillator (AROONOSC) values.
type AROONOSC url.Values

// NewAROONOSC creates a new AROONOSC query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
func NewAROONOSC(apikey, symbol, interval string, timePeriod int) AROONOSC {
	return AROONOSC{
		KeyFunction:   []string{FunctionAROONOSC},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q AROONOSC) MonthString(month string) AROONOSC {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q AROONOSC) Month(year int, month time.Month) AROONOSC {
	return encodeMonth(q, year, month)
}

func (q AROONOSC) DataType(o DataTypeOption) AROONOSC { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q AROONOSC) DataTypeCSV() AROONOSC { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q AROONOSC) DataTypeJSON() AROONOSC { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q AROONOSC) Validate() error {
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
func (q AROONOSC) Encode() string { return encode(q) }
