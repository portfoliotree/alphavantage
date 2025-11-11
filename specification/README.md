# API Specification

This directory contains JSON specifications for all AlphaVantage API functions. These specifications drive the code generation for both the Go client library and CLI tool.

## Complete API Coverage

This library supports **all 92 AlphaVantage API functions** organized into the following categories:

### Core Stock APIs (11 functions)
- `GLOBAL_QUOTE` - Latest price and volume information
- `MARKET_STATUS` - Global market status
- `REALTIME_BULK_QUOTES` - Bulk quotes for up to 100 symbols
- `SYMBOL_SEARCH` - Search for stock symbols
- `TIME_SERIES_DAILY` - Daily time series
- `TIME_SERIES_DAILY_ADJUSTED` - Daily adjusted for splits & dividends
- `TIME_SERIES_INTRADAY` - Intraday time series (1min, 5min, 15min, 30min, 60min)
- `TIME_SERIES_MONTHLY` - Monthly time series
- `TIME_SERIES_MONTHLY_ADJUSTED` - Monthly adjusted
- `TIME_SERIES_WEEKLY` - Weekly time series
- `TIME_SERIES_WEEKLY_ADJUSTED` - Weekly adjusted

### Fundamental Data (12 functions)
- `BALANCE_SHEET` - Annual & quarterly balance sheets
- `CASH_FLOW` - Annual & quarterly cash flow
- `DIVIDENDS` - Dividend history
- `EARNINGS` - Quarterly & annual EPS
- `EARNINGS_CALENDAR` - Upcoming earnings
- `EARNINGS_ESTIMATES` - Analyst earnings estimates
- `ETF_PROFILE` - ETF holdings and allocations
- `INCOME_STATEMENT` - Annual & quarterly income statements
- `IPO_CALENDAR` - Upcoming IPOs
- `LISTING_STATUS` - Listing & delisting status
- `OVERVIEW` - Company information and financials
- `SPLITS` - Stock split history

### Technical Indicators (56 functions)

**Moving Averages (10):**
- `DEMA` - Double Exponential Moving Average
- `EMA` - Exponential Moving Average
- `KAMA` - Kaufman Adaptive Moving Average
- `MAMA` - MESA Adaptive Moving Average
- `SMA` - Simple Moving Average
- `T3` - Triple Exponential Moving Average
- `TEMA` - Triple Exponential Moving Average
- `TRIMA` - Triangular Moving Average
- `VWAP` - Volume Weighted Average Price
- `WMA` - Weighted Moving Average

**Oscillators & Momentum (13):**
- `ADX` - Average Directional Movement Index
- `ADXR` - Average Directional Movement Index Rating
- `APO` - Absolute Price Oscillator
- `AROON` - Aroon
- `AROONOSC` - Aroon Oscillator
- `BOP` - Balance Of Power
- `CCI` - Commodity Channel Index
- `CMO` - Chande Momentum Oscillator
- `DX` - Directional Movement Index
- `MACD` - Moving Average Convergence/Divergence
- `MACDEXT` - MACD with controllable MA type
- `MOM` - Momentum
- `PPO` - Percentage Price Oscillator

**Stochastic (4):**
- `RSI` - Relative Strength Index
- `STOCH` - Stochastic
- `STOCHF` - Stochastic Fast
- `STOCHRSI` - Stochastic Relative Strength Index

**Directional (4):**
- `MINUS_DI` - Minus Directional Indicator
- `MINUS_DM` - Minus Directional Movement
- `PLUS_DI` - Plus Directional Indicator
- `PLUS_DM` - Plus Directional Movement

**Volatility (5):**
- `ATR` - Average True Range
- `BBANDS` - Bollinger Bands
- `NATR` - Normalized Average True Range
- `SAR` - Parabolic SAR
- `TRANGE` - True Range

**Volume (5):**
- `AD` - Chaikin A/D Line
- `ADOSC` - Chaikin A/D Oscillator
- `MFI` - Money Flow Index
- `OBV` - On Balance Volume
- `TRIX` - 1-day Rate-Of-Change of a Triple Smooth EMA

**Price (2):**
- `MIDPOINT` - Midpoint over period
- `MIDPRICE` - Midpoint Price over period

**Hilbert Transform (6):**
- `HT_DCPERIOD` - Hilbert Transform - Dominant Cycle Period
- `HT_DCPHASE` - Hilbert Transform - Dominant Cycle Phase
- `HT_PHASOR` - Hilbert Transform - Phasor Components
- `HT_SINE` - Hilbert Transform - SineWave
- `HT_TRENDLINE` - Hilbert Transform - Instantaneous Trendline
- `HT_TRENDMODE` - Hilbert Transform - Trend vs Cycle Mode

**Other (7):**
- `ROC` - Rate of change
- `ROCR` - Rate of change Ratio
- `ULTOSC` - Ultimate Oscillator
- `WILLR` - Williams' %R

### Economic Indicators (10 functions)
- `CPI` - Consumer Price Index
- `DURABLES` - Durable goods orders
- `FEDERAL_FUNDS_RATE` - Federal funds rate
- `INFLATION` - Inflation rates
- `NONFARM_PAYROLL` - Non-farm payroll
- `REAL_GDP` - US Real GDP
- `REAL_GDP_PER_CAPITA` - GDP per capita
- `RETAIL_SALES` - Retail sales
- `TREASURY_YIELD` - Treasury yields
- `UNEMPLOYMENT` - Unemployment rate

### Forex (5 functions)
- `CURRENCY_EXCHANGE_RATE` - Real-time exchange rates
- `FX_DAILY` - Daily forex rates
- `FX_INTRADAY` - Intraday forex rates
- `FX_MONTHLY` - Monthly forex rates
- `FX_WEEKLY` - Weekly forex rates

### Cryptocurrency (4 functions)
- `CRYPTO_INTRADAY` - Intraday crypto prices
- `DIGITAL_CURRENCY_DAILY` - Daily crypto prices
- `DIGITAL_CURRENCY_MONTHLY` - Monthly crypto prices
- `DIGITAL_CURRENCY_WEEKLY` - Weekly crypto prices

### Commodities (10 functions)
- `ALUMINUM` - Aluminum prices
- `ALL_COMMODITIES` - Global commodity index
- `BRENT` - Brent crude oil
- `COFFEE` - Coffee prices
- `COPPER` - Copper prices
- `CORN` - Corn prices
- `COTTON` - Cotton prices
- `NATURAL_GAS` - Natural gas prices
- `SUGAR` - Sugar prices
- `WHEAT` - Wheat prices
- `WTI` - West Texas Intermediate crude oil

### Intelligence & Analytics (6 functions)
- `ANALYTICS_FIXED_WINDOW` - Time series analytics (fixed window)
- `ANALYTICS_SLIDING_WINDOW` - Time series analytics (sliding window)
- `EARNINGS_CALL_TRANSCRIPT` - Earnings call transcripts
- `INSIDER_TRANSACTIONS` - Insider trading data
- `NEWS_SENTIMENT` - News with sentiment analysis
- `TOP_GAINERS_LOSERS` - Market movers

### Options (2 functions)
- `HISTORICAL_OPTIONS` - Historical options data
- `REALTIME_OPTIONS` - Real-time options chain

## Specification Structure

The specifications are organized as follows:

- **`query_parameters.json`** - Shared query parameters with single source of truth
- **`functions/*.json`** - Function specifications organized by category
- **`identifiers.json`** - Maps API strings to Go identifiers (public and private)
- **`testdata/examples/`** - Cached API responses for testing

### Files by Category

- `functions/time_series.json` - Core stock APIs
- `functions/fundamental.json` - Fundamental data
- `functions/technical_*.json` - Technical indicators (8 files)
- `functions/economic.json` - Economic indicators
- `functions/forex.json` - Forex data
- `functions/crypto.json` - Cryptocurrency
- `functions/commodities.json` - Commodities
- `functions/intelligence.json` - Intelligence & analytics
- `functions/options.json` - Options data

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
