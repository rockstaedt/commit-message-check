package mocks

import (
	"github.com/stretchr/testify/mock"
)

type FakeWriterMock struct {
	mock.Mock
}

func (fwm *FakeWriterMock) ResetCalls() {
	fwm.Calls = []mock.Call{}
	fwm.ExpectedCalls = []*mock.Call{}
}

func (fwm *FakeWriterMock) Write(p []byte) (int, error) {
	args := fwm.Called(p)

	return 0, args.Error(1)
}
