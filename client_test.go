package alphavantage_test

import (
	"io/ioutil"
	"net/http"
	"testing"

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
		buf, err := ioutil.ReadFile("test_data/query_time_series_weekly_ibm.json")
		if err != nil {
			t.Fatal(err)
		}
		quotes, err := alphavantage.ParseQuotesResponse(buf)
		if err != nil {
			t.Fatal(err)
		}
		if len(quotes) != 1043 {
			t.Fail()
		}

		if v := quotes[len(quotes)-1].Volume; v != 22560100 {
			t.Errorf("expected 250 got %f", v)
		}

		if v := quotes[0].Open; v != 104.4400 {
			t.Errorf("expected 104.4400 got %f", v)
		}
	})

	t.Run("when quote keys are intraday", func(t *testing.T) {
		buf, err := ioutil.ReadFile("test_data/query_time_series_intraday_ibm.json")
		if err != nil {
			t.Fatal(err)
		}
		quotes, err := alphavantage.ParseQuotesResponse(buf)
		if err != nil {
			t.Fatal(err)
		}

		if v := quotes[len(quotes)-1].Time.Hour(); v != 18 {
			t.Errorf("expected 18 got %d", v)
		}
	})
}
