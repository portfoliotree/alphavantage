package query

import (
	"net/url"
	"time"
)

// HilbertTransformTrendLine builds query parameters for the HT_TRENDLINE API.
// Returns the Hilbert transform, instantaneous trendline values.
type HilbertTransformTrendLine url.Values

// NewHilbertTransformTrendLine creates a new HT_TRENDLINE query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	seriesType: The desired price type (close, open, high, low)
func NewHilbertTransformTrendLine(apikey, symbol string, interval IDWMIntervalOption, seriesType SeriesTypeOption) HilbertTransformTrendLine {
	return HilbertTransformTrendLine{
		KeyFunction:   []string{FunctionHilbertTransformTrendLine},
		KeySymbol:     []string{symbol},
		KeyInterval:   interval.values(),
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q HilbertTransformTrendLine) MonthString(month string) HilbertTransformTrendLine {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q HilbertTransformTrendLine) Month(year int, month time.Month) HilbertTransformTrendLine {
	return encodeMonth(q, year, month)
}

func (q HilbertTransformTrendLine) DataType(o DataTypeOption) HilbertTransformTrendLine {
	return dataType(q, o)
}

// DataTypeCSV sets the response format to CSV.
func (q HilbertTransformTrendLine) DataTypeCSV() HilbertTransformTrendLine { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q HilbertTransformTrendLine) DataTypeJSON() HilbertTransformTrendLine { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q HilbertTransformTrendLine) Validate() error {
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
func (q HilbertTransformTrendLine) Encode() string { return encode(q) }

// HilbertTransformSine builds query parameters for the HT_SINE API.
// Returns the Hilbert transform, sine wave values.
type HilbertTransformSine url.Values

// NewHilbertTransformSine creates a new HT_SINE query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	seriesType: The desired price type (close, open, high, low)
func NewHilbertTransformSine(apikey, symbol string, interval IDWMIntervalOption, seriesType SeriesTypeOption) HilbertTransformSine {
	return HilbertTransformSine{
		KeyFunction:   []string{FunctionHilbertTransformSine},
		KeySymbol:     []string{symbol},
		KeyInterval:   interval.values(),
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q HilbertTransformSine) MonthString(month string) HilbertTransformSine {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q HilbertTransformSine) Month(year int, month time.Month) HilbertTransformSine {
	return encodeMonth(q, year, month)
}

func (q HilbertTransformSine) DataType(o DataTypeOption) HilbertTransformSine { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q HilbertTransformSine) DataTypeCSV() HilbertTransformSine { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q HilbertTransformSine) DataTypeJSON() HilbertTransformSine { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q HilbertTransformSine) Validate() error {
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
func (q HilbertTransformSine) Encode() string { return encode(q) }

// HilbertTransformTrendMode builds query parameters for the HT_TRENDMODE API.
// Returns the Hilbert transform, trend vs cycle mode values.
type HilbertTransformTrendMode url.Values

// NewHilbertTransformTrendMode creates a new HT_TRENDMODE query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	seriesType: The desired price type (close, open, high, low)
func NewHilbertTransformTrendMode(apikey, symbol string, interval IDWMIntervalOption, seriesType SeriesTypeOption) HilbertTransformTrendMode {
	return HilbertTransformTrendMode{
		KeyFunction:   []string{FunctionHilbertTransformTrendMode},
		KeySymbol:     []string{symbol},
		KeyInterval:   interval.values(),
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q HilbertTransformTrendMode) MonthString(month string) HilbertTransformTrendMode {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q HilbertTransformTrendMode) Month(year int, month time.Month) HilbertTransformTrendMode {
	return encodeMonth(q, year, month)
}

func (q HilbertTransformTrendMode) DataType(o DataTypeOption) HilbertTransformTrendMode {
	return dataType(q, o)
}

// DataTypeCSV sets the response format to CSV.
func (q HilbertTransformTrendMode) DataTypeCSV() HilbertTransformTrendMode { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q HilbertTransformTrendMode) DataTypeJSON() HilbertTransformTrendMode { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q HilbertTransformTrendMode) Validate() error {
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
func (q HilbertTransformTrendMode) Encode() string { return encode(q) }

// HilbertTransformDCPeriod builds query parameters for the HT_DCPERIOD API.
// Returns the Hilbert transform, dominant cycle period values.
type HilbertTransformDCPeriod url.Values

// NewHilbertTransformDCPeriod creates a new HT_DCPERIOD query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	seriesType: The desired price type (close, open, high, low)
func NewHilbertTransformDCPeriod(apikey, symbol string, interval IDWMIntervalOption, seriesType SeriesTypeOption) HilbertTransformDCPeriod {
	return HilbertTransformDCPeriod{
		KeyFunction:   []string{FunctionHilbertTransformDCPeriod},
		KeySymbol:     []string{symbol},
		KeyInterval:   interval.values(),
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}

}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q HilbertTransformDCPeriod) MonthString(month string) HilbertTransformDCPeriod {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q HilbertTransformDCPeriod) Month(year int, month time.Month) HilbertTransformDCPeriod {
	return encodeMonth(q, year, month)
}

func (q HilbertTransformDCPeriod) DataType(o DataTypeOption) HilbertTransformDCPeriod {
	return dataType(q, o)
}

// DataTypeCSV sets the response format to CSV.
func (q HilbertTransformDCPeriod) DataTypeCSV() HilbertTransformDCPeriod { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q HilbertTransformDCPeriod) DataTypeJSON() HilbertTransformDCPeriod { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q HilbertTransformDCPeriod) Validate() error {
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
func (q HilbertTransformDCPeriod) Encode() string { return encode(q) }

// HilbertTransformDCPhase builds query parameters for the HT_DCPHASE API.
// Returns the Hilbert transform, dominant cycle phase values.
type HilbertTransformDCPhase url.Values

// NewHilbertTransformDCPhase creates a new HT_DCPHASE query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	seriesType: The desired price type (close, open, high, low)
func NewHilbertTransformDCPhase(apikey, symbol string, interval IDWMIntervalOption, seriesType SeriesTypeOption) HilbertTransformDCPhase {
	return HilbertTransformDCPhase{
		KeyFunction:   []string{FunctionHilbertTransformDCPhase},
		KeySymbol:     []string{symbol},
		KeyInterval:   interval.values(),
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q HilbertTransformDCPhase) MonthString(month string) HilbertTransformDCPhase {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q HilbertTransformDCPhase) Month(year int, month time.Month) HilbertTransformDCPhase {
	return encodeMonth(q, year, month)
}

func (q HilbertTransformDCPhase) DataType(o DataTypeOption) HilbertTransformDCPhase {
	return dataType(q, o)
}

// DataTypeCSV sets the response format to CSV.
func (q HilbertTransformDCPhase) DataTypeCSV() HilbertTransformDCPhase { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q HilbertTransformDCPhase) DataTypeJSON() HilbertTransformDCPhase { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q HilbertTransformDCPhase) Validate() error {
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
func (q HilbertTransformDCPhase) Encode() string { return encode(q) }

// HilbertTransformPhasor builds query parameters for the HT_PHASOR API.
// Returns the Hilbert transform, phasor components values.
type HilbertTransformPhasor url.Values

// NewHilbertTransformPhasor creates a new HT_PHASOR query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The name of the ticker (e.g., "IBM")
//	interval: Time interval between data points
//	seriesType: The desired price type (close, open, high, low)
func NewHilbertTransformPhasor(apikey, symbol string, interval IDWMIntervalOption, seriesType SeriesTypeOption) HilbertTransformPhasor {
	return HilbertTransformPhasor{
		KeyFunction:   []string{FunctionHilbertTransformPhasor},
		KeySymbol:     []string{symbol},
		KeyInterval:   interval.values(),
		KeySeriesType: seriesType.values(),
		KeyAPIKey:     []string{apikey},
	}
}

// MonthString sets the month parameter directly from a string in YYYY-MM format.
func (q HilbertTransformPhasor) MonthString(month string) HilbertTransformPhasor {
	q[KeyMonth] = []string{month}
	return q
}

// Month sets the month parameter (YYYY-MM format).
func (q HilbertTransformPhasor) Month(year int, month time.Month) HilbertTransformPhasor {
	return encodeMonth(q, year, month)
}

func (q HilbertTransformPhasor) DataType(o DataTypeOption) HilbertTransformPhasor {
	return dataType(q, o)
}

// DataTypeCSV sets the response format to CSV.
func (q HilbertTransformPhasor) DataTypeCSV() HilbertTransformPhasor { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q HilbertTransformPhasor) DataTypeJSON() HilbertTransformPhasor { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q HilbertTransformPhasor) Validate() error {
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
func (q HilbertTransformPhasor) Encode() string { return encode(q) }
