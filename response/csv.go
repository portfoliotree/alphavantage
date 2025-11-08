package response

type AD struct {
	Date  RawTime   `column-name:"time" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"Chaikin A/D"`
}
type ADOSC struct {
	Date  RawTime   `column-name:"time" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"ADOSC"`
}
type ADX struct {
	Date  RawTime   `column-name:"time" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"ADX"`
}
type ADXR struct {
	Date  RawTime   `column-name:"time" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"ADXR"`
}
type AllCommodities struct {
	Date  RawTime   `column-name:"timestamp" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"value"`
}
type Aluminum struct {
	Date  RawTime   `column-name:"timestamp" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"value"`
}
type APO struct {
	Date  RawTime   `column-name:"timestamp" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"APO"`
}
type AROON struct {
	Date RawTime   `column-name:"time" time-layout:"2006-01-02"`
	Down RawNumber `column-name:"Aroon Down"`
	Up   RawNumber `column-name:"Aroon Up"`
}
type AROONOSC struct {
	Date  RawTime   `column-name:"timestamp" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"AROONOSC"`
}
type ATR struct {
	Date  RawTime   `column-name:"timestamp" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"ATR"`
}
type BollingerBands struct {
	Date           RawTime   `column-name:"time" time-layout:"2006-01-02"`
	RealLowerBand  RawNumber `column-name:"Real Lower Band"`
	RealMiddleBand RawNumber `column-name:"Real Middle Band"`
	RealUpperBand  RawNumber `column-name:"Real Upper Band"`
}
type BRENT struct {
	Date  RawTime   `column-name:"timestamp" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"value"`
}
type CCI struct {
	Date  RawTime   `column-name:"time" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"CCI"`
}
type CMO struct {
	Date  RawTime   `column-name:"time" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"CMO"`
}
type Coffee struct {
	Date  RawTime   `column-name:"timestamp" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"value"`
}
type Copper struct {
	Date  RawTime   `column-name:"timestamp" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"value"`
}
type Corn struct {
	Date  RawTime   `column-name:"timestamp" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"value"`
}
type Cotton struct {
	Date  RawTime   `column-name:"timestamp" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"value"`
}
type CPI struct {
	Date  RawTime   `column-name:"timestamp" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"value"`
}
type CryptoIntraday struct {
	Time   RawTime   `column-name:"timestamp" time-layout:"2006-01-02 15:04:05"`
	Open   RawNumber `column-name:"open"`
	High   RawNumber `column-name:"high"`
	Low    RawNumber `column-name:"low"`
	Close  RawNumber `column-name:"close"`
	Volume RawNumber `column-name:"volume"`
}
type DEMA struct {
	Date  RawTime   `column-name:"time"`
	Value RawNumber `column-name:"DEMA"`
}
type DigitalCurrencyDaily struct {
	Date   RawTime   `column-name:"timestamp" time-layout:"2006-01-02"`
	Open   RawNumber `column-name:"open"`
	High   RawNumber `column-name:"high"`
	Low    RawNumber `column-name:"low"`
	Close  RawNumber `column-name:"close"`
	Volume RawNumber `column-name:"volume"`
}
type DigitalCurrencyMonthly struct {
	Date   RawTime   `column-name:"timestamp" time-layout:"2006-01-02"`
	Open   RawNumber `column-name:"open"`
	High   RawNumber `column-name:"high"`
	Low    RawNumber `column-name:"low"`
	Close  RawNumber `column-name:"close"`
	Volume RawNumber `column-name:"volume"`
}
type DigitalCurrencyWeekly struct {
	Date   RawTime   `column-name:"timestamp" time-layout:"2006-01-02"`
	Open   RawNumber `column-name:"open"`
	High   RawNumber `column-name:"high"`
	Low    RawNumber `column-name:"low"`
	Close  RawNumber `column-name:"close"`
	Volume RawNumber `column-name:"volume"`
}
type Dividends struct {
	ExDividendDate  RawTime   `colum-name:"ex_dividend_date" time-layout:"2006-01-02"`
	DeclarationDate RawTime   `colum-name:"declaration_date"`
	RecordDate      RawTime   `colum-name:"record_date" time-layout:"2006-01-02"`
	PaymentDate     RawTime   `colum-name:"payment_date" time-layout:"2006-01-02"`
	Amount          RawNumber `column-name:"amount"`
}
type Durables struct {
	Date  RawTime   `column-name:"timestamp" time-layout:"2006-01-02"`
	Value RawNumber `column-name:"value"`
}
type DX struct{}
type EarningsCalendar struct{}
type EMA struct{}
type FederalFundsRate struct{}
type FXDaily struct{}
type FXIntraday struct{}
type FXMonthly struct{}
type FXWeekly struct{}
type HistoricalOptions struct{}
type HilbertTransformDCPeriod struct{}
type HilbertTransformDCPhase struct{}
type HilbertTransformPhasor struct{}
type HilbertTransformSine struct{}
type HilbertTransformTrendLine struct{}
type HilbertTransformTrendMode struct{}
type Inflation struct{}
type IPOCalendar struct{}
type KAMA struct{}
type ListingStatus struct{}
type MACD struct{}
type MACDEXT struct{}
type MAMA struct{}
type MFI struct{}
type Midpoint struct{}
type Midprice struct{}
type MinusDirectionalIndicator struct{}
type MinusDirectionalMovement struct{}
type MOM struct{}
type NATR struct{}
type NaturalGas struct{}
type NonfarmPayroll struct{}
type OBV struct{}
type PLUS_DI struct{}
type PLUS_DM struct{}
type PPO struct{}
type RealGDP struct{}
type RealGDPPerCapita struct{}
type RetailSales struct{}
type ROC struct{}
type ROCR struct{}
type RSI struct{}
type SAR struct{}
type SharesOutstanding struct{}
type SMA struct{}
type Splits struct{}
type STOCH struct{}
type STOCHF struct{}
type STOCHRSI struct{}
type Sugar struct{}
type SymbolSearch struct{}
type T3 struct{}
type TEMA struct{}
type TimeSeriesDaily struct{}
type TimeSeriesDailyAdjusted struct{}
type TimeSeriesIntraday struct{}
type TimeSeriesMonthly struct{}
type TimeSeriesMonthlyAdjusted struct{}
type TimeSeriesWeekly struct{}
type TimeSeriesWeeklyAdjusted struct{}
type TRANGE struct{}
type TreasuryYield struct{}
type TRIMA struct{}
type TRIX struct{}
type ULTOSC struct{}
type Unemployment struct{}
type VWAP struct{}
type Wheat struct{}
type WILLR struct{}
type WMA struct{}
type WTI struct{}
