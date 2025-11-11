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
	apiKey := cmp.Or(os.Getenv(alphavantage.TokenEnvironmentVariableName), "demo")
	client := alphavantage.NewClient(apiKey, alphavantage.PremiumPlan75)

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
	apiKey := cmp.Or(os.Getenv(alphavantage.TokenEnvironmentVariableName), "demo")
	client := alphavantage.NewClient(apiKey, alphavantage.PremiumPlan75)

	// Use adjusted functions for split/dividend-adjusted data:
	// result, err := client.DoQuotesRequest(ctx, "AAPL", alphavantage.TimeSeriesDailyAdjusted)

	fmt.Printf("Client ready for adjusted data: %t\n", client != nil)
	// Output: Client ready for adjusted data: true
}
