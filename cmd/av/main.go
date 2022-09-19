package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/portfoliotree/alphavantage"
)

func main() {
	token := "demo"
	if t := os.Getenv(alphavantage.StandardTokenEnvironmentVariableName); t != "" {
		token = t
	}
	var tokenFlag string
	flag.StringVar(&tokenFlag, "token", "", "api authentication token")

	flag.Parse()

	if tokenFlag != "" {
		token = tokenFlag
	}

	cmd := flag.Arg(0)
	switch cmd {
	case "quotes":
		err := quotes(token, flag.Args()[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "status":
		err := listingStatus(token, flag.Args()[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "symbol-search":
		err := symbolSearch(token, flag.Args()[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "help":
		_ = help("", nil)
	default:
		if cmd != "" {
			fmt.Printf("ERROR: unknown command: %s\n\n", cmd)
		} else if cmd == "" {
			fmt.Printf("ERROR: missing command\n\n")
		}

		_ = help("", nil)
		os.Exit(1)
	}
}

func help(string, []string) error {
	fmt.Println("av - An AlphaVantage CLI in Go")
	fmt.Println()
	fmt.Println("Global Flags:")
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  quotes\n\tFetch time series stock quotes.\n\thttps://www.alphavantage.co/documentation/#time-series-data")
	fmt.Println("  status\n\tFetch listing & de-listing status.\n\thttps://www.alphavantage.co/documentation/#listing-status")
	fmt.Println("  symbol-search\n\tWrites symbol search results to stdout.\n\thttps://www.alphavantage.co/documentation/#symbolsearch")
	fmt.Println()
	return nil
}

func quotes(token string, args []string) error {
	flags := flag.NewFlagSet("quotes", flag.ContinueOnError)

	var function string
	flags.StringVar(&function, "function", string(alphavantage.TimeSeriesDailyAdjusted), "enter one of the stock time series functions with the TIME_ prefix")
	if err := flags.Parse(args); err != nil {
		return err
	}
	symbols := flags.Args()
	if len(args) == 0 {
		return nil
	}

	client := alphavantage.NewClient(token)

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
	f, err := os.Create(symbol + ".csv")
	if err != nil {
		return err
	}
	defer closeAndIgnoreError(f)

	rc, err := client.DoQuotesRequest(ctx, symbol, alphavantage.QuoteFunction(function))
	if err != nil {
		return err
	}
	defer closeAndIgnoreError(rc)
	_, err = io.Copy(f, rc)
	return err
}

func listingStatus(token string, args []string) error {
	flags := flag.NewFlagSet("status", flag.ContinueOnError)

	var status bool
	flags.BoolVar(&status, "listed", true, "listing status")
	if err := flags.Parse(args); err != nil {
		return err
	}

	client := alphavantage.NewClient(token)

	ctx := context.TODO()

	f, err := os.Create(fmt.Sprintf("status_listed_%t.csv", status))
	if err != nil {
		return err
	}
	defer closeAndIgnoreError(f)

	rc, err := client.DoListingStatusRequest(ctx, status)
	if err != nil {
		return fmt.Errorf("failed fetching listing status: %w", err)
	}
	defer closeAndIgnoreError(rc)
	_, err = io.Copy(f, rc)
	return nil
}

func symbolSearch(token string, args []string) error {
	flags := flag.NewFlagSet("search", flag.ContinueOnError)
	err := flags.Parse(args)
	if err != nil {
		return err
	}

	client := alphavantage.NewClient(token)

	ctx := context.TODO()

	result, err := client.DoSymbolSearchRequest(ctx, flags.Arg(0))
	if err != nil {
		return err
	}
	defer closeAndIgnoreError(result)

	_, _ = io.Copy(os.Stdout, result)

	_, _ = fmt.Fprintln(os.Stdout)

	return err
}

func closeAndIgnoreError(c io.Closer) {
	_ = c.Close()
}
