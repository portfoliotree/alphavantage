package main

import (
	"context"
	"fmt"
	"os"

	"github.com/portfoliotree/alphavantage"
)

func main() {
	client := alphavantage.NewClient()

	ctx := context.Background()
	query := alphavantage.QueryTimeSeriesDaily(client.APIKey, "AAPL")
	rows, err := client.GetTimeSeriesDailyCSVRows(ctx, query)
	if err != nil {
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
