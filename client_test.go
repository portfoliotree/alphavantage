package alphavantage_test

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/portfoliotree/alphavantage"
)

const apiKeyTestValue = "demo"

type doerFunc func(*http.Request) (*http.Response, error)

func (fn doerFunc) Do(req *http.Request) (*http.Response, error) { return fn(req) }

type waitFunc func(ctx context.Context) error

func (wf waitFunc) Wait(ctx context.Context) error {
	return wf(ctx)
}

func TestParse(t *testing.T) {
	t.Run("nil data", func(t *testing.T) {
		assert.Panics(t, func() {
			_ = alphavantage.ParseCSV(bytes.NewReader(nil), (*[]alphavantage.TimeSeriesDailyQuery)(nil), nil)
		})
	})

	t.Run("real data", func(t *testing.T) {
		var someFolks []struct {
			ID           int       `column-name:"id"`
			FirstInitial string    `column-name:"first_initial"`
			BirthDate    time.Time `column-name:"birth_date" time-layout:"2006/01/02"`
			Mass         float64   `column-name:"mass"`
		}

		err := alphavantage.ParseCSV(strings.NewReader(panthersCSV), &someFolks, nil)
		require.NoError(t, err)
		assert.Len(t, someFolks, 3)

		assert.Equal(t, 1, someFolks[0].ID)
		assert.Equal(t, "N", someFolks[0].FirstInitial)
		assert.Equal(t, mustParseDate(t, "2020-02-17"), someFolks[0].BirthDate)
		assert.Equal(t, 70.0, someFolks[0].Mass)

		assert.Equal(t, 2, someFolks[1].ID)
		assert.Equal(t, "S", someFolks[1].FirstInitial)
		assert.Equal(t, mustParseDate(t, "2020-10-22"), someFolks[1].BirthDate)
		assert.Equal(t, 68.2, someFolks[1].Mass)

		assert.Equal(t, 3, someFolks[2].ID)
		assert.Equal(t, "C", someFolks[2].FirstInitial)
		assert.Equal(t, mustParseDate(t, "2021-08-31"), someFolks[2].BirthDate)
		assert.Equal(t, 72.9, someFolks[2].Mass)
	})
}

const panthersCSV = `id,first_initial,birth_date,mass
1, N, 2020/02/17, 70
2, S, 2020/10/22, 68.2
3, C, 2021/08/31, 72.9
`

func mustParseDate(t *testing.T, date string) time.Time {
	tm, err := time.ParseInLocation(alphavantage.DefaultDateFormat, date, time.UTC)
	if err != nil {
		t.Fatal(err)
	}
	return tm
}

func TestClient_ETFProfile(t *testing.T) {
	f, err := os.Open("testdata/SPY_etf_profile.json")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()
	buf, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	var (
		avReq         *http.Request
		waitCallCount = 0
	)

	profile, err := (&alphavantage.Client{
		Client: doerFunc(func(request *http.Request) (*http.Response, error) {
			avReq = request
			return &http.Response{
				Body:       io.NopCloser(bytes.NewReader(buf)),
				StatusCode: http.StatusOK,
			}, nil
		}),
		APIKey: apiKeyTestValue,
		Limiter: waitFunc(func(ctx context.Context) error {
			waitCallCount++
			return nil
		}),
	}).ETFProfile(ctx, alphavantage.QueryETFProfile(apiKeyTestValue, "SPY"))
	require.NoError(t, err)

	assert.Equal(t, "ETF_PROFILE", avReq.URL.Query().Get("function"))
	assert.Equal(t, "SPY", avReq.URL.Query().Get("symbol"))

	assert.Equal(t, "654800000000", profile.NetAssets)
	assert.Equal(t, "0.000945", profile.NetExpenseRatio)
	assert.Equal(t, "0.03", profile.PortfolioTurnover)
	assert.Equal(t, "0.0108", profile.DividendYield)
	assert.Equal(t, "1993-01-22", profile.InceptionDate)
	assert.Equal(t, "NO", profile.Leveraged)
	assert.NotEmpty(t, profile.Sectors)
	assert.NotEmpty(t, profile.Holdings)

	// Check first sector
	assert.Equal(t, "INFORMATION TECHNOLOGY", profile.Sectors[0].Sector)
	assert.Equal(t, "0.337", profile.Sectors[0].Weight)

	// Check first holding
	assert.Equal(t, "NVDA", profile.Holdings[0].Symbol)
	assert.Equal(t, "NVIDIA CORP", profile.Holdings[0].Description)
	assert.Equal(t, "0.076", profile.Holdings[0].Weight)
}

// StockPrice represents a daily stock price with struct tags for CSV parsing.
// Supported field types: string, int, float64, time.Time
type StockPrice struct {
	Date   time.Time `column-name:"timestamp"`
	Open   float64   `column-name:"open"`
	High   float64   `column-name:"high"`
	Low    float64   `column-name:"low"`
	Close  float64   `column-name:"close"`
	Volume int       `column-name:"volume"`
}

// SearchResult represents a symbol search result with struct tags.
type SearchResult struct {
	Symbol      string  `column-name:"symbol"`
	Name        string  `column-name:"name"`
	Type        string  `column-name:"type"`
	Region      string  `column-name:"region"`
	MarketOpen  string  `column-name:"marketOpen"`
	MarketClose string  `column-name:"marketClose"`
	Timezone    string  `column-name:"timezone"`
	Currency    string  `column-name:"currency"`
	MatchScore  float64 `column-name:"matchScore"`
}

// ExampleParseCSV demonstrates parsing CSV data into a slice of structs.
// The struct fields must be tagged with "column-name" matching the CSV headers.
func ExampleParseCSV() {
	// Sample CSV data (truncated from testdata)
	csvData := `timestamp,open,high,low,close,volume
2020-08-21,13.2600,13.3200,13.1500,13.2500,751279
2020-08-20,13.4700,13.4750,13.2200,13.3800,854559
2020-08-19,13.5700,13.7100,13.4700,13.5000,521089`

	var prices []StockPrice
	err := alphavantage.ParseCSV(strings.NewReader(csvData), &prices, time.UTC)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Parsed %d prices\n", len(prices))
	fmt.Printf("First price: %s - Close: %.2f, Volume: %d\n",
		prices[0].Date.Format("2006-01-02"), prices[0].Close, prices[0].Volume)

	// Output: Parsed 3 prices
	// First price: 2020-08-21 - Close: 13.25, Volume: 751279
}

// ExampleParseCSVRows demonstrates streaming CSV parsing using an iterator.
// This is memory-efficient for large datasets as it processes one row at a time.
func ExampleParseCSVRows() {
	csvData := `symbol,name,type,region,marketOpen,marketClose,timezone,currency,matchScore
BA,Boeing Company,Equity,United States,09:30,16:00,UTC-04,USD,1.0000
BAB,Invesco Taxable Municipal Bond ETF,ETF,United States,09:30,16:00,UTC-04,USD,0.8000`

	count := 0
	for result := range alphavantage.ParseCSVRows[SearchResult](strings.NewReader(csvData), time.UTC, func(err error) bool {
		fmt.Printf("Parse error: %v\n", err)
		return false // stop on error
	}) {
		count++
		fmt.Printf("Symbol: %s, Match: %.1f\n", result.Symbol, result.MatchScore)
	}

	fmt.Printf("Processed %d results\n", count)

	// Output: Symbol: BA, Match: 1.0
	// Symbol: BAB, Match: 0.8
	// Processed 2 results
}

//go:embed testdata/intraday_5min_IBM.csv
var intradayIBM []byte

func TestTimeSeriesIntraday(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping unit test in short mode")
	}

	ctx := context.Background()

	// Mock client that intercepts HTTP requests
	mockClient := &alphavantage.Client{
		Client: doerFunc(func(req *http.Request) (*http.Response, error) {
			// Verify the request
			assert.Equal(t, "/query", req.URL.Path)
			assert.Equal(t, "TIME_SERIES_INTRADAY", req.URL.Query().Get("function"))
			assert.Equal(t, "IBM", req.URL.Query().Get("symbol"))
			assert.Equal(t, "csv", req.URL.Query().Get("datatype"))
			assert.Equal(t, "15min", req.URL.Query().Get("interval"))
			assert.Equal(t, "true", req.URL.Query().Get("extended_hours"))
			assert.Equal(t, "compact", req.URL.Query().Get("outputsize"))
			assert.Equal(t, "test-key", req.URL.Query().Get("apikey"))

			// Return test data from embedded CSV
			return &http.Response{
				StatusCode: 200,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader(intradayIBM)),
			}, nil
		}),
		Limiter: waitFunc(func(ctx context.Context) error { return nil }),
		APIKey:  "test-key",
	}

	result, err := mockClient.GetTimeSeriesIntradayCSVRows(ctx, alphavantage.QueryTimeSeriesIntraday("test-key", "IBM", alphavantage.IntervalOption15min).ExtendedHours(true).OutputSizeCompact())
	require.NoError(t, err)

	// Verify we got the expected number of quotes
	assert.Len(t, result, 100)

	// Verify first quote details
	assert.Equal(t, "2020-08-21 19:40:00 +0000 UTC", result[0].TimeStamp.String())
	assert.Equal(t, 123.1700, result[0].Open)
	assert.Equal(t, 123.1700, result[0].High)
	assert.Equal(t, 123.1700, result[0].Low)
	assert.Equal(t, 123.1700, result[0].Close)
	assert.Equal(t, 825, result[0].Volume)

	// Verify last quote details (from previous day)
	assert.Equal(t, "2020-08-20 17:10:00 +0000 UTC", result[99].TimeStamp.String())
	assert.Equal(t, 123.1500, result[99].Open)
	assert.Equal(t, 123.1500, result[99].High)
	assert.Equal(t, 123.1500, result[99].Low)
	assert.Equal(t, 123.1500, result[99].Close)
	assert.Equal(t, 2916, result[99].Volume)
}

func TestClient_CompanyOverview(t *testing.T) {
	f, err := os.Open(filepath.FromSlash("testdata/company_overview.json"))
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()
	buf, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	var (
		avReq         *http.Request
		waitCallCount = 0
	)

	overview, err := (&alphavantage.Client{
		Client: doerFunc(func(request *http.Request) (*http.Response, error) {
			avReq = request
			return &http.Response{
				Body:       io.NopCloser(bytes.NewReader(buf)),
				StatusCode: http.StatusOK,
			}, nil
		}),
		APIKey: apiKeyTestValue,
		Limiter: waitFunc(func(ctx context.Context) error {
			waitCallCount++
			return nil
		}),
	}).CompanyOverview(ctx, alphavantage.QueryOverview(apiKeyTestValue, "IBM"))
	require.NoError(t, err)

	require.NotNil(t, avReq)
	assert.Equal(t, "OVERVIEW", avReq.URL.Query().Get("function"))
	assert.Equal(t, "IBM", avReq.URL.Query().Get("symbol"))

	assert.Equal(t, "IBM", overview.Symbol)
	assert.Equal(t, "Common Stock", overview.AssetType)
	assert.Equal(t, "International Business Machines Corporation", overview.Name)
	assert.Equal(t, "International Business Machines Corporation (IBM) is an American multinational technology company headquartered in Armonk, New York, with operations in over 170 countries. The company began in 1911, founded in Endicott, New York, as the Computing-Tabulating-Recording Company (CTR) and was renamed International Business Machines in 1924. IBM is incorporated in New York. IBM produces and sells computer hardware, middleware and software, and provides hosting and consulting services in areas ranging from mainframe computers to nanotechnology. IBM is also a major research organization, holding the record for most annual U.S. patents generated by a business (as of 2020) for 28 consecutive years. Inventions by IBM include the automated teller machine (ATM), the floppy disk, the hard disk drive, the magnetic stripe card, the relational database, the SQL programming language, the UPC barcode, and dynamic random-access memory (DRAM). The IBM mainframe, exemplified by the System/360, was the dominant computing platform during the 1960s and 1970s.", overview.Description)
	assert.Equal(t, "51143", overview.CIK)
	assert.Equal(t, "NYSE", overview.Exchange)
	assert.Equal(t, "USD", overview.Currency)
	assert.Equal(t, "USA", overview.Country)
	assert.Equal(t, "TECHNOLOGY", overview.Sector)
	assert.Equal(t, "COMPUTER & OFFICE EQUIPMENT", overview.Industry)
	assert.Equal(t, "1 NEW ORCHARD ROAD, ARMONK, NY, US", overview.Address)
	assert.Equal(t, "December", overview.FiscalYearEnd)
	assert.Equal(t, mustParseDate(t, "2021-06-30"), overview.LatestQuarter)
	assert.Equal(t, 119909687000, overview.MarketCapitalization)
	assert.Equal(t, 15992001000, overview.EBITDA)
	assert.Equal(t, 22.61, overview.PERatio)
	assert.Equal(t, 1.397, overview.PEGRatio)
	assert.Equal(t, 24.48, overview.BookValue)
	assert.Equal(t, 6.53, overview.DividendPerShare)
	assert.Equal(t, 0.0486, overview.DividendYield)
	assert.Equal(t, 5.92, overview.EPS)
	assert.Equal(t, 83.3, overview.RevenuePerShareTTM)
	assert.Equal(t, 0.0717, overview.ProfitMargin)
	assert.Equal(t, 0.124, overview.OperatingMarginTTM)
	assert.Equal(t, 0.0385, overview.ReturnOnAssetsTTM)
	assert.Equal(t, 0.245, overview.ReturnOnEquityTTM)
	assert.Equal(t, 74400997000, overview.RevenueTTM)
	assert.Equal(t, 35575000000, overview.GrossProfitTTM)
	assert.Equal(t, 5.92, overview.DilutedEPSTTM)
	assert.Equal(t, -0.032, overview.QuarterlyEarningsGrowthYOY)
	assert.Equal(t, 0.034, overview.QuarterlyRevenueGrowthYOY)
	assert.Equal(t, 150.0, overview.AnalystTargetPrice)
	assert.Equal(t, 22.61, overview.TrailingPE)
	assert.Equal(t, 11.04, overview.ForwardPE)
	assert.Equal(t, 1.612, overview.PriceToSalesRatioTTM)
	assert.Equal(t, 5.49, overview.PriceToBookRatio)
	assert.Equal(t, 2.326, overview.EVToRevenue)
	assert.Equal(t, 12.81, overview.EVToEBITDA)
	assert.Equal(t, 1.212, overview.Beta)
	assert.Equal(t, 151.1, overview.FiftyTwoWeekHigh)
	assert.Equal(t, 100.73, overview.FiftyTwoWeekLow)
	assert.Equal(t, 139.81, overview.FiftyDayMovingAverage)
	assert.Equal(t, 140.06, overview.TwoHundredDayMovingAverage)
	assert.Equal(t, 896320000, overview.SharesOutstanding)
	assert.Equal(t, 894743000, overview.SharesFloat)
	assert.Equal(t, 25087600, overview.SharesShort)
	assert.Equal(t, 25615000, overview.SharesShortPriorMonth)
	assert.Equal(t, 7.92, overview.ShortRatio)
	assert.Equal(t, 0.03, overview.ShortPercentOutstanding)
	assert.Equal(t, 0.028, overview.ShortPercentFloat)
	assert.Equal(t, 0.133, overview.PercentInsiders)
	assert.Equal(t, 57.72, overview.PercentInstitutions)
	assert.Equal(t, 6.56, overview.ForwardAnnualDividendRate)
	assert.Equal(t, 0.0488, overview.ForwardAnnualDividendYield)
	assert.Equal(t, 0.747, overview.PayoutRatio)
	assert.Equal(t, mustParseDate(t, "2021-09-10"), overview.DividendDate)
	assert.Equal(t, mustParseDate(t, "2021-08-09"), overview.ExDividendDate)
	assert.Equal(t, "2:1", overview.LastSplitFactor)
	assert.Equal(t, mustParseDate(t, "1999-05-27"), overview.LastSplitDate)
}

func TestSearch(t *testing.T) {
	f, err := os.Open(filepath.FromSlash("testdata/search_results.csv"))
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()

	ctx := context.Background()

	var (
		avReq         *http.Request
		waitCallCount = 0
	)

	results, err := (&alphavantage.Client{
		Client: doerFunc(func(request *http.Request) (*http.Response, error) {
			avReq = request
			return &http.Response{
				Body:       io.NopCloser(f),
				StatusCode: http.StatusOK,
			}, nil
		}),
		APIKey: apiKeyTestValue,
		Limiter: waitFunc(func(ctx context.Context) error {
			waitCallCount++
			return nil
		}),
	}).GetSymbolSearchCSVRows(ctx, alphavantage.QuerySymbolSearch("demo", "BA"))
	require.NoError(t, err)
	assert.Len(t, results, 10)

	assert.Equal(t, "www.alphavantage.co", avReq.Host)
	assert.Equal(t, "https", avReq.URL.Scheme)
	assert.Equal(t, "/query", avReq.URL.Path)
	assert.Equal(t, "SYMBOL_SEARCH", avReq.URL.Query().Get("function"))
	assert.Equal(t, "BA", avReq.URL.Query().Get("keywords"))
	assert.Equal(t, apiKeyTestValue, avReq.URL.Query().Get("apikey"))
	assert.Equal(t, "csv", avReq.URL.Query().Get("datatype"))
	assert.Equal(t, 1, waitCallCount)
}
