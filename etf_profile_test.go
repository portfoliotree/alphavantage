package alphavantage_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/portfoliotree/alphavantage"
)

func TestClient_ETFProfile(t *testing.T) {
	f, err := os.Open("testdata/SPY_etf_profile.json")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()
	buf, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	var (
		avReq         *http.Request
		waitCallCount = 0
	)

	profile, err := (&alphavantage.Client{
		Client: doerFunc(func(request *http.Request) (*http.Response, error) {
			avReq = request
			return &http.Response{
				Body:       io.NopCloser(bytes.NewReader(buf)),
				StatusCode: http.StatusOK,
			}, nil
		}),
		APIKey: "demo",
		Limiter: waitFunc(func(ctx context.Context) error {
			waitCallCount++
			return nil
		}),
	}).ETFProfile(ctx, "SPY")
	require.NoError(t, err)

	assert.Equal(t, "ETF_PROFILE", avReq.URL.Query().Get("function"))
	assert.Equal(t, "SPY", avReq.URL.Query().Get("symbol"))

	assert.Equal(t, "654800000000", profile.NetAssets)
	assert.Equal(t, "0.000945", profile.NetExpenseRatio)
	assert.Equal(t, "0.03", profile.PortfolioTurnover)
	assert.Equal(t, "0.0108", profile.DividendYield)
	assert.Equal(t, "1993-01-22", profile.InceptionDate)
	assert.Equal(t, "NO", profile.Leveraged)
	assert.NotEmpty(t, profile.Sectors)
	assert.NotEmpty(t, profile.Holdings)

	// Check first sector
	assert.Equal(t, "INFORMATION TECHNOLOGY", profile.Sectors[0].Sector)
	assert.Equal(t, "0.337", profile.Sectors[0].Weight)

	// Check first holding
	assert.Equal(t, "NVDA", profile.Holdings[0].Symbol)
	assert.Equal(t, "NVIDIA CORP", profile.Holdings[0].Description)
	assert.Equal(t, "0.076", profile.Holdings[0].Weight)
}
