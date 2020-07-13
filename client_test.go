package alphavantage_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	Ω "github.com/onsi/gomega"

	"github.com/crhntr/alphavantage"
	"github.com/crhntr/alphavantage/fakes"
)

func TestService_Do(t *testing.T) {
	fakeClient := &fakes.Doer{}
	client := alphavantage.Service{
		Client: fakeClient,
	}
	req, _ := http.NewRequest(http.MethodGet, "/query?foo=bar", nil)

	_, err := client.Do(req)

	if err != nil {
		t.Error("it should not error")
		t.Logf("got: %s", err)
	}
	if fakeClient.CallCount != 1 {
		t.Error("it calls the underlying client")
	}
	if fakeClient.Recieves.Req.URL.Scheme != "https" {
		t.Error("it calls the correct scheme")
		t.Logf("got: %q", fakeClient.Recieves.Req.URL.Scheme)
	}
	if fakeClient.Recieves.Req.URL.Host != "www.alphavantage.co" {
		t.Error("it calls the correct host")
		t.Logf("got: %q", fakeClient.Recieves.Req.URL.Host)
	}
	if fakeClient.Recieves.Req.URL.Path != "/query" {
		t.Error("it calls the endpoint")
		t.Logf("got: %q", fakeClient.Recieves.Req.URL.Path)
	}
}

func TestService_QueryStocks(t *testing.T) {
	t.Run("when quote keys are daily", func(t *testing.T) {
		g := Ω.NewWithT(t)
		buf, err := ioutil.ReadFile("test_data/query_time_series_weekly_ibm.json")
		g.Expect(err).NotTo(Ω.HaveOccurred())
		quotes, err := alphavantage.ParseQuotesResponse(buf)
		g.Expect(err).NotTo(Ω.HaveOccurred())

		g.Expect(quotes).To(Ω.HaveLen(1043))

		lastQuote := quotes[len(quotes)-1]

		g.Expect(lastQuote.Time.Hour()).To(Ω.Equal(0))
		g.Expect(lastQuote.Volume).To(Ω.Equal(22560100.0))

		firstQuote := quotes[0]

		g.Expect(firstQuote.Open).To(Ω.Equal(104.4400))
	})

	t.Run("when quote keys are intraday", func(t *testing.T) {
		g := Ω.NewWithT(t)
		buf, err := ioutil.ReadFile("test_data/query_time_series_intraday_ibm.json")
		g.Expect(err).NotTo(Ω.HaveOccurred())
		quotes, err := alphavantage.ParseQuotesResponse(buf)
		g.Expect(err).NotTo(Ω.HaveOccurred())
		g.Expect(quotes[len(quotes)-1].Time.Hour()).To(Ω.Equal(18))
	})
}
