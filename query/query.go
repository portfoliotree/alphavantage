package query

import (
	"fmt"
	"net/url"
	"time"
)

//goland:noinspection SpellCheckingInspection
const (
	KeyOutputSize = `outputsize`
	KeyDataType   = `dataType`
	KeyAPIKey     = `apikey`
	KeyFunction   = `function`
	KeyInterval   = `interval`
	KeySymbol     = `symbol`
	KeySeriesType = `series_type`
	KeyTimePeriod = `time_period`
	KeyMonth      = `month`
	KeyMarket     = `market`

	KeyFromCurrency = `from_currency`
	KeyToCurrency   = `to_currency`
	KeyFromSymbol   = `from_symbol`
	KeyToSymbol     = `to_symbol`
	KeyOHLC         = `OHLC`
	KeyState        = `state`
	KeyHorizon      = `horizon`
	KeySort         = `sort`
	KeyMaturity     = `maturity`
	KeyDate         = `date`

	KeyFastPeriod   = `fastperiod`
	KeySlowPeriod   = `slowperiod`
	KeySignalPeriod = `signalperiod`
	KeyFastMAType   = `fastmatype`
	KeySlowMAType   = `slowmatype`
	KeyMAType       = `matype`

	KeyFastKPeriod  = `fastkperiod`
	KeySlowKPeriod  = `slowkperiod`
	KeyFastDPeriod  = `fastdperiod`
	KeySlowDPeriod  = `slowdperiod`
	KeyFastDMAType  = `fastdmatype`
	KeySlowDMAType  = `slowdmatype`
	KeySlowKMAType  = `slowkmatype`
	KeySignalMAType = "signalmatype"
	KeySlowLimit    = `slowlimit`
	KeyFastLimit    = `fastlimit`

	KeyNbdevup      = `nbdevup`
	KeyNbdevdn      = `nbdevdn`
	KeyAcceleration = `acceleration`
	KeyMaximum      = `maximum`

	KeyTimePeriod1 = `timeperiod1`
	KeyTimePeriod2 = `timeperiod2`
	KeyTimePeriod3 = `timeperiod3`

	KeySymbols           = `SYMBOLS`
	KeyRange             = `RANGE`
	KeyIntervalAnalytics = `INTERVAL` // Uppercase INTERVAL for analytics endpoints (different from lowercase KeyInterval)
	KeyWindowSize        = `WINDOW_SIZE`
	KeyCalculations      = `CALCULATIONS`

	KeyQuarter       = `quarter`
	KeyTickers       = `tickers`
	KeyTopics        = `topics`
	KeyTimeFrom      = `time_from`
	KeyTimeTo        = `time_to`
	KeyLimit         = `limit`
	KeyContract      = `contract`
	KeyRequireGreeks = `require_greeks`

	KeyAdjusted      = `adjusted`
	KeyExtendedHours = `extended_hours`
	KeyKeywords      = `keywords`

	dataTypeValueJSON = `json`
	dataTypeValueCSV  = `csv`

	booleanValueTrue  = `true`
	booleanValueFalse = `false`

	outputSizeValueCompact = `compact`
	outputSizeValueFull    = `full`

	seriesTypeValueClose = `close`
	seriesTypeValueOpen  = `open`
	seriesTypeValueHigh  = `high`
	seriesTypeValueLow   = `low`

	stateValueActive   = `active`
	stateValueDelisted = `delisted`

	horizonValue3Month  = `3month`
	horizonValue6Month  = `6month`
	horizonValue12Month = `12month`

	sortValueLatest    = `LATEST`
	sortValueEarliest  = `EARLIEST`
	sortValueRelevance = `RELEVANCE`

	maturityValue3Month = `3month`
	maturityValue2Year  = `2year`
	maturityValue5Year  = `5year`
	maturityValue7Year  = `7year`
	maturityValue10Year = `10year`
	maturityValue30Year = `30year`

	intervalOption1min       = "1min"
	intervalOption5min       = "5min"
	intervalOption15min      = "15min"
	intervalOption30min      = "30min"
	intervalOption60min      = "60min"
	intervalOptionDaily      = "daily"
	intervalOptionWeekly     = "weekly"
	intervalOptionMonthly    = "monthly"
	intervalOptionQuarterly  = "quarterly"
	intervalOptionSemiannual = "semiannual"
	intervalOptionAnnual     = "annual"
)

type urlValues = map[string][]string

var _ = urlValues(url.Values{})

func encode[T ~urlValues](q T) string {
	return url.Values(q).Encode()
}

func dataType[T ~urlValues](q T, dt DataTypeOption) T {
	q[KeyDataType] = dt.values()
	return q
}

func dataTypeJSON[T interface {
	~urlValues
	DataType(DataTypeOption) T
}](q T) T {
	return q.DataType(DatatypeOptionJSON)
}

func dataTypeCSV[T interface {
	~urlValues
	DataType(DataTypeOption) T
}](q T) T {
	return q.DataType(DatatypeOptionCSV)
}

func boolean[T ~urlValues](q T, k string, b bool) T {
	if b {
		q[k] = []string{booleanValueTrue}
	} else {
		q[k] = []string{booleanValueFalse}
	}
	return q
}

func encodeMonth[T interface {
	~urlValues
	MonthString(string) T
}](q T, year int, m time.Month) T {
	return q.MonthString(fmt.Sprintf("%04d-%02d", year, m))
}
