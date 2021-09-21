package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/crhntr/alphavantage"
)

func main() {
	token := "demo"
	if t := os.Getenv("ALPHA_VANTAGE_TOKEN"); t != "" {
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
		err := fetchQuotes(token, flag.Args()[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "help":
		flag.Usage()
	default:
		fmt.Printf("unknown command: %s", cmd)
		flag.Usage()
		os.Exit(1)
	}
}

func fetchQuotes(token string, args []string) error {
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
