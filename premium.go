package alphavantage

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

// PremiumPlan helps you set up an HTTP request rate limiter.
// See plan details here: https://www.alphavantage.co/premium/
type PremiumPlan int

const (
	PremiumPlan75   PremiumPlan = 75
	PremiumPlan150  PremiumPlan = 150
	PremiumPlan300  PremiumPlan = 300
	PremiumPlan600  PremiumPlan = 600
	PremiumPlan1200 PremiumPlan = 1200
)

func (plan PremiumPlan) String() string {
	return fmt.Sprintf("premium plan with %d requests per minute", plan)
}

func (plan PremiumPlan) Limit() rate.Limit {
	return rate.Every(time.Minute / time.Duration(plan))
}
