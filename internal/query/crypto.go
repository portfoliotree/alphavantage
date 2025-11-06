package query

import (
	"net/url"
)

// CryptoIntraday builds query parameters for the CRYPTO_INTRADAY API.
// Returns intraday time series (timestamp, open, high, low, close, volume) of the
// cryptocurrency specified, updated realtime.
type CryptoIntraday url.Values

// NewCryptoIntraday creates a new CRYPTO_INTRADAY query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The digital/crypto currency symbol (e.g., "ETH", "BTC")
//	market: The exchange market (e.g., "USD", "EUR")
//	interval: Minimum time interval between quotes
func NewCryptoIntraday(apikey, symbol, market string, interval IIntervalOption) CryptoIntraday {
	return CryptoIntraday{
		KeyFunction: []string{FunctionCryptoIntraday},
		KeySymbol:   []string{symbol},
		KeyMarket:   []string{market},
		KeyInterval: interval.values(),
		KeyAPIKey:   []string{apikey},
	}
}

// OutputSize sets the data output size.
// Valid values: "compact", "full"
func (q CryptoIntraday) OutputSize(outputSize OutputSizeOption) CryptoIntraday {
	q[KeyOutputSize] = outputSize.values()
	return q
}

// OutputSizeCompact sets the output size to compact (latest 100 data points).
func (q CryptoIntraday) OutputSizeCompact() CryptoIntraday {
	return q.OutputSize(OutputSizeOptionCompact)
}

// OutputSizeFull sets the output size to full (full-length time series).
func (q CryptoIntraday) OutputSizeFull() CryptoIntraday {
	return q.OutputSize(OutputSizeOptionFull)
}

func (q CryptoIntraday) DataType(o DataTypeOption) CryptoIntraday { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q CryptoIntraday) DataTypeCSV() CryptoIntraday { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q CryptoIntraday) DataTypeJSON() CryptoIntraday { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q CryptoIntraday) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeySymbol, KeyMarket, KeyInterval); err != nil {
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
func (q CryptoIntraday) Encode() string { return encode(q) }

// DigitalCurrencyDaily builds query parameters for the DIGITAL_CURRENCY_DAILY API.
// Returns the daily historical time series for a digital currency traded on a specific market,
// refreshed daily at midnight (UTC).
type DigitalCurrencyDaily url.Values

// NewDigitalCurrencyDaily creates a new DIGITAL_CURRENCY_DAILY query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The digital/crypto currency symbol (e.g., "BTC", "ETH")
//	market: The exchange market (e.g., "EUR", "USD")
func NewDigitalCurrencyDaily(apikey, symbol, market string) DigitalCurrencyDaily {
	return DigitalCurrencyDaily{
		KeyFunction: []string{FunctionDigitalCurrencyDaily},
		KeySymbol:   []string{symbol},
		KeyMarket:   []string{market},
		KeyAPIKey:   []string{apikey},
	}
}

func (q DigitalCurrencyDaily) DataType(o DataTypeOption) DigitalCurrencyDaily { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q DigitalCurrencyDaily) DataTypeCSV() DigitalCurrencyDaily { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q DigitalCurrencyDaily) DataTypeJSON() DigitalCurrencyDaily { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q DigitalCurrencyDaily) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeySymbol, KeyMarket); err != nil {
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
func (q DigitalCurrencyDaily) Encode() string { return encode(q) }

// DigitalCurrencyWeekly builds query parameters for the DIGITAL_CURRENCY_WEEKLY API.
// Returns the weekly historical time series for a digital currency traded on a specific market,
// refreshed daily at midnight (UTC).
type DigitalCurrencyWeekly url.Values

// NewDigitalCurrencyWeekly creates a new DIGITAL_CURRENCY_WEEKLY query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The digital/crypto currency symbol (e.g., "BTC", "ETH")
//	market: The exchange market (e.g., "EUR", "USD")
func NewDigitalCurrencyWeekly(apikey, symbol, market string) DigitalCurrencyWeekly {
	return DigitalCurrencyWeekly{
		KeyFunction: []string{FunctionDigitalCurrencyWeekly},
		KeySymbol:   []string{symbol},
		KeyMarket:   []string{market},
		KeyAPIKey:   []string{apikey},
	}
}

func (q DigitalCurrencyWeekly) DataType(o DataTypeOption) DigitalCurrencyWeekly {
	return dataType(q, o)
}

// DataTypeCSV sets the response format to CSV.
func (q DigitalCurrencyWeekly) DataTypeCSV() DigitalCurrencyWeekly { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q DigitalCurrencyWeekly) DataTypeJSON() DigitalCurrencyWeekly { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q DigitalCurrencyWeekly) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeySymbol, KeyMarket); err != nil {
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
func (q DigitalCurrencyWeekly) Encode() string { return encode(q) }

// DigitalCurrencyMonthly builds query parameters for the DIGITAL_CURRENCY_MONTHLY API.
// Returns the monthly historical time series for a digital currency traded on a specific market,
// refreshed daily at midnight (UTC).
type DigitalCurrencyMonthly url.Values

// NewDigitalCurrencyMonthly creates a new DIGITAL_CURRENCY_MONTHLY query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The digital/crypto currency symbol (e.g., "BTC", "ETH")
//	market: The exchange market (e.g., "EUR", "USD")
func NewDigitalCurrencyMonthly(apikey, symbol, market string) DigitalCurrencyMonthly {
	return DigitalCurrencyMonthly{
		KeyFunction: []string{FunctionDigitalCurrencyMonthly},
		KeySymbol:   []string{symbol},
		KeyMarket:   []string{market},
		KeyAPIKey:   []string{apikey},
	}
}

func (q DigitalCurrencyMonthly) DataType(o DataTypeOption) DigitalCurrencyMonthly {
	return dataType(q, o)
}

// DataTypeCSV sets the response format to CSV.
func (q DigitalCurrencyMonthly) DataTypeCSV() DigitalCurrencyMonthly { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q DigitalCurrencyMonthly) DataTypeJSON() DigitalCurrencyMonthly { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q DigitalCurrencyMonthly) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeySymbol, KeyMarket); err != nil {
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
func (q DigitalCurrencyMonthly) Encode() string { return encode(q) }
