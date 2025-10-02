# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is an unofficial Go client library and CLI for the AlphaVantage REST API (https://www.alphavantage.co). The project provides both a Go package for programmatic access and a command-line tool called `av` for fetching financial data.

**Development Philosophy**: This project follows Test-Driven Development (TDD) principles where each new feature starts with a failing test, followed by minimal implementation to make it pass, then refactoring for quality.

## Development Goals

### Current API Coverage Gaps
Based on AlphaVantage's full API, this package is missing:
- Economic Indicators (GDP, CPI, unemployment, etc.)
- Technical Indicators (SMA, EMA, RSI, etc.)
- Forex (FX) data
- Cryptocurrency data
- Commodities data
- Quote endpoint (GLOBAL_QUOTE)
- Weekly/Monthly adjusted time series
- News & Sentiment data
- Earnings data (statements, calendar)
- Market status endpoint

### TDD Implementation Strategy
1. **Red**: Write a failing test that describes the desired functionality
2. **Green**: Implement minimal code to make the test pass
3. **Refactor**: Clean up code while keeping tests passing

### Documentation Standards
Following [Diataxis](http://diataxis.fr) principles:
- **Tutorials**: Step-by-step learning-oriented guides
- **How-to Guides**: Problem-solving oriented directions
- **Reference**: Information-oriented API documentation
- **Explanation**: Understanding-oriented discussions

## Build and Test Commands

```bash
# Build the project
go build -v ./...

# Run all tests (unit + integration)
go test -v ./...

# Run only unit tests
go test -v -short ./...

# Run integration tests (requires ALPHA_VANTAGE_TOKEN)
go test -v -run Integration ./...

# Run CLI tests with scripttest
go test -v ./cmd/av

# Install the CLI tool
go install github.com/portfoliotree/alphavantage/cmd/av@latest

# Build the CLI locally
go build -o av ./cmd/av
```

## Core Architecture

### Client Structure
- `Client` struct in `client.go` is the main HTTP client with rate limiting (5 requests per minute)
- Uses `golang.org/x/time/rate` for API rate limiting
- Supports custom HTTP client and limiter interfaces for testability
- Error handling propagates original API error messages to end users

### API Endpoints
The library is organized by API endpoint types:

- **Time Series Data** (`global_quote.go`): Stock quotes with functions like `TIME_SERIES_DAILY`, `TIME_SERIES_MONTHLY`
- **Company Data** (`overview.go`): Company overview information
- **Listing Status** (`listing_status.go`): Stock listing and delisting status
- **Symbol Search** (`symbol_search.go`): Search for stock symbols

### Response Handling
- Most responses return `io.ReadCloser` for CSV data streaming
- Company overview uses JSON parsing to `CompanyOverview` struct
- Row iterator pattern in `io.go` for processing CSV responses line by line
- Error responses from API are parsed and propagated with original messages

### CLI Tool
- Located in `cmd/av/main.go`
- Uses `pflag` library for POSIX-compliant command-line flags
- Supports commands: `quotes`, `listing-status`, `symbol-search`, `help`
- Uses `ALPHA_VANTAGE_TOKEN` environment variable for authentication
- Outputs data to CSV files or stdout

## Authentication

Set the `ALPHA_VANTAGE_TOKEN` environment variable with your AlphaVantage API key. The CLI also accepts a `--token` flag.

## Testing

### Unit Tests
- Use `testify` library for assertions
- Mock HTTP responses with test data in `testdata/` directory
- Fast execution, no external dependencies

### Integration Tests
- Use `rsc.io/script/scripttest` for CLI testing
- Make real API calls with rate limiting respect
- Require `ALPHA_VANTAGE_TOKEN` environment variable
- Handle free tier API limits (5 requests/minute)
- Tests automatically wait between requests to respect rate limits

### Test Organization
- Unit tests: `*_test.go` files with `-short` flag support
- Integration tests: `*_integration_test.go` files
- CLI tests: Located in `cmd/av/` directory using scripttest