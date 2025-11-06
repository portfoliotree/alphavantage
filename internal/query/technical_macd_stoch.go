package query

import (
	"net/url"
	"strconv"
	"time"
)

// MACD builds query parameters for the MACD API.
// Returns the moving average convergence / divergence (MACD) values.
type MACD url.Values

// NewMACD creates a new MACD query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	seriesType: The desired price type (close, open, high, low)
func NewMACD(apikey, symbol, interval string, seriesType SeriesTypeOption) MACD {
	return MACD{
		KeyFunction:   []string{FunctionMACD},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q MACD) MonthString(month string) MACD {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q MACD) Month(year int, month time.Month) MACD {
	return encodeMonth(q, year, month)
}

// FastPeriod sets the fast period parameter.
// Positive integers are accepted. By default, fastperiod=12.
func (q MACD) FastPeriod(p int) MACD {
	q[KeyFastPeriod] = []string{strconv.Itoa(p)}
	return q
}

// SlowPeriod sets the slow period parameter.
// Positive integers are accepted. By default, slowperiod=26.
func (q MACD) SlowPeriod(p int) MACD {
	q[KeySlowPeriod] = []string{strconv.Itoa(p)}
	return q
}

// SignalPeriod sets the signal period parameter.
// Positive integers are accepted. By default, signalperiod=9.
func (q MACD) SignalPeriod(p int) MACD {
	q[KeySignalPeriod] = []string{strconv.Itoa(p)}
	return q
}

func (q MACD) DataType(o DataTypeOption) MACD { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q MACD) DataTypeCSV() MACD { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q MACD) DataTypeJSON() MACD { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q MACD) Validate() error {
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
func (q MACD) Encode() string { return encode(q) }

// MACDEXT builds query parameters for the MACDEXT API.
// Returns the moving average convergence / divergence values with controllable moving average type.
type MACDEXT url.Values

// NewMACDEXT creates a new MACDEXT query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	seriesType: The desired price type (close, open, high, low)
func NewMACDEXT(apikey, symbol, interval string, seriesType SeriesTypeOption) MACDEXT {
	return MACDEXT{
		KeyFunction:   []string{FunctionMACDEXT},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q MACDEXT) MonthString(month string) MACDEXT {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q MACDEXT) Month(year int, month time.Month) MACDEXT {
	return encodeMonth(q, year, month)
}

// FastPeriod sets the fast period parameter.
// Positive integers are accepted. By default, fastperiod=12.
func (q MACDEXT) FastPeriod(p int) MACDEXT {
	q[KeyFastPeriod] = []string{strconv.Itoa(p)}
	return q
}

// SlowPeriod sets the slow period parameter.
// Positive integers are accepted. By default, slowperiod=26.
func (q MACDEXT) SlowPeriod(p int) MACDEXT {
	q[KeySlowPeriod] = []string{strconv.Itoa(p)}
	return q
}

// SignalPeriod sets the signal period parameter.
// Positive integers are accepted. By default, signalperiod=9.
func (q MACDEXT) SignalPeriod(p int) MACDEXT {
	q[KeySignalPeriod] = []string{strconv.Itoa(p)}
	return q
}

// FastMAType sets the moving average type for the faster moving average.
// By default, fastmatype=0. Integers 0-8 are accepted.
func (q MACDEXT) FastMAType(t int) MACDEXT {
	q[KeyFastMAType] = []string{strconv.Itoa(t)}
	return q
}

// SlowMAType sets the moving average type for the slower moving average.
// By default, slowmatype=0. Integers 0-8 are accepted.
func (q MACDEXT) SlowMAType(t int) MACDEXT {
	q[KeySlowMAType] = []string{strconv.Itoa(t)}
	return q
}

// SignalMAType sets the moving average type for the signal moving average.
// By default, signalmatype=0. Integers 0-8 are accepted.
func (q MACDEXT) SignalMAType(t int) MACDEXT {
	q[KeySignalMAType] = []string{strconv.Itoa(t)}
	return q
}

func (q MACDEXT) DataType(o DataTypeOption) MACDEXT { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q MACDEXT) DataTypeCSV() MACDEXT { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q MACDEXT) DataTypeJSON() MACDEXT { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q MACDEXT) Validate() error {
	if err := validateRequired(q, KeyAPIKey, KeySymbol, KeyInterval, KeySeriesType); err != nil {
		return err
	}
	if seriesType, ok := q[KeySeriesType]; ok {
		if err := validateSeriesType(seriesType); err != nil {
			return err
		}
	}
	if t, ok := q[KeyFastMAType]; ok {
		if err := validateMAType(t, KeyFastMAType); err != nil {
			return err
		}
	}
	if t, ok := q[KeySlowMAType]; ok {
		if err := validateMAType(t, KeySlowMAType); err != nil {
			return err
		}
	}
	if t, ok := q[KeySignalMAType]; ok {
		if err := validateMAType(t, KeySignalMAType); err != nil {
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
func (q MACDEXT) Encode() string { return encode(q) }

// STOCH builds query parameters for the STOCH API.
// Returns the stochastic oscillator (STOCH) values.
type STOCH url.Values

// NewSTOCH creates a new STOCH query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
func NewSTOCH(apikey, symbol, interval string) STOCH {
	return STOCH{
		KeyFunction: []string{FunctionSTOCH},
		KeySymbol:   []string{symbol},
		KeyInterval: []string{interval},
		KeyAPIKey:   []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q STOCH) MonthString(month string) STOCH {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q STOCH) Month(year int, month time.Month) STOCH {
	return encodeMonth(q, year, month)
}

// FastKPeriod sets the time period of the fastk moving average.
// Positive integers are accepted. By default, fastkperiod=5.
func (q STOCH) FastKPeriod(p int) STOCH {
	q[KeyFastKPeriod] = []string{strconv.Itoa(p)}
	return q
}

// SlowKPeriod sets the time period of the slowk moving average.
// Positive integers are accepted. By default, slowkperiod=3.
func (q STOCH) SlowKPeriod(p int) STOCH {
	q[KeySlowKPeriod] = []string{strconv.Itoa(p)}
	return q
}

// SlowDPeriod sets the time period of the slowd moving average.
// Positive integers are accepted. By default, slowdperiod=3.
func (q STOCH) SlowDPeriod(p int) STOCH {
	q[KeySlowDPeriod] = []string{strconv.Itoa(p)}
	return q
}

// SlowKMAType sets the moving average type for the slowk moving average.
// By default, slowkmatype=0. Integers 0-8 are accepted.
func (q STOCH) SlowKMAType(t int) STOCH {
	q[KeySlowKMAType] = []string{strconv.Itoa(t)}
	return q
}

// SlowDMAType sets the moving average type for the slowd moving average.
// By default, slowdmatype=0. Integers 0-8 are accepted.
func (q STOCH) SlowDMAType(t int) STOCH {
	q[KeySlowDMAType] = []string{strconv.Itoa(t)}
	return q
}

func (q STOCH) DataType(o DataTypeOption) STOCH { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q STOCH) DataTypeCSV() STOCH { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q STOCH) DataTypeJSON() STOCH { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q STOCH) Validate() error {
	if err := validateRequired(q, KeyAPIKey, KeySymbol, KeyInterval); err != nil {
		return err
	}
	if t, ok := q[KeySlowKMAType]; ok {
		if err := validateMAType(t, KeySlowKMAType); err != nil {
			return err
		}
	}
	if t, ok := q[KeySlowDMAType]; ok {
		if err := validateMAType(t, KeySlowDMAType); err != nil {
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
func (q STOCH) Encode() string { return encode(q) }

// STOCHF builds query parameters for the STOCHF API.
// Returns the stochastic fast (STOCHF) values.
type STOCHF url.Values

// NewSTOCHF creates a new STOCHF query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
func NewSTOCHF(apikey, symbol, interval string) STOCHF {
	return STOCHF{
		KeyFunction: []string{FunctionSTOCHF},
		KeySymbol:   []string{symbol},
		KeyInterval: []string{interval},
		KeyAPIKey:   []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q STOCHF) MonthString(month string) STOCHF {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q STOCHF) Month(year int, month time.Month) STOCHF {
	return encodeMonth(q, year, month)
}

// FastKPeriod sets the time period of the fastk moving average.
// Positive integers are accepted. By default, fastkperiod=5.
func (q STOCHF) FastKPeriod(t int) STOCHF {
	q[KeyFastKPeriod] = []string{strconv.Itoa(t)}
	return q
}

// FastDPeriod sets the time period of the fastd moving average.
// Positive integers are accepted. By default, fastdperiod=3.
func (q STOCHF) FastDPeriod(t int) STOCHF {
	q[KeyFastDPeriod] = []string{strconv.Itoa(t)}
	return q
}

// FastDMAType sets the moving average type for the fastd moving average.
// By default, fastdmatype=0. Integers 0-8 are accepted.
func (q STOCHF) FastDMAType(t int) STOCHF {
	q[KeyFastDMAType] = []string{strconv.Itoa(t)}
	return q
}

func (q STOCHF) DataType(o DataTypeOption) STOCHF { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q STOCHF) DataTypeCSV() STOCHF { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q STOCHF) DataTypeJSON() STOCHF { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q STOCHF) Validate() error {
	if err := validateRequired(q, KeyAPIKey, KeySymbol, KeyInterval); err != nil {
		return err
	}
	if t, ok := q[KeyFastDMAType]; ok {
		if err := validateMAType(t, "fastdmatype"); err != nil {
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
func (q STOCHF) Encode() string { return encode(q) }

// RSI builds query parameters for the RSI API.
// Returns the relative strength index (RSI) values.
type RSI url.Values

// NewRSI creates a new RSI query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each RSI value
//	seriesType: The desired price type (close, open, high, low)
func NewRSI(apikey, symbol, interval string, timePeriod int, seriesType SeriesTypeOption) RSI {
	return RSI{
		KeyFunction:   []string{FunctionRSI},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q RSI) MonthString(month string) RSI {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q RSI) Month(year int, month time.Month) RSI {
	return encodeMonth(q, year, month)
}

func (q RSI) DataType(o DataTypeOption) RSI { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q RSI) DataTypeCSV() RSI { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q RSI) DataTypeJSON() RSI { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q RSI) Validate() error {
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
func (q RSI) Encode() string { return encode(q) }

// STOCHRSI builds query parameters for the STOCHRSI API.
// Returns the stochastic relative strength index (STOCHRSI) values.
type STOCHRSI url.Values

// NewSTOCHRSI creates a new STOCHRSI query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each STOCHRSI value
//	seriesType: The desired price type (close, open, high, low)
func NewSTOCHRSI(apikey, symbol, interval string, timePeriod int, seriesType SeriesTypeOption) STOCHRSI {
	return STOCHRSI{
		KeyFunction:   []string{FunctionSTOCHRSI},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q STOCHRSI) MonthString(month string) STOCHRSI {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q STOCHRSI) Month(year int, month time.Month) STOCHRSI {
	return encodeMonth(q, year, month)
}

// FastKPeriod sets the time period of the fastk moving average.
// Positive integers are accepted. By default, fastkperiod=5.
func (q STOCHRSI) FastKPeriod(t int) STOCHRSI {
	q[KeyFastKPeriod] = []string{strconv.Itoa(t)}
	return q
}

// FastDPeriod sets the time period of the fastd moving average.
// Positive integers are accepted. By default, fastdperiod=3.
func (q STOCHRSI) FastDPeriod(t int) STOCHRSI {
	q[KeyFastDPeriod] = []string{strconv.Itoa(t)}
	return q
}

// FastDMAType sets the moving average type for the fastd moving average.
// By default, fastdmatype=0. Integers 0-8 are accepted.
func (q STOCHRSI) FastDMAType(t int) STOCHRSI {
	q[KeyFastDMAType] = []string{strconv.Itoa(t)}
	return q
}

func (q STOCHRSI) DataType(o DataTypeOption) STOCHRSI { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q STOCHRSI) DataTypeCSV() STOCHRSI { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q STOCHRSI) DataTypeJSON() STOCHRSI { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q STOCHRSI) Validate() error {
	if err := validateRequired(q, KeyAPIKey, KeySymbol, KeyInterval, KeyTimePeriod, KeySeriesType); err != nil {
		return err
	}
	if seriesType, ok := q[KeySeriesType]; ok {
		if err := validateSeriesType(seriesType); err != nil {
			return err
		}
	}
	if t, ok := q[KeyFastDMAType]; ok {
		if err := validateMAType(t, "fastdmatype"); err != nil {
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
func (q STOCHRSI) Encode() string { return encode(q) }
