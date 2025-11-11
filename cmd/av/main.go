package main

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"

	"github.com/portfoliotree/alphavantage"
)

func main() {
	var (
		apikey            = "demo"
		perMinuteRequests = int(alphavantage.PremiumPlan75)
	)
	if t, ok := os.LookupEnv(alphavantage.TokenEnvironmentVariableName); ok {
		apikey = t
	}
	if val, ok := os.LookupEnv(alphavantage.RequestsPerMinuteEnvironmentVariableName); ok {
		n, err := alphavantage.NewRequestsPerMinute(val)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		perMinuteRequests = int(n)
	}

	var cmd string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "help", "":
		help()
	case "version":
		info, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("failed to read build info")
			os.Exit(1)
		}
		fmt.Println(info.Main.Version)
	default:
		client := alphavantage.NewClient(apikey, alphavantage.RequestsPerMinute(perMinuteRequests))

		err := RunFunction(client, cmd, os.Args[1:], os.Stdout)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}
	}
}

func help() {
	fmt.Println("av - An AlphaVantage CLI in Go")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  av <function> [flags]")
	fmt.Println()
	fmt.Println("Special Commands:")
	fmt.Println("  help       Show this help message")
	fmt.Println("  version    Show version information")
	fmt.Println()
	fmt.Println("API Functions:")
	fmt.Println("  Use any AlphaVantage function name (e.g., TIME_SERIES_DAILY, GLOBAL_QUOTE, SMA)")
	fmt.Println("  Run 'av <function> --help' to see function-specific flags")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  av GLOBAL_QUOTE --symbol=IBM")
	fmt.Println("  av TIME_SERIES_DAILY --symbol=AAPL --outputsize=full -o aapl.csv")
	fmt.Println("  av SMA --symbol=MSFT --interval=daily --time-period=20 --series-type=close")
	fmt.Println()
	fmt.Println("Documentation: https://www.alphavantage.co/documentation/")
}

func closeAndIgnoreError(c io.Closer) {
	_ = c.Close()
}
