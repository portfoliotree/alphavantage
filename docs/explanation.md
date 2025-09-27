# Understanding AlphaVantage API Integration

This document explains the design decisions and concepts behind the AlphaVantage Go client.

## Architecture Overview

### Why CSV over JSON?

The AlphaVantage API supports both JSON and CSV output formats. This client primarily uses CSV because:

1. **Efficiency**: CSV data is more compact than JSON for large time series datasets
2. **Streaming**: CSV can be processed line-by-line without loading entire datasets into memory
3. **Parsing Speed**: CSV parsing is generally faster than JSON for tabular data
4. **Memory Usage**: Lower memory footprint for large historical datasets

### Rate Limiting Design

The client implements automatic rate limiting to respect AlphaVantage's API limits:

- **Free Tier**: 5 requests per minute
- **Automatic Backoff**: Uses `golang.org/x/time/rate` for precise timing
- **Context Support**: Respects context cancellation during rate limiting waits

```go
// Default rate limiter
rate.NewLimiter(rate.Every(time.Minute/5), 5)
```

### Error Handling Philosophy

The client follows Go's explicit error handling patterns while providing meaningful error messages:

1. **API Error Propagation**: Original AlphaVantage error messages are preserved and returned
2. **HTTP Error Context**: HTTP status codes and response bodies are included in errors
3. **Validation**: Input validation happens client-side to catch errors early

### Client Interface Design

The client uses dependency injection for HTTP client and rate limiter:

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

This design enables:
- **Testing**: Easy mocking of HTTP requests and rate limiting
- **Customization**: Users can provide custom HTTP clients or rate limiters
- **Flexibility**: Support for different rate limiting strategies

## Data Flow

### Request Pipeline

1. **Input Validation**: Validate symbols and function parameters
2. **URL Construction**: Build API URL with proper query parameters
3. **Rate Limiting**: Wait if necessary to respect API limits
4. **HTTP Request**: Execute HTTP request with context support
5. **Error Checking**: Parse response for API errors (JSON error responses)
6. **Stream Return**: Return io.ReadCloser for efficient data processing

### Response Processing

The client returns `io.ReadCloser` interfaces to allow:
- **Streaming**: Process large datasets without loading everything into memory
- **Flexible Processing**: Users can choose CSV parsing, direct file writing, or custom processing
- **Resource Management**: Proper cleanup with defer statements

## Testing Strategy

### Test-Driven Development

The project follows TDD principles:

1. **Red**: Write failing tests that specify desired behavior
2. **Green**: Implement minimal code to make tests pass
3. **Refactor**: Improve code quality while maintaining test coverage

### Test Types

- **Unit Tests**: Fast tests with mocked HTTP responses
- **Integration Tests**: Real API calls with rate limiting (requires API key)
- **CLI Tests**: End-to-end testing of command-line interface

### Rate-Limited Integration Testing

Integration tests respect API rate limits to avoid overwhelming the service:

```go
func waitForRateLimit() {
    const minInterval = 12 * time.Second // 5 requests per minute
    if time.Since(lastAPICall) < minInterval {
        time.Sleep(minInterval - time.Since(lastAPICall))
    }
    lastAPICall = time.Now()
}
```

## CLI Design Philosophy

### POSIX Compliance

The CLI uses `pflag` for POSIX-compliant flag parsing:
- Supports both short (`-h`) and long (`--help`) flags
- Standard GNU-style flag behavior
- Better integration with shell completion

### File Output Strategy

Commands save data to files with predictable naming:
- `global-quote IBM` → `IBM_quote.csv`
- `quotes IBM` → `IBM.csv`
- `listing-status --listed=false` → `status_listed_false.csv`

This approach:
- Avoids output collisions
- Makes automation easier
- Provides clear file organization

## Performance Considerations

### Memory Efficiency

- **Streaming Responses**: Use io.ReadCloser to avoid loading large datasets into memory
- **CSV Processing**: Line-by-line processing for large time series data
- **Minimal Allocations**: Reuse buffers and minimize object creation

### Network Efficiency

- **Connection Reuse**: Default HTTP client enables connection pooling
- **Compression**: HTTP client automatically handles gzip compression
- **Request Batching**: CLI supports multiple symbols in single command

## Security Considerations

### API Key Management

- **Environment Variables**: Preferred method for API key storage
- **No Logging**: API keys are never logged or printed in error messages
- **Secure Transport**: All requests use HTTPS

### Input Validation

- **Symbol Validation**: Basic validation of stock symbols
- **Function Validation**: Enum-based validation for API functions
- **URL Construction**: Proper URL encoding to prevent injection attacks

## Future Extensibility

The current architecture supports adding new endpoints through:

1. **New Function Constants**: Add to QuoteFunction enum
2. **New Client Methods**: Follow existing patterns for consistency
3. **New CLI Commands**: Extend switch statement in main function
4. **New Data Types**: Add structs for JSON responses (like CompanyOverview)