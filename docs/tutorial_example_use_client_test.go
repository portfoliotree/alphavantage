package docs_test

import (
	"cmp"
	"fmt"
	"os"

	"github.com/portfoliotree/alphavantage"
)

// ExampleClient_GlobalQuote demonstrates the basic usage pattern shown in the tutorial
// for getting a stock quote.
func ExampleClient_GetGlobalQuote() {
	// Step 1: Create a client with your API key from environment variable
	apiKey := cmp.Or(os.Getenv(alphavantage.TokenEnvironmentVariableName), "demo")
	client := alphavantage.NewClient(apiKey, alphavantage.PremiumPlan75)

	// Step 2: The basic pattern for API calls:
	// ctx := context.Background()
	// result, err := client.GlobalQuote(ctx, "IBM")
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer result.Close()
	//
	// // Step 3: Process the returned CSV data
	// data, err := io.ReadAll(result)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// fmt.Println(string(data))

	fmt.Printf("Tutorial client setup: %t\n", client != nil)
	// Output: Tutorial client setup: true
}
