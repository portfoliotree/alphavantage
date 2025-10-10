package alphavantage_test

import (
	"cmp"
	"fmt"
	"os"

	"github.com/portfoliotree/alphavantage"
)

// ExampleClient demonstrates how to create a new AlphaVantage client.
func ExampleClient() {
	// Get API key from environment variable
	apiKey := cmp.Or(os.Getenv(alphavantage.StandardTokenEnvironmentVariableName), "demo")

	client := alphavantage.NewClient(apiKey)
	fmt.Printf("Client created: %t\n", client != nil)
	// Output: Client created: true
}

// ExampleQuoteFunction_Validate demonstrates how to validate QuoteFunction constants.
func ExampleQuoteFunction_Validate() {
	// Valid function
	err := alphavantage.TimeSeriesDaily.Validate()
	fmt.Printf("Daily function is valid: %t\n", err == nil)

	// Valid weekly function
	err = alphavantage.TimeSeriesWeekly.Validate()
	fmt.Printf("Weekly function is valid: %t\n", err == nil)

	// Invalid function
	invalidFunction := alphavantage.QuoteFunction("INVALID_FUNCTION")
	err = invalidFunction.Validate()
	fmt.Printf("Invalid function rejected: %t\n", err != nil)

	// Output: Daily function is valid: true
	// Weekly function is valid: true
	// Invalid function rejected: true
}

// ExampleNewQuotesURL demonstrates URL construction for time series requests.
func ExampleNewQuotesURL() {
	url, err := alphavantage.NewQuotesURL("", "IBM", alphavantage.TimeSeriesDaily)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("URL contains symbol: %t\n", len(url) > 0)
	// Output: URL contains symbol: true
}
