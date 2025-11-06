package query

import (
	"net/url"
)

// RealGDP builds query parameters for the REAL_GDP API.
// Returns the annual and quarterly Real GDP of the United States.
type RealGDP url.Values

// NewRealGDP creates a new REAL_GDP query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewRealGDP(apikey string) RealGDP {
	return RealGDP{
		KeyFunction: []string{FunctionRealGDP},
		KeyAPIKey:   []string{apikey},
	}
}

// Interval sets the time interval.
// Valid values: "annual", "quarterly"
func (q RealGDP) Interval(interval QAIntervalOption) RealGDP {
	q[KeyInterval] = interval.values()
	return q
}

func (q RealGDP) DataType(o DataTypeOption) RealGDP { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q RealGDP) DataTypeCSV() RealGDP { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q RealGDP) DataTypeJSON() RealGDP { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q RealGDP) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	// Validate optional fields if present
	if values, ok := q[KeyInterval]; ok {
		if err := validateEnum(values, KeyInterval, QAIntervalOptions()); err != nil {
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
func (q RealGDP) Encode() string { return encode(q) }

// RealGDPPerCapita builds query parameters for the REAL_GDP_PER_CAPITA API.
// Returns the quarterly Real GDP per Capita data of the United States.
type RealGDPPerCapita url.Values

// NewRealGDPPerCapita creates a new REAL_GDP_PER_CAPITA query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewRealGDPPerCapita(apikey string) RealGDPPerCapita {
	return RealGDPPerCapita{
		KeyFunction: []string{FunctionRealGDPPerCapita},
		KeyAPIKey:   []string{apikey},
	}
}

func (q RealGDPPerCapita) DataType(o DataTypeOption) RealGDPPerCapita { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q RealGDPPerCapita) DataTypeCSV() RealGDPPerCapita { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q RealGDPPerCapita) DataTypeJSON() RealGDPPerCapita { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q RealGDPPerCapita) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey); err != nil {
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
func (q RealGDPPerCapita) Encode() string { return encode(q) }

// TreasuryYield builds query parameters for the TREASURY_YIELD API.
// Returns the daily, weekly, and monthly US treasury yield of a given maturity timeline.
type TreasuryYield url.Values

// NewTreasuryYield creates a new TREASURY_YIELD query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewTreasuryYield(apikey string) TreasuryYield {
	return TreasuryYield{
		KeyFunction: []string{FunctionTreasuryYield},
		KeyAPIKey:   []string{apikey},
	}
}

// Interval sets the time interval.
// Valid values: "daily", "weekly", "monthly"
func (q TreasuryYield) Interval(interval DWMIntervalOption) TreasuryYield {
	q[KeyInterval] = interval.values()
	return q
}

// Maturity sets the maturity timeline.
// Valid values: "3month", "2year", "5year", "7year", "10year", "30year"
func (q TreasuryYield) Maturity(maturity MaturityOption) TreasuryYield {
	q[KeyMaturity] = maturity.values()
	return q
}

func (q TreasuryYield) DataType(o DataTypeOption) TreasuryYield { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q TreasuryYield) DataTypeCSV() TreasuryYield { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q TreasuryYield) DataTypeJSON() TreasuryYield { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q TreasuryYield) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	// Validate optional fields if present
	if values, ok := q[KeyInterval]; ok {
		if err := validateEnum(values, KeyInterval, DWMIntervalOptions()); err != nil {
			return err
		}
	}

	if maturity, ok := q[KeyMaturity]; ok {
		if err := validateEnum(maturity, KeyMaturity, MaturityOptions()); err != nil {
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
func (q TreasuryYield) Encode() string { return encode(q) }

// FederalFundsRate builds query parameters for the FEDERAL_FUNDS_RATE API.
// Returns the daily, weekly, and monthly federal funds rate (interest rate) of the United States.
type FederalFundsRate url.Values

// NewFederalFundsRate creates a new FEDERAL_FUNDS_RATE query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewFederalFundsRate(apikey string) FederalFundsRate {
	return FederalFundsRate{
		KeyFunction: []string{FunctionFederalFundsRate},
		KeyAPIKey:   []string{apikey},
	}
}

// Interval sets the time interval.
// Valid values: "daily", "weekly", "monthly"
func (q FederalFundsRate) Interval(interval DWMIntervalOption) FederalFundsRate {
	q[KeyInterval] = interval.values()
	return q
}

func (q FederalFundsRate) DataType(o DataTypeOption) FederalFundsRate { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q FederalFundsRate) DataTypeCSV() FederalFundsRate { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q FederalFundsRate) DataTypeJSON() FederalFundsRate { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q FederalFundsRate) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	// Validate optional fields if present
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
func (q FederalFundsRate) Encode() string { return encode(q) }

// CPI builds query parameters for the CPI API.
// Returns the monthly and semiannual consumer price index (CPI) of the United States.
type CPI url.Values

// NewCPI creates a new CPI query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewCPI(apikey string) CPI {
	return CPI{
		KeyFunction: []string{FunctionCPI},
		KeyAPIKey:   []string{apikey},
	}
}

// Interval sets the time interval.
// Valid values: "monthly", "semiannual"
func (q CPI) Interval(interval MSIntervalOption) CPI {
	q[KeyInterval] = interval.values()
	return q
}

func (q CPI) DataType(o DataTypeOption) CPI { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q CPI) DataTypeCSV() CPI { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q CPI) DataTypeJSON() CPI { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q CPI) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	// Validate optional fields if present
	if interval, ok := q[KeyInterval]; ok {
		if err := validateEnum(interval, KeyInterval, MSIntervalOptions()); err != nil {
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
func (q CPI) Encode() string { return encode(q) }

// InflationQuery builds query parameters for the INFLATION API.
// Returns the annual inflation rates (consumer prices) of the United States.
type InflationQuery url.Values

// NewInflation creates a new INFLATION query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewInflation(apikey string) InflationQuery {
	return InflationQuery{
		KeyFunction: []string{FunctionInflation},
		KeyAPIKey:   []string{apikey},
	}
}

func (q InflationQuery) DataType(o DataTypeOption) InflationQuery { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q InflationQuery) DataTypeCSV() InflationQuery { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q InflationQuery) DataTypeJSON() InflationQuery { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q InflationQuery) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey); err != nil {
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
func (q InflationQuery) Encode() string { return encode(q) }

// RetailSalesQuery builds query parameters for the RETAIL_SALES API.
// Returns the monthly Advance Retail Sales: Retail Trade data of the United States.
type RetailSalesQuery url.Values

// NewRetailSales creates a new RETAIL_SALES query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewRetailSales(apikey string) RetailSalesQuery {
	return RetailSalesQuery{
		KeyFunction: []string{FunctionRetailSales},
		KeyAPIKey:   []string{apikey},
	}
}

func (q RetailSalesQuery) DataType(o DataTypeOption) RetailSalesQuery { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q RetailSalesQuery) DataTypeCSV() RetailSalesQuery { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q RetailSalesQuery) DataTypeJSON() RetailSalesQuery { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q RetailSalesQuery) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey); err != nil {
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
func (q RetailSalesQuery) Encode() string { return encode(q) }

// DurablesQuery builds query parameters for the DURABLES API.
// Returns the monthly manufacturers' new orders of durable goods in the United States.
type DurablesQuery url.Values

// NewDurables creates a new DURABLES query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewDurables(apikey string) DurablesQuery {
	return DurablesQuery{
		KeyFunction: []string{FunctionDurables},
		KeyAPIKey:   []string{apikey},
	}
}

func (q DurablesQuery) DataType(o DataTypeOption) DurablesQuery { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q DurablesQuery) DataTypeCSV() DurablesQuery { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q DurablesQuery) DataTypeJSON() DurablesQuery { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q DurablesQuery) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey); err != nil {
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
func (q DurablesQuery) Encode() string { return encode(q) }

// Unemployment builds query parameters for the UNEMPLOYMENT API.
// Returns the monthly unemployment data of the United States.
type Unemployment url.Values

// NewUnemployment creates a new UNEMPLOYMENT query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewUnemployment(apikey string) Unemployment {
	return Unemployment{
		KeyFunction: []string{FunctionUnemployment},
		KeyAPIKey:   []string{apikey},
	}
}

func (q Unemployment) DataType(o DataTypeOption) Unemployment { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q Unemployment) DataTypeCSV() Unemployment { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q Unemployment) DataTypeJSON() Unemployment { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q Unemployment) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey); err != nil {
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
func (q Unemployment) Encode() string { return encode(q) }

// NonfarmPayroll builds query parameters for the NONFARM_PAYROLL API.
// Returns the monthly US All Employees: Total Nonfarm (commonly known as Total Nonfarm Payroll).
type NonfarmPayroll url.Values

// NewNonfarmPayroll creates a new NONFARM_PAYROLL query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewNonfarmPayroll(apikey string) NonfarmPayroll {
	return NonfarmPayroll{
		KeyFunction: []string{FunctionNonfarmPayroll},
		KeyAPIKey:   []string{apikey},
	}
}

func (q NonfarmPayroll) DataType(o DataTypeOption) NonfarmPayroll { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q NonfarmPayroll) DataTypeCSV() NonfarmPayroll { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q NonfarmPayroll) DataTypeJSON() NonfarmPayroll { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q NonfarmPayroll) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey); err != nil {
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
func (q NonfarmPayroll) Encode() string { return encode(q) }
