package response

import (
	"cmp"
	"encoding/json"
	"time"
)

type RawTime string

func (t RawTime) Time(format string) (time.Time, error) {
	format = cmp.Or(format, "2006-01-02")
	return time.Parse(format, string(t))
}

type RawNumber json.Number

func (n RawNumber) Int64() (int64, error)     { return json.Number(n).Int64() }
func (n RawNumber) String() string            { return json.Number(n).String() }
func (n RawNumber) Float64() (float64, error) { return json.Number(n).Float64() }
