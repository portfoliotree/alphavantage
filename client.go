package alphavantage

import (
	"net/http"
	"net/url"
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type Service struct {
	Client Doer

	APIKey string
}

func (service Service) Do(req *http.Request) (*http.Response, error) {
	if service.Client == nil {
		service.Client = http.DefaultClient
	}
	req.URL, _ = url.Parse("https://www.alphavantage.co")
	req.URL.Query().Set("apikey", service.APIKey)
	return service.Client.Do(req)
}

// func (service Service) QueryStocks(symbol string) ([]StockQuotes, error) {
// 	res, _ := service.Do(fmt.Sprintf("/query?symbol=%s", symbol))
// 	buf, _ := ioutil.ReadAll(res)
// 	var quotes []StockQuote
// 	json.Unmarshal(buf, &quotes)
// 	return quotes, nil
// }
