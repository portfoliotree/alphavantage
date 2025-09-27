# How-to Guides

This document provides solutions to common tasks with the AlphaVantage Go client.

## Time Series Data

### How to fetch daily stock prices

**Example:** [how_to_example_time_series_test.go - ExampleHowTo_FetchDailyStockPrices](how_to_example_time_series_test.go)

Use `DoQuotesRequest` with the `TimeSeriesDaily` function to get raw CSV data.

### How to get adjusted prices (with dividends and splits)

**Example:** [how_to_example_time_series_test.go - ExampleHowTo_GetAdjustedPrices](how_to_example_time_series_test.go)

Use adjusted functions like `TimeSeriesDailyAdjusted` for split/dividend-adjusted data.

### How to fetch weekly or monthly data

**Example:** [how_to_example_time_series_test.go - ExampleHowTo_FetchWeeklyOrMonthlyData](how_to_example_time_series_test.go)

Use `TimeSeriesWeekly`, `TimeSeriesWeeklyAdjusted`, `TimeSeriesMonthly`, or `TimeSeriesMonthlyAdjusted` functions.

## Fundamental Data

### How to get company overview

**Example:** [how_to_example_time_series_test.go - ExampleHowTo_GetCompanyOverview](how_to_example_time_series_test.go)

Use `CompanyOverview` to get detailed company information including financials.

### How to check listing status

**Example:** [how_to_example_time_series_test.go - ExampleHowTo_CheckListingStatus](how_to_example_time_series_test.go)

Use `DoListingStatusRequest` with `true` for listed companies or `false` for delisted companies.

## Rate Limiting

### How to handle API rate limits

The client automatically handles rate limiting for free tier (5 requests/minute). No additional code is needed.

### How to customize rate limiting

**Example:** [how_to_example_time_series_test.go - ExampleHowTo_CustomRateLimiting](how_to_example_time_series_test.go)

Create a custom `Client` struct with your own `Limiter` implementation.

## Error Handling

### How to handle API errors

The client automatically parses and propagates API error messages. Standard Go error handling applies - check the returned error for API-specific messages.

## CLI Usage

### How to save data to specific files

```bash
# Files are automatically saved with symbol names
av global-quote IBM    # Creates IBM_quote.csv
av quotes --function=TIME_SERIES_DAILY AAPL  # Creates AAPL.csv
```

### How to search for symbols

```bash
# Search writes to stdout by default
av symbol-search "Apple Inc"

# Save search results to file
av symbol-search -O "tech companies"  # Creates "tech companies.csv"
```

### How to use custom API keys

```bash
# Using environment variable (recommended)
export ALPHA_VANTAGE_TOKEN="your-key"
av global-quote IBM

# Using command line flag
av --token="your-key" global-quote IBM
```