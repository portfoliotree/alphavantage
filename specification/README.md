# API Specification

This directory contains JSON specifications for all AlphaVantage API functions. These specifications drive the code generation for both the Go client library and CLI tool.

## Navigating the Specifications

The specifications mirror the structure of the [AlphaVantage API documentation](https://www.alphavantage.co/documentation/).

### Directory Structure

```
specification/
├── query_parameters.json    # Shared parameters (apikey, datatype, etc.)
├── identifiers.json         # API string to Go identifier mappings
├── functions/               # Function specifications by category
│   ├── time_series.json     # Stock Time Series
│   └── ...                  # Options
└── testdata/examples/       # Cached API responses for tests
```

### Mapping to AlphaVantage Documentation

Each JSON file in `functions/` corresponds to a section in the [AlphaVantage documentation](https://www.alphavantage.co/documentation/).
### Finding a Specific API Function

1. **Check the AlphaVantage docs** at https://www.alphavantage.co/documentation/ to find the API function name (e.g., `TIME_SERIES_DAILY`)
2. **Determine the category** from the docs sidebar (e.g., "Stock Time Series APIs")
3. **Open the corresponding JSON file** in `functions/` (e.g., `time_series.json`)
4. **Search for the function** by its `name` field in the JSON array

Example: To find the `CPI` (Consumer Price Index) function:
- AlphaVantage docs → "Economic Indicators" section
- Open `functions/economic.json`
- Search for `"name": "CPI"`

## Code Generation

The specifications in this directory drive code generation for:

1. **Go Client Methods** - Type-safe query builders and response parsers
2. **CLI Handlers** - Command-line interface for all functions
3. **Documentation** - API reference and examples

To regenerate code from specifications:

```bash
go generate ./...
```

This uses `cmd/generate/main.go` which reads the JSON specifications and generates:
- `*.go` files in the root package (time_series.go, fundamental.go, etc.)
- `cmd/av/functions.go` - CLI handlers

## Validation

The specifications are validated by tests in `specification_test.go`:

- Structural integrity of JSON files
- Referential integrity between files
- Required fields presence
- Valid Go identifier mappings

Run validation:

```bash
go test ./specification
```

## Adding New Functions

To add support for a new AlphaVantage API function:

1. Add the function specification to the appropriate `functions/*.json` file
2. Add any new parameters to `query_parameters.json` if needed
3. Add identifier mappings to `identifiers.json`
4. Run `go generate ./...` to regenerate code
5. Add test data to `testdata/examples/`
6. Run tests to verify

The generated code will automatically include the new function in both the Go library and CLI tool.
