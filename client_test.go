package alphavantage_test

import (
	"bytes"
	"context"
	_ "embed"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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
		assert.Panics(t, func() {
			_ = alphavantage.ParseCSV(bytes.NewReader(nil), nil, nil)
		})
	})

	t.Run("non pointer data", func(t *testing.T) {
		assert.Panics(t, func() {
			_ = alphavantage.ParseCSV(bytes.NewReader(nil), struct{}{}, nil)
		})
	})

	t.Run("real data", func(t *testing.T) {

		var someFolks []struct {
			ID           int       `column-name:"id"`
			FirstInitial string    `column-name:"first_initial"`
			BirthDate    time.Time `column-name:"birth_date" time-layout:"2006/01/02"`
			Mass         float64   `column-name:"mass"`
		}

		err := alphavantage.ParseCSV(strings.NewReader(panthersCSV), &someFolks, nil)
		require.NoError(t, err)
		assert.Len(t, someFolks, 3)

		assert.Equal(t, 1, someFolks[0].ID)
		assert.Equal(t, "N", someFolks[0].FirstInitial)
		assert.Equal(t, mustParseDate(t, "2020-02-17"), someFolks[0].BirthDate)
		assert.Equal(t, 70.0, someFolks[0].Mass)

		assert.Equal(t, 2, someFolks[1].ID)
		assert.Equal(t, "S", someFolks[1].FirstInitial)
		assert.Equal(t, mustParseDate(t, "2020-10-22"), someFolks[1].BirthDate)
		assert.Equal(t, 68.2, someFolks[1].Mass)

		assert.Equal(t, 3, someFolks[2].ID)
		assert.Equal(t, "C", someFolks[2].FirstInitial)
		assert.Equal(t, mustParseDate(t, "2021-08-31"), someFolks[2].BirthDate)
		assert.Equal(t, 72.9, someFolks[2].Mass)
	})
}

const panthersCSV = `id,first_initial,birth_date,mass
1, N, 2020/02/17, 70
2, S, 2020/10/22, 68.2
3, C, 2021/08/31, 72.9
`

func mustParseDate(t *testing.T, date string) time.Time {
	tm, err := time.ParseInLocation(alphavantage.DefaultDateFormat, date, time.UTC)
	if err != nil {
		t.Fatal(err)
	}
	return tm
}
