# Understanding AlphaVantage API Integration

This document explains the design decisions and concepts behind the AlphaVantage Go client.

## Architecture Overview

### Code Generation Philosophy

Version 0.4 introduced comprehensive code generation for both the client library and CLI tool. This decision was made to:

1. **Ensure Consistency** - All 92 API functions follow identical patterns
2. **Reduce Errors** - Generated code is less error-prone than hand-written code
3. **Simplify Maintenance** - Adding new API functions is just adding JSON specifications
4. **Improve Testability** - Consistent patterns make testing systematic
5. **Enable Rapid Development** - New AlphaVantage endpoints can be added quickly

The code generator (`cmd/generate/main.go`) uses Go's `go/ast` package to build complete Abstract Syntax Trees, then formats them into idiomatic Go code. This ensures perfect formatting and allows for sophisticated code generation patterns.

### Query Builder Pattern

The library uses a query builder pattern for all API calls:

```go
query := alphavantage.QueryTimeSeriesDaily(client.APIKey, "IBM").
    OutputSizeFull().
    DataTypeCSV()

rows, err := client.GetTimeSeriesDailyCSVRows(ctx, query)
```

**Why query builders?**

1. **Type Safety** - Each query type has specific methods that match its API parameters
2. **Fluent API** - Method chaining makes code readable and self-documenting
3. **Compile-Time Validation** - Missing required parameters cause compile errors
4. **Optional Parameters** - Easy to add optional parameters without breaking existing code
5. **Testability** - Queries can be built and inspected without making API calls

Under the hood, query types are just `url.Values` with type-safe methods:

```go
type TimeSeriesDailyQuery url.Values

func (q TimeSeriesDailyQuery) OutputSizeFull() TimeSeriesDailyQuery {
    q["outputsize"] = []string{"full"}
    return q
}
```

### Typed Response Structs

Some CSV responses are parsed into typed structs (most still are specified as strings):

```go
type TimeSeriesDailyRow struct {
    Timestamp time.Time `column-name:"timestamp"`
    Open      float64   `column-name:"open"`
    High      float64   `column-name:"high"`
    Low       float64   `column-name:"low"`
    Close     float64   `column-name:"close"`
    Volume    int       `column-name:"volume"`
}
```

The `ParseCSV` function uses reflection to map CSV columns to struct fields based on `column-name` tags.

## Why CSV over JSON?

The AlphaVantage API supports both JSON and CSV output formats. This client primarily uses CSV because:

### Efficiency

CSV is more compact than JSON for tabular data:

```json
// JSON (verbose)
{
  "Time Series (Daily)": {
    "2024-01-15": {
      "1. open": "185.89",
      "2. high": "187.33",
      "3. low": "185.46",
      "4. close": "186.64",
      "5. volume": "45739600"
    }
  }
}
```

```csv
// CSV (compact)
timestamp,open,high,low,close,volume
2024-01-15,185.89,187.33,185.46,186.64,45739600
```

For 20+ years of daily data, CSV is significantly smaller.

## Rate Limiting Design

### Automatic Rate Limiting

The client implements automatic rate limiting using `golang.org/x/time/rate`:

```go
type Client struct {
    Limiter interface {
        Wait(ctx context.Context) error
    }
    // ...
}
```

This is mostly when using the client in a long-lived process (like a server). The CLI only makes one API call per invocation.

**How it works:**

1. Before each request, the client calls `Limiter.Wait(ctx)`
2. The limiter blocks until the request can be made within rate limits
3. Context cancellation is respected during waiting

## Error Handling Philosophy

### Transparent Error Propagation

AlphaVantage error messages are preserved and returned as Go errors:

```go
func checkError(rc io.ReadCloser) (io.ReadCloser, error) {
    // Peek at first byte to detect JSON error responses
    if firstByte == '{' {
        var message struct {
            ErrorMessage string `json:"Error Message,omitempty"`
            Note         string `json:"Note,omitempty"`
        }
        // Parse and return AlphaVantage's error message
        return nil, fmt.Errorf("alphavantage: %s", message.ErrorMessage)
    }
    return rc, nil
}
```

**Benefits:**

1. **Original Messages** - Users see the actual AlphaVantage error
2. **No Translation** - Don't hide important details
3. **Debuggable** - Errors can be traced to API issues
4. **Actionable** - Users know how to fix the problem

### HTTP Error Context

HTTP errors include status codes and response bodies:

```go
if res.StatusCode >= 300 {
    body, _ := io.ReadAll(io.LimitReader(res.Body, 1024))
    return nil, fmt.Errorf("request failed with status %d: %s",
        res.StatusCode, string(body))
}
```

This helps debug network issues and API problems.

### Parse Errors

CSV parsing errors include row and column information:

```go
return nil, fmt.Errorf("failed to parse float64 value %q on row %d column %d (%s): %w",
    value, rowIndex, columnIndex, headerName, err)
```

Users know exactly which data caused the parsing failure.

## Client Interface Design

### Dependency Injection

The client uses interfaces for HTTP client and rate limiter:

```go
type Client struct {
    Limiter interface {
        Wait(ctx context.Context) error
    }
    Client interface {
        Do(*http.Request) (*http.Response, error)
    }
    APIKey string
}
```

**Benefits:**

1. **Testability** - Mock HTTP responses and rate limiting in tests
2. **Customization** - Users can provide custom HTTP clients
3. **Flexibility** - Different rate limiting strategies are possible
4. **Composition** - Clients can be wrapped or decorated

### Constructor Convenience

The `NewClient` function provides sensible defaults:

```go
func NewClient(apiKey string, reqPerMin RequestsPerMinute) *Client {
    return &Client{
        Client:  http.DefaultClient,
        Limiter: rate.NewLimiter(reqPerMin.Limit(), 5),
        APIKey:  cmp.Or(apiKey, os.Getenv("ALPHA_VANTAGE_TOKEN"), "demo"),
    }
}
```

- Uses `http.DefaultClient` for HTTP
- Creates appropriate rate limiter
- Falls back to environment variable or "demo" key

Users can still create custom clients for advanced use cases.

## Data Flow Architecture

### Request Pipeline

1. **Query Construction** - User builds typed query with required/optional parameters
2. **URL Building** - Query converts to url.Values and constructs request URL
3. **Rate Limiting** - Client waits if necessary to respect API limits
4. **HTTP Request** - Execute request with context support
5. **Error Detection** - Check for HTTP and API errors
6. **Response Processing** - Parse CSV/JSON into typed structures
7. **Resource Cleanup** - Close response body

### Response Processing Options

The library offers three levels of response processing:

**1. Raw HTTP Response**
```go
resp, err := client.GetTimeSeriesDaily(ctx, query)
defer resp.Body.Close()
// Process resp.Body manually
```

**2. Automatic CSV Parsing (Recommended)**
```go
rows, err := client.GetTimeSeriesDailyCSVRows(ctx, query)
// rows is []TimeSeriesDailyRow - already parsed
```

**3. Streaming with Iterator**
```go
for row := range ParseCSVRows[TimeSeriesDailyRow](resp.Body, ...) {
    // Process row by row without loading all into memory
}
```

This flexibility allows users to choose the right level of abstraction for their use case.

## Testing Strategy

### Specification-Driven Testing

Tests use cached API responses stored in `specification/testdata/`:

```
specification/
├── functions/          # JSON specs for each API function
│   ├── time_series.json
│   ├── technical_*.json
│   └── ...
└── testdata/
    └── examples/      # Cached real API responses
        ├── time_series/
        │   ├── TIME_SERIES_DAILY_*.csv
        │   └── ...
        └── index.json # Maps queries to responses
```

**Benefits:**

1. **Fast Tests** - No actual API calls during testing
2. **Reliable** - Tests don't depend on network or API availability
3. **Realistic** - Tests use real API response data
4. **Comprehensive** - Can test edge cases and error conditions
5. **Regression Prevention** - Cached responses catch API changes

### Test Organization

```
.
├── *_test.go              # Unit tests with mocked responses
├── client_internal_test.go # Internal implementation tests
├── cmd/av/main_test.go    # CLI integration tests
└── specification/
    └── specification_test.go  # Specification validation tests
```

Tests are organized by what they test:
- **Unit tests** - Test individual functions with mocked data
- **Integration tests** - Test full request/response cycle
- **CLI tests** - Test command-line interface end-to-end
- **Specification tests** - Validate JSON specifications

### Test-Driven Development Workflow

1. **Red** - Write failing test with expected behavior
2. **Green** - Implement minimum code to make test pass
3. **Refactor** - Improve code quality while keeping tests green
4. **Document** - Update docs to reflect new functionality

This ensures code is testable and requirements are clear.

## CLI Design Philosophy

### Function-Based Command Structure

Version 0.4 changed from subcommands to function-based commands:

**v0.3 (subcommands):**
```bash
av quotes --function=TIME_SERIES_DAILY --symbol=IBM
av global-quote IBM
```

**v0.4 (functions):**
```bash
av TIME_SERIES_DAILY --symbol=IBM
av GLOBAL_QUOTE --symbol=IBM
```

**Why this change?**

1. **Consistency** - Matches API function names exactly
2. **Discoverability** - `av help` lists all 92 functions
3. **Extensibility** - New functions are automatically available
4. **Documentation** - Function help matches API documentation
5. **Simplicity** - No mapping between command names and API functions

### Generated CLI Handlers

CLI handlers are generated from the same specifications as client methods:

```go
func handleGLOBALQUOTE(args []string, token, output string) error {
    // Parse flags
    var symbol string
    flags.StringVar(&symbol, "symbol", "", "Stock symbol")
    flags.Parse(args)

    // Validate required parameters
    if symbol == "" {
        return fmt.Errorf("--symbol is required")
    }

    // Create client and query
    client := alphavantage.NewClient(token, alphavantage.FreePlan)
    query := alphavantage.QueryGlobalQuote(client.APIKey, symbol)

    // Execute and save results
    // ...
}
```

This ensures CLI and library APIs stay in sync.

### Flag Naming Conventions

All flags use lowercase with hyphens (kebab-case):

```bash
av TIME_SERIES_INTRADAY \
    --symbol=IBM \
    --interval=5min \
    --outputsize=compact \
    --extended-hours=true
```

This matches common CLI conventions (git, docker, etc.).

## Performance Considerations

### Memory Efficiency

**Streaming Responses:**
```go
// Bad: Loads entire dataset into memory
data, _ := io.ReadAll(resp.Body)
rows := parseCSV(data)

// Good: Processes line by line
for row := range ParseCSVRows[T](resp.Body, ...) {
    process(row)  // Memory used per row, not entire dataset
}
```

**CSV Processing:**
- Single-pass parsing
- Minimal allocations
- No intermediate data structures
- Reflection cached per type

### Network Efficiency

**Connection Reuse:**
```go
Client: http.DefaultClient  // Enables connection pooling
```

**Compression:**
HTTP client automatically handles gzip compression.

**Request Batching:**
Some endpoints support bulk operations:
```go
// Single request for multiple symbols
symbols := "AAPL,MSFT,GOOGL,AMZN"
query := alphavantage.QueryRealtimeBulkQuotes(apiKey, symbols)
```

### Code Generation Efficiency

Generated code is:
- **Consistent** - Same patterns everywhere, easier to optimize
- **Minimal** - No unnecessary abstractions
- **Inlineable** - Simple methods can be inlined by compiler
- **Type-Safe** - No reflection in generated code paths

## Security Considerations

### API Key Management

**Never log or expose API keys:**
```go
// API keys not included in error messages
fmt.Errorf("request failed: %s", errorMessage)  // No API key

// Environment variable support
APIKey: cmp.Or(apiKey, os.Getenv("ALPHA_VANTAGE_TOKEN"))
```

**Secure defaults:**
- HTTPS only (no HTTP fallback)
- No API key in URLs (query parameters, not path)
- Environment variables for key storage

### Input Validation

**Symbol validation:**
```go
if symbol == "" {
    return fmt.Errorf("symbol is required")
}
```

**URL encoding:**
```go
apiURL.RawQuery = values.Encode()  // Proper URL encoding
```

This prevents injection attacks and malformed requests.

### Dependency Security

Minimal dependencies:
- `golang.org/x/time/rate` - Official Go package
- `github.com/spf13/pflag` - Well-maintained, widely-used

All dependencies are regularly updated and security-scanned.

## Future Extensibility

### Adding New API Functions

1. Add JSON specification to `specification/functions/`
2. Run `go generate`
3. New functions automatically available in library and CLI

No hand-written code needed.

### Custom Response Parsing

Users can define custom structs for specialized parsing:

```go
type CustomRow struct {
    Date   time.Time `column-name:"timestamp"`
    Price  float64   `column-name:"close"`
    Custom string    `column-name:"custom_field"`
}

var rows []CustomRow
err := alphavantage.ParseCSV(resp.Body, &rows, time.UTC)
```

### API Version Changes

If AlphaVantage changes their API:

1. Update JSON specifications
2. Regenerate code
3. Add migration notes to docs/migration-guide.md
4. Tag new major version if breaking

### Alternative HTTP Clients

Users can provide custom HTTP clients for:
- Custom timeouts
- Proxy support
- Custom TLS configuration
- Request/response logging
- Metrics collection

```go
client := &alphavantage.Client{
    Client: myCustomHTTPClient,
    // ...
}
```

## Design Trade-offs

### Code Generation vs Hand-Written

**Pros:**
- Consistency across all 92 functions
- Easy to add new functions
- Less error-prone
- Systematic testing

**Cons:**
- More complex build process
- Generated code can be verbose
- Changes require regeneration

**Decision:** Benefits outweigh costs for a library with 92 similar functions.

### CSV vs JSON Default

**Pros:**
- More efficient for tabular data
- Faster parsing
- Lower memory usage
- Streaming support

**Cons:**
- Not suitable for nested data
- Manual handling for non-tabular responses

**Decision:** CSV for time series, JSON for complex structures.

### Query Builders vs Simple Functions

**Pros:**
- Type-safe optional parameters
- Fluent, readable API
- Compile-time validation
- Future extensibility

**Cons:**
- More verbose than simple functions
- Steeper learning curve

**Decision:** Type safety and extensibility worth the verbosity.

## Summary

The AlphaVantage Go client is designed around these principles:

1. **Type Safety** - Catch errors at compile time
2. **Code Generation** - Consistency and maintainability
3. **Performance** - Efficient CSV parsing and streaming
4. **Testability** - Specification-driven tests with cached responses
5. **Simplicity** - Sensible defaults, minimal configuration
6. **Extensibility** - Easy to add new functions and customize behavior

These design decisions result in a library that is:
- Easy to use for common cases
- Flexible for advanced use cases
- Maintainable as the API evolves
- Well-tested and reliable
