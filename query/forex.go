package query

import (
	"net/url"
)

// CurrencyExchangeRate builds query parameters for the CURRENCY_EXCHANGE_RATE API.
// Returns the realtime exchange rate for any pair of digital currency (e.g., Bitcoin)
// or physical currency (e.g., USD).
type CurrencyExchangeRate url.Values

// NewCurrencyExchangeRate creates a new CURRENCY_EXCHANGE_RATE query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	fromCurrency: The currency you would like to get the exchange rate for (e.g., "USD", "BTC")
//	toCurrency: The destination currency for the exchange rate (e.g., "EUR", "BTC")
func NewCurrencyExchangeRate(apikey, fromCurrency, toCurrency string) CurrencyExchangeRate {
	return CurrencyExchangeRate{
		KeyFunction:     []string{FunctionCurrencyExchangeRate},
		KeyFromCurrency: []string{fromCurrency},
		KeyToCurrency:   []string{toCurrency},
		KeyAPIKey:       []string{apikey},
	}
}

// Validate checks if all required parameters are present.
func (q CurrencyExchangeRate) Validate() error {
	return validateRequired(q, KeyAPIKey, KeyFromCurrency, KeyToCurrency)
}

// Encode returns the URL-encoded query string.
func (q CurrencyExchangeRate) Encode() string { return encode(q) }

// FXIntraday builds query parameters for the FX_INTRADAY API.
// Returns intraday time series (timestamp, open, high, low, close) of the FX currency pair
// specified, updated realtime.
type FXIntraday url.Values

// NewFXIntraday creates a new FX_INTRADAY query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	fromSymbol: A three-letter symbol from the forex currency list (e.g., "EUR")
//	toSymbol: A three-letter symbol from the forex currency list (e.g., "USD")
//	interval: Minimum time interval between results
func NewFXIntraday(apikey, fromSymbol, toSymbol string, interval IIntervalOption) FXIntraday {
	return FXIntraday{
		KeyFunction:   []string{FunctionFXIntraday},
		KeyFromSymbol: []string{fromSymbol},
		KeyToSymbol:   []string{toSymbol},
		KeyInterval:   interval.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// OutputSize sets the data output size.
// Valid values: "compact", "full"
func (q FXIntraday) OutputSize(outputSize OutputSizeOption) FXIntraday {
	q[KeyOutputSize] = outputSize.values()
	return q
}

// OutputSizeCompact sets the output size to compact (latest 100 data points).
func (q FXIntraday) OutputSizeCompact() FXIntraday {
	return q.OutputSize(OutputSizeOptionCompact)
}

// OutputSizeFull sets the output size to full (full-length time series).
func (q FXIntraday) OutputSizeFull() FXIntraday {
	return q.OutputSize(OutputSizeOptionFull)
}

func (q FXIntraday) DataType(o DataTypeOption) FXIntraday { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q FXIntraday) DataTypeCSV() FXIntraday { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q FXIntraday) DataTypeJSON() FXIntraday { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q FXIntraday) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeyFromSymbol, KeyToSymbol, KeyInterval); err != nil {
		return err
	}

	// Validate interval
	if err := validateEnum(q[KeyInterval], KeyInterval, IIntervalOptions()); err != nil {
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
func (q FXIntraday) Encode() string { return encode(q) }

// FXDaily builds query parameters for the FX_DAILY API.
// Returns the daily time series (timestamp, open, high, low, close) of the FX currency pair
// specified, updated realtime.
type FXDaily url.Values

// NewFXDaily creates a new FX_DAILY query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	fromSymbol: A three-letter symbol from the forex currency list (e.g., "EUR")
//	toSymbol: A three-letter symbol from the forex currency list (e.g., "USD")
func NewFXDaily(apikey, fromSymbol, toSymbol string) FXDaily {
	return FXDaily{
		KeyFunction:   []string{FunctionFXDaily},
		KeyFromSymbol: []string{fromSymbol},
		KeyToSymbol:   []string{toSymbol},
		KeyAPIKey:     []string{apikey},
	}
}

// OutputSize sets the data output size.
// Valid values: "compact", "full"
func (q FXDaily) OutputSize(outputSize OutputSizeOption) FXDaily {
	q[KeyOutputSize] = outputSize.values()
	return q
}

// OutputSizeCompact sets the output size to compact (latest 100 data points).
func (q FXDaily) OutputSizeCompact() FXDaily {
	return q.OutputSize(OutputSizeOptionCompact)
}

// OutputSizeFull sets the output size to full (full-length time series).
func (q FXDaily) OutputSizeFull() FXDaily {
	return q.OutputSize(OutputSizeOptionFull)
}

func (q FXDaily) DataType(o DataTypeOption) FXDaily { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q FXDaily) DataTypeCSV() FXDaily { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q FXDaily) DataTypeJSON() FXDaily { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q FXDaily) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeyFromSymbol, KeyToSymbol); err != nil {
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
func (q FXDaily) Encode() string { return encode(q) }

// FXWeekly builds query parameters for the FX_WEEKLY API.
// Returns the weekly time series (timestamp, open, high, low, close) of the FX currency pair
// specified, updated realtime.
type FXWeekly url.Values

// NewFXWeekly creates a new FX_WEEKLY query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	fromSymbol: A three-letter symbol from the forex currency list (e.g., "EUR")
//	toSymbol: A three-letter symbol from the forex currency list (e.g., "USD")
func NewFXWeekly(apikey, fromSymbol, toSymbol string) FXWeekly {
	return FXWeekly{
		KeyFunction:   []string{FunctionFXWeekly},
		KeyFromSymbol: []string{fromSymbol},
		KeyToSymbol:   []string{toSymbol},
		KeyAPIKey:     []string{apikey},
	}
}

func (q FXWeekly) DataType(o DataTypeOption) FXWeekly { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q FXWeekly) DataTypeCSV() FXWeekly { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q FXWeekly) DataTypeJSON() FXWeekly { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q FXWeekly) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeyFromSymbol, KeyToSymbol); err != nil {
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
func (q FXWeekly) Encode() string { return encode(q) }

// FXMonthly builds query parameters for the FX_MONTHLY API.
// Returns the monthly time series (timestamp, open, high, low, close) of the FX currency pair
// specified, updated realtime.
type FXMonthly url.Values

// NewFXMonthly creates a new FX_MONTHLY query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	fromSymbol: A three-letter symbol from the forex currency list (e.g., "EUR")
//	toSymbol: A three-letter symbol from the forex currency list (e.g., "USD")
func NewFXMonthly(apikey, fromSymbol, toSymbol string) FXMonthly {
	return FXMonthly{
		KeyFunction:   []string{FunctionFXMonthly},
		KeyFromSymbol: []string{fromSymbol},
		KeyToSymbol:   []string{toSymbol},
		KeyAPIKey:     []string{apikey},
	}
}

func (q FXMonthly) DataType(o DataTypeOption) FXMonthly { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q FXMonthly) DataTypeCSV() FXMonthly { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q FXMonthly) DataTypeJSON() FXMonthly { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q FXMonthly) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeyFromSymbol, KeyToSymbol); err != nil {
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
func (q FXMonthly) Encode() string { return encode(q) }
