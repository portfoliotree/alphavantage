package alphavantage

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkError(t *testing.T) {
	t.Run("error message", func(t *testing.T) {
		rc := io.NopCloser(bytes.NewBufferString(`{"Error Message": "the parameter apikey is invalid or missing. Please claim your free API key on (https://www.alphavantage.co/support/#api-key). It should take less than 20 seconds."}`))

		_, err := checkError(rc)
		require.ErrorContains(t, err, "the parameter apikey")
	})

	t.Run("detail", func(t *testing.T) {
		rc := io.NopCloser(bytes.NewBufferString(`{"detail": "Could not satisfy the request Accept header."}`))
		_, err := checkError(rc)
		require.ErrorContains(t, err, "Could not satisfy")
	})
}
