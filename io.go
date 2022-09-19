package alphavantage

import "io"

func closeAndIgnoreError(c io.Closer) {
	_ = c.Close()
}

type multiReadCloser struct {
	io.Reader
	close func() error
}

func (mrc multiReadCloser) Close() error {
	return mrc.close()
}
