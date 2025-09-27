package alphavantage_test

import (
	"cmp"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/portfoliotree/alphavantage"
)

// StockPrice represents a daily stock price with struct tags for CSV parsing.
// Supported field types: string, int, float64, time.Time
type StockPrice struct {
	Date   time.Time `column-name:"timestamp"`
	Open   float64   `column-name:"open"`
	High   float64   `column-name:"high"`
	Low    float64   `column-name:"low"`
	Close  float64   `column-name:"close"`
	Volume int       `column-name:"volume"`
}

// SearchResult represents a symbol search result with struct tags.
type SearchResult struct {
	Symbol      string  `column-name:"symbol"`
	Name        string  `column-name:"name"`
	Type        string  `column-name:"type"`
	Region      string  `column-name:"region"`
	MarketOpen  string  `column-name:"marketOpen"`
	MarketClose string  `column-name:"marketClose"`
	Timezone    string  `column-name:"timezone"`
	Currency    string  `column-name:"currency"`
	MatchScore  float64 `column-name:"matchScore"`
}

// ExampleParseCSV demonstrates parsing CSV data into a slice of structs.
// The struct fields must be tagged with "column-name" matching the CSV headers.
func ExampleParseCSV() {
	// Sample CSV data (truncated from testdata)
	csvData := `timestamp,open,high,low,close,volume
2020-08-21,13.2600,13.3200,13.1500,13.2500,751279
2020-08-20,13.4700,13.4750,13.2200,13.3800,854559
2020-08-19,13.5700,13.7100,13.4700,13.5000,521089`

	var prices []StockPrice
	err := alphavantage.ParseCSV(strings.NewReader(csvData), &prices, time.UTC)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Parsed %d prices\n", len(prices))
	fmt.Printf("First price: %s - Close: %.2f, Volume: %d\n",
		prices[0].Date.Format("2006-01-02"), prices[0].Close, prices[0].Volume)

	// Output: Parsed 3 prices
	// First price: 2020-08-21 - Close: 13.25, Volume: 751279
}

// ExampleParseCSVRows demonstrates streaming CSV parsing using an iterator.
// This is memory-efficient for large datasets as it processes one row at a time.
func ExampleParseCSVRows() {
	csvData := `symbol,name,type,region,marketOpen,marketClose,timezone,currency,matchScore
BA,Boeing Company,Equity,United States,09:30,16:00,UTC-04,USD,1.0000
BAB,Invesco Taxable Municipal Bond ETF,ETF,United States,09:30,16:00,UTC-04,USD,0.8000`

	count := 0
	for result := range alphavantage.ParseCSVRows[SearchResult](strings.NewReader(csvData), time.UTC, func(err error) bool {
		fmt.Printf("Parse error: %v\n", err)
		return false // stop on error
	}) {
		count++
		fmt.Printf("Symbol: %s, Match: %.1f\n", result.Symbol, result.MatchScore)
	}

	fmt.Printf("Processed %d results\n", count)

	// Output: Symbol: BA, Match: 1.0
	// Symbol: BAB, Match: 0.8
	// Processed 2 results
}

// ExampleClient_DoQuotesRequest_parseCSV demonstrates fetching real data and parsing it.
func ExampleClient_DoQuotesRequest_parseCSV() {
	apiKey := cmp.Or(os.Getenv(alphavantage.StandardTokenEnvironmentVariableName), "demo")
	client := alphavantage.NewClient(apiKey)

	// This example shows the pattern but doesn't make a real API call
	// ctx := context.Background()
	// result, err := client.DoQuotesRequest(ctx, "IBM", alphavantage.TimeSeriesDaily)
	// if err != nil {
	//     fmt.Printf("Error: %v\n", err)
	//     return
	// }
	// defer result.Close()
	//
	// var prices []StockPrice
	// err = alphavantage.ParseCSV(result, &prices, time.UTC)
	// if err != nil {
	//     fmt.Printf("Parse error: %v\n", err)
	//     return
	// }
	//
	// fmt.Printf("Fetched %d daily prices for IBM\n", len(prices))

	fmt.Printf("Client ready for CSV parsing: %t\n", client != nil)
	// Output: Client ready for CSV parsing: true
}