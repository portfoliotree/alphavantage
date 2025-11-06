package query

import (
	"net/url"
	"strconv"
	"time"
)

// BBANDS builds query parameters for the BBANDS API.
// Returns the Bollinger bands (BBANDS) values.
type BBANDS url.Values

// NewBBANDS creates a new BBANDS query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
//	seriesType: The desired price type (close, open, high, low)
func NewBBANDS(apikey, symbol, interval string, timePeriod int, seriesType SeriesTypeOption) BBANDS {
	return BBANDS{
		KeyFunction:   []string{FunctionBBANDS},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q BBANDS) MonthString(month string) BBANDS {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q BBANDS) Month(year int, month time.Month) BBANDS {
	return encodeMonth(q, year, month)
}

// NBDEVUP sets the standard deviation multiplier of the upper band.
func (q BBANDS) NBDEVUP(v int) BBANDS {
	q[KeyNbdevup] = []string{strconv.Itoa(v)}
	return q
}

// NBDEVDN sets the standard deviation multiplier of the lower band.
func (q BBANDS) NBDEVDN(v int) BBANDS {
	q[KeyNbdevdn] = []string{strconv.Itoa(v)}
	return q
}

// MAType sets the moving average type.
// Valid values: 0-8 (0=SMA, 1=EMA, 2=WMA, 3=DEMA, 4=TEMA, 5=TRIMA, 6=KAMA, 7=MAMA, 8=T3)
func (q BBANDS) MAType(t int) BBANDS {
	q[KeyMAType] = []string{strconv.Itoa(t)}
	return q
}

func (q BBANDS) DataType(o DataTypeOption) BBANDS { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q BBANDS) DataTypeCSV() BBANDS { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q BBANDS) DataTypeJSON() BBANDS { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q BBANDS) Validate() error {
	if err := validateRequired(q, KeyAPIKey, KeySymbol, KeyInterval, KeyTimePeriod, KeySeriesType); err != nil {
		return err
	}
	if seriesType, ok := q[KeySeriesType]; ok {
		if err := validateSeriesType(seriesType); err != nil {
			return err
		}
	}
	if t, ok := q[KeyMAType]; ok {
		if err := validateMAType(t, KeyMAType); err != nil {
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
func (q BBANDS) Encode() string { return encode(q) }

// TRANGE builds query parameters for the TRANGE API.
// Returns the true range (TRANGE) values.
type TRANGE url.Values

// NewTRANGE creates a new TRANGE query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
func NewTRANGE(apikey, symbol, interval string) TRANGE {
	return TRANGE{
		KeyFunction: []string{FunctionTRANGE},
		KeySymbol:   []string{symbol},
		KeyInterval: []string{interval},
		KeyAPIKey:   []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q TRANGE) MonthString(month string) TRANGE {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q TRANGE) Month(year int, month time.Month) TRANGE {
	return encodeMonth(q, year, month)
}

func (q TRANGE) DataType(o DataTypeOption) TRANGE { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q TRANGE) DataTypeCSV() TRANGE { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q TRANGE) DataTypeJSON() TRANGE { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q TRANGE) Validate() error {
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
func (q TRANGE) Encode() string { return encode(q) }

// ATR builds query parameters for the ATR API.
// Returns the average true range (ATR) values.
type ATR url.Values

// NewATR creates a new ATR query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
func NewATR(apikey, symbol, interval string, timePeriod int) ATR {
	return ATR{
		KeyFunction:   []string{FunctionATR},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q ATR) MonthString(month string) ATR {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q ATR) Month(year int, month time.Month) ATR {
	return encodeMonth(q, year, month)
}

func (q ATR) DataType(o DataTypeOption) ATR { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q ATR) DataTypeCSV() ATR { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q ATR) DataTypeJSON() ATR { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q ATR) Validate() error {
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
func (q ATR) Encode() string { return encode(q) }

// NATR builds query parameters for the NATR API.
// Returns the normalized average true range (NATR) values.
type NATR url.Values

// NewNATR creates a new NATR query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each value
func NewNATR(apikey, symbol, interval string, timePeriod int) NATR {
	return NATR{
		KeyFunction:   []string{FunctionNATR},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q NATR) MonthString(month string) NATR {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q NATR) Month(year int, month time.Month) NATR {
	return encodeMonth(q, year, month)
}

func (q NATR) DataType(o DataTypeOption) NATR { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q NATR) DataTypeCSV() NATR { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q NATR) DataTypeJSON() NATR { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q NATR) Validate() error {
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
func (q NATR) Encode() string { return encode(q) }

// SAR builds query parameters for the SAR API.
// Returns the parabolic SAR (SAR) values.
//
// Required: symbol, interval, apikey
type SAR url.Values

// NewSAR creates a new SAR query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
func NewSAR(apikey, symbol, interval string) SAR {
	return SAR{
		KeyFunction: []string{FunctionSAR},
		KeySymbol:   []string{symbol},
		KeyInterval: []string{interval},
		KeyAPIKey:   []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q SAR) MonthString(month string) SAR {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q SAR) Month(year int, month time.Month) SAR {
	return encodeMonth(q, year, month)
}

// Acceleration sets the acceleration factor.
// Positive floats are accepted. By default, 0.01.
func (q SAR) Acceleration(v float64) SAR {
	q.AccelerationString(strconv.FormatFloat(v, 'f', -1, 64))
	return q
}

func (q SAR) AccelerationString(v string) SAR {
	q[KeyAcceleration] = []string{v}
	return q
}

// Maximum sets the acceleration factor maximum value.
// Positive floats are accepted. By default, 0.2.
func (q SAR) Maximum(v float64) SAR {
	q.MaximumString(strconv.FormatFloat(v, 'f', -1, 64))
	return q
}

func (q SAR) MaximumString(v string) SAR {
	q[KeyMaximum] = []string{v}
	return q
}

func (q SAR) DataType(o DataTypeOption) SAR { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q SAR) DataTypeCSV() SAR { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q SAR) DataTypeJSON() SAR { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q SAR) Validate() error {
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
func (q SAR) Encode() string { return encode(q) }

// ULTOSC builds query parameters for the ULTOSC API.
// Returns the ultimate oscillator (ULTOSC) values.
type ULTOSC url.Values

// NewULTOSC creates a new ULTOSC query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
func NewULTOSC(apikey, symbol, interval string) ULTOSC {
	return ULTOSC{
		KeyFunction: []string{FunctionULTOSC},
		KeySymbol:   []string{symbol},
		KeyInterval: []string{interval},
		KeyAPIKey:   []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q ULTOSC) MonthString(month string) ULTOSC {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q ULTOSC) Month(year int, month time.Month) ULTOSC {
	return encodeMonth(q, year, month)
}

// TimePeriod1 sets the first time period for the indicator.
func (q ULTOSC) TimePeriod1(p int) ULTOSC {
	q[KeyTimePeriod1] = []string{strconv.Itoa(p)}
	return q
}

// TimePeriod2 sets the second time period for the indicator.
func (q ULTOSC) TimePeriod2(p int) ULTOSC {
	q[KeyTimePeriod2] = []string{strconv.Itoa(p)}
	return q
}

// TimePeriod3 sets the third time period for the indicator.
func (q ULTOSC) TimePeriod3(p int) ULTOSC {
	q[KeyTimePeriod3] = []string{strconv.Itoa(p)}
	return q
}

func (q ULTOSC) DataType(o DataTypeOption) ULTOSC { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q ULTOSC) DataTypeCSV() ULTOSC { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q ULTOSC) DataTypeJSON() ULTOSC { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q ULTOSC) Validate() error {
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
func (q ULTOSC) Encode() string { return encode(q) }
