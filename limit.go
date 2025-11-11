package alphavantage

import (
	"cmp"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/time/rate"
)

// RequestsPerMinute helps you set up an HTTP request rate limiter.
// See plan details here: https://www.alphavantage.co/premium/
type RequestsPerMinute int

const (
	PremiumPlan75   RequestsPerMinute = 75
	PremiumPlan150  RequestsPerMinute = 150
	PremiumPlan300  RequestsPerMinute = 300
	PremiumPlan600  RequestsPerMinute = 600
	PremiumPlan1200 RequestsPerMinute = 1200
)

func (plan RequestsPerMinute) String() string {
	return fmt.Sprintf("rate limit with %d requests per minute", plan)
}

func (plan RequestsPerMinute) Limit() rate.Limit {
	return rate.Every(time.Minute / time.Duration(cmp.Or(plan, 1)))
}

func NewRequestsPerMinute(s string) (RequestsPerMinute, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse rate limit %q: %w", s, err)
	}
	return RequestsPerMinute(n), nil
}
