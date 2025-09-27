# Getting Started with AlphaVantage Go Client

This tutorial will walk you through using the AlphaVantage Go client library to fetch financial data from the AlphaVantage API.

## Prerequisites

- A recent Go SDK installed
- An AlphaVantage API key (free tier available at https://www.alphavantage.co)

## Installation

Install the package:

```bash
go get github.com/portfoliotree/alphavantage
```

Or install the CLI tool:

```bash
go install github.com/portfoliotree/alphavantage/cmd/av@latest
```

## Setting up Authentication

First, get your API key from AlphaVantage and set it as an environment variable:

```bash
export ALPHA_VANTAGE_TOKEN="your-api-key-here"
```

## Your First Request - Getting a Stock Quote

Let's start with getting the latest quote for a stock.

**Example:** [tutorial_example_use_client_test.go](tutorial_example_use_client_test.go)

The basic pattern is:
1. Create a client with your API key
2. Call the appropriate method with a context and symbol
3. Process the returned CSV data
4. Always close the response when done

## Using the CLI Tool

The CLI tool makes it easy to fetch data without writing Go code:

```bash
# Get latest quote
av global-quote IBM

# Get daily time series data
av quotes --function=TIME_SERIES_DAILY IBM

# Search for symbols
av symbol-search "International Business"
```

## Next Steps

- Learn about [time series data](how-to-guides.md#time-series-data)
- Explore [company fundamentals](how-to-guides.md#fundamental-data)
- Check the [API reference](https://pkg.go.dev/github.com/portfoliotree/alphavantage) for all available methods