package alphavantage

import (
	"bytes"
	"io"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_checkError(t *testing.T) {
	t.Run("error message", func(t *testing.T) {
		rc := io.NopCloser(bytes.NewBufferString(`{"Error Message": "the parameter apikey is invalid or missing. Please claim your free API key on (https://www.alphavantage.co/support/#api-key). It should take less than 20 seconds."}`))

		please := NewWithT(t)
		_, err := checkError(rc)
		please.Expect(err).To(MatchError(ContainSubstring("the parameter apikey")))
	})

	t.Run("detail", func(t *testing.T) {
		rc := io.NopCloser(bytes.NewBufferString(`{"detail": "Could not satisfy the request Accept header."}`))
		please := NewWithT(t)
		_, err := checkError(rc)
		please.Expect(err).To(MatchError(ContainSubstring(": Could not satisfy")))
	})
}
