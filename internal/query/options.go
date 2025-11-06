package query

import (
	"net/url"
)

// RealtimeOptions builds query parameters for the REALTIME_OPTIONS API.
// Returns realtime US options data with full market coverage.
type RealtimeOptions url.Values

// NewRealtimeOptions creates a new REALTIME_OPTIONS query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The stock symbol (e.g., "IBM")
func NewRealtimeOptions(apikey, symbol string) RealtimeOptions {
	return RealtimeOptions{
		KeyFunction: []string{FunctionRealtimeOptions},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

// RequireGreeks sets whether to include Greeks and implied volatility.
func (q RealtimeOptions) RequireGreeks(requireGreeks bool) RealtimeOptions {
	return boolean(q, KeyRequireGreeks, requireGreeks)
}

// Contract sets the specific US options contract ID to query.
// Example: "IBM270115C00390000"
func (q RealtimeOptions) Contract(contract string) RealtimeOptions {
	q[KeyContract] = []string{contract}
	return q
}
func (q RealtimeOptions) DataType(o DataTypeOption) RealtimeOptions { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q RealtimeOptions) DataTypeCSV() RealtimeOptions { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q RealtimeOptions) DataTypeJSON() RealtimeOptions { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q RealtimeOptions) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey, KeySymbol); err != nil {
		return err
	}

	// Validate optional fields if present
	if requireGreeks, ok := q[KeyRequireGreeks]; ok {
		if err := validateBoolean(requireGreeks, KeyRequireGreeks); err != nil {
			return err
		}
	}

	if dt, ok := q[KeyDataType]; ok {
		if err := validateDatatype(dt); err != nil {
			return err
		}
	}

	// Note: contract is a free-form string, no validation

	return nil
}

// Encode returns the URL-encoded query string.
func (q RealtimeOptions) Encode() string { return encode(q) }

// HistoricalOptions builds query parameters for the HISTORICAL_OPTIONS API.
// Returns the full historical options chain for a specific symbol on a specific date.
type HistoricalOptions url.Values

// NewHistoricalOptions creates a new HISTORICAL_OPTIONS query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The stock symbol (e.g., "IBM")
func NewHistoricalOptions(apikey, symbol string) HistoricalOptions {
	return HistoricalOptions{
		KeyFunction: []string{FunctionHistoricalOptions},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

// Date sets the specific date to query options data.
// Format: YYYY-MM-DD (e.g., "2017-11-15")
// By default, returns data for the previous trading session.
func (q HistoricalOptions) Date(date string) HistoricalOptions {
	q[KeyDate] = []string{date}
	return q
}

func (q HistoricalOptions) DataType(o DataTypeOption) HistoricalOptions { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q HistoricalOptions) DataTypeCSV() HistoricalOptions { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q HistoricalOptions) DataTypeJSON() HistoricalOptions { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q HistoricalOptions) Validate() error {
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

	// Note: date is not validated (date format - for future MR)

	return nil
}

// Encode returns the URL-encoded query string.
func (q HistoricalOptions) Encode() string { return encode(q) }
