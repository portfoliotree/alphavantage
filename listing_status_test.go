package alphavantage_test

import (
	"context"
	"io"
	"net/http"
	"os"
	"testing"

	. "github.com/onsi/gomega"

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
	}).ListingStatus(ctx, true)
	o.Expect(err).NotTo(HaveOccurred())
	o.Expect(results).To(HaveLen(8))

	o.Expect(avReq.Host).To(Equal("www.alphavantage.co"))
	o.Expect(avReq.URL.Scheme).To(Equal("https"))
	o.Expect(avReq.URL.Path).To(Equal("/query"))
	o.Expect(avReq.URL.Query().Get("function")).To(Equal("LISTING_STATUS"))
	o.Expect(avReq.URL.Query().Get("state")).To(Equal("active"))
	o.Expect(avReq.URL.Query().Get("apikey")).To(Equal("demo"))
	o.Expect(avReq.URL.Query().Get("datatype")).To(Equal("csv"))
	o.Expect(waitCallCount).To(Equal(1))
}

func TestClient_ListingStatus_delisted(t *testing.T) {
	f, err := os.Open("test_data/listing_status.csv")
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
	o.Expect(err).NotTo(HaveOccurred())
	o.Expect(avReq.URL.Query().Get("state")).To(Equal("delisted"))
}
