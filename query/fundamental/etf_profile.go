package fundamental

type ETFProfile struct {
	Symbol            string       `json:"symbol,omitempty"`
	NetAssets         string       `json:"net_assets,omitempty"`
	NetExpenseRatio   string       `json:"net_expense_ratio,omitempty"`
	PortfolioTurnover string       `json:"portfolio_turnover,omitempty"`
	DividendYield     string       `json:"dividend_yield,omitempty"`
	InceptionDate     string       `json:"inception_date,omitempty"`
	Leveraged         string       `json:"leveraged,omitempty"`
	Sectors           []ETFSector  `json:"sectors,omitempty"`
	Holdings          []ETFHolding `json:"holdings,omitempty"`
}

type ETFSector struct {
	Sector string `json:"sector,omitempty"`
	Weight string `json:"weight,omitempty"`
}

type ETFHolding struct {
	Symbol      string `json:"symbol,omitempty"`
	Description string `json:"description,omitempty"`
	Weight      string `json:"weight,omitempty"`
}
