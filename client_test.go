package alphavantage_test

import (
	"bytes"
	"testing"

	Ω "github.com/onsi/gomega"

	"github.com/crhntr/alphavantage"
)

func TestService_ParseQueryResponse(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		please := Ω.NewGomegaWithT(t)
		const responseText = `timestamp,open,high,low,close,volume
2020-08-21,13.2600,13.3200,13.1500,13.2500,751279
2020-08-20,13.4700,13.4750,13.2200,13.3800,854559
2020-08-19,13.5700,13.7100,13.4700,13.5000,521089
2020-08-18,13.8100,13.8700,13.5400,13.5700,571445
`

		_, err := alphavantage.ParseStockQuery(bytes.NewReader([]byte(responseText)))
		please.Expect(err).NotTo(Ω.HaveOccurred())
	})

	t.Run("intra-day", func(t *testing.T) {
		please := Ω.NewGomegaWithT(t)
		const responseText = `timestamp,open,high,low,close,volume
2020-08-21 19:40:00,123.1700,123.1700,123.1700,123.1700,825
2020-08-21 19:20:00,123.2000,123.2000,123.2000,123.2000,200
2020-08-21 18:50:00,123.1700,123.1700,123.1700,123.1700,115
2020-08-21 17:30:00,123.0200,123.0200,123.0200,123.0200,200`
		_, err := alphavantage.ParseStockQuery(bytes.NewReader([]byte(responseText)))
		please.Expect(err).NotTo(Ω.HaveOccurred())
	})

	t.Run("unexpected column", func(t *testing.T) {
		please := Ω.NewGomegaWithT(t)
		const responseText = `timestamp,open,high,low,close,volume,unexpected
2020-08-21 19:40:00,123.1700,123.1700,123.1700,123.1700,825,123456789
2020-08-21 19:20:00,123.2000,123.2000,123.2000,123.2000,200,123456789
2020-08-21 18:50:00,123.1700,123.1700,123.1700,123.1700,115,123456789
2020-08-21 17:30:00,123.0200,123.0200,123.0200,123.0200,200,123456789`
		_, err := alphavantage.ParseStockQuery(bytes.NewReader([]byte(responseText)))
		please.Expect(err).NotTo(Ω.HaveOccurred())
	})
}
