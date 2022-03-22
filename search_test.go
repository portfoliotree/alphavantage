package alphavantage_test

import (
	"context"
	_ "embed"
	"io"
	"net/http"
	"os"
	"testing"

	Ω "github.com/onsi/gomega"

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

	please := Ω.NewWithT(t)

	ctx := context.TODO()

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
	}).Search(ctx, "BA")
	please.Expect(err).NotTo(Ω.HaveOccurred())
	please.Expect(results).To(Ω.HaveLen(10))

	please.Expect(avReq.Host).To(Ω.Equal("www.alphavantage.co"))
	please.Expect(avReq.URL.Scheme).To(Ω.Equal("https"))
	please.Expect(avReq.URL.Path).To(Ω.Equal("/query"))
	please.Expect(avReq.URL.Query().Get("function")).To(Ω.Equal("SYMBOL_SEARCH"))
	please.Expect(avReq.URL.Query().Get("keywords")).To(Ω.Equal("BA"))
	please.Expect(avReq.URL.Query().Get("apikey")).To(Ω.Equal("demo"))
	please.Expect(avReq.URL.Query().Get("datatype")).To(Ω.Equal("csv"))
	please.Expect(waitCallCount).To(Ω.Equal(1))
}

func TestParseSearchQuery(t *testing.T) {
	f, err := os.Open("test_data/search_results.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()

	please := Ω.NewWithT(t)
	results, err := alphavantage.ParseSearchQuery(f)
	please.Expect(err).NotTo(Ω.HaveOccurred())
	please.Expect(results).To(Ω.HaveLen(10))
	please.Expect(results[:2]).To(Ω.Equal([]alphavantage.SearchResult{
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
