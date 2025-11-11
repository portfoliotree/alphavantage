# Getting Started with AlphaVantage Go Client

This tutorial will walk you through using the AlphaVantage Go client library to fetch financial data from the AlphaVantage API.

## Prerequisites

- Go 1.24 or later installed
- An AlphaVantage API key (free tier available at https://www.alphavantage.co/support/#api-key)

## Installation

### Installing the Go Library

Add the library to your project:

```bash
go get github.com/portfoliotree/alphavantage
```

### Installing the CLI Tool

Install the command-line tool globally:

```bash
go install github.com/portfoliotree/alphavantage/cmd/av@latest
```

Verify the installation:

```bash
av help
```

## Setting up Authentication

Get your API key from AlphaVantage and set it as an environment variable:

```bash
export ALPHA_VANTAGE_API_KEY="your-api-key-here"
# or use the legacy variable name
export ALPHA_VANTAGE_TOKEN="your-api-key-here"
```

The `NewClient()` function automatically loads your API key from these environment variables.

## Your First Request - Getting a Stock Quote

Let's start by getting the latest quote for a stock using the Go library.

### Example 1: Latest Stock Quote

See [examples/getting_started/01_stock_quote.go](examples/getting_started/01_stock_quote.go) for a complete runnable example.

**Key concepts:**

1. **Create a client** with `NewClient()`
2. **Build a query** using `QueryGlobalQuote()`
3. **Get parsed results** with `GetGlobalQuoteCSVRows()`
4. **Always handle errors** - API calls can fail
5. **Use a context** for timeout and cancellation support

### Example 2: Historical Stock Prices

Let's fetch daily historical prices.

See [examples/getting_started/02_historical_prices.go](examples/getting_started/02_historical_prices.go) for a complete runnable example.

**Note:** The response is automatically parsed into `[]TimeSeriesDailyRow` with typed fields like `Timestamp` (time.Time), `Open` (float64), etc.

### Example 3: Using the Query Builder

Queries support a fluent builder pattern for optional parameters.

See [examples/getting_started/03_query_builder.go](examples/getting_started/03_query_builder.go) for a complete runnable example.

## Using the CLI Tool

The CLI provides quick access to all 92 API functions without writing Go code.

### Getting a Quote

```bash
av GLOBAL_QUOTE --symbol=IBM
```

This creates a CSV file with the quote data.

### Getting Daily Prices

```bash
# Get daily prices (compact - last 100 days)
av TIME_SERIES_DAILY --symbol=AAPL --outputsize=compact

# Get full history (20+ years)
av TIME_SERIES_DAILY --symbol=AAPL --outputsize=full
```

### Getting Intraday Data

```bash
# 5-minute intervals
av TIME_SERIES_INTRADAY --symbol=MSFT --interval=5min

# Specific month
av TIME_SERIES_INTRADAY --symbol=MSFT --interval=15min --month=2024-01
```

### Searching for Symbols

```bash
av SYMBOL_SEARCH --keywords="Apple Inc"
```

### Getting Help

```bash
# List all available functions
av help

# Get help for a specific function
av TIME_SERIES_DAILY --help
```

## Understanding Rate Limits

AlphaVantage enforces rate limits based on your subscription tier.

**Free Tier**: The free tier has a limit of 25 requests per day. This is too restrictive for automatic rate limiting - you should implement application-level logic to stay under this limit.

**Premium Tiers**: 75, 150, 300, 600, or 1200 requests per minute

Configure rate limiting via environment variable or by setting the limiter directly:

```go
// Option 1: Configure via environment variable
// export ALPHA_VANTAGE_REQUEST_PER_MINUTE=75
client := alphavantage.NewClient()

// Option 2: Set limiter directly on client
client := alphavantage.NewClient()
client.Limiter = alphavantage.PremiumPlan75.Limiter()

// Option 3: Set manually without limiter (no automatic rate limiting)
client := alphavantage.NewClient()
client.Limiter = nil  // No automatic rate limiting
```

When a limiter is configured, the client **automatically handles rate limiting** - you don't need to add delays or throttling yourself.

## Response Types

### CSV Responses (Recommended)

Most functions return CSV data that's automatically parsed into typed structs:

```go
// Time series returns []TimeSeriesDailyRow
rows, err := client.GetTimeSeriesDailyCSVRows(ctx, query)

// Each row has typed fields
for _, row := range rows {
	fmt.Printf("Date: %s, Close: %.2f\n",
		row.Timestamp.Format("2006-01-02"),
		row.Close)
}
```

### JSON Responses

Some functions (like company overview) return JSON:

```go
query := alphavantage.QueryOverview(client.APIKey, "IBM")
resp, err := client.GetOverview(ctx, query)
if err != nil {
	log.Fatal(err)
}
defer resp.Body.Close()

var overview struct {
	Symbol      string `json:"Symbol"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Sector      string `json:"Sector"`
	// ... more fields
}

err = json.NewDecoder(resp.Body).Decode(&overview)
fmt.Printf("%s - %s\n", overview.Symbol, overview.Name)
```

### Manual CSV Control

You can also choose between CSV and JSON for most endpoints:

```go
// Force CSV format
query := alphavantage.QueryGlobalQuote(client.APIKey, "IBM").DataTypeCSV()

// Force JSON format
query := alphavantage.QueryGlobalQuote(client.APIKey, "IBM").DataTypeJSON()
```

## Common Patterns

### Error Handling

Always check errors - API calls can fail for many reasons:

```go
rows, err := client.GetTimeSeriesDailyCSVRows(ctx, query)
if err != nil {
	// Handle the error appropriately
	log.Printf("Failed to fetch data: %v", err)
	return
}
```
