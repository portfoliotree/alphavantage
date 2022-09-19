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
	fmt.Println("  status\n\tFetch listing & delisting status.\n\thttps://www.alphavantage.co/documentation/#listing-status")
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
		err := client.QuotesRequest(ctx, symbol, alphavantage.QuoteFunction(function), func(r io.Reader) error {
			f, err := os.Create(symbol + ".csv")
			if err != nil {
				return err
			}
			defer func() {
				_ = f.Close()
			}()

			_, err = io.Copy(f, r)
			return err
		})
		if err != nil {
			return fmt.Errorf("failed saving quotes for %q: %w", symbol, err)
		}
	}
	return nil
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

	err := client.ListingStatusRequest(ctx, status, func(r io.Reader) error {
		f, err := os.Create(fmt.Sprintf("status_listed_%t.csv", status))
		if err != nil {
			return err
		}
		defer func() {
			_ = f.Close()
		}()

		_, err = io.Copy(f, r)
		return err
	})
	if err != nil {
		return fmt.Errorf("failed fetching listing status: %w", err)
	}

	return nil
}
