package query

import (
	"net/url"
)

// WTI builds query parameters for the WTI API.
// Returns the West Texas Intermediate (WTI) crude oil prices.
type WTI url.Values

// NewWTI creates a new WTI query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewWTI(apikey string) WTI {
	return WTI{
		KeyFunction: []string{FunctionWTI},
		KeyAPIKey:   []string{apikey},
	}
}

// Interval sets the time interval.
// Valid values: "daily", "weekly", "monthly" (default: "monthly")
func (q WTI) Interval(interval DWMIntervalOption) WTI {
	q[KeyInterval] = interval.values()
	return q
}

func (q WTI) DataType(o DataTypeOption) WTI { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q WTI) DataTypeCSV() WTI { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q WTI) DataTypeJSON() WTI { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q WTI) Validate() error {
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	if values, ok := q[KeyInterval]; ok {
		if err := validateEnum(values, KeyInterval, DWMIntervalOptions()); err != nil {
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
func (q WTI) Encode() string { return encode(q) }

// Brent builds query parameters for the BRENT API.
// Returns the Brent (Europe) crude oil prices.
type Brent url.Values

// NewBrent creates a new BRENT query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewBrent(apikey string) Brent {
	return Brent{
		KeyFunction: []string{FunctionBrent},
		KeyAPIKey:   []string{apikey},
	}
}

// Interval sets the time interval.
// Valid values: "daily", "weekly", "monthly" (default: "monthly")
func (q Brent) Interval(interval DWMIntervalOption) Brent {
	q[KeyInterval] = interval.values()
	return q
}

func (q Brent) DataType(o DataTypeOption) Brent { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q Brent) DataTypeCSV() Brent { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q Brent) DataTypeJSON() Brent { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q Brent) Validate() error {
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	if values, ok := q[KeyInterval]; ok {
		if err := validateEnum(values, KeyInterval, DWMIntervalOptions()); err != nil {
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
func (q Brent) Encode() string { return encode(q) }

// NaturalGas builds query parameters for the NATURAL_GAS API.
// Returns the Henry Hub natural gas spot prices.
type NaturalGas url.Values

// NewNaturalGas creates a new NATURAL_GAS query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewNaturalGas(apikey string) NaturalGas {
	return NaturalGas{
		KeyFunction: []string{FunctionNaturalGas},
		KeyAPIKey:   []string{apikey},
	}
}

// Interval sets the time interval.
// Valid values: "daily", "weekly", "monthly" (default: "monthly")
func (q NaturalGas) Interval(interval DWMIntervalOption) NaturalGas {
	q[KeyInterval] = interval.values()
	return q
}

func (q NaturalGas) DataType(o DataTypeOption) NaturalGas { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q NaturalGas) DataTypeCSV() NaturalGas { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q NaturalGas) DataTypeJSON() NaturalGas { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q NaturalGas) Validate() error {
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	if values, ok := q[KeyInterval]; ok {
		if err := validateEnum(values, KeyInterval, DWMIntervalOptions()); err != nil {
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
func (q NaturalGas) Encode() string { return encode(q) }

// Copper builds query parameters for the COPPER API.
// Returns the global price of copper.
type Copper url.Values

// NewCopper creates a new COPPER query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewCopper(apikey string) Copper {
	return Copper{
		KeyFunction: []string{FunctionCopper},
		KeyAPIKey:   []string{apikey},
	}
}

// Interval sets the time interval.
// Valid values: "monthly", "quarterly", "annual" (default: "monthly")
func (q Copper) Interval(interval MQAIntervalOption) Copper {
	q[KeyInterval] = interval.values()
	return q
}

func (q Copper) DataType(o DataTypeOption) Copper { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q Copper) DataTypeCSV() Copper { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q Copper) DataTypeJSON() Copper { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q Copper) Validate() error {
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	if values, ok := q[KeyInterval]; ok {
		if err := validateEnum(values, KeyInterval, MQAIntervalOptions()); err != nil {
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
func (q Copper) Encode() string { return encode(q) }

// Aluminum builds query parameters for the ALUMINUM API.
// Returns the global price of aluminum.
type Aluminum url.Values

// NewAluminum creates a new ALUMINUM query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewAluminum(apikey string) Aluminum {
	return Aluminum{
		KeyFunction: []string{FunctionAluminum},
		KeyAPIKey:   []string{apikey},
	}
}

// Interval sets the time interval.
// Valid values: "monthly", "quarterly", "annual" (default: "monthly")
func (q Aluminum) Interval(interval MQAIntervalOption) Aluminum {
	q[KeyInterval] = interval.values()
	return q
}

func (q Aluminum) DataType(o DataTypeOption) Aluminum { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q Aluminum) DataTypeCSV() Aluminum { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q Aluminum) DataTypeJSON() Aluminum { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q Aluminum) Validate() error {
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	if values, ok := q[KeyInterval]; ok {
		if err := validateEnum(values, KeyInterval, MQAIntervalOptions()); err != nil {
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
func (q Aluminum) Encode() string { return encode(q) }

// Wheat builds query parameters for the WHEAT API.
// Returns the global price of wheat.
type Wheat url.Values

// NewWheat creates a new WHEAT query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewWheat(apikey string) Wheat {
	return Wheat{
		KeyFunction: []string{FunctionWheat},
		KeyAPIKey:   []string{apikey},
	}
}

// Interval sets the time interval.
// Valid values: "monthly", "quarterly", "annual" (default: "monthly")
func (q Wheat) Interval(interval MQAIntervalOption) Wheat {
	q[KeyInterval] = interval.values()
	return q
}

func (q Wheat) DataType(o DataTypeOption) Wheat { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q Wheat) DataTypeCSV() Wheat { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q Wheat) DataTypeJSON() Wheat { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q Wheat) Validate() error {
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	if values, ok := q[KeyInterval]; ok {
		if err := validateEnum(values, KeyInterval, MQAIntervalOptions()); err != nil {
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
func (q Wheat) Encode() string { return encode(q) }

// Corn builds query parameters for the CORN API.
// Returns the global price of corn.
type Corn url.Values

// NewCorn creates a new CORN query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewCorn(apikey string) Corn {
	return Corn{
		KeyFunction: []string{FunctionCorn},
		KeyAPIKey:   []string{apikey},
	}
}

// Interval sets the time interval.
// Valid values: "monthly", "quarterly", "annual" (default: "monthly")
func (q Corn) Interval(interval MQAIntervalOption) Corn {
	q[KeyInterval] = interval.values()
	return q
}

func (q Corn) DataType(o DataTypeOption) Corn { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q Corn) DataTypeCSV() Corn { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q Corn) DataTypeJSON() Corn { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q Corn) Validate() error {
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	if values, ok := q[KeyInterval]; ok {
		if err := validateEnum(values, KeyInterval, MQAIntervalOptions()); err != nil {
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
func (q Corn) Encode() string { return encode(q) }

// Cotton builds query parameters for the COTTON API.
// Returns the global price of cotton.
type Cotton url.Values

// NewCotton creates a new COTTON query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewCotton(apikey string) Cotton {
	return Cotton{
		KeyFunction: []string{FunctionCotton},
		KeyAPIKey:   []string{apikey},
	}
}

// Interval sets the time interval.
// Valid values: "monthly", "quarterly", "annual" (default: "monthly")
func (q Cotton) Interval(interval MQAIntervalOption) Cotton {
	q[KeyInterval] = interval.values()
	return q
}

func (q Cotton) DataType(o DataTypeOption) Cotton { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q Cotton) DataTypeCSV() Cotton { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q Cotton) DataTypeJSON() Cotton { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q Cotton) Validate() error {
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	if values, ok := q[KeyInterval]; ok {
		if err := validateEnum(values, KeyInterval, MQAIntervalOptions()); err != nil {
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
func (q Cotton) Encode() string { return encode(q) }

// Sugar builds query parameters for the SUGAR API.
// Returns the global price of sugar.
type Sugar url.Values

// NewSugar creates a new SUGAR query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewSugar(apikey string) Sugar {
	return Sugar{
		KeyFunction: []string{FunctionSugar},
		KeyAPIKey:   []string{apikey},
	}
}

// Interval sets the time interval.
// Valid values: "monthly", "quarterly", "annual" (default: "monthly")
func (q Sugar) Interval(interval MQAIntervalOption) Sugar {
	q[KeyInterval] = interval.values()
	return q
}

func (q Sugar) DataType(o DataTypeOption) Sugar { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q Sugar) DataTypeCSV() Sugar { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q Sugar) DataTypeJSON() Sugar { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q Sugar) Validate() error {
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	if values, ok := q[KeyInterval]; ok {
		if err := validateEnum(values, KeyInterval, MQAIntervalOptions()); err != nil {
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
func (q Sugar) Encode() string { return encode(q) }

// Coffee builds query parameters for the COFFEE API.
// Returns the global price of coffee.
type Coffee url.Values

// NewCoffee creates a new COFFEE query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewCoffee(apikey string) Coffee {
	return Coffee{
		KeyFunction: []string{FunctionCoffee},
		KeyAPIKey:   []string{apikey},
	}
}

// Interval sets the time interval.
// Valid values: "monthly", "quarterly", "annual" (default: "monthly")
func (q Coffee) Interval(interval MQAIntervalOption) Coffee {
	q[KeyInterval] = interval.values()
	return q
}

func (q Coffee) DataType(o DataTypeOption) Coffee { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q Coffee) DataTypeCSV() Coffee { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q Coffee) DataTypeJSON() Coffee { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q Coffee) Validate() error {
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	if values, ok := q[KeyInterval]; ok {
		if err := validateEnum(values, KeyInterval, MQAIntervalOptions()); err != nil {
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
func (q Coffee) Encode() string { return encode(q) }

// AllCommodities builds query parameters for the ALL_COMMODITIES API.
// Returns the global price index of all commodities.
type AllCommodities url.Values

// NewAllCommodities creates a new ALL_COMMODITIES query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewAllCommodities(apikey string) AllCommodities {
	return AllCommodities{
		KeyFunction: []string{FunctionAllCommodities},
		KeyAPIKey:   []string{apikey},
	}
}

// Interval sets the time interval.
// Valid values: "monthly", "quarterly", "annual" (default: "monthly")
func (q AllCommodities) Interval(interval MQAIntervalOption) AllCommodities {
	q[KeyInterval] = interval.values()
	return q
}

func (q AllCommodities) DataType(o DataTypeOption) AllCommodities { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q AllCommodities) DataTypeCSV() AllCommodities { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q AllCommodities) DataTypeJSON() AllCommodities {
	return dataTypeJSON(q)
}

// Validate checks if all required parameters are present and enum values are valid.
func (q AllCommodities) Validate() error {
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	if values, ok := q[KeyInterval]; ok {
		if err := validateEnum(values, KeyInterval, MQAIntervalOptions()); err != nil {
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
func (q AllCommodities) Encode() string { return encode(q) }
