package alphavantage_test

import (
	"context"
	_ "embed"
	"io"
	"net/http"
	"os"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/portfoliotree/alphavantage"
)

func TestSearch(t *testing.T) {
	f, err := os.Open("test_data/search_results.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()

	o := NewWithT(t)

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
		APIKey: "demo",
		Limiter: waitFunc(func(ctx context.Context) error {
			waitCallCount++
			return nil
		}),
	}).SymbolSearch(ctx, "BA")
	o.Expect(err).NotTo(HaveOccurred())
	o.Expect(results).To(HaveLen(10))

	o.Expect(avReq.Host).To(Equal("www.alphavantage.co"))
	o.Expect(avReq.URL.Scheme).To(Equal("https"))
	o.Expect(avReq.URL.Path).To(Equal("/query"))
	o.Expect(avReq.URL.Query().Get("function")).To(Equal("SYMBOL_SEARCH"))
	o.Expect(avReq.URL.Query().Get("keywords")).To(Equal("BA"))
	o.Expect(avReq.URL.Query().Get("apikey")).To(Equal("demo"))
	o.Expect(avReq.URL.Query().Get("datatype")).To(Equal("csv"))
	o.Expect(waitCallCount).To(Equal(1))
}

func TestParseSearchQuery(t *testing.T) {
	f, err := os.Open("test_data/search_results.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()

	please := NewWithT(t)
	results, err := alphavantage.ParseSymbolSearchQuery(f)
	please.Expect(err).NotTo(HaveOccurred())
	please.Expect(results).To(HaveLen(10))
	please.Expect(results[:2]).To(Equal([]alphavantage.SymbolSearchResult{
		{
			Symbol:      "BA",
			Name:        "Boeing Company",
			Type:        "Equity",
			Region:      "United States",
			MarketOpen:  "09:30",
			MarketClose: "16:00",
			TimeZone:    "UTC-04",
			Currency:    "USD",
			MatchScore:  1,
		},
		{
			Symbol:      "BAB",
			Name:        "Invesco Taxable Municipal Bond ETF",
			Type:        "ETF",
			Region:      "United States",
			MarketOpen:  "09:30",
			MarketClose: "16:00",
			TimeZone:    "UTC-04",
			Currency:    "USD",
			MatchScore:  0.8,
		},
	}))
}
