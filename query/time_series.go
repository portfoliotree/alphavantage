package query

import (
	"net/url"
	"time"
)

// TimeSeriesIntraday builds query parameters for the TIME_SERIES_INTRADAY API.
// Returns current and 20+ years of historical intraday OHLCV time series of the equity specified,
// covering pre-market and post-market hours when applicable.
type TimeSeriesIntraday url.Values

// NewTimeSeriesIntraday creates a new TIME_SERIES_INTRADAY query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the equity (e.g., "IBM")
//	interval: Time interval between consecutive data points
func NewTimeSeriesIntraday(apikey, symbol string, interval IIntervalOption) TimeSeriesIntraday {
	return TimeSeriesIntraday{
		KeyFunction: []string{FunctionTimeSeriesIntraday},
		KeySymbol:   []string{symbol},
		KeyInterval: interval.values(),
		KeyAPIKey:   []string{apikey},
	}
}

// Adjusted sets whether to adjust for historical split and dividend events.
// Default: true
func (q TimeSeriesIntraday) Adjusted(adjusted bool) TimeSeriesIntraday {
	return boolean(q, KeyAdjusted, adjusted)
}

// ExtendedHours sets whether to include extended trading hours.
// Default: true
func (q TimeSeriesIntraday) ExtendedHours(extendedHours bool) TimeSeriesIntraday {
	return boolean(q, KeyExtendedHours, extendedHours)
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q TimeSeriesIntraday) MonthString(month string) TimeSeriesIntraday {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the specific month to query from year and month values.
func (q TimeSeriesIntraday) Month(year int, month time.Month) TimeSeriesIntraday {
	return encodeMonth(q, year, month)
}

// OutputSize sets the output size.
// Valid values: "compact" (latest 100 data points), "full" (full-length time series)
func (q TimeSeriesIntraday) OutputSize(outputSize OutputSizeOption) TimeSeriesIntraday {
	q[KeyOutputSize] = outputSize.values()
	return q
}

// OutputSizeCompact sets the output size to compact (latest 100 data points).
func (q TimeSeriesIntraday) OutputSizeCompact() TimeSeriesIntraday {
	return q.OutputSize(OutputSizeOptionCompact)
}

// OutputSizeFull sets the output size to full (full-length time series).
func (q TimeSeriesIntraday) OutputSizeFull() TimeSeriesIntraday {
	return q.OutputSize(OutputSizeOptionFull)
}

// Datatype sets the response format.
func (q TimeSeriesIntraday) Datatype(dt DataTypeOption) TimeSeriesIntraday {
	return dataType(q, dt)
}

func (q TimeSeriesIntraday) DataType(o DataTypeOption) TimeSeriesIntraday { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q TimeSeriesIntraday) DataTypeCSV() TimeSeriesIntraday { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q TimeSeriesIntraday) DataTypeJSON() TimeSeriesIntraday { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q TimeSeriesIntraday) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeySymbol, KeyInterval); err != nil {
		return err
	}

	// Validate interval
	if values, ok := q[KeyInterval]; ok {
		if err := validateEnum(values, KeyInterval, IIntervalOptions()); err != nil {
			return err
		}
	}

	// Validate optional fields if present
	if adjusted, ok := q[KeyAdjusted]; ok {
		if err := validateBoolean(adjusted, KeyAdjusted); err != nil {
			return err
		}
	}

	if extendedHours, ok := q[KeyExtendedHours]; ok {
		if err := validateBoolean(extendedHours, KeyExtendedHours); err != nil {
			return err
		}
	}

	if value, ok := q[KeyOutputSize]; ok {
		if err := validateOutputSize(value); err != nil {
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
func (q TimeSeriesIntraday) Encode() string { return encode(q) }

// TimeSeriesDaily builds query parameters for the TIME_SERIES_DAILY API.
// Returns raw (as-traded) daily time series (OHLCV) of the global equity specified,
// covering 20+ years of historical data.
type TimeSeriesDaily url.Values

// NewTimeSeriesDaily creates a new TIME_SERIES_DAILY query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the equity (e.g., "IBM")
func NewTimeSeriesDaily(apikey, symbol string) TimeSeriesDaily {
	return TimeSeriesDaily{
		KeyFunction: []string{FunctionTimeSeriesDaily},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

// OutputSize sets the output size.
// Valid values: "compact" (latest 100 data points), "full" (20+ years of historical data)
func (q TimeSeriesDaily) OutputSize(outputSize OutputSizeOption) TimeSeriesDaily {
	q[KeyOutputSize] = outputSize.values()
	return q
}

// OutputSizeCompact sets the output size to compact (latest 100 data points).
func (q TimeSeriesDaily) OutputSizeCompact() TimeSeriesDaily {
	return q.OutputSize(OutputSizeOptionCompact)
}

// OutputSizeFull sets the output size to full (20+ years of historical data).
func (q TimeSeriesDaily) OutputSizeFull() TimeSeriesDaily {
	return q.OutputSize(OutputSizeOptionFull)
}

// Datatype sets the response format.
func (q TimeSeriesDaily) Datatype(dt DataTypeOption) TimeSeriesDaily {
	return dataType(q, dt)
}

func (q TimeSeriesDaily) DataType(o DataTypeOption) TimeSeriesDaily { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q TimeSeriesDaily) DataTypeCSV() TimeSeriesDaily { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q TimeSeriesDaily) DataTypeJSON() TimeSeriesDaily { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q TimeSeriesDaily) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeySymbol); err != nil {
		return err
	}

	// Validate optional fields if present
	if value, ok := q[KeyOutputSize]; ok {
		if err := validateOutputSize(value); err != nil {
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
func (q TimeSeriesDaily) Encode() string { return encode(q) }

// TimeSeriesDailyAdjusted builds query parameters for the TIME_SERIES_DAILY_ADJUSTED API.
// Returns raw daily OHLCV values, adjusted close values, and historical split/dividend events.
type TimeSeriesDailyAdjusted url.Values

// NewTimeSeriesDailyAdjusted creates a new TIME_SERIES_DAILY_ADJUSTED query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the equity (e.g., "IBM")
func NewTimeSeriesDailyAdjusted(apikey, symbol string) TimeSeriesDailyAdjusted {
	return TimeSeriesDailyAdjusted{
		KeyFunction: []string{FunctionTimeSeriesDailyAdjusted},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

// OutputSize sets the output size.
// Valid values: "compact" (latest 100 data points), "full" (20+ years of historical data)
func (q TimeSeriesDailyAdjusted) OutputSize(outputSize OutputSizeOption) TimeSeriesDailyAdjusted {
	q[KeyOutputSize] = outputSize.values()
	return q
}

// OutputSizeCompact sets the output size to compact (latest 100 data points).
func (q TimeSeriesDailyAdjusted) OutputSizeCompact() TimeSeriesDailyAdjusted {
	return q.OutputSize(OutputSizeOptionCompact)
}

// OutputSizeFull sets the output size to full (20+ years of historical data).
func (q TimeSeriesDailyAdjusted) OutputSizeFull() TimeSeriesDailyAdjusted {
	return q.OutputSize(OutputSizeOptionFull)
}

func (q TimeSeriesDailyAdjusted) DataType(o DataTypeOption) TimeSeriesDailyAdjusted {
	return dataType(q, o)
}

// DataTypeCSV sets the response format to CSV.
func (q TimeSeriesDailyAdjusted) DataTypeCSV() TimeSeriesDailyAdjusted { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q TimeSeriesDailyAdjusted) DataTypeJSON() TimeSeriesDailyAdjusted { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q TimeSeriesDailyAdjusted) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeySymbol); err != nil {
		return err
	}

	// Validate optional fields if present
	if value, ok := q[KeyOutputSize]; ok {
		if err := validateOutputSize(value); err != nil {
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
func (q TimeSeriesDailyAdjusted) Encode() string { return encode(q) }

// TimeSeriesWeekly builds query parameters for the TIME_SERIES_WEEKLY API.
// Returns weekly time series (last trading day of each week, OHLCV) covering 20+ years
// of historical data.
type TimeSeriesWeekly url.Values

// NewTimeSeriesWeekly creates a new TIME_SERIES_WEEKLY query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the equity (e.g., "IBM")
func NewTimeSeriesWeekly(apikey, symbol string) TimeSeriesWeekly {
	return TimeSeriesWeekly{
		KeyFunction: []string{FunctionTimeSeriesWeekly},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

func (q TimeSeriesWeekly) DataType(o DataTypeOption) TimeSeriesWeekly { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q TimeSeriesWeekly) DataTypeCSV() TimeSeriesWeekly { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q TimeSeriesWeekly) DataTypeJSON() TimeSeriesWeekly { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q TimeSeriesWeekly) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeySymbol); err != nil {
		return err
	}

	// Validate optional fields if present
	if dt, ok := q[KeyDataType]; ok {
		if err := validateDatatype(dt); err != nil {
			return err
		}
	}

	return nil
}

// Encode returns the URL-encoded query string.
func (q TimeSeriesWeekly) Encode() string { return encode(q) }

// TimeSeriesWeeklyAdjusted builds query parameters for the TIME_SERIES_WEEKLY_ADJUSTED API.
// Returns weekly adjusted time series (last trading day of each week, OHLCV, adjusted close,
// volume, dividend) covering 20+ years.
type TimeSeriesWeeklyAdjusted url.Values

// NewTimeSeriesWeeklyAdjusted creates a new TIME_SERIES_WEEKLY_ADJUSTED query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the equity (e.g., "IBM")
func NewTimeSeriesWeeklyAdjusted(apikey, symbol string) TimeSeriesWeeklyAdjusted {
	return TimeSeriesWeeklyAdjusted{
		KeyFunction: []string{FunctionTimeSeriesWeeklyAdjusted},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

func (q TimeSeriesWeeklyAdjusted) DataType(o DataTypeOption) TimeSeriesWeeklyAdjusted {
	return dataType(q, o)
}

// DataTypeCSV sets the response format to CSV.
func (q TimeSeriesWeeklyAdjusted) DataTypeCSV() TimeSeriesWeeklyAdjusted { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q TimeSeriesWeeklyAdjusted) DataTypeJSON() TimeSeriesWeeklyAdjusted { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q TimeSeriesWeeklyAdjusted) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeySymbol); err != nil {
		return err
	}

	// Validate optional fields if present
	if dt, ok := q[KeyDataType]; ok {
		if err := validateDatatype(dt); err != nil {
			return err
		}
	}

	return nil
}

// Encode returns the URL-encoded query string.
func (q TimeSeriesWeeklyAdjusted) Encode() string { return encode(q) }

// TimeSeriesMonthly builds query parameters for the TIME_SERIES_MONTHLY API.
// Returns monthly time series (last trading day of each month, OHLCV) covering 20+ years
// of historical data.
type TimeSeriesMonthly url.Values

// NewTimeSeriesMonthly creates a new TIME_SERIES_MONTHLY query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the equity (e.g., "IBM")
func NewTimeSeriesMonthly(apikey, symbol string) TimeSeriesMonthly {
	return TimeSeriesMonthly{
		KeyFunction: []string{FunctionTimeSeriesMonthly},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

func (q TimeSeriesMonthly) DataType(o DataTypeOption) TimeSeriesMonthly { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q TimeSeriesMonthly) DataTypeCSV() TimeSeriesMonthly { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q TimeSeriesMonthly) DataTypeJSON() TimeSeriesMonthly { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q TimeSeriesMonthly) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeySymbol); err != nil {
		return err
	}

	// Validate optional fields if present
	if dt, ok := q[KeyDataType]; ok {
		if err := validateDatatype(dt); err != nil {
			return err
		}
	}

	return nil
}

// Encode returns the URL-encoded query string.
func (q TimeSeriesMonthly) Encode() string { return encode(q) }

// TimeSeriesMonthlyAdjusted builds query parameters for the TIME_SERIES_MONTHLY_ADJUSTED API.
// Returns monthly adjusted time series (last trading day of each month, OHLCV, adjusted close,
// volume, dividend) covering 20+ years.
type TimeSeriesMonthlyAdjusted url.Values

// NewTimeSeriesMonthlyAdjusted creates a new TIME_SERIES_MONTHLY_ADJUSTED query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the equity (e.g., "IBM")
func NewTimeSeriesMonthlyAdjusted(apikey, symbol string) TimeSeriesMonthlyAdjusted {
	return TimeSeriesMonthlyAdjusted{
		KeyFunction: []string{FunctionTimeSeriesMonthlyAdjusted},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

func (q TimeSeriesMonthlyAdjusted) DataType(o DataTypeOption) TimeSeriesMonthlyAdjusted {
	return dataType(q, o)
}

// DataTypeCSV sets the response format to CSV.
func (q TimeSeriesMonthlyAdjusted) DataTypeCSV() TimeSeriesMonthlyAdjusted { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q TimeSeriesMonthlyAdjusted) DataTypeJSON() TimeSeriesMonthlyAdjusted { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q TimeSeriesMonthlyAdjusted) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeySymbol); err != nil {
		return err
	}

	// Validate optional fields if present
	if dt, ok := q[KeyDataType]; ok {
		if err := validateDatatype(dt); err != nil {
			return err
		}
	}

	return nil
}

// Encode returns the URL-encoded query string.
func (q TimeSeriesMonthlyAdjusted) Encode() string { return encode(q) }

// RealtimeBulkQuotes builds query parameters for the REALTIME_BULK_QUOTES API.
// Returns realtime quotes for US-traded symbols in bulk, accepting up to 100 symbols
// per request.
type RealtimeBulkQuotes url.Values

// NewRealtimeBulkQuotes creates a new REALTIME_BULK_QUOTES query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: Up to 100 symbols separated by comma (e.g., "MSFT,AAPL,IBM")
func NewRealtimeBulkQuotes(apikey, symbol string) RealtimeBulkQuotes {
	return RealtimeBulkQuotes{
		KeyFunction: []string{FunctionRealtimeBulkQuotes},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

func (q RealtimeBulkQuotes) DataType(o DataTypeOption) RealtimeBulkQuotes { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q RealtimeBulkQuotes) DataTypeCSV() RealtimeBulkQuotes { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q RealtimeBulkQuotes) DataTypeJSON() RealtimeBulkQuotes { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q RealtimeBulkQuotes) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeySymbol); err != nil {
		return err
	}

	// Validate optional fields if present
	if dt, ok := q[KeyDataType]; ok {
		if err := validateDatatype(dt); err != nil {
			return err
		}
	}

	return nil
}

// Encode returns the URL-encoded query string.
func (q RealtimeBulkQuotes) Encode() string { return encode(q) }

// SymbolSearch builds query parameters for the SYMBOL_SEARCH API.
// Returns best-matching symbols and market information based on keywords.
type SymbolSearch url.Values

// NewSymbolSearch creates a new SYMBOL_SEARCH query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	keywords: A text string of your choice (e.g., "microsoft")
func NewSymbolSearch(apikey, keywords string) SymbolSearch {
	return SymbolSearch{
		KeyFunction: []string{FunctionSymbolSearch},
		KeyKeywords: []string{keywords},
		KeyAPIKey:   []string{apikey},
	}
}

func (q SymbolSearch) DataType(o DataTypeOption) SymbolSearch { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q SymbolSearch) DataTypeCSV() SymbolSearch { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q SymbolSearch) DataTypeJSON() SymbolSearch { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q SymbolSearch) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeyKeywords); err != nil {
		return err
	}

	// Validate optional fields if present
	if dt, ok := q[KeyDataType]; ok {
		if err := validateDatatype(dt); err != nil {
			return err
		}
	}

	return nil
}

// Encode returns the URL-encoded query string.
func (q SymbolSearch) Encode() string { return encode(q) }

// MarketStatus builds query parameters for the MARKET_STATUS API.
// Returns the current market status (open vs. closed) of major trading venues worldwide.
type MarketStatus url.Values

// NewMarketStatus creates a new MARKET_STATUS query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewMarketStatus(apikey string) MarketStatus {
	return MarketStatus{
		KeyFunction: []string{FunctionMarketStatus},
		KeyAPIKey:   []string{apikey},
	}
}

// Validate checks if all required parameters are present.
func (q MarketStatus) Validate() error {
	return validateRequired(q, KeyAPIKey)
}

// Encode returns the URL-encoded query string.
func (q MarketStatus) Encode() string { return encode(q) }

// GlobalQuote builds query parameters for the GLOBAL_QUOTE API.
// Returns the latest price and volume information for a ticker of your choice.
type GlobalQuote url.Values

// NewGlobalQuote creates a new GLOBAL_QUOTE query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The symbol of the global ticker (e.g., "IBM")
func NewGlobalQuote(apikey, symbol string) GlobalQuote {
	return GlobalQuote{
		KeyFunction: []string{FunctionGlobalQuote},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

func (q GlobalQuote) DataType(o DataTypeOption) GlobalQuote { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q GlobalQuote) DataTypeCSV() GlobalQuote { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q GlobalQuote) DataTypeJSON() GlobalQuote { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q GlobalQuote) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeySymbol); err != nil {
		return err
	}

	// Validate optional fields if present
	if dt, ok := q[KeyDataType]; ok {
		if err := validateDatatype(dt); err != nil {
			return err
		}
	}

	return nil
}

// Encode returns the URL-encoded query string.
func (q GlobalQuote) Encode() string { return encode(q) }
