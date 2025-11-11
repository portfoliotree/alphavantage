# How-to Guides

This document provides solutions to common tasks with the AlphaVantage Go client.

## Runnable Examples

For complete, runnable code examples, see the [examples directory](examples/). Each example is a standalone Go program demonstrating specific functionality:

- **Getting Started**: [examples/getting_started/](examples/getting_started/)
- **Stock Data**: [examples/stock_data/](examples/stock_data/)
- **Technical Indicators**: [examples/technical_indicators/](examples/technical_indicators/)
- **Economic Data**: [examples/economic_data/](examples/economic_data/)
- **Forex**: [examples/forex/](examples/forex/)
- **Cryptocurrency**: [examples/crypto/](examples/crypto/)
- **Commodities**: [examples/commodities/](examples/commodities/)
- **Fundamental Data**: [examples/fundamental_data/](examples/fundamental_data/)
- **Advanced Patterns**: [examples/advanced/](examples/advanced/)

See [examples/README.md](examples/README.md) for a complete catalog of all examples.

## Stock Data

### How to get a realtime stock quote

See [examples/getting_started/01_stock_quote.go](examples/getting_started/01_stock_quote.go)

### How to fetch daily stock prices

See [examples/getting_started/02_historical_prices.go](examples/getting_started/02_historical_prices.go)

### How to get adjusted prices (with dividends and splits)

```go
query := alphavantage.QueryTimeSeriesDailyAdjusted(client.APIKey, "MSFT")
rows, err := client.GetTimeSeriesDailyAdjustedCSVRows(ctx, query)
if err != nil {
    log.Fatal(err)
}

for _, row := range rows[:5] {
    fmt.Printf("%s: Adjusted Close=$%.2f Dividend=%.2f\n",
        row.Timestamp.Format("2006-01-02"),
        row.AdjustedClose,
        row.DividendAmount)
}
```

### How to fetch intraday data

See [examples/stock_data/02_intraday.go](examples/stock_data/02_intraday.go) and [examples/getting_started/03_query_builder.go](examples/getting_started/03_query_builder.go)

### How to get weekly or monthly data

See [examples/stock_data/03_weekly_monthly.go](examples/stock_data/03_weekly_monthly.go)

### How to search for stock symbols

```go
query := alphavantage.QuerySymbolSearch(client.APIKey, "International Business")
rows, err := client.GetSymbolSearchCSVRows(ctx, query)
if err != nil {
    log.Fatal(err)
}

for _, row := range rows {
    fmt.Printf("%s - %s (%s)\n", row.Symbol, row.Name, row.Type)
}
```

### How to get bulk quotes for multiple symbols

```go
// Up to 100 symbols comma-separated
symbols := "AAPL,MSFT,GOOGL,AMZN,TSLA"
query := alphavantage.QueryRealtimeBulkQuotes(client.APIKey, symbols)
resp, err := client.GetRealtimeBulkQuotes(ctx, query)
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

// Parse JSON response
var quotes []struct {
    Symbol string  `json:"symbol"`
    Price  float64 `json:"price"`
    Volume int64   `json:"volume"`
}
json.NewDecoder(resp.Body).Decode(&quotes)
```

## Fundamental Data

### How to get company overview

```go
query := alphavantage.QueryOverview(client.APIKey, "IBM")
resp, err := client.GetOverview(ctx, query)
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

var overview struct {
    Symbol             string `json:"Symbol"`
    Name               string `json:"Name"`
    Description        string `json:"Description"`
    Sector             string `json:"Sector"`
    Industry           string `json:"Industry"`
    MarketCapitalization string `json:"MarketCapitalization"`
    PERatio            string `json:"PERatio"`
    DividendYield      string `json:"DividendYield"`
}
err = json.NewDecoder(resp.Body).Decode(&overview)
fmt.Printf("%s - %s\nSector: %s\nP/E: %s\n",
    overview.Symbol, overview.Name, overview.Sector, overview.PERatio)
```

### How to get financial statements

```go
// Income Statement
query := alphavantage.QueryIncomeStatement(client.APIKey, "AAPL")
resp, err := client.GetIncomeStatement(ctx, query)

// Balance Sheet
query = alphavantage.QueryBalanceSheet(client.APIKey, "AAPL")
resp, err = client.GetBalanceSheet(ctx, query)

// Cash Flow
query = alphavantage.QueryCashFlow(client.APIKey, "AAPL")
resp, err = client.GetCashFlow(ctx, query)

// All return JSON - parse with json.Decoder
```

### How to get earnings data

```go
// Historical earnings
query := alphavantage.QueryEarnings(client.APIKey, "MSFT")
resp, err := client.GetEarnings(ctx, query)

// Analyst estimates
query = alphavantage.QueryEarningsEstimates(client.APIKey, "MSFT")
resp, err = client.GetEarningsEstimates(ctx, query)
```

### How to get dividend and split history

```go
// Dividends
query := alphavantage.QueryDividends(client.APIKey, "JNJ")
rows, err := client.GetDividendsCSVRows(ctx, query)
for _, row := range rows {
    fmt.Printf("%s: $%s\n", row.ExDate, row.Amount)
}

// Stock splits
query = alphavantage.QuerySplits(client.APIKey, "AAPL")
rows, err = client.GetSplitsCSVRows(ctx, query)
for _, row := range rows {
    fmt.Printf("%s: %s\n", row.EffectiveDate, row.SplitFactor)
}
```

### How to check listing status

```go
// Active listings
query := alphavantage.QueryListingStatus(client.APIKey).StateActive()
rows, err := client.GetListingStatusCSVRows(ctx, query)

// Delisted companies
query = alphavantage.QueryListingStatus(client.APIKey).StateDelisted()
rows, err = client.GetListingStatusCSVRows(ctx, query)

// Specific date
query = alphavantage.QueryListingStatus(client.APIKey).Date("2024-01-15")
rows, err = client.GetListingStatusCSVRows(ctx, query)
```

### How to get upcoming earnings

```go
// All upcoming earnings
query := alphavantage.QueryEarningsCalendar(client.APIKey)
rows, err := client.GetEarningsCalendarCSVRows(ctx, query)

// For specific symbol
query = alphavantage.QueryEarningsCalendar(client.APIKey).Symbol("AAPL")
rows, err = client.GetEarningsCalendarCSVRows(ctx, query)

// Within time horizon
query = alphavantage.QueryEarningsCalendar(client.APIKey).Horizon("3month")
rows, err = client.GetEarningsCalendarCSVRows(ctx, query)
```

### How to get ETF profile

```go
query := alphavantage.QueryETFProfile(client.APIKey, "QQQ")
resp, err := client.GetETFProfile(ctx, query)
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

var profile struct {
    Symbol string `json:"symbol"`
    Name   string `json:"name"`
    AssetClass string `json:"asset_class"`
    Holdings []struct {
        Symbol string `json:"symbol"`
        Name   string `json:"name"`
        Weight string `json:"weight"`
    } `json:"holdings"`
}
json.NewDecoder(resp.Body).Decode(&profile)
```

## Technical Indicators

### How to calculate moving averages

See [examples/technical_indicators/02_sma.go](examples/technical_indicators/02_sma.go)

Available moving average indicators: SMA, EMA, WMA, DEMA, TEMA, TRIMA, KAMA, MAMA, T3, VWAP

### How to calculate RSI

See [examples/technical_indicators/01_rsi.go](examples/technical_indicators/01_rsi.go)

### How to calculate MACD

See [examples/technical_indicators/03_macd.go](examples/technical_indicators/03_macd.go)

### How to calculate Bollinger Bands

```go
query := alphavantage.QueryBBANDS(client.APIKey, "GOOGL", "daily", 20, "close")
rows, err := client.GetBBANDSCSVRows(ctx, query)

for _, row := range rows[:5] {
    fmt.Printf("%s: Upper=%.2f Middle=%.2f Lower=%.2f\n",
        row.Time, row.RealUpperBand, row.RealMiddleBand, row.RealLowerBand)
}
```

### How to calculate Stochastic Oscillator

```go
query := alphavantage.QuerySTOCH(client.APIKey, "TSLA", "daily")
rows, err := client.GetSTOCHCSVRows(ctx, query)

for _, row := range rows[:10] {
    fmt.Printf("%s: %%K=%.2f %%D=%.2f\n", row.Time, row.SlowK, row.SlowD)
}
```

### How to use custom indicator parameters

```go
// MACD with custom parameters
query := alphavantage.QueryMACDEXT(client.APIKey, "AAPL", "daily", "close").
    FastPeriod(8).
    SlowPeriod(21).
    SignalPeriod(5).
    FastMAType(0)  // SMA=0, EMA=1, etc.
rows, err := client.GetMACDEXTCSVRows(ctx, query)

// Stochastic with custom settings
query = alphavantage.QuerySTOCH(client.APIKey, "IBM", "daily").
    FastKPeriod(5).
    SlowKPeriod(3).
    SlowDPeriod(3).
    SlowKMAType(0).
    SlowDMAType(0)
rows, err = client.GetSTOCHCSVRows(ctx, query)
```

## Economic Data

### How to get GDP data

```go
// Real GDP
query := alphavantage.QueryRealGDP(client.APIKey)
rows, err := client.GetRealGDPCSVRows(ctx, query)

// Quarterly data
query = query.IntervalQuarterly()
rows, err = client.GetRealGDPCSVRows(ctx, query)

// GDP per capita
query = alphavantage.QueryRealGDPPerCapita(client.APIKey)
rows, err = client.GetRealGDPPerCapitaCSVRows(ctx, query)
```

### How to get unemployment data

```go
query := alphavantage.QueryUnemployment(client.APIKey)
rows, err := client.GetUnemploymentCSVRows(ctx, query)
if err != nil {
    log.Fatal(err)
}

for _, row := range rows[:12] {
    fmt.Printf("%s: %.1f%%\n", row.Date, row.Value)
}
```

### How to get inflation data

```go
// Consumer Price Index
query := alphavantage.QueryCPI(client.APIKey)
rows, err := client.GetCPICSVRows(ctx, query)

// With specific interval
query = query.IntervalMonthly()
rows, err = client.GetCPICSVRows(ctx, query)

// Inflation rate
query = alphavantage.QueryInflation(client.APIKey)
rows, err = client.GetInflationCSVRows(ctx, query)
```

### How to get interest rates

```go
// Federal Funds Rate
query := alphavantage.QueryFederalFundsRate(client.APIKey)
rows, err := client.GetFederalFundsRateCSVRows(ctx, query)

// With interval
query = query.IntervalDaily()
rows, err = client.GetFederalFundsRateCSVRows(ctx, query)

// Treasury Yields
query = alphavantage.QueryTreasuryYield(client.APIKey)
query = query.Maturity("10year").IntervalMonthly()
rows, err = client.GetTreasuryYieldCSVRows(ctx, query)
```

### How to get retail and manufacturing data

```go
// Retail Sales
query := alphavantage.QueryRetailSales(client.APIKey)
rows, err := client.GetRetailSalesCSVRows(ctx, query)

// Durable Goods Orders
query = alphavantage.QueryDurables(client.APIKey)
rows, err = client.GetDurablesCSVRows(ctx, query)

// Nonfarm Payroll
query = alphavantage.QueryNonfarmPayroll(client.APIKey)
rows, err = client.GetNonfarmPayrollCSVRows(ctx, query)
```

## Forex Data

### How to get currency exchange rates

See [examples/forex/01_exchange_rate.go](examples/forex/01_exchange_rate.go)

### How to get historical forex data

See [examples/forex/02_fx_daily.go](examples/forex/02_fx_daily.go)

Available functions: `QueryFXDaily`, `QueryFXIntraday`, `QueryFXWeekly`, `QueryFXMonthly`

## Cryptocurrency Data

### How to get crypto prices

```go
// Intraday crypto
query := alphavantage.QueryCryptoIntraday(client.APIKey, "BTC", "USD", "5min")
rows, err := client.GetCryptoIntradayCSVRows(ctx, query)

// Daily crypto
query = alphavantage.QueryDigitalCurrencyDaily(client.APIKey, "ETH", "USD")
rows, err = client.GetDigitalCurrencyDailyCSVRows(ctx, query)

// Weekly crypto
query = alphavantage.QueryDigitalCurrencyWeekly(client.APIKey, "BTC", "EUR")
rows, err = client.GetDigitalCurrencyWeeklyCSVRows(ctx, query)

// Monthly crypto
query = alphavantage.QueryDigitalCurrencyMonthly(client.APIKey, "LTC", "USD")
rows, err = client.GetDigitalCurrencyMonthlyCSVRows(ctx, query)
```

## Commodities Data

### How to get oil prices

```go
// WTI Crude Oil
query := alphavantage.QueryWTI(client.APIKey)
rows, err := client.GetWTICSVRows(ctx, query)

// With interval
query = query.IntervalMonthly()
rows, err = client.GetWTICSVRows(ctx, query)

// Brent Crude Oil
query = alphavantage.QueryBrent(client.APIKey).IntervalWeekly()
rows, err = client.GetBrentCSVRows(ctx, query)
```

### How to get natural gas prices

```go
query := alphavantage.QueryNaturalGas(client.APIKey)
rows, err := client.GetNaturalGasCSVRows(ctx, query)

// Daily data
query = query.IntervalDaily()
rows, err = client.GetNaturalGasCSVRows(ctx, query)
```

### How to get metals and agriculture prices

```go
// Copper
query := alphavantage.QueryCopper(client.APIKey).IntervalMonthly()
rows, err := client.GetCopperCSVRows(ctx, query)

// Aluminum
query = alphavantage.QueryAluminum(client.APIKey).IntervalQuarterly()
rows, err = client.GetAluminumCSVRows(ctx, query)

// Wheat
query = alphavantage.QueryWheat(client.APIKey).IntervalMonthly()
rows, err = client.GetWheatCSVRows(ctx, query)

// Corn
query = alphavantage.QueryCorn(client.APIKey)
rows, err = client.GetCornCSVRows(ctx, query)

// Coffee, Cotton, Sugar also available
```

### How to get commodity index

```go
query := alphavantage.QueryAllCommodities(client.APIKey)
rows, err := client.GetAllCommoditiesCSVRows(ctx, query)

// With interval
query = query.IntervalQuarterly()
rows, err = client.GetAllCommoditiesCSVRows(ctx, query)
```

## Intelligence & Analytics

### How to get news with sentiment

```go
query := alphavantage.QueryNewsSentiment(client.APIKey)

// Filter by tickers
query = query.Tickers("AAPL,MSFT")

// Filter by topics
query = query.Topics("technology,earnings")

// Limit results
query = query.Limit(50)

// Time range
query = query.TimeFrom("20240101T0000").TimeTo("20240131T2359")

// Sort order
query = query.SortRelevance()  // or SortLatest(), SortEarliest()

resp, err := client.GetNewsSentiment(ctx, query)
// Returns JSON with news articles and sentiment scores
```

### How to get top gainers and losers

```go
query := alphavantage.QueryTopGainersLosers(client.APIKey)
resp, err := client.GetTopGainersLosers(ctx, query)
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

var data struct {
    TopGainers []struct {
        Ticker        string `json:"ticker"`
        Price         string `json:"price"`
        ChangePercent string `json:"change_percentage"`
    } `json:"top_gainers"`
    TopLosers []struct {
        Ticker        string `json:"ticker"`
        Price         string `json:"price"`
        ChangePercent string `json:"change_percentage"`
    } `json:"top_losers"`
    MostActivelyTraded []struct {
        Ticker string `json:"ticker"`
        Price  string `json:"price"`
        Volume string `json:"volume"`
    } `json:"most_actively_traded"`
}
json.NewDecoder(resp.Body).Decode(&data)
```

### How to get insider transactions

```go
query := alphavantage.QueryInsiderTransactions(client.APIKey, "AAPL")
resp, err := client.GetInsiderTransactions(ctx, query)
// Returns JSON with insider trading activity
```

### How to get earnings call transcripts

```go
query := alphavantage.QueryEarningsCallTranscript(client.APIKey, "IBM", "2024Q1")
resp, err := client.GetEarningsCallTranscript(ctx, query)
// Returns JSON with transcript text
```

## Options Data

### How to get realtime options chain

```go
query := alphavantage.QueryRealtimeOptions(client.APIKey, "AAPL")

// With contract filter
query = query.Contract("AAPL250117C00150000")

// Include Greeks and IV
query = query.RequireGreeks(true)

resp, err := client.GetRealtimeOptions(ctx, query)
// Returns JSON with options data
```

### How to get historical options

```go
query := alphavantage.QueryHistoricalOptions(client.APIKey, "MSFT")

// Specific date
query = query.Date("2024-01-15")

resp, err := client.GetHistoricalOptions(ctx, query)
// Returns CSV with historical options chain
```

## Advanced Usage

### How to handle API rate limits

See [examples/advanced/04_rate_limit_plans.go](examples/advanced/04_rate_limit_plans.go)

The client automatically handles rate limiting. Available plans: `FreePlan`, `PremiumPlan15`, `PremiumPlan30`, `PremiumPlan75`, `PremiumPlan120`, `PremiumPlan300`, `PremiumPlan600`, `PremiumPlan1200`

### How to use a custom HTTP client

See [examples/advanced/03_custom_http_client.go](examples/advanced/03_custom_http_client.go)

### How to add request timeouts

See [examples/advanced/02_context_timeout.go](examples/advanced/02_context_timeout.go)

### How to stream large datasets

See [examples/advanced/01_streaming.go](examples/advanced/01_streaming.go)

### How to parse custom CSV structures

```go
// Define custom struct
type MyCustomRow struct {
    Date   time.Time `column-name:"timestamp"`
    Price  float64   `column-name:"close"`
    Volume int       `column-name:"volume"`
}

// Parse CSV into custom type
query := alphavantage.QueryTimeSeriesDaily(client.APIKey, "IBM")
resp, err := client.GetTimeSeriesDaily(ctx, query.DataTypeCSV())
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

var rows []MyCustomRow
err = alphavantage.ParseCSV(resp.Body, &rows, time.UTC)
```

### How to use the CLI for automation

```bash
#!/bin/bash

# Fetch daily data for multiple symbols
for symbol in AAPL MSFT GOOGL AMZN; do
    av TIME_SERIES_DAILY --symbol=$symbol --outputsize=compact
    sleep 15  # Respect rate limits (5 per minute)
done

# Process the results
for file in *.csv; do
    echo "Processing $file"
    # Your data processing here
done
```

### How to check market status

```go
query := alphavantage.QueryMarketStatus(client.APIKey)
resp, err := client.GetMarketStatus(ctx, query)
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

var status struct {
    Markets []struct {
        MarketType   string `json:"market_type"`
        Region       string `json:"region"`
        PrimaryExchanges string `json:"primary_exchanges"`
        CurrentStatus string `json:"current_status"`
    } `json:"markets"`
}
json.NewDecoder(resp.Body).Decode(&status)

for _, market := range status.Markets {
    fmt.Printf("%s (%s): %s\n",
        market.MarketType, market.Region, market.CurrentStatus)
}
```

## Troubleshooting

### API returns error messages

AlphaVantage error messages are automatically parsed and returned as Go errors:

```go
rows, err := client.GetTimeSeriesDailyCSVRows(ctx, query)
if err != nil {
    // Error will contain AlphaVantage's error message
    log.Printf("API Error: %v", err)
    return
}
```

### Handling rate limit errors

If you exceed rate limits, the client will wait automatically. If you see errors:

```go
// Make sure you specified the correct plan
client := alphavantage.NewClient(apiKey, alphavantage.FreePlan)

// Add a longer timeout if needed
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
defer cancel()
```

### Working with different data formats

```go
// Some endpoints only support JSON
query := alphavantage.QueryOverview(client.APIKey, "IBM")
resp, err := client.GetOverview(ctx, query)
// Always JSON, parse accordingly

// Most endpoints support both
query = alphavantage.QueryGlobalQuote(client.APIKey, "IBM")
query = query.DataTypeCSV()   // Force CSV (automatic parsing)
query = query.DataTypeJSON()  // Force JSON (manual parsing)
```
