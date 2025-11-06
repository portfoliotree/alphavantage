package query

import (
	"net/url"
	"strconv"
	"time"
)

// WILLR builds query parameters for the WILLR API.
// Returns the Williams' %R (WILLR) values.
type WILLR url.Values

// NewWILLR creates a new WILLR query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each WILLR value
func NewWILLR(apikey, symbol, interval string, timePeriod int) WILLR {
	return WILLR{
		KeyFunction:   []string{FunctionWILLR},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q WILLR) MonthString(month string) WILLR {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q WILLR) Month(year int, month time.Month) WILLR {
	return encodeMonth(q, year, month)
}

func (q WILLR) DataType(o DataTypeOption) WILLR { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q WILLR) DataTypeCSV() WILLR { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q WILLR) DataTypeJSON() WILLR { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q WILLR) Validate() error {
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
func (q WILLR) Encode() string { return encode(q) }

// APO builds query parameters for the APO API.
// Returns the absolute price oscillator (APO) values.
type APO url.Values

// NewAPO creates a new APO query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	seriesType: The desired price type (close, open, high, low)
func NewAPO(apikey, symbol, interval string, seriesType SeriesTypeOption) APO {
	return APO{
		KeyFunction:   []string{FunctionAPO},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q APO) MonthString(month string) APO {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q APO) Month(year int, month time.Month) APO {
	return encodeMonth(q, year, month)
}

// FastPeriod sets the fast period parameter.
// Positive integers are accepted. By default, fastperiod=12.
func (q APO) FastPeriod(p int) APO {
	q[KeyFastPeriod] = []string{strconv.Itoa(p)}
	return q
}

// SlowPeriod sets the slow period parameter.
// Positive integers are accepted. By default, slowperiod=26.
func (q APO) SlowPeriod(p int) APO {
	q[KeySlowPeriod] = []string{strconv.Itoa(p)}
	return q
}

// MAType sets the moving average type.
// By default, matype=0. Integers 0-8 are accepted.
func (q APO) MAType(t int) APO {
	q[KeyMAType] = []string{strconv.Itoa(t)}
	return q
}

func (q APO) DataType(o DataTypeOption) APO { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q APO) DataTypeCSV() APO { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q APO) DataTypeJSON() APO { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q APO) Validate() error {
	if err := validateRequired(q, KeyAPIKey, KeySymbol, KeyInterval, KeySeriesType); err != nil {
		return err
	}
	if t, ok := q[KeySeriesType]; ok {
		if err := validateSeriesType(t); err != nil {
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
func (q APO) Encode() string { return encode(q) }

// PPO builds query parameters for the PPO API.
// Returns the percentage price oscillator (PPO) values.
type PPO url.Values

// NewPPO creates a new PPO query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	seriesType: The desired price type (close, open, high, low)
func NewPPO(apikey, symbol, interval string, seriesType SeriesTypeOption) PPO {
	return PPO{
		KeyFunction:   []string{FunctionPPO},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q PPO) MonthString(month string) PPO {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q PPO) Month(year int, month time.Month) PPO {
	return encodeMonth(q, year, month)
}

// FastPeriod sets the fast period parameter.
// Positive integers are accepted. By default, fastperiod=12.
func (q PPO) FastPeriod(p int) PPO {
	q[KeyFastPeriod] = []string{strconv.Itoa(p)}
	return q
}

// SlowPeriod sets the slow period parameter.
// Positive integers are accepted. By default, slowperiod=26.
func (q PPO) SlowPeriod(p int) PPO {
	q[KeySlowPeriod] = []string{strconv.Itoa(p)}
	return q
}

// MAType sets the moving average type.
// By default, matype=0. Integers 0-8 are accepted.
func (q PPO) MAType(t int) PPO {
	q[KeyMAType] = []string{strconv.Itoa(t)}
	return q
}

func (q PPO) DataType(o DataTypeOption) PPO { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q PPO) DataTypeCSV() PPO { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q PPO) DataTypeJSON() PPO { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q PPO) Validate() error {
	if err := validateRequired(q, KeyAPIKey, KeySymbol, KeyInterval, KeySeriesType); err != nil {
		return err
	}
	if t, ok := q[KeySeriesType]; ok {
		if err := validateSeriesType(t); err != nil {
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
func (q PPO) Encode() string { return encode(q) }

// MOM builds query parameters for the MOM API.
// Returns the momentum (MOM) values.
type MOM url.Values

// NewMOM creates a new MOM query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each MOM value
//	seriesType: The desired price type (close, open, high, low)
func NewMOM(apikey, symbol, interval string, timePeriod int, seriesType SeriesTypeOption) MOM {
	return MOM{
		KeyFunction:   []string{FunctionMOM},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q MOM) MonthString(month string) MOM {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q MOM) Month(year int, month time.Month) MOM {
	return encodeMonth(q, year, month)
}

func (q MOM) DataType(o DataTypeOption) MOM { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q MOM) DataTypeCSV() MOM { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q MOM) DataTypeJSON() MOM { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q MOM) Validate() error {
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
func (q MOM) Encode() string { return encode(q) }

// BOP builds query parameters for the BOP API.
// Returns the balance of power (BOP) values.
type BOP url.Values

// NewBOP creates a new BOP query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
func NewBOP(apikey, symbol, interval string) BOP {
	return BOP{
		KeyFunction: []string{FunctionBOP},
		KeySymbol:   []string{symbol},
		KeyInterval: []string{interval},
		KeyAPIKey:   []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q BOP) MonthString(month string) BOP {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q BOP) Month(year int, month time.Month) BOP {
	return encodeMonth(q, year, month)
}

func (q BOP) DataType(o DataTypeOption) BOP { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q BOP) DataTypeCSV() BOP { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q BOP) DataTypeJSON() BOP { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q BOP) Validate() error {
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
func (q BOP) Encode() string { return encode(q) }

// CCI builds query parameters for the CCI API.
// Returns the commodity channel index (CCI) values.
type CCI url.Values

// NewCCI creates a new CCI query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each CCI value
func NewCCI(apikey, symbol, interval string, timePeriod int) CCI {
	return CCI{
		KeyFunction:   []string{FunctionCCI},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q CCI) MonthString(month string) CCI {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q CCI) Month(year int, month time.Month) CCI {
	return encodeMonth(q, year, month)
}

func (q CCI) DataType(o DataTypeOption) CCI { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q CCI) DataTypeCSV() CCI { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q CCI) DataTypeJSON() CCI { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q CCI) Validate() error {
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
func (q CCI) Encode() string { return encode(q) }

// CMO builds query parameters for the CMO API.
// Returns the Chande momentum oscillator (CMO) values.
type CMO url.Values

// NewCMO creates a new CMO query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each CMO value
//	seriesType: The desired price type (close, open, high, low)
func NewCMO(apikey, symbol, interval string, timePeriod int, seriesType SeriesTypeOption) CMO {
	return CMO{
		KeyFunction:   []string{FunctionCMO},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q CMO) MonthString(month string) CMO {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q CMO) Month(year int, month time.Month) CMO {
	return encodeMonth(q, year, month)
}

func (q CMO) DataType(o DataTypeOption) CMO { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q CMO) DataTypeCSV() CMO { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q CMO) DataTypeJSON() CMO { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q CMO) Validate() error {
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
func (q CMO) Encode() string { return encode(q) }

// ROC builds query parameters for the ROC API.
// Returns the rate of change (ROC) values.
type ROC url.Values

// NewROC creates a new ROC query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each ROC value
//	seriesType: The desired price type (close, open, high, low)
func NewROC(apikey, symbol, interval string, timePeriod int, seriesType SeriesTypeOption) ROC {
	return ROC{
		KeyFunction:   []string{FunctionROC},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q ROC) MonthString(month string) ROC {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q ROC) Month(year int, month time.Month) ROC {
	return encodeMonth(q, year, month)
}

func (q ROC) DataType(o DataTypeOption) ROC { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q ROC) DataTypeCSV() ROC { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q ROC) DataTypeJSON() ROC { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q ROC) Validate() error {
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
func (q ROC) Encode() string { return encode(q) }

// ROCR builds query parameters for the ROCR API.
// Returns the rate of change ratio (ROCR) values.
type ROCR url.Values

// NewROCR creates a new ROCR query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	timePeriod: Number of data points used to calculate each ROCR value
//	seriesType: The desired price type (close, open, high, low)
func NewROCR(apikey, symbol, interval string, timePeriod int, seriesType SeriesTypeOption) ROCR {
	return ROCR{
		KeyFunction:   []string{FunctionROCR},
		KeySymbol:     []string{symbol},
		KeyInterval:   []string{interval},
		KeyTimePeriod: []string{strconv.Itoa(timePeriod)},
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q ROCR) MonthString(month string) ROCR {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q ROCR) Month(year int, month time.Month) ROCR {
	return encodeMonth(q, year, month)
}

func (q ROCR) DataType(o DataTypeOption) ROCR { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q ROCR) DataTypeCSV() ROCR { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q ROCR) DataTypeJSON() ROCR { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q ROCR) Validate() error {
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
func (q ROCR) Encode() string { return encode(q) }
