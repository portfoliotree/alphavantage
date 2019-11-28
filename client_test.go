package alphavantage_test

import (
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
		t.Error("it calls the correct host")
		t.Logf("got: %q", fakeClient.Recieves.Req.URL.Host)
	}
}

func TestService_QueryStocks(t *testing.T) {

}
