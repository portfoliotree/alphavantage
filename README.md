# AlphaVantage Go Client

[![Go Reference](https://pkg.go.dev/badge/github.com/portfoliotree/alphavantage.svg)](https://pkg.go.dev/github.com/portfoliotree/alphavantage)
[![Test](https://github.com/portfoliotree/alphavantage/actions/workflows/ci.yml/badge.svg)](https://github.com/portfoliotree/alphavantage/actions/workflows/ci.yml)

An unofficial Go client library and CLI for the [AlphaVantage](https://www.alphavantage.co) financial data API.

**Comprehensive API coverage** - Supports all current (Nov. 2025) AlphaVantage API functions including stocks, forex, crypto, technical indicators, economic data, commodities, and more. See [specification/README.md](specification/README.md) for complete API function coverage.

## Documentation

- **[Tutorial](docs/tutorial.md)** - Get started with your first API calls
- **[How-to Guides](docs/how-to-guides.md)** - Solve specific problems and tasks
- **[API Reference](https://pkg.go.dev/github.com/portfoliotree/alphavantage)** - Complete Go API documentation
- **[Explanation](docs/explanation.md)** - Understand the design and architecture
- **[Migration Guide](docs/migration-guide.md)** - Upgrade from v0.3 to v0.4

## Quick Start

### Installation

#### Go Library

```bash
go get github.com/portfoliotree/alphavantage
```

#### CLI Tool

```bash
go install github.com/portfoliotree/alphavantage/cmd/av@latest
```

Verify installation:
```bash
av help
```

### Authentication

Set your AlphaVantage API key as an environment variable:

```bash
export ALPHA_VANTAGE_TOKEN="your-api-key-here"
```

Get a free API key at https://www.alphavantage.co/support/#api-key

## Usage Example

### Go Library

```go
package main

import (
    "context"
    "fmt"
    "github.com/portfoliotree/alphavantage"
)

func main() {
    // Create client with your API key and rate limit plan
    client := alphavantage.NewClient("your-api-key", alphavantage.PremiumPlan75)
    ctx := context.Background()

    // Get daily stock prices with query builder
    query := alphavantage.QueryTimeSeriesDaily(client.APIKey, "IBM")
    rows, err := client.GetTimeSeriesDailyCSVRows(ctx, query)
    if err != nil {
        panic(err)
    }

    // Print recent closing prices (data already parsed into typed structs)
    for _, row := range rows[:5] {
        fmt.Printf("%s: $%.2f\n", row.Timestamp.Format("2006-01-02"), row.Close)
    }
}
```

See [cmd/example/main.go](cmd/example/main.go) and [docs/tutorial.md](docs/tutorial.md) for more examples.

### CLI Tool

```bash
export ALPHA_VANTAGE_API_KEY="your-api-key"

# Get stock quote
av GLOBAL_QUOTE --symbol=IBM

# Get daily time series data
av TIME_SERIES_DAILY --symbol=AAPL --outputsize=compact

# Technical indicator (RSI)
av RSI --symbol=IBM --interval=daily --time-period=14 --series-type=close

# Economic data
av REAL_GDP --interval=annual

# Get help for any function
av TIME_SERIES_INTRADAY --help
```

See [docs/how-to-guides.md](docs/how-to-guides.md) for comprehensive CLI usage examples.

## Contributing

Contributions are welcome! Please:

1. **Report bugs** via GitHub Issues
2. **Suggest features** via GitHub Issues
3. **Submit pull requests** with tests
4. Follow the [TDD](https://en.wikipedia.org/wiki/Test-driven_development) workflow

This project follows [Diataxis](https://diataxis.fr) for documentation structure.

## Development

### Building from Source

```bash
git clone https://github.com/portfoliotree/alphavantage.git
cd alphavantage
go mod download
```

### Running Tests

```bash
# Unit tests
go test ./...

# All tests (requires ALPHA_VANTAGE_TOKEN)
export ALPHA_VANTAGE_TOKEN="your-api-key"
go test ./...
```

### Code Generation

This project uses code generation for client methods and CLI commands:

```bash
go generate ./...
```

See [docs/explanation.md](docs/explanation.md) for details on the code generation architecture.

## License

This project is provided as-is under the MIT License. See [LICENSE](LICENSE) for details.

---

**Note**: This is an unofficial library and is not affiliated with AlphaVantage.
