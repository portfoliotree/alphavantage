# AlphaVantage Go Client

[![Go Reference](https://pkg.go.dev/badge/github.com/portfoliotree/alphavantage.svg)](https://pkg.go.dev/github.com/portfoliotree/alphavantage)
[![Test](https://github.com/portfoliotree/alphavantage/actions/workflows/test.yml/badge.svg)](https://github.com/portfoliotree/alphavantage/actions/workflows/test.yml)

This is an unofficial REST API wrapper for https://www.alphavantage.co

## Documentation

This project follows the [Diataxis](https://diataxis.fr/) documentation framework:

- **[Tutorial](docs/tutorial.md)** - Get started with your first API calls
- **[How-to Guides](docs/how-to-guides.md)** - Solve specific problems and tasks
- **[API Reference](https://pkg.go.dev/github.com/portfoliotree/alphavantage)** - Complete Go API documentation
- **[Explanation](docs/explanation.md)** - Understand the design and architecture

## Quick Start

### Installation

1. Ensure you have a [recent version of Go](https://go.dev/doc/devel/release) installed.
2. Execute `go install github.com/portfoliotree/alphavantage/cmd/av@latest`
3. Check to see if it installed by executing `av help`

### Authentication

Set an environment variable `ALPHA_VANTAGE_TOKEN` with your token.


### Commands

#### `global-quote`

Fetches latest price and volume information for equity symbols and writes the CSV result to a file.

- Given `av` is installed and the environment variable `ALPHA_VANTAGE_TOKEN` is set
- When you run `av global-quote IBM` and it succeeds
- Then you will see a file "IBM_quote.csv" in your current directory

#### `quotes`

Fetches time series data for given symbols and writes the CSV result to a file.

- Given `av` is installed and the environment variable `ALPHA_VANTAGE_TOKEN` is set
- When you run `av quotes --function=TIME_SERIES_MONTHLY IBM` and it succeeds
- Then you will see a file "IBM.csv" in your current directory

The `--function` may have any of the values: TIME_SERIES_INTRADAY, TIME_SERIES_DAILY, TIME_SERIES_DAILY_ADJUSTED, TIME_SERIES_WEEKLY, TIME_SERIES_WEEKLY_ADJUSTED, TIME_SERIES_MONTHLY, TIME_SERIES_MONTHLY_ADJUSTED.

#### `listing-status`

Fetches listing status and writes the result to a CSV file.

- Given `av` is installed and the environment variable `ALPHA_VANTAGE_TOKEN` is set
- When you run `av listing-status --listed=false` and it succeeds
- Then you will see a file "listing_status_false.csv" in your current directory

#### `symbol-search`

Queries the search endpoint and writes the output to stdout (standard out not a file).

- Given `av` is installed and the environment variable `ALPHA_VANTAGE_TOKEN` is set
- When you run `av symbol-search 'VMware'` and it succeeds
- Then you will see the result printed to standard out formatted as CSV.

#### `help`

Documents the above commands.

The output looks something like this.
```
av - An AlphaVantage CLI in Go

Global Flags:
      --token string   api authentication token

Commands:
  global-quote
	Fetch latest price and volume information for equity.
	https://www.alphavantage.co/documentation/#latestprice
  listing-status
	Fetch listing & de-listing status.
	https://www.alphavantage.co/documentation/#listing-status
  quotes
	Fetch time series stock quotes.
	https://www.alphavantage.co/documentation/#time-series-data
  symbol-search
	Writes symbol search results to stdout.
	https://www.alphavantage.co/documentation/#symbolsearch

```


## Supported Endpoints

Not all endpoints are supported yet. Make an Issue or PR if you'd like to have more endpoints covered.

I have found the CSV encoding easier to parse.

### [Core Stock APIs](https://www.alphavantage.co/documentation/#fundamentals)

- [x] TIME_SERIES_INTRADAY
- [x] TIME_SERIES_DAILY
- [x] TIME_SERIES_DAILY_ADJUSTED
- [x] TIME_SERIES_WEEKLY
- [x] TIME_SERIES_WEEKLY_ADJUSTED
- [x] TIME_SERIES_MONTHLY
- [x] TIME_SERIES_MONTHLY_ADJUSTED
- [x] Quote (GLOBAL_QUOTE)
- [x] Search Endpoint

### [Fundamental Data](https://www.alphavantage.co/documentation/#fx)

- [x] Company Overview
- [ ] Income Statement
- [ ] Balance Sheet
- [ ] Cashflow
- [ ] Earnings
- [x] Listing & Delisting Status
- [ ] Earnings Calendar
- [ ] IPO Calendar

### [Digital & Crypto Currencies](https://www.alphavantage.co/documentation/#digital-currency)

None implemented.

### [Economic Indicators](https://www.alphavantage.co/documentation/#fx)

None implemented (I generally just use FRED directly).

### [Technical Indicators](https://www.alphavantage.co/documentation/#technical-indicators)

None implemented.
