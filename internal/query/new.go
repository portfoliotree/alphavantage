package query

import "net/url"

type Values interface {
	Values() url.Values
	Encode() string
	Validate() error
}

func New(function string) Values {
	switch function {
	case FunctionAD:
		return make(AD)
	case FunctionADOSC:
		return make(ADOSC)
	case FunctionADX:
		return make(ADX)
	case FunctionADXR:
		return make(ADXR)
	case FunctionAPO:
		return make(APO)
	case FunctionAROON:
		return make(AROON)
	case FunctionAROONOSC:
		return make(AROONOSC)
	case FunctionATR:
		return make(ATR)
	case FunctionAllCommodities:
		return make(AllCommodities)
	case FunctionAluminum:
		return make(Aluminum)
	case FunctionAnalyticsFixedWindow:
		return make(AnalyticsFixedWindow)
	case FunctionAnalyticsSlidingWindow:
		return make(AnalyticsSlidingWindow)
	case FunctionBBANDS:
		return make(BBANDS)
	case FunctionBOP:
		return make(BOP)
	case FunctionBalanceSheet:
		return make(BalanceSheet)
	case FunctionBrent:
		return make(Brent)
	case FunctionCCI:
		return make(CCI)
	case FunctionCMO:
		return make(CMO)
	case FunctionCPI:
		return make(CPI)
	case FunctionCashFlow:
		return make(CashFlow)
	case FunctionCoffee:
		return make(Coffee)
	case FunctionCopper:
		return make(Copper)
	case FunctionCorn:
		return make(Corn)
	case FunctionCotton:
		return make(Cotton)
	case FunctionCryptoIntraday:
		return make(CryptoIntraday)
	case FunctionCurrencyExchangeRate:
		return make(CurrencyExchangeRate)
	case FunctionDEMA:
		return make(DEMA)
	case FunctionDX:
		return make(DX)
	case FunctionDigitalCurrencyDaily:
		return make(DigitalCurrencyDaily)
	case FunctionDigitalCurrencyMonthly:
		return make(DigitalCurrencyMonthly)
	case FunctionDigitalCurrencyWeekly:
		return make(DigitalCurrencyWeekly)
	case FunctionDividends:
		return make(Dividends)
	case FunctionDurables:
		return make(DurablesQuery)
	case FunctionEMA:
		return make(EMA)
	case FunctionETFProfile:
		return make(ETFProfile)
	case FunctionEarnings:
		return make(Earnings)
	case FunctionEarningsCalendar:
		return make(EarningsCalendar)
	case FunctionEarningsCallTranscript:
		return make(EarningsCallTranscript)
	case FunctionEarningsEstimates:
		return make(EarningsEstimates)
	case FunctionFXDaily:
		return make(FXDaily)
	case FunctionFXIntraday:
		return make(FXIntraday)
	case FunctionFXMonthly:
		return make(FXMonthly)
	case FunctionFXWeekly:
		return make(FXWeekly)
	case FunctionFederalFundsRate:
		return make(FederalFundsRate)
	case FunctionGlobalQuote:
		return make(GlobalQuote)
	case FunctionHilbertTransformDCPeriod:
		return make(HilbertTransformDCPeriod)
	case FunctionHilbertTransformDCPhase:
		return make(HilbertTransformDCPhase)
	case FunctionHilbertTransformPhasor:
		return make(HilbertTransformPhasor)
	case FunctionHilbertTransformSine:
		return make(HilbertTransformSine)
	case FunctionHilbertTransformTrendLine:
		return make(HilbertTransformTrendLine)
	case FunctionHilbertTransformTrendMode:
		return make(HilbertTransformTrendMode)
	case FunctionHistoricalOptions:
		return make(HistoricalOptions)
	case FunctionIPOCalendar:
		return make(IPOCalendar)
	case FunctionIncomeStatement:
		return make(IncomeStatement)
	case FunctionInflation:
		return make(InflationQuery)
	case FunctionInsiderTransactions:
		return make(InsiderTransactions)
	case FunctionKAMA:
		return make(KAMA)
	case FunctionListingStatus:
		return make(ListingStatus)
	case FunctionMACD:
		return make(MACD)
	case FunctionMACDEXT:
		return make(MACDEXT)
	case FunctionMAMA:
		return make(MAMA)
	case FunctionMFI:
		return make(MFI)
	case FunctionMidpoint:
		return make(Midpoint)
	case FunctionMIDPRICE:
		return make(MIDPRICE)
	case FunctionMINUSDI:
		return make(MINUSDI)
	case FunctionMINUSDM:
		return make(MINUSDM)
	case FunctionMOM:
		return make(MOM)
	case FunctionMarketStatus:
		return make(MarketStatus)
	case FunctionNATR:
		return make(NATR)
	case FunctionNaturalGas:
		return make(NaturalGas)
	case FunctionNewsSentiment:
		return make(NewsSentiment)
	case FunctionNonfarmPayroll:
		return make(NonfarmPayroll)
	case FunctionOBV:
		return make(OBV)
	case FunctionOverview:
		return make(Overview)
	case FunctionPLUSDI:
		return make(PLUSDI)
	case FunctionPLUSDM:
		return make(PLUSDM)
	case FunctionPPO:
		return make(PPO)
	case FunctionROC:
		return make(ROC)
	case FunctionROCR:
		return make(ROCR)
	case FunctionRSI:
		return make(RSI)
	case FunctionRealGDP:
		return make(RealGDP)
	case FunctionRealGDPPerCapita:
		return make(RealGDPPerCapita)
	case FunctionRealtimeBulkQuotes:
		return make(RealtimeBulkQuotes)
	case FunctionRealtimeOptions:
		return make(RealtimeOptions)
	case FunctionRetailSales:
		return make(RetailSalesQuery)
	case FunctionSAR:
		return make(SAR)
	case FunctionSMA:
		return make(SMA)
	case FunctionSTOCH:
		return make(STOCH)
	case FunctionSTOCHF:
		return make(STOCHF)
	case FunctionSTOCHRSI:
		return make(STOCHRSI)
	case FunctionSharesOutstanding:
		return make(SharesOutstanding)
	case FunctionSplits:
		return make(Splits)
	case FunctionSugar:
		return make(Sugar)
	case FunctionSymbolSearch:
		return make(SymbolSearch)
	case FunctionT3:
		return make(T3)
	case FunctionTEMA:
		return make(TEMA)
	case FunctionTRANGE:
		return make(TRANGE)
	case FunctionTRIMA:
		return make(TRIMA)
	case FunctionTRIX:
		return make(TRIX)
	case FunctionTimeSeriesDaily:
		return make(TimeSeriesDaily)
	case FunctionTimeSeriesDailyAdjusted:
		return make(TimeSeriesDailyAdjusted)
	case FunctionTimeSeriesIntraday:
		return make(TimeSeriesIntraday)
	case FunctionTimeSeriesMonthly:
		return make(TimeSeriesMonthly)
	case FunctionTimeSeriesMonthlyAdjusted:
		return make(TimeSeriesMonthlyAdjusted)
	case FunctionTimeSeriesWeekly:
		return make(TimeSeriesWeekly)
	case FunctionTimeSeriesWeeklyAdjusted:
		return make(TimeSeriesWeeklyAdjusted)
	case FunctionTopGainersLosers:
		return make(TopGainersLosers)
	case FunctionTreasuryYield:
		return make(TreasuryYield)
	case FunctionULTOSC:
		return make(ULTOSC)
	case FunctionUnemployment:
		return make(Unemployment)
	case FunctionVWAP:
		return make(VWAP)
	case FunctionWILLR:
		return make(WILLR)
	case FunctionWMA:
		return make(WMA)
	case FunctionWTI:
		return make(WTI)
	case FunctionWheat:
		return make(Wheat)
	default:
		return nil
	}
}

func FunctionOptions() []string {
	return []string{
		FunctionAD,
		FunctionADOSC,
		FunctionADX,
		FunctionADXR,
		FunctionAllCommodities,
		FunctionAluminum,
		FunctionAnalyticsFixedWindow,
		FunctionAnalyticsSlidingWindow,
		FunctionAPO,
		FunctionAROON,
		FunctionAROONOSC,
		FunctionATR,
		FunctionBalanceSheet,
		FunctionBBANDS,
		FunctionBOP,
		FunctionBrent,
		FunctionCashFlow,
		FunctionCCI,
		FunctionCMO,
		FunctionCoffee,
		FunctionCopper,
		FunctionCorn,
		FunctionCotton,
		FunctionCPI,
		FunctionCryptoIntraday,
		FunctionCurrencyExchangeRate,
		FunctionDEMA,
		FunctionDigitalCurrencyDaily,
		FunctionDigitalCurrencyMonthly,
		FunctionDigitalCurrencyWeekly,
		FunctionDividends,
		FunctionDurables,
		FunctionDX,
		FunctionEarnings,
		FunctionEarningsCalendar,
		FunctionEarningsCallTranscript,
		FunctionEarningsEstimates,
		FunctionEMA,
		FunctionETFProfile,
		FunctionFederalFundsRate,
		FunctionFXDaily,
		FunctionFXIntraday,
		FunctionFXMonthly,
		FunctionFXWeekly,
		FunctionGlobalQuote,
		FunctionHistoricalOptions,
		FunctionHilbertTransformDCPeriod,
		FunctionHilbertTransformDCPhase,
		FunctionHilbertTransformPhasor,
		FunctionHilbertTransformSine,
		FunctionHilbertTransformTrendLine,
		FunctionHilbertTransformTrendMode,
		FunctionIncomeStatement,
		FunctionInflation,
		FunctionInsiderTransactions,
		FunctionIPOCalendar,
		FunctionKAMA,
		FunctionListingStatus,
		FunctionMACD,
		FunctionMACDEXT,
		FunctionMAMA,
		FunctionMarketStatus,
		FunctionMFI,
		FunctionMidpoint,
		FunctionMIDPRICE,
		FunctionMINUSDI,
		FunctionMINUSDM,
		FunctionMOM,
		FunctionNATR,
		FunctionNaturalGas,
		FunctionNewsSentiment,
		FunctionNonfarmPayroll,
		FunctionOBV,
		FunctionOverview,
		FunctionPLUSDI,
		FunctionPLUSDM,
		FunctionPPO,
		FunctionRealtimeBulkQuotes,
		FunctionRealtimeOptions,
		FunctionRealGDP,
		FunctionRealGDPPerCapita,
		FunctionRetailSales,
		FunctionROC,
		FunctionROCR,
		FunctionRSI,
		FunctionSAR,
		FunctionSharesOutstanding,
		FunctionSMA,
		FunctionSplits,
		FunctionSTOCH,
		FunctionSTOCHF,
		FunctionSTOCHRSI,
		FunctionSugar,
		FunctionSymbolSearch,
		FunctionT3,
		FunctionTEMA,
		FunctionTimeSeriesDaily,
		FunctionTimeSeriesDailyAdjusted,
		FunctionTimeSeriesIntraday,
		FunctionTimeSeriesMonthly,
		FunctionTimeSeriesMonthlyAdjusted,
		FunctionTimeSeriesWeekly,
		FunctionTimeSeriesWeeklyAdjusted,
		FunctionTopGainersLosers,
		FunctionTRANGE,
		FunctionTreasuryYield,
		FunctionTRIMA,
		FunctionTRIX,
		FunctionULTOSC,
		FunctionUnemployment,
		FunctionVWAP,
		FunctionWheat,
		FunctionWILLR,
		FunctionWMA,
		FunctionWTI,
	}
}
