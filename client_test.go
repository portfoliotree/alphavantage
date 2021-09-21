package alphavantage_test

import (
	"bytes"
	"context"
	_ "embed"
	"net/http"
	"testing"

	Ω "github.com/onsi/gomega"

	"github.com/crhntr/alphavantage"
)

type doerFunc func(*http.Request) (*http.Response, error)

func (fn doerFunc) Do(req *http.Request) (*http.Response, error) { return fn(req) }

type waitFunc func(ctx context.Context) error

func (wf waitFunc) Wait(ctx context.Context) error {
	return wf(ctx)
}

func TestParse(t *testing.T) {
	t.Run("nil data", func(t *testing.T) {
		please := Ω.NewWithT(t)
		please.Expect(func() {
			_ = alphavantage.ParseCSV(bytes.NewReader(nil), nil)
		}).To(Ω.Panic())
	})

	t.Run("non pointer data", func(t *testing.T) {
		please := Ω.NewWithT(t)
		please.Expect(func() {
			_ = alphavantage.ParseCSV(bytes.NewReader(nil), struct{}{})
		}).To(Ω.Panic())
	})
}
