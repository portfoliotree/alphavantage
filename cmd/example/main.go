package main

import (
	"context"
	"fmt"
	"os"

	"github.com/portfoliotree/alphavantage"
	"github.com/portfoliotree/alphavantage/api"
	"github.com/portfoliotree/alphavantage/query/timeseries"
)

func main() {
	client := alphavantage.NewClient()

	ctx := context.Background()
	query := timeseries.QueryDaily(client.APIKey, "AAPL")
	result, err := client.Query(ctx, query)
	if err != nil {
		printErrorAndExit(err)
	}

	var rows []timeseries.DailyRow
	if err := api.ParseCSV(result.Body, &rows, nil); err != nil {
		printErrorAndExit(err)
	}

	for _, row := range rows[:10] {
		fmt.Println(row.Close - row.Open)
	}
}

func printErrorAndExit(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
