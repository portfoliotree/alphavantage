package main

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"

	"github.com/spf13/pflag"

	"github.com/portfoliotree/alphavantage"
)

func main() {
	var (
		functionName      string
		outputFile        string
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

	flagSet := pflag.NewFlagSet("av", pflag.ExitOnError)
	flagSet.StringVar(&outputFile, "output", "", "output filename")
	flagSet.StringVar(&functionName, "function", "", "function name")
	flagSet.StringVar(&apikey, "apikey", "demo", "API key")
	flagSet.IntVar(&perMinuteRequests, "requests-per-minute", perMinuteRequests, "requests per minute to permit")
	args := os.Args[1:]
	if err := flagSet.Parse(args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	functionArgs := flagSet.Args()

	switch functionName {
	case "help":
		help()
	case "version":
		info, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("failed to read build info")
			os.Exit(1)
		}
		fmt.Println(info.Main.Version)
	default:
		var output io.Writer = os.Stdout
		if outputFile != "" {
			f, err := os.Create(outputFile)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "ERROR: failed to create output file: %v\n", err)
				os.Exit(1)
			}
			defer closeAndIgnoreError(f)
			output = f
		}

		client := alphavantage.NewClient(apikey, alphavantage.RequestsPerMinute(perMinuteRequests))

		err := RunFunction(client, functionName, functionArgs, output)
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
	fmt.Println("  av [global-flags] <function> [flags]")
	fmt.Println()
	fmt.Println("Global Flags:")
	fmt.Println("  -o, --output string   output file (defaults to stdout)")
	fmt.Println("      --token string    API authentication token")
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
