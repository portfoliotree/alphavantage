package fakes

import (
	"net/http"
)

type Doer struct {
	CallCount int
	Recieves  struct {
		Req *http.Request
	}
	Returns struct {
		Res *http.Response
		Err error
	}
}

func (mock *Doer) Do(req *http.Request) (*http.Response, error) {
	mock.CallCount++

	mock.Recieves.Req = req
	return mock.Returns.Res, mock.Returns.Err
}
