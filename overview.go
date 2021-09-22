package alphavantage

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func (client *Client) CompanyOverview(ctx context.Context, symbol string) (CompanyOverview, error) {
	u := url.URL{
		Scheme: "https",
		Host:   "www.alphavantage.co",
		Path:   "/query",
		RawQuery: url.Values{
			"function": []string{"OVERVIEW"},
			"symbol":   []string{symbol},
		}.Encode(),
	}

	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		u.String(),
		nil,
	)
	if err != nil {
		return CompanyOverview{}, fmt.Errorf("failed to create listing status request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return CompanyOverview{}, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return CompanyOverview{}, err
	}

	var result CompanyOverview
	return result, json.Unmarshal(buf, &result)
}

type CompanyOverview struct {
	Symbol                     string  `json:"Symbol"`
	AssetType                  string  `json:"AssetType"`
	Name                       string  `json:"Name"`
	Description                string  `json:"Description"`
	CIK                        int     `json:"CIK,string"`
	Exchange                   string  `json:"Exchange"`
	Currency                   string  `json:"Currency"`
	Country                    string  `json:"Country"`
	Sector                     string  `json:"Sector"`
	Industry                   string  `json:"Industry"`
	Address                    string  `json:"Address"`
	FiscalYearEnd              string  `json:"FiscalYearEnd"`
	LatestQuarter              Date    `json:"LatestQuarter"`
	MarketCapitalization       int     `json:"MarketCapitalization,string"`
	EBITDA                     int     `json:"EBITDA,string"`
	PERatio                    float64 `json:"PERatio,string"`
	PEGRatio                   float64 `json:"PEGRatio,string"`
	BookValue                  float64 `json:"BookValue,string"`
	DividendPerShare           float64 `json:"DividendPerShare,string"`
	DividendYield              float64 `json:"DividendYield,string"`
	EPS                        float64 `json:"EPS,string"`
	RevenuePerShareTTM         float64 `json:"RevenuePerShareTTM,string"`
	ProfitMargin               float64 `json:"ProfitMargin,string"`
	OperatingMarginTTM         float64 `json:"OperatingMarginTTM,string"`
	ReturnOnAssetsTTM          float64 `json:"ReturnOnAssetsTTM,string"`
	ReturnOnEquityTTM          float64 `json:"ReturnOnEquityTTM,string"`
	RevenueTTM                 int     `json:"RevenueTTM,string"`
	GrossProfitTTM             int     `json:"GrossProfitTTM,string"`
	DilutedEPSTTM              float64 `json:"DilutedEPSTTM,string"`
	QuarterlyEarningsGrowthYOY float64 `json:"QuarterlyEarningsGrowthYOY,string"`
	QuarterlyRevenueGrowthYOY  float64 `json:"QuarterlyRevenueGrowthYOY,string"`
	AnalystTargetPrice         float64 `json:"AnalystTargetPrice,string"`
	TrailingPE                 float64 `json:"TrailingPE,string"`
	ForwardPE                  float64 `json:"ForwardPE,string"`
	PriceToSalesRatioTTM       float64 `json:"PriceToSalesRatioTTM,string"`
	PriceToBookRatio           float64 `json:"PriceToBookRatio,string"`
	EVToRevenue                float64 `json:"EVToRevenue,string"`
	EVToEBITDA                 float64 `json:"EVToEBITDA,string"`
	Beta                       float64 `json:"Beta,string"`
	FiftyTwoWeekHigh           float64 `json:"52WeekHigh,string"`
	FiftyTwoWeekLow            float64 `json:"52WeekLow,string"`
	FiftyDayMovingAverage      float64 `json:"50DayMovingAverage,string"`
	TwoHundredDayMovingAverage float64 `json:"200DayMovingAverage,string"`
	SharesOutstanding          int     `json:"SharesOutstanding,string"`
	SharesFloat                int     `json:"SharesFloat,string"`
	SharesShort                int     `json:"SharesShort,string"`
	SharesShortPriorMonth      int     `json:"SharesShortPriorMonth,string"`
	ShortRatio                 float64 `json:"ShortRatio,string"`
	ShortPercentOutstanding    float64 `json:"ShortPercentOutstanding,string"`
	ShortPercentFloat          float64 `json:"ShortPercentFloat,string"`
	PercentInsiders            float64 `json:"PercentInsiders,string"`
	PercentInstitutions        float64 `json:"PercentInstitutions,string"`
	ForwardAnnualDividendRate  float64 `json:"ForwardAnnualDividendRate,string"`
	ForwardAnnualDividendYield float64 `json:"ForwardAnnualDividendYield,string"`
	PayoutRatio                float64 `json:"PayoutRatio,string"`
	DividendDate               Date    `json:"DividendDate"`
	ExDividendDate             Date    `json:"ExDividendDate"`
	LastSplitFactor            string  `json:"LastSplitFactor"`
	LastSplitDate              Date    `json:"LastSplitDate"`
}

type Date time.Time

func (d Date) Time() time.Time {
	return time.Time(d)
}

func (d *Date) UnmarshalJSON(in []byte) error {
	var s string
	err := json.Unmarshal(in, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation(DefaultDateFormat, s, easternTimezone)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format(DefaultDateFormat))
}

func (d Date) String() string {
	return time.Time(d).Format(DefaultDateFormat)
}
