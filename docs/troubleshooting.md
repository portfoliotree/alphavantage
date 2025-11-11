# Troubleshooting

## "rate limit exceeded" error

You're making requests too fast. The client should handle this automatically. If you see this error:
- Verify you specified the correct rate plan in `NewClient()`
- Check if you're creating multiple clients (each needs its own limiter)

## "invalid API key" error

- Verify your API key is correct
- Make sure `ALPHA_VANTAGE_TOKEN` environment variable is set
- Check your API key is active at https://www.alphavantage.co

## Empty results

- Verify the symbol exists using `SYMBOL_SEARCH`
- Check if you're requesting data outside market hours for realtime data
- Some symbols may not have full historical data available

## Parse errors

- Ensure you're using the correct data type (CSV vs JSON)
- Check the API documentation for the expected response format
- Some endpoints return JSON only (like `OVERVIEW`)
