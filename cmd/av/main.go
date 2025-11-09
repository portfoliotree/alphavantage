package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/debug"

	"github.com/spf13/pflag"

	"github.com/portfoliotree/alphavantage"
)

func main() {
	token := "demo"
	if t := os.Getenv(alphavantage.StandardTokenEnvironmentVariableName); t != "" {
		token = t
	}
	var tokenFlag string
	pflag.StringVar(&tokenFlag, "token", "", "api authentication token")

	pflag.Parse()

	if tokenFlag != "" {
		token = tokenFlag
	}

	cmd := pflag.Arg(0)
	switch cmd {
	case "quotes":
		err := quotes(token, pflag.Args()[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "listing-status":
		err := listingStatus(token, pflag.Args()[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "symbol-search":
		err := symbolSearch(token, pflag.Args()[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "global-quote":
		err := globalQuote(token, pflag.Args()[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "etf-profile":
		err := etfProfile(token, flag.Args()[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "help":
		_ = help("", nil)
	case "version":
		info, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("failed to read build info")
			os.Exit(1)
		}
		fmt.Println(info.Main.Version)
	default:
		if cmd != "" {
			fmt.Printf("ERROR: unknown command: %s\n\n", cmd)
		} else {
			fmt.Printf("ERROR: missing command\n\n")
		}

		_ = help("", nil)
		os.Exit(1)
	}
}

func help(_ string, _ []string) error {
	fmt.Println("av - An AlphaVantage CLI in Go")
	fmt.Println()
	fmt.Println("Global Flags:")
	pflag.PrintDefaults()
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  global-quote\n\tFetch latest price and volume information for equity.\n\thttps://www.alphavantage.co/documentation/#latestprice")
	fmt.Println("  listing-status\n\tFetch listing & de-listing status.\n\thttps://www.alphavantage.co/documentation/#listing-status")
	fmt.Println("  quotes\n\tFetch time series stock quotes.\n\thttps://www.alphavantage.co/documentation/#time-series-data")
	fmt.Println("  symbol-search\n\tWrites symbol search results to stdout.\n\thttps://www.alphavantage.co/documentation/#symbolsearch")
	fmt.Println("  etf-profile\n\tFetch ETF profile data.\n\thttps://www.alphavantage.co/documentation/#etf")
	fmt.Println()
	return nil
}

func quotes(token string, args []string) error {
	flags := pflag.NewFlagSet("quotes", pflag.ContinueOnError)

	var function string
	flags.StringVar(&function, "function", string(alphavantage.TimeSeriesDailyAdjusted), "enter one of the stock time series functions with the TIME_ prefix")
	if err := flags.Parse(args); err != nil {
		return err
	}
	symbols := flags.Args()
	if len(args) == 0 {
		return nil
	}

	client := alphavantage.NewClient(token, alphavantage.PremiumPlan75)

	ctx := context.TODO()

	for _, symbol := range symbols {
		err := requestQuotes(ctx, client, function, symbol)
		if err != nil {
			return fmt.Errorf("failed saving quotes for %q: %w", symbol, err)
		}
	}
	return nil
}

func requestQuotes(ctx context.Context, client *alphavantage.Client, function, symbol string) error {
	fileName := symbol + ".csv"
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer closeAndIgnoreError(f)

	rc, err := client.DoQuotesRequest(ctx, symbol, alphavantage.QuoteFunction(function))
	if err != nil {
		_ = os.Remove(fileName)
		return err
	}
	defer closeAndIgnoreError(rc)

	fmt.Printf("writing quotes for %q to file %s\n", symbol, fileName)

	_, err = io.Copy(f, rc)
	return err
}

func listingStatus(token string, args []string) error {
	flags := pflag.NewFlagSet("status", pflag.ContinueOnError)

	var status bool
	flags.BoolVar(&status, "listed", true, "listing status")
	if err := flags.Parse(args); err != nil {
		return err
	}

	client := alphavantage.NewClient(token, alphavantage.PremiumPlan75)

	ctx := context.TODO()

	fileName := fmt.Sprintf("status_listed_%t.csv", status)
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer closeAndIgnoreError(f)

	rc, err := client.DoListingStatusRequest(ctx, status)
	if err != nil {
		_ = os.Remove(fileName)
		return fmt.Errorf("failed fetching listing status: %w", err)
	}
	defer closeAndIgnoreError(rc)

	fmt.Printf("writing listing statuses to file %s\n", fileName)

	_, err = io.Copy(f, rc)
	return nil
}

func symbolSearch(token string, args []string) error {
	flags := pflag.NewFlagSet("search", pflag.ContinueOnError)
	var writeToFile bool
	flags.BoolVar(&writeToFile, "O", false, "write to files instead of stdout")
	err := flags.Parse(args)
	if err != nil {
		return err
	}

	client := alphavantage.NewClient(token, alphavantage.PremiumPlan75)

	ctx := context.TODO()

	for _, arg := range flags.Args() {
		err = singleSymbolSearch(ctx, client, arg, writeToFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func singleSymbolSearch(ctx context.Context, client *alphavantage.Client, symbol string, writeToFile bool) error {
	result, err := client.DoSymbolSearchRequest(ctx, symbol)
	if err != nil {
		return err
	}
	defer closeAndIgnoreError(result)

	if !writeToFile {
		_, _ = io.Copy(os.Stdout, result)
		_, _ = fmt.Fprintln(os.Stdout)
		return nil
	}

	buffer, err := io.ReadAll(result)
	if err != nil {
		return err
	}

	fileName := symbol + ".csv"
	fmt.Println("writing file", fileName)
	return os.WriteFile(fileName, buffer, 0o666)
}

func globalQuote(token string, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("symbol is required")
	}

	client := alphavantage.NewClient(token, alphavantage.PremiumPlan75)
	ctx := context.TODO()

	for _, symbol := range args {
		err := requestGlobalQuote(ctx, client, symbol)
		if err != nil {
			return fmt.Errorf("failed getting quote for %q: %w", symbol, err)
		}
	}
	return nil
}

func requestGlobalQuote(ctx context.Context, client *alphavantage.Client, symbol string) error {
	fileName := symbol + "_quote.csv"
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer closeAndIgnoreError(f)

	rc, err := client.GetGlobalQuote(ctx, alphavantage.QueryGlobalQuote(client.APIKey, symbol))
	if err != nil {
		_ = os.Remove(fileName)
		return err
	}
	defer closeAndIgnoreError(rc.Body)

	fmt.Printf("writing global quote for %q to file %s\n", symbol, fileName)

	_, err = io.Copy(f, rc.Body)
	return err
}

func closeAndIgnoreError(c io.Closer) {
	_ = c.Close()
}

func etfProfile(token string, args []string) error {
	flags := flag.NewFlagSet("etf-profile", flag.ContinueOnError)
	var writeToFile bool
	flags.BoolVar(&writeToFile, "O", false, "write to files instead of stdout")
	err := flags.Parse(args)
	if err != nil {
		return err
	}

	symbols := flags.Args()
	if len(symbols) == 0 {
		return fmt.Errorf("no symbols provided")
	}

	client := alphavantage.NewClient(token, alphavantage.PremiumPlan75)
	ctx := context.TODO()

	for _, symbol := range symbols {
		profile, err := client.ETFProfile(ctx, alphavantage.QueryETFProfile(client.APIKey, symbol))
		if err != nil {
			return fmt.Errorf("failed to get ETF profile for %s: %w", symbol, err)
		}

		jsonData, err := json.MarshalIndent(profile, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal ETF profile: %w", err)
		}

		if !writeToFile {
			fmt.Println(string(jsonData))
			continue
		}

		fileName := symbol + "_etf_profile.json"
		fmt.Printf("writing ETF profile for %q to file %s\n", symbol, fileName)
		err = os.WriteFile(fileName, jsonData, 0o644)
		if err != nil {
			return fmt.Errorf("failed to write file %s: %w", fileName, err)
		}
	}
	return nil
}
