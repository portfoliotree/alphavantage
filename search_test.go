package alphavantage_test

import (
	"bytes"
	_ "embed"
	"testing"

	Ω "github.com/onsi/gomega"

	"github.com/crhntr/alphavantage"
)

//go:embed test_data/search_results.csv
var searchResults []byte

func TestParseSearchQuery(t *testing.T) {
	please := Ω.NewWithT(t)
	results, err := alphavantage.ParseSearchQuery(bytes.NewReader(searchResults))
	please.Expect(err).NotTo(Ω.HaveOccurred())
	please.Expect(results).To(Ω.HaveLen(9))
	please.Expect(results[:2]).To(Ω.Equal([]alphavantage.SearchResult{
		{
			Symbol:      "GDX",
			Name:        "VanEck Vectors Gold Miners ETF",
			Type:        "ETF",
			Region:      "United States",
			MarketOpen:  "09:30",
			MarketClose: "16:00",
			TimeZone:    "UTC-04",
			Currency:    "USD",
			MatchScore:  1.0,
		}, {
			Symbol:      "GDXD",
			Name:        "Bank of Montreal",
			Type:        "Equity",
			Region:      "United States",
			MarketOpen:  "09:30",
			MarketClose: "16:00",
			TimeZone:    "UTC-04",
			Currency:    "USD",
			MatchScore:  0.8571,
		},
	}))
}
