# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is an unofficial Go client library and CLI for the AlphaVantage REST API (https://www.alphavantage.co). The project provides both a Go package for programmatic access and a command-line tool called `av` for fetching financial data.

**Version**: v0.4.x - Complete API rewrite with comprehensive coverage

**Development Philosophy**: This project follows Test-Driven Development (TDD) principles where each new feature starts with a failing test, followed by minimal implementation to make it pass, then refactoring for quality.

## API Coverage

**Complete Coverage**: This library supports **all AlphaVantage API functions** across 9 categories:

- ✅ Core Stock APIs (11 functions) - TIME_SERIES_*, GLOBAL_QUOTE, MARKET_STATUS, etc.
- ✅ Fundamental Data (12 functions) - OVERVIEW, INCOME_STATEMENT, BALANCE_SHEET, EARNINGS, etc.
- ✅ Technical Indicators (56 functions) - SMA, EMA, RSI, MACD, BBANDS, etc.
- ✅ Economic Indicators (10 functions) - REAL_GDP, UNEMPLOYMENT, CPI, INFLATION, etc.
- ✅ Forex Data (5 functions) - CURRENCY_EXCHANGE_RATE, FX_DAILY, FX_INTRADAY, etc.
- ✅ Cryptocurrency (4 functions) - CRYPTO_INTRADAY, DIGITAL_CURRENCY_DAILY, etc.
- ✅ Commodities (10 functions) - WTI, BRENT, NATURAL_GAS, COPPER, WHEAT, etc.
- ✅ Intelligence & Analytics (6 functions) - NEWS_SENTIMENT, TOP_GAINERS_LOSERS, INSIDER_TRANSACTIONS, etc.
- ✅ Options Data (2 functions) - REALTIME_OPTIONS, HISTORICAL_OPTIONS

See [specification/README.md](specification/README.md) for complete function list.

## Code Generation

**Important**: Most client code is **auto-generated** from JSON specifications.

### Code Generation Architecture
- Specifications: `specification/functions/*.json` and `specification/query_parameters.json`
- Generator: `cmd/generate/main.go` uses `go/ast` to generate code
- Generated files: `*.go` files in root package (time_series.go, fundamental.go, etc.)
- CLI handlers: `cmd/av/functions.go`

### Regenerating Code
```bash
go generate ./...
```

### Adding New API Functions
1. Add specification to `specification/functions/*.json`
2. Add identifiers to `specification/identifiers.json`
3. Run `go generate ./...`
4. Add test data to `specification/testdata/examples/`

## Development Philosophy

### TDD Implementation Strategy
1. **Red**: Write a failing test that describes the desired functionality
2. **Green**: Implement minimal code to make the test pass
3. **Refactor**: Clean up code while keeping tests passing

### Documentation Standards
Following [Diataxis](http://diataxis.fr) principles:
- **Tutorials**: Step-by-step learning-oriented guides (docs/tutorial.md)
- **How-to Guides**: Problem-solving oriented directions (docs/how-to-guides.md)
- **Reference**: Information-oriented API documentation (pkg.go.dev)
- **Explanation**: Understanding-oriented discussions (docs/explanation.md)
- **Examples**: Runnable code examples (docs/examples/)

## Navigating Documentation

### For Learning (New Users)
**Start here**: [docs/tutorial.md](docs/tutorial.md)
- Step-by-step introduction to the library
- Basic examples with explanations
- Links to runnable code in `docs/examples/`

### For Solving Problems (Developers)
**Start here**: [docs/how-to-guides.md](docs/how-to-guides.md)
- Task-oriented guides for common scenarios
- References to runnable examples for each task
- Quick reference for all API categories

### For Understanding Design (Contributors)
**Start here**: [docs/explanation.md](docs/explanation.md)
- Why query builder pattern?
- Why CSV over JSON?
- Code generation philosophy using go/ast
- Architecture decisions and trade-offs

### For API Reference
- Complete API docs: https://pkg.go.dev/github.com/portfoliotree/alphavantage
- Function catalog: [specification/README.md](specification/README.md)
- Runnable examples: [docs/examples/README.md](docs/examples/README.md)

### For Migration from v0.3
**Start here**: [docs/migration-guide.md](docs/migration-guide.md)

## Build and Test Commands

```bash
# Regenerate code from specifications
go generate ./...

# Build the project
go build -v ./...

# Run all tests (unit + integration)
go test -v ./...

# Run only unit tests
go test -v -short ./...

# Install the CLI tool
go install github.com/portfoliotree/alphavantage/cmd/av@latest
```

## Core Architecture

### Code Generation System

**Critical**: Do NOT manually edit generated files. Always edit specifications and regenerate.

#### Specification Files
```
specification/
├── functions/              # API function specs (one file per category)
│   ├── time_series.json   # Stock time series APIs
│   ├── fundamental.json   # Fundamental data APIs
│   ├── technical_*.json   # Technical indicators (8 files)
│   └── ...
├── query_parameters.json  # Shared query parameters
├── identifiers.json       # Go identifier mappings
└── testdata/examples/     # Cached API responses for testing
```

#### Generated Files
- **Client methods**: `time_series.go`, `fundamental.go`, `technical_*.go`, etc. (in root package)
- **Query types**: Type-safe query builders for each API function
- **Response structs**: CSV row types with `column-name` tags
- **CLI handlers**: `cmd/av/functions.go` with all 92 commands

#### Code Generator
- **Location**: `cmd/generate/main.go`
- **Uses**: `go/ast` package to build Abstract Syntax Trees
- **Output**: Formatted Go code with proper imports and documentation

#### How Code Generation Works
1. **Read** JSON specifications from `specification/`
2. **Build** Go AST nodes for structs, functions, methods
3. **Format** using `go/format` for idiomatic code
4. **Write** to target files

**Example**: Adding a new API function never requires writing Go code manually - just add JSON spec and regenerate.

### Client Structure
- `Client` struct in `client.go` with automatic rate limiting
- Uses `golang.org/x/time/rate` for API throttling
- Supports all rate limit tiers: Free (5/min) through Premium (1200/min)
- Error handling propagates original API error messages

### Query Builder Pattern
Every API function has a typed query builder:

```go
// Build query with type-safe methods
query := alphavantage.QueryTimeSeriesDaily(client.APIKey, "IBM").
    OutputSizeFull().
    DataTypeCSV()

// Execute and get typed results
rows, err := client.GetTimeSeriesDailyCSVRows(ctx, query)
```

**Implementation**: Query types are `url.Values` with type-safe methods added by code generation.

### Response Handling
- **CSV responses**: Automatically parsed into typed structs (e.g., `[]TimeSeriesDailyRow`)
- **JSON responses**: Return `io.ReadCloser` for manual decoding
- **Streaming**: `ParseCSVRows` iterator for memory-efficient processing
- **Error detection**: Automatic detection of API error responses

### CLI Tool
- **Function-based commands**: `av TIME_SERIES_DAILY --symbol=IBM`
- **Coverage**: All 92 API functions available as commands
- **Generation**: Auto-generated from same specifications as Go client
- **Authentication**: Uses `ALPHA_VANTAGE_TOKEN` environment variable

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