package mocks

import "errors"

type FakeWriter struct {
}

func (fw FakeWriter) Write(p []byte) (int, error) {
	return 0, errors.New("error at writing")
}
