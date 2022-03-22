package alphavantage_test

import (
	"bytes"
	"context"
	_ "embed"
	"net/http"
	"strings"
	"testing"
	"time"

	Ω "github.com/onsi/gomega"

	"github.com/portfoliotree/alphavantage"
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
			_ = alphavantage.ParseCSV(bytes.NewReader(nil), nil, nil)
		}).To(Ω.Panic())
	})

	t.Run("non pointer data", func(t *testing.T) {
		please := Ω.NewWithT(t)
		please.Expect(func() {
			_ = alphavantage.ParseCSV(bytes.NewReader(nil), struct{}{}, nil)
		}).To(Ω.Panic())
	})

	t.Run("real data", func(t *testing.T) {
		please := Ω.NewWithT(t)

		var someFolks []struct {
			ID           int       `column-name:"id"`
			FirstInitial string    `column-name:"first_initial"`
			BirthDate    time.Time `column-name:"birth_date" time-layout:"2006/01/02"`
			Mass         float64   `column-name:"mass"`
		}

		err := alphavantage.ParseCSV(strings.NewReader(panthersCSV), &someFolks, nil)
		please.Expect(err).NotTo(Ω.HaveOccurred())
		please.Expect(someFolks).To(Ω.HaveLen(3))

		please.Expect(someFolks[0].ID).To(Ω.Equal(1))
		please.Expect(someFolks[0].FirstInitial).To(Ω.Equal("N"))
		please.Expect(someFolks[0].BirthDate).To(Ω.Equal(mustParseDate(t, "2020-02-17")))
		please.Expect(someFolks[0].Mass).To(Ω.Equal(70.0))

		please.Expect(someFolks[1].ID).To(Ω.Equal(2))
		please.Expect(someFolks[1].FirstInitial).To(Ω.Equal("S"))
		please.Expect(someFolks[1].BirthDate).To(Ω.Equal(mustParseDate(t, "2020-10-22")))
		please.Expect(someFolks[1].Mass).To(Ω.Equal(68.2))

		please.Expect(someFolks[2].ID).To(Ω.Equal(3))
		please.Expect(someFolks[2].FirstInitial).To(Ω.Equal("C"))
		please.Expect(someFolks[2].BirthDate).To(Ω.Equal(mustParseDate(t, "2021-08-31")))
		please.Expect(someFolks[2].Mass).To(Ω.Equal(72.9))
	})
}

const panthersCSV = `id,first_initial,birth_date,mass
1, N, 2020/02/17, 70
2, S, 2020/10/22, 68.2
3, C, 2021/08/31, 72.9
`

func mustParseDate(t *testing.T, date string) time.Time {
	tz, err := time.LoadLocation("US/Eastern")
	if err != nil {
		t.Fatal(err)
	}
	tm, err := time.ParseInLocation("2006-01-02", date, tz)
	if err != nil {
		t.Fatal(err)
	}
	return tm
}
