package alphavantage_test

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"net/http"
	"testing"

	Ω "github.com/onsi/gomega"

	"github.com/crhntr/alphavantage"
)

//go:embed test_data/monthly_IBM.csv
var monthlyIBM []byte

func TestQuotes(t *testing.T) {
	please := Ω.NewWithT(t)

	ctx := context.TODO()

	var avReq *http.Request

	quotes, err := alphavantage.Quotes(ctx, doerFunc(func(request *http.Request) (*http.Response, error) {
		avReq = request
		return &http.Response{
			Body:       io.NopCloser(bytes.NewReader(monthlyIBM)),
			StatusCode: http.StatusOK,
		}, nil
	}), "demo", "IBM", alphavantage.TimeSeriesMonthly)

	please.Expect(err).NotTo(Ω.HaveOccurred())
	please.Expect(quotes).To(Ω.HaveLen(260))
	please.Expect(avReq.Host).To(Ω.Equal("www.alphavantage.co"))
	please.Expect(avReq.URL.Scheme).To(Ω.Equal("https"))
	please.Expect(avReq.URL.Path).To(Ω.Equal("/query"))
	please.Expect(avReq.URL.Query().Get("function")).To(Ω.Equal("TIME_SERIES_MONTHLY"))
	please.Expect(avReq.URL.Query().Get("symbol")).To(Ω.Equal("IBM"))
	please.Expect(avReq.URL.Query().Get("apikey")).To(Ω.Equal("demo"))
	please.Expect(avReq.URL.Query().Get("datatype")).To(Ω.Equal("csv"))
}

func TestSearch(t *testing.T) {
	please := Ω.NewWithT(t)

	ctx := context.TODO()

	var avReq *http.Request

	quotes, err := alphavantage.Search(ctx, doerFunc(func(request *http.Request) (*http.Response, error) {
		avReq = request
		return &http.Response{
			Body:       io.NopCloser(bytes.NewReader(searchResults)),
			StatusCode: http.StatusOK,
		}, nil
	}), "demo", "GDX")

	please.Expect(err).NotTo(Ω.HaveOccurred())
	please.Expect(quotes).To(Ω.HaveLen(9))
	please.Expect(avReq.Host).To(Ω.Equal("www.alphavantage.co"))
	please.Expect(avReq.URL.Scheme).To(Ω.Equal("https"))
	please.Expect(avReq.URL.Path).To(Ω.Equal("/query"))
	please.Expect(avReq.URL.Query().Get("function")).To(Ω.Equal("SYMBOL_SEARCH"))
	please.Expect(avReq.URL.Query().Get("keywords")).To(Ω.Equal("GDX"))
	please.Expect(avReq.URL.Query().Get("apikey")).To(Ω.Equal("demo"))
	please.Expect(avReq.URL.Query().Get("datatype")).To(Ω.Equal("csv"))
}

type doerFunc func(*http.Request) (*http.Response, error)

func (fn doerFunc) Do(req *http.Request) (*http.Response, error) { return fn(req) }
