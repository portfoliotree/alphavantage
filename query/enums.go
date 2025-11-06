package query

// IIntervalOption enumerates intraday interval options
// Valid values:  "1min", "5min", "15min", "30min", "60min"
type IIntervalOption string

const (
	IIntervalOption1min  IIntervalOption = intervalOption1min
	IIntervalOption5min  IIntervalOption = intervalOption5min
	IIntervalOption15min IIntervalOption = intervalOption15min
	IIntervalOption30min IIntervalOption = intervalOption30min
	IIntervalOption60min IIntervalOption = intervalOption60min
)

func IIntervalOptions() []IIntervalOption {
	return []IIntervalOption{
		IIntervalOption1min,
		IIntervalOption5min,
		IIntervalOption15min,
		IIntervalOption30min,
		IIntervalOption60min,
	}
}

func (o IIntervalOption) values() []string { return []string{string(o)} }

// DWMIntervalOption enumerates interval options
// Valid values: "daily", "weekly", "monthly"
type DWMIntervalOption string

const (
	DWMIntervalOptionDaily   DWMIntervalOption = intervalOptionDaily
	DWMIntervalOptionWeekly  DWMIntervalOption = intervalOptionWeekly
	DWMIntervalOptionMonthly DWMIntervalOption = intervalOptionMonthly
)

func DWMIntervalOptions() []DWMIntervalOption {
	return []DWMIntervalOption{
		DWMIntervalOptionDaily,
		DWMIntervalOptionWeekly,
		DWMIntervalOptionMonthly,
	}
}

func (o DWMIntervalOption) values() []string { return []string{string(o)} }

// MQAIntervalOption enumerates interval options
// Valid values: "annual", "quarterly", "monthly"
type MQAIntervalOption string

const (
	MQAIntervalOptionAnnual    MQAIntervalOption = intervalOptionAnnual
	MQAIntervalOptionQuarterly MQAIntervalOption = intervalOptionQuarterly
	MQAIntervalOptionMonthly   MQAIntervalOption = intervalOptionMonthly
)

func MQAIntervalOptions() []MQAIntervalOption {
	return []MQAIntervalOption{MQAIntervalOptionAnnual, MQAIntervalOptionQuarterly, MQAIntervalOptionMonthly}
}

func (o MQAIntervalOption) values() []string { return []string{string(o)} }

// QAIntervalOption enumerates interval options
// Valid values: "annual", "quarterly"
type QAIntervalOption string

const (
	QAIntervalOptionAnnual    QAIntervalOption = intervalOptionAnnual
	QAIntervalOptionQuarterly QAIntervalOption = intervalOptionQuarterly
)

func QAIntervalOptions() []QAIntervalOption {
	return []QAIntervalOption{QAIntervalOptionAnnual, QAIntervalOptionQuarterly}
}

func (o QAIntervalOption) values() []string { return []string{string(o)} }

// MSIntervalOption enumerates interval options
// Valid values: "monthly", "semiannual"
type MSIntervalOption string

const (
	MSIntervalOptionMonthly    MSIntervalOption = intervalOptionMonthly
	MSIntervalOptionSemiannual MSIntervalOption = intervalOptionSemiannual
)

func MSIntervalOptions() []MSIntervalOption {
	return []MSIntervalOption{MSIntervalOptionMonthly, MSIntervalOptionSemiannual}
}

func (o MSIntervalOption) values() []string { return []string{string(o)} }

// MaturityOption enumerates maturity timeline options for treasury yields
// Valid values: "3month", "2year", "5year", "7year", "10year", "30year"
type MaturityOption string

const (
	MaturityOption3Month MaturityOption = maturityValue3Month
	MaturityOption2Year  MaturityOption = maturityValue2Year
	MaturityOption5Year  MaturityOption = maturityValue5Year
	MaturityOption7Year  MaturityOption = maturityValue7Year
	MaturityOption10Year MaturityOption = maturityValue10Year
	MaturityOption30Year MaturityOption = maturityValue30Year
)

func MaturityOptions() []MaturityOption {
	return []MaturityOption{
		MaturityOption3Month,
		MaturityOption2Year,
		MaturityOption5Year,
		MaturityOption7Year,
		MaturityOption10Year,
		MaturityOption30Year,
	}
}

func (o MaturityOption) values() []string { return []string{string(o)} }

// SeriesTypeOption enumerates series type options
// Valid values: "close", "open", "high", "low"
type SeriesTypeOption string

const (
	SeriesTypeOptionClose SeriesTypeOption = seriesTypeValueClose
	SeriesTypeOptionOpen  SeriesTypeOption = seriesTypeValueOpen
	SeriesTypeOptionHigh  SeriesTypeOption = seriesTypeValueHigh
	SeriesTypeOptionLow   SeriesTypeOption = seriesTypeValueLow
)

func SeriesTypeOptions() []SeriesTypeOption {
	return []SeriesTypeOption{
		SeriesTypeOptionClose,
		SeriesTypeOptionOpen,
		SeriesTypeOptionHigh,
		SeriesTypeOptionLow,
	}
}

func (o SeriesTypeOption) values() []string { return []string{string(o)} }

// OutputSizeOption enumerates output size options
// Valid values: "compact", "full"
type OutputSizeOption string

const (
	OutputSizeOptionCompact OutputSizeOption = outputSizeValueCompact
	OutputSizeOptionFull    OutputSizeOption = outputSizeValueFull
)

func OutputSizeOptions() []OutputSizeOption {
	return []OutputSizeOption{OutputSizeOptionCompact, OutputSizeOptionFull}
}

func (o OutputSizeOption) values() []string { return []string{string(o)} }

// StateOption enumerates state options for listing status
// Valid values: "active", "delisted"
type StateOption string

const (
	StateOptionActive   StateOption = stateValueActive
	StateOptionDelisted StateOption = stateValueDelisted
)

func StateOptions() []StateOption {
	return []StateOption{StateOptionActive, StateOptionDelisted}
}

func (o StateOption) values() []string { return []string{string(o)} }

// HorizonOption enumerates horizon options for earnings calendar
// Valid values: "3month", "6month", "12month"
type HorizonOption string

const (
	HorizonOption3Month  HorizonOption = horizonValue3Month
	HorizonOption6Month  HorizonOption = horizonValue6Month
	HorizonOption12Month HorizonOption = horizonValue12Month
)

func HorizonOptions() []HorizonOption {
	return []HorizonOption{
		HorizonOption3Month,
		HorizonOption6Month,
		HorizonOption12Month,
	}
}

func (o HorizonOption) values() []string { return []string{string(o)} }

// SortOption enumerates sort options for news sentiment
// Valid values: "LATEST", "EARLIEST", "RELEVANCE"
type SortOption string

const (
	SortOptionLatest    SortOption = sortValueLatest
	SortOptionEarliest  SortOption = sortValueEarliest
	SortOptionRelevance SortOption = sortValueRelevance
)

func SortOptions() []SortOption {
	return []SortOption{
		SortOptionLatest,
		SortOptionEarliest,
		SortOptionRelevance,
	}
}

func (o SortOption) values() []string { return []string{string(o)} }

// DataTypeOption enumerates dataType options
// Valid values: "json", "csv"
type DataTypeOption string

const (
	DatatypeOptionJSON DataTypeOption = dataTypeValueJSON
	DatatypeOptionCSV  DataTypeOption = dataTypeValueCSV
)

func DataTypeOptions() []DataTypeOption {
	return []DataTypeOption{DatatypeOptionJSON, DatatypeOptionCSV}
}

func (o DataTypeOption) values() []string { return []string{string(o)} }

type IDWMIntervalOption string

const (
	IDWMIntervalOption1min    IDWMIntervalOption = intervalOption1min
	IDWMIntervalOption5min    IDWMIntervalOption = intervalOption5min
	IDWMIntervalOption15min   IDWMIntervalOption = intervalOption15min
	IDWMIntervalOption30min   IDWMIntervalOption = intervalOption30min
	IDWMIntervalOption60min   IDWMIntervalOption = intervalOption60min
	IDWMIntervalOptionDaily   IDWMIntervalOption = intervalOptionDaily
	IDWMIntervalOptionWeekly  IDWMIntervalOption = intervalOptionWeekly
	IDWMIntervalOptionMonthly IDWMIntervalOption = intervalOptionMonthly
)

func IDWMIntervalOptions() []IDWMIntervalOption {
	return []IDWMIntervalOption{
		IDWMIntervalOption1min,
		IDWMIntervalOption5min,
		IDWMIntervalOption15min,
		IDWMIntervalOption30min,
		IDWMIntervalOption60min,
		IDWMIntervalOptionDaily,
		IDWMIntervalOptionWeekly,
		IDWMIntervalOptionMonthly,
	}
}

func (o IDWMIntervalOption) values() []string { return []string{string(o)} }
