package mocks

import "github.com/stretchr/testify/mock"

type headerGetterMock struct {
	mock.Mock
}

func NewHeaderGetterMock() *headerGetterMock {
	return &headerGetterMock{}
}

func (m *headerGetterMock) Get(key string, defaultValue ...string) string {
	args := m.Called(key, defaultValue)

	return args.Get(0).(string)
}

type ctxLocalerMock struct {
	mock.Mock
}

func NewCtxLocalerMock() *ctxLocalerMock {
	return &ctxLocalerMock{}
}

func (m *ctxLocalerMock) Locals(key interface{}, value ...interface{}) interface{} {
	args := m.Called(key, value)

	return args.Get(0)
}
