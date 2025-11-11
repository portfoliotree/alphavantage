# Migration Guide: v0.3.x to v0.4.x

This guide helps you migrate from alphavantage v0.3.x to v0.4.x. Version 0.4 introduces significant improvements including code generation, comprehensive API coverage, and a redesigned CLI.

## Overview of Changes

### Major Improvements

- **All API functions supported** (up from about 7 functions in v0.3)
- **Code-generated client methods** for consistency and maintainability
- **Query builder pattern** for fluent, type-safe API calls
- **Unified CLI** with all functions accessible via flags
- **Comprehensive test coverage** with cached API responses
- **Better error handling** and response parsing

### Breaking Changes

1. **Client method signatures changed** - All methods now use Query types
2. **CLI command structure changed** - From subcommands to function-based invocation
3. **Removed legacy methods** - Old quote/time series methods deprecated
4. **Response types renamed** - More consistent naming conventions

## Client API Migration

### Old Pattern (v0.3.x)

```go
import "github.com/portfoliotree/alphavantage"

client := alphavantage.NewClient(apiKey)

// Time series (deprecated methods)
resp, err := client.DoQuotesRequest(ctx,
    alphavantage.TimeSeriesDaily,
    "IBM")

// Company overview
overview, err := client.CompanyOverview(ctx, "IBM")

// Global quote
resp, err := client.DoGlobalQuoteRequest(ctx, "IBM")
```

### New Pattern (v0.4.x)

```go
import "github.com/portfoliotree/alphavantage"

client := alphavantage.NewClient(apiKey, alphavantage.FreePlan)

// Time series with query builder
query := alphavantage.QueryTimeSeriesDaily(client.APIKey, "IBM")
rows, err := client.GetTimeSeriesDailyCSVRows(ctx, query)

// Company overview
query := alphavantage.QueryOverview(client.APIKey, "IBM")
resp, err := client.GetOverview(ctx, query)

// Global quote
query := alphavantage.QueryGlobalQuote(client.APIKey, "IBM")
rows, err := client.GetGlobalQuoteCSVRows(ctx, query)
```

### Client Constructor Changes

**v0.3.x:**
```go
client := alphavantage.NewClient(apiKey)
```

**v0.4.x:**
```go
// Specify rate limit plan
client := alphavantage.NewClient(apiKey, alphavantage.FreePlan)      // 5 req/min
client := alphavantage.NewClient(apiKey, alphavantage.PremiumPlan75) // 75 req/min
```

Available rate limit plans:
- `PremiumPlan75` - 75 requests per minute
- `PremiumPlan159` - 120 requests per minute
- `PremiumPlan300` - 300 requests per minute
- `PremiumPlan600` - 600 requests per minute
- `PremiumPlan1200` - 1200 requests per minute

## Detailed Method Migration

### Time Series Data

**v0.3.x:**
```go
// Using DoQuotesRequest with function constants
resp, err := client.DoQuotesRequest(ctx,
    alphavantage.TimeSeriesDaily, "AAPL")
defer resp.Close()

resp, err := client.DoQuotesRequest(ctx,
    alphavantage.TimeSeriesIntraday, "AAPL")
```

**v0.4.x:**
```go
// Daily data with typed query
query := alphavantage.QueryTimeSeriesDaily(client.APIKey, "AAPL")
rows, err := client.GetTimeSeriesDailyCSVRows(ctx, query)
// rows is []TimeSeriesDailyRow - already parsed!

// Intraday with required interval parameter
query := alphavantage.QueryTimeSeriesIntraday(client.APIKey, "AAPL", "5min")
query = query.Month("2024-01") // Optional parameters via fluent API
rows, err := client.GetTimeSeriesIntradayCSVRows(ctx, query)

// Adjusted time series
query := alphavantage.QueryTimeSeriesDailyAdjusted(client.APIKey, "AAPL")
rows, err := client.GetTimeSeriesDailyAdjustedCSVRows(ctx, query)
```

### Company Overview

**v0.3.x:**
```go
overview, err := client.CompanyOverview(ctx, "IBM")
fmt.Println(overview.Name)
fmt.Println(overview.MarketCapitalization)
```

**v0.4.x:**
```go
query := alphavantage.QueryOverview(client.APIKey, "IBM")
resp, err := client.GetOverview(ctx, query)
defer resp.Body.Close()

var overview alphavantage.CompanyOverview
err = json.NewDecoder(resp.Body).Decode(&overview)
fmt.Println(overview.Name)
fmt.Println(overview.MarketCapitalization)
```

### Global Quote

**v0.3.x:**
```go
resp, err := client.DoGlobalQuoteRequest(ctx, "IBM")
defer resp.Close()
// Parse CSV manually
```

**v0.4.x:**
```go
query := alphavantage.QueryGlobalQuote(client.APIKey, "IBM")
rows, err := client.GetGlobalQuoteCSVRows(ctx, query)
// rows is []GlobalQuoteRow - already parsed!

fmt.Printf("Symbol: %s, Price: %s\n", rows[0].Symbol, rows[0].Price)
```

### Symbol Search

**v0.3.x:**
```go
resp, err := client.DoSymbolSearchRequest(ctx, "Microsoft")
defer resp.Close()
```

**v0.4.x:**
```go
query := alphavantage.QuerySymbolSearch(client.APIKey, "Microsoft")
rows, err := client.GetSymbolSearchCSVRows(ctx, query)

for _, row := range rows {
    fmt.Printf("%s - %s\n", row.Symbol, row.Name)
}
```

### Listing Status

**v0.3.x:**
```go
resp, err := client.DoListingStatusRequest(ctx, true) // listed
resp, err := client.DoListingStatusRequest(ctx, false) // delisted
defer resp.Close()
```

**v0.4.x:**
```go
query := alphavantage.QueryListingStatus(client.APIKey)
query = query.StateActive()  // For listed companies
// OR
query = query.StateDelisted() // For delisted companies

rows, err := client.GetListingStatusCSVRows(ctx, query)
```

## New API Categories (v0.4.x)

Version 0.4 adds comprehensive support for many new AlphaVantage API categories. See Go Doc for new methods and functions.

## CLI Migration

### Command Structure Changes

**v0.3.x - Subcommand-based:**
```bash
# Old subcommand approach
av global-quote IBM
av quotes --function=TIME_SERIES_DAILY IBM
av listing-status --listed=true
av symbol-search "Microsoft"
```

**v0.4.x - Function-based:**
```bash
# New unified approach - all functions directly accessible
av GLOBAL_QUOTE --symbol=IBM
av TIME_SERIES_DAILY --symbol=IBM
av LISTING_STATUS --state=active
av SYMBOL_SEARCH --keywords="Microsoft"

# Technical indicators
av RSI --symbol=IBM --interval=daily --time-period=14 --series-type=close
av SMA --symbol=AAPL --interval=weekly --time-period=50 --series-type=close

# Economic indicators
av REAL_GDP --interval=annual
av UNEMPLOYMENT
av CPI --interval=monthly

# Forex
av FX_DAILY --from-symbol=EUR --to-symbol=USD

# Commodities
av WTI --interval=monthly
av COPPER
```

### Flag Names

All flags now use lowercase with hyphens (kebab-case):

**v0.3.x:**
```bash
av quotes --function=TIME_SERIES_DAILY IBM
```

**v0.4.x:**
```bash
av TIME_SERIES_DAILY --symbol=IBM --datatype=csv --outputsize=compact
```

### Help System

**v0.3.x:**
```bash
av help              # General help
av global-quote -h   # Subcommand help
```

**v0.4.x:**
```bash
av help              # Lists all 92 functions
av TIME_SERIES_DAILY --help  # Function-specific help with all parameters
```

## Query Builder Pattern

Version 0.4 introduces a fluent query builder pattern for all API calls:

```go
// Build queries with method chaining
query := alphavantage.QueryTimeSeriesIntraday(client.APIKey, "IBM", "5min").
    Month("2024-01").
    OutputSizeFull().
    Adjusted(true).
    ExtendedHours(true)

rows, err := client.GetTimeSeriesIntradayCSVRows(ctx, query)

// Each query type has specific builder methods
query := alphavantage.QueryRSI(client.APIKey, "AAPL", "daily", 14, "close").
    Month("2024-01")

// Data type selection
query := alphavantage.QueryGlobalQuote(client.APIKey, "IBM").
    DataTypeCSV()  // or DataTypeJSON()
```

## Response Parsing

### CSV Response Handling

**v0.3.x - Manual parsing:**
```go
resp, err := client.DoQuotesRequest(ctx, alphavantage.TimeSeriesDaily, "IBM")
defer resp.Close()

reader := csv.NewReader(resp)
for {
    record, err := reader.Read()
    // Manual CSV parsing...
}
```

**v0.4.x - Automatic parsing:**
```go
query := alphavantage.QueryTimeSeriesDaily(client.APIKey, "IBM")
rows, err := client.GetTimeSeriesDailyCSVRows(ctx, query)

for _, row := range rows {
    fmt.Printf("Date: %s, Close: %f\n", row.Timestamp, row.Close)
}
```

### Typed Response Structs

Some CSV responses now have typed row structs:

```go
type TimeSeriesDailyRow struct {
    Timestamp time.Time `column-name:"timestamp"`
    Open      float64   `column-name:"open"`
    High      float64   `column-name:"high"`
    Low       float64   `column-name:"low"`
    Close     float64   `column-name:"close"`
    Volume    int       `column-name:"volume"`
}

type GlobalQuoteRow struct {
    Symbol        string `column-name:"symbol"`
    Open          string `column-name:"open"`
    High          string `column-name:"high"`
    Low           string `column-name:"low"`
    Price         string `column-name:"price"`
    Volume        string `column-name:"volume"`
    LatestDay     string `column-name:"latestDay"`
    PreviousClose string `column-name:"previousClose"`
    Change        string `column-name:"change"`
    ChangePercent string `column-name:"changePercent"`
}
```

_Help me add typed fields by changing the type field in csv_columns arrays in the ./specification/functions files.

## Error Handling

Error handling remains similar, but error messages are more descriptive:

```go
// Both versions
rows, err := client.GetTimeSeriesDailyCSVRows(ctx, query)
if err != nil {
    // v0.4 provides more context in error messages
    log.Fatal(err)
}
```

## Deprecated Methods

The following methods from v0.3.x are **removed** in v0.4.x:

- `DoQuotesRequest()` → Use `QueryTimeSeriesXXX()` + `GetTimeSeriesXXXCSVRows()`
- `DoGlobalQuoteRequest()` → Use `QueryGlobalQuote()` + `GetGlobalQuoteCSVRows()`
- `DoSymbolSearchRequest()` → Use `QuerySymbolSearch()` + `GetSymbolSearchCSVRows()`
- `DoListingStatusRequest()` → Use `QueryListingStatus()` + `GetListingStatusCSVRows()`
- `CompanyOverview()` → Use `QueryOverview()` + `GetOverview()`
