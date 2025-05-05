package alphavantage_test

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/portfoliotree/alphavantage"
)

//go:embed test_data/monthly_IBM.csv
var monthlyIBM []byte

func TestQuotes(t *testing.T) {
	ctx := context.Background()

	var avReq *http.Request

	waitCallCount := 0

	quotes, err := (&alphavantage.Client{
		Client: doerFunc(func(request *http.Request) (*http.Response, error) {
			avReq = request
			return &http.Response{
				Body:       io.NopCloser(bytes.NewReader(monthlyIBM)),
				StatusCode: http.StatusOK,
			}, nil
		}),
		APIKey: "demo",
		Limiter: waitFunc(func(ctx context.Context) error {
			waitCallCount++
			return nil
		}),
	}).Quotes(ctx, "IBM", alphavantage.TimeSeriesMonthly, nil)

	require.NoError(t, err)
	assert.Len(t, quotes, 260)
	assert.Equal(t, "www.alphavantage.co", avReq.Host)
	assert.Equal(t, "https", avReq.URL.Scheme)
	assert.Equal(t, "/query", avReq.URL.Path)
	assert.Equal(t, "TIME_SERIES_MONTHLY", avReq.URL.Query().Get("function"))
	assert.Equal(t, "IBM", avReq.URL.Query().Get("symbol"))
	assert.Equal(t, "demo", avReq.URL.Query().Get("apikey"))
	assert.Equal(t, "csv", avReq.URL.Query().Get("datatype"))
	assert.Equal(t, 1, waitCallCount)
}

func TestService_ParseQueryResponse(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		const responseText = `timestamp,open,high,low,close,volume
2020-08-21,13.2600,13.3200,13.1500,13.2500,751279
2020-08-20,13.4700,13.4750,13.2200,13.3800,854559
2020-08-19,13.5700,13.7100,13.4700,13.5000,521089
2020-08-18,13.8100,13.8700,13.5400,13.5700,571445
`

		_, err := alphavantage.ParseQuotes(bytes.NewReader([]byte(responseText)), nil)
		require.NoError(t, err)
	})

	t.Run("intra-day", func(t *testing.T) {
		const responseText = `timestamp,open,high,low,close,volume
2020-08-21 19:40:00,123.1700,123.1700,123.1700,123.1700,825
2020-08-21 19:20:00,123.2000,123.2000,123.2000,123.2000,200
2020-08-21 18:50:00,123.1700,123.1700,123.1700,123.1700,115
2020-08-21 17:30:00,123.0200,123.0200,123.0200,123.0200,200`
		_, err := alphavantage.ParseIntraDayQuotes(bytes.NewReader([]byte(responseText)), nil)
		require.NoError(t, err)
	})

	t.Run("unexpected column", func(t *testing.T) {
		const responseText = `timestamp,open,high,low,close,volume,unexpected
2020-08-21 19:40:00,123.1700,123.1700,123.1700,123.1700,825,123456789
2020-08-21 19:20:00,123.2000,123.2000,123.2000,123.2000,200,123456789
2020-08-21 18:50:00,123.1700,123.1700,123.1700,123.1700,115,123456789
2020-08-21 17:30:00,123.0200,123.0200,123.0200,123.0200,200,123456789`
		_, err := alphavantage.ParseIntraDayQuotes(bytes.NewReader([]byte(responseText)), nil)
		require.NoError(t, err)
	})

	t.Run("split_coefficient", func(t *testing.T) {
		const responseText = `timestamp,open,high,low,close,adjusted_close,volume,dividend_amount,split_coefficient
2020-08-31,444.6100,500.1400,440.1100,498.3200,498.3200,115847020,0.0000,5.0000
`

		quotes, err := alphavantage.ParseQuotes(bytes.NewReader([]byte(responseText)), nil)
		require.NoError(t, err)
		assert.Len(t, quotes, 1)
		assert.Equal(t, "5.0000", quotes[0].SplitCoefficient)
	})

	t.Run("dividend", func(t *testing.T) {
		const responseText = `timestamp,open,high,low,close,adjusted_close,volume,dividend_amount,split_coefficient
2020-08-31,444.6100,500.1400,440.1100,498.3200,498.3200,115847020,4.20,5.0000
`

		quotes, err := alphavantage.ParseQuotes(bytes.NewReader([]byte(responseText)), nil)
		require.NoError(t, err)
		assert.Len(t, quotes, 1)
		assert.Equal(t, "4.20", quotes[0].DividendAmount)
	})
}
