package alphavantage

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
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

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return CompanyOverview{}, err
	}

	var result CompanyOverview
	err = json.Unmarshal(buf, &result)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

type CompanyOverview struct {
	CIK                        string    `av-json:"CIK"`
	Symbol                     string    `av-json:"Symbol"`
	AssetType                  string    `av-json:"AssetType"`
	Name                       string    `av-json:"Name"`
	Description                string    `av-json:"Description"`
	Exchange                   string    `av-json:"Exchange"`
	Currency                   string    `av-json:"Currency"`
	Country                    string    `av-json:"Country"`
	Sector                     string    `av-json:"Sector"`
	Industry                   string    `av-json:"Industry"`
	Address                    string    `av-json:"Address"`
	FiscalYearEnd              string    `av-json:"FiscalYearEnd"`
	LatestQuarter              time.Time `av-json:"LatestQuarter"`
	MarketCapitalization       int       `av-json:"MarketCapitalization"`
	EBITDA                     int       `av-json:"EBITDA"`
	PERatio                    string    `av-json:"PERatio"`
	PEGRatio                   string    `av-json:"PEGRatio"`
	BookValue                  string    `av-json:"BookValue"`
	DividendPerShare           string    `av-json:"DividendPerShare"`
	DividendYield              string    `av-json:"DividendYield"`
	EPS                        string    `av-json:"EPS"`
	RevenuePerShareTTM         string    `av-json:"RevenuePerShareTTM"`
	ProfitMargin               string    `av-json:"ProfitMargin"`
	OperatingMarginTTM         string    `av-json:"OperatingMarginTTM"`
	ReturnOnAssetsTTM          string    `av-json:"ReturnOnAssetsTTM"`
	ReturnOnEquityTTM          string    `av-json:"ReturnOnEquityTTM"`
	RevenueTTM                 int       `av-json:"RevenueTTM"`
	GrossProfitTTM             int       `av-json:"GrossProfitTTM"`
	DilutedEPSTTM              string    `av-json:"DilutedEPSTTM"`
	QuarterlyEarningsGrowthYOY string    `av-json:"QuarterlyEarningsGrowthYOY"`
	QuarterlyRevenueGrowthYOY  string    `av-json:"QuarterlyRevenueGrowthYOY"`
	AnalystTargetPrice         string    `av-json:"AnalystTargetPrice"`
	TrailingPE                 string    `av-json:"TrailingPE"`
	ForwardPE                  string    `av-json:"ForwardPE"`
	PriceToSalesRatioTTM       string    `av-json:"PriceToSalesRatioTTM"`
	PriceToBookRatio           string    `av-json:"PriceToBookRatio"`
	EVToRevenue                string    `av-json:"EVToRevenue"`
	EVToEBITDA                 string    `av-json:"EVToEBITDA"`
	Beta                       string    `av-json:"Beta"`
	FiftyTwoWeekHigh           string    `av-json:"52WeekHigh"`
	FiftyTwoWeekLow            string    `av-json:"52WeekLow"`
	FiftyDayMovingAverage      string    `av-json:"50DayMovingAverage"`
	TwoHundredDayMovingAverage string    `av-json:"200DayMovingAverage"`
	SharesOutstanding          int       `av-json:"SharesOutstanding"`
	SharesFloat                int       `av-json:"SharesFloat"`
	SharesShort                int       `av-json:"SharesShort"`
	SharesShortPriorMonth      int       `av-json:"SharesShortPriorMonth"`
	ShortRatio                 string    `av-json:"ShortRatio"`
	ShortPercentOutstanding    string    `av-json:"ShortPercentOutstanding"`
	ShortPercentFloat          string    `av-json:"ShortPercentFloat"`
	PercentInsiders            string    `av-json:"PercentInsiders"`
	PercentInstitutions        string    `av-json:"PercentInstitutions"`
	ForwardAnnualDividendRate  string    `av-json:"ForwardAnnualDividendRate"`
	ForwardAnnualDividendYield string    `av-json:"ForwardAnnualDividendYield"`
	PayoutRatio                string    `av-json:"PayoutRatio"`
	DividendDate               time.Time `av-json:"DividendDate"`
	ExDividendDate             time.Time `av-json:"ExDividendDate"`
	LastSplitFactor            string    `av-json:"LastSplitFactor"`
	LastSplitDate              time.Time `av-json:"LastSplitDate"`
}

func (c *CompanyOverview) UnmarshalJSON(in []byte) error {
	var data map[string]string

	err := json.Unmarshal(in, &data)
	if err != nil {
		return err
	}

	rv := reflect.ValueOf(c)

	numFields := rv.Type().Elem().NumField()
	for i := 0; i < numFields; i++ {
		ft := rv.Elem().Type().Field(i)
		fv := rv.Elem().Field(i)
		jsonKey := ft.Tag.Get("av-json")

		v, ok := data[jsonKey]
		if !ok || v == "" || v == "None" {
			continue
		}

		switch fv.Interface().(type) {
		case string:
			fv.SetString(v)
		case int:
			in, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse %s: %w", jsonKey, err)
			}
			fv.SetInt(in)
		case float64:
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return fmt.Errorf("failed to parse %s: %w", jsonKey, err)
			}
			fv.SetFloat(f)
		case time.Time:
			if v == "0000-00-00" {
				continue
			}
			t, err := time.ParseInLocation(DefaultDateFormat, v, time.UTC)
			if err != nil {
				return err
			}
			fv.Set(reflect.ValueOf(t))
		default:
			return fmt.Errorf("unsupported type %T", fv.Interface())
		}
	}

	return nil
}
