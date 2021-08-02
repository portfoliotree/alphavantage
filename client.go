package alphavantage

import (
	"time"
)

var timezone *time.Location

func init() {
	var err error
	timezone, err = time.LoadLocation("US/Eastern")
	if err != nil {
		panic(err)
	}
}
