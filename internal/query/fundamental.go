package query

import (
	"net/url"
)

// Overview builds query parameters for the OVERVIEW API.
// Returns company information, financial ratios, and key metrics for the specified equity.
type Overview url.Values

// NewOverview creates a new OVERVIEW query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The symbol of the ticker (e.g., "IBM")
func NewOverview(apikey, symbol string) Overview {
	return Overview{
		KeyFunction: []string{FunctionOverview},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

// Validate checks if all required parameters are present.
func (q Overview) Validate() error {
	return validateRequired(q, KeyAPIKey, KeySymbol)
}

// Encode returns the URL-encoded query string.
func (q Overview) Encode() string { return encode(q) }

// ETFProfile builds query parameters for the ETF_PROFILE API.
// Returns key ETF metrics and holdings with allocation by asset types and sectors.
type ETFProfile url.Values

// NewETFProfile creates a new ETF_PROFILE query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The symbol of the ticker (e.g., "QQQ")
func NewETFProfile(apikey, symbol string) ETFProfile {
	return ETFProfile{
		KeyFunction: []string{FunctionETFProfile},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

// Validate checks if all required parameters are present.
func (q ETFProfile) Validate() error {
	return validateRequired(q, KeyAPIKey, KeySymbol)
}

// Encode returns the URL-encoded query string.
func (q ETFProfile) Encode() string { return encode(q) }

// Dividends builds query parameters for the DIVIDENDS API.
// Returns historical and future (declared) dividend distributions.
type Dividends url.Values

// NewDividends creates a new DIVIDENDS query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The symbol of the ticker (e.g., "IBM")
func NewDividends(apikey, symbol string) Dividends {
	return Dividends{
		KeyFunction: []string{FunctionDividends},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

func (q Dividends) DataType(o DataTypeOption) Dividends { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q Dividends) DataTypeCSV() Dividends { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q Dividends) DataTypeJSON() Dividends { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q Dividends) Validate() error {
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
func (q Dividends) Encode() string { return encode(q) }

// Splits builds query parameters for the SPLITS API.
// Returns historical split events.
type Splits url.Values

// NewSplits creates a new SPLITS query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The symbol of the ticker (e.g., "IBM")
func NewSplits(apikey, symbol string) Splits {
	return Splits{
		KeyFunction: []string{FunctionSplits},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

func (q Splits) DataType(o DataTypeOption) Splits { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q Splits) DataTypeCSV() Splits { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q Splits) DataTypeJSON() Splits { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q Splits) Validate() error {
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
func (q Splits) Encode() string { return encode(q) }

// IncomeStatement builds query parameters for the INCOME_STATEMENT API.
// Returns annual and quarterly income statements with normalized fields.
type IncomeStatement url.Values

// NewIncomeStatement creates a new INCOME_STATEMENT query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The symbol of the ticker (e.g., "IBM")
func NewIncomeStatement(apikey, symbol string) IncomeStatement {
	return IncomeStatement{
		KeyFunction: []string{FunctionIncomeStatement},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

// Validate checks if all required parameters are present.
func (q IncomeStatement) Validate() error {
	return validateRequired(q, KeyAPIKey, KeySymbol)
}

// Encode returns the URL-encoded query string.
func (q IncomeStatement) Encode() string { return encode(q) }

// BalanceSheet builds query parameters for the BALANCE_SHEET API.
// Returns annual and quarterly balance sheets with normalized fields.
type BalanceSheet url.Values

// NewBalanceSheet creates a new BALANCE_SHEET query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The symbol of the ticker (e.g., "IBM")
func NewBalanceSheet(apikey, symbol string) BalanceSheet {
	return BalanceSheet{
		KeyFunction: []string{FunctionBalanceSheet},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

// Validate checks if all required parameters are present.
func (q BalanceSheet) Validate() error {
	return validateRequired(q, KeyAPIKey, KeySymbol)
}

// Encode returns the URL-encoded query string.
func (q BalanceSheet) Encode() string { return encode(q) }

// CashFlow builds query parameters for the CASH_FLOW API.
// Returns annual and quarterly cash flow with normalized fields.
type CashFlow url.Values

// NewCashFlow creates a new CASH_FLOW query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The symbol of the ticker (e.g., "IBM")
func NewCashFlow(apikey, symbol string) CashFlow {
	return CashFlow{
		KeyFunction: []string{FunctionCashFlow},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

// Validate checks if all required parameters are present.
func (q CashFlow) Validate() error {
	return validateRequired(q, KeyAPIKey, KeySymbol)
}

// Encode returns the URL-encoded query string.
func (q CashFlow) Encode() string { return encode(q) }

// SharesOutstanding builds query parameters for the SHARES_OUTSTANDING API.
// Returns shares outstanding data for the specified equity.
type SharesOutstanding url.Values

// NewSharesOutstanding creates a new SHARES_OUTSTANDING query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The symbol of the ticker (e.g., "MSFT")
func NewSharesOutstanding(apikey, symbol string) SharesOutstanding {
	return SharesOutstanding{
		KeyFunction: []string{FunctionSharesOutstanding},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

func (q SharesOutstanding) DataType(o DataTypeOption) SharesOutstanding { return dataType(q, o) }

// DataTypeCSV sets the response format to CSV.
func (q SharesOutstanding) DataTypeCSV() SharesOutstanding { return dataTypeCSV(q) }

// DataTypeJSON sets the response format to JSON.
func (q SharesOutstanding) DataTypeJSON() SharesOutstanding { return dataTypeJSON(q) }

// Validate checks if all required parameters are present and enum values are valid.
func (q SharesOutstanding) Validate() error {
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
func (q SharesOutstanding) Encode() string { return encode(q) }

// Earnings builds query parameters for the EARNINGS API.
// Returns annual and quarterly earnings (EPS) for the company.
type Earnings url.Values

// NewEarnings creates a new EARNINGS query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The symbol of the ticker (e.g., "IBM")
func NewEarnings(apikey, symbol string) Earnings {
	return Earnings{
		KeyFunction: []string{FunctionEarnings},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

// Validate checks if all required parameters are present.
func (q Earnings) Validate() error {
	return validateRequired(q, KeyAPIKey, KeySymbol)
}

// Encode returns the URL-encoded query string.
func (q Earnings) Encode() string { return encode(q) }

// EarningsEstimates builds query parameters for the EARNINGS_ESTIMATES API.
// Returns annual and quarterly EPS and revenue estimates with analyst data.
type EarningsEstimates url.Values

// NewEarningsEstimates creates a new EARNINGS_ESTIMATES query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
//	symbol: The symbol of the ticker (e.g., "IBM")
func NewEarningsEstimates(apikey, symbol string) EarningsEstimates {
	return EarningsEstimates{
		KeyFunction: []string{FunctionEarningsEstimates},
		KeySymbol:   []string{symbol},
		KeyAPIKey:   []string{apikey},
	}
}

// Validate checks if all required parameters are present.
func (q EarningsEstimates) Validate() error {
	return validateRequired(q, KeyAPIKey, KeySymbol)
}

// Encode returns the URL-encoded query string.
func (q EarningsEstimates) Encode() string { return encode(q) }

// ListingStatus builds query parameters for the LISTING_STATUS API.
// Returns a list of active or delisted US stocks and ETFs.
type ListingStatus url.Values

// NewListingStatus creates a new LISTING_STATUS query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewListingStatus(apikey string) ListingStatus {
	return ListingStatus{
		KeyFunction: []string{FunctionListingStatus},
		KeyAPIKey:   []string{apikey},
	}
}

// Date sets the date filter.
// If not set, returns symbols as of the latest trading day.
func (q ListingStatus) Date(date string) ListingStatus {
	q[KeyDate] = []string{date}
	return q
}

// State sets the listing state filter.
// Valid values: "active", "delisted"
func (q ListingStatus) State(state StateOption) ListingStatus {
	q[KeyState] = state.values()
	return q
}

// StateActive sets the state filter to active listings.
func (q ListingStatus) StateActive() ListingStatus {
	return q.State(StateOptionActive)
}

// StateDelisted sets the state filter to delisted securities.
func (q ListingStatus) StateDelisted() ListingStatus {
	return q.State(StateOptionDelisted)
}

// Validate checks if all required parameters are present and enum values are valid.
func (q ListingStatus) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	// Validate optional fields if present
	if state, ok := q[KeyState]; ok {
		if err := validateState(state); err != nil {
			return err
		}
	}

	return nil
}

// Encode returns the URL-encoded query string.
func (q ListingStatus) Encode() string { return encode(q) }

// EarningsCalendar builds query parameters for the EARNINGS_CALENDAR API.
// Returns a list of company earnings expected in the next 3, 6, or 12 months.
type EarningsCalendar url.Values

// NewEarningsCalendar creates a new EARNINGS_CALENDAR query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewEarningsCalendar(apikey string) EarningsCalendar {
	return EarningsCalendar{
		KeyFunction: []string{FunctionEarningsCalendar},
		KeyAPIKey:   []string{apikey},
	}
}

// Symbol sets the symbol filter.
// If not set, returns full list of scheduled earnings.
func (q EarningsCalendar) Symbol(symbol string) EarningsCalendar {
	q[KeySymbol] = []string{symbol}
	return q
}

// Horizon sets the time horizon.
// Valid values: "3month", "6month", "12month"
func (q EarningsCalendar) Horizon(horizon HorizonOption) EarningsCalendar {
	q[KeyHorizon] = horizon.values()
	return q
}

// Horizon3Month sets the time horizon to 3 months.
func (q EarningsCalendar) Horizon3Month() EarningsCalendar {
	return q.Horizon(HorizonOption3Month)
}

// Horizon6Month sets the time horizon to 6 months.
func (q EarningsCalendar) Horizon6Month() EarningsCalendar {
	return q.Horizon(HorizonOption6Month)
}

// Horizon12Month sets the time horizon to 12 months.
func (q EarningsCalendar) Horizon12Month() EarningsCalendar {
	return q.Horizon(HorizonOption12Month)
}

// Validate checks if all required parameters are present and enum values are valid.
func (q EarningsCalendar) Validate() error {
	// Check required fields
	if err := validateRequired(q, KeyAPIKey); err != nil {
		return err
	}

	// Validate optional fields if present
	if horizon, ok := q[KeyHorizon]; ok {
		if err := validateHorizon(horizon); err != nil {
			return err
		}
	}

	return nil
}

// Encode returns the URL-encoded query string.
func (q EarningsCalendar) Encode() string { return encode(q) }

// IPOCalendar builds query parameters for the IPO_CALENDAR API.
// Returns a list of IPOs expected in the next 3 months.
type IPOCalendar url.Values

// NewIPOCalendar creates a new IPO_CALENDAR query.
//
// Parameters:
//
//	apikey: Your Alpha Vantage API key
func NewIPOCalendar(apikey string) IPOCalendar {
	return IPOCalendar{
		KeyFunction: []string{FunctionIPOCalendar},
		KeyAPIKey:   []string{apikey},
	}
}

// Validate checks if all required parameters are present.
func (q IPOCalendar) Validate() error {
	return validateRequired(q, KeyAPIKey)
}

// Encode returns the URL-encoded query string.
func (q IPOCalendar) Encode() string { return encode(q) }
