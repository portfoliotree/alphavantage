package alphavantage_test

import (
	"context"
	"io"
	"net/http"
	"os"
	"testing"

	Ω "github.com/onsi/gomega"
	"github.com/portfoliotree/alphavantage"
)

func TestClient_ListingStatus_listed(t *testing.T) {
	f, err := os.Open("test_data/listing_status.csv")
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
	}).ListingStatus(ctx, true)
	please.Expect(err).NotTo(Ω.HaveOccurred())
	please.Expect(results).To(Ω.HaveLen(8))

	please.Expect(avReq.Host).To(Ω.Equal("www.alphavantage.co"))
	please.Expect(avReq.URL.Scheme).To(Ω.Equal("https"))
	please.Expect(avReq.URL.Path).To(Ω.Equal("/query"))
	please.Expect(avReq.URL.Query().Get("function")).To(Ω.Equal("LISTING_STATUS"))
	please.Expect(avReq.URL.Query().Get("state")).To(Ω.Equal("active"))
	please.Expect(avReq.URL.Query().Get("apikey")).To(Ω.Equal("demo"))
	please.Expect(avReq.URL.Query().Get("datatype")).To(Ω.Equal("csv"))
	please.Expect(waitCallCount).To(Ω.Equal(1))
}

func TestClient_ListingStatus_delisted(t *testing.T) {
	f, err := os.Open("test_data/listing_status.csv")
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

	_, err = (&alphavantage.Client{
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
	}).ListingStatus(ctx, false)
	please.Expect(err).NotTo(Ω.HaveOccurred())
	please.Expect(avReq.URL.Query().Get("state")).To(Ω.Equal("delisted"))
}
