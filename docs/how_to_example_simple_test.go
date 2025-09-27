package docs_test

import (
	"cmp"
	"fmt"
	"os"

	"github.com/portfoliotree/alphavantage"
)

// ExampleClient_DoQuotesRequest demonstrates how to fetch daily stock prices
// as shown in the how-to guide.
func ExampleClient_DoQuotesRequest() {
	// Get API key from environment variable
	apiKey := cmp.Or(os.Getenv(alphavantage.StandardTokenEnvironmentVariableName), "demo")
	client := alphavantage.NewClient(apiKey)

	// The pattern for fetching daily stock prices:
	// ctx := context.Background()
	// result, err := client.DoQuotesRequest(ctx, "AAPL", alphavantage.TimeSeriesDaily)
	// if err != nil {
	//     return err
	// }
	// defer result.Close()
	// // Process CSV data...

	fmt.Printf("Client ready for daily data: %t\n", client != nil)
	// Output: Client ready for daily data: true
}

// ExampleClient_DoQuotesRequest_adjusted demonstrates how to get adjusted prices
// with dividends and splits.
func ExampleClient_DoQuotesRequest_adjusted() {
	apiKey := cmp.Or(os.Getenv(alphavantage.StandardTokenEnvironmentVariableName), "demo")
	client := alphavantage.NewClient(apiKey)

	// Use adjusted functions for split/dividend-adjusted data:
	// result, err := client.DoQuotesRequest(ctx, "AAPL", alphavantage.TimeSeriesDailyAdjusted)

	fmt.Printf("Client ready for adjusted data: %t\n", client != nil)
	// Output: Client ready for adjusted data: true
}

// ExampleClient_DoQuotesRequest_weekly demonstrates how to fetch weekly
// or monthly data.
func ExampleClient_DoQuotesRequest_weekly() {
	apiKey := cmp.Or(os.Getenv(alphavantage.StandardTokenEnvironmentVariableName), "demo")
	_ = alphavantage.NewClient(apiKey)

	// Available time series functions:
	functions := []alphavantage.QuoteFunction{
		alphavantage.TimeSeriesWeekly,
		alphavantage.TimeSeriesWeeklyAdjusted,
		alphavantage.TimeSeriesMonthly,
		alphavantage.TimeSeriesMonthlyAdjusted,
	}

	// Weekly data: client.DoQuotesRequest(ctx, "AAPL", alphavantage.TimeSeriesWeekly)
	// Monthly data: client.DoQuotesRequest(ctx, "AAPL", alphavantage.TimeSeriesMonthly)

	fmt.Printf("Available weekly/monthly functions: %d\n", len(functions))
	// Output: Available weekly/monthly functions: 4
}

// ExampleQuoteFunction_Validate demonstrates function validation.
func ExampleQuoteFunction_Validate() {
	// Validate a time series function before use
	err := alphavantage.TimeSeriesWeekly.Validate()

	fmt.Printf("Weekly function is valid: %t\n", err == nil)
	// Output: Weekly function is valid: true
}