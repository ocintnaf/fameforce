package mocks

import (
	"context"

	"github.com/ocintnaf/fameforce/types"
	"github.com/stretchr/testify/mock"
)

type idTokenVerifierMock struct {
	mock.Mock
}

func NewIDTokenVerifierMock() *idTokenVerifierMock {
	return &idTokenVerifierMock{}
}

func (m *idTokenVerifierMock) VerifyIDToken(ctx context.Context, idToken string) (*types.BaseIDToken, error) {
	args := m.Called(ctx, idToken)

	return args.Get(0).(*types.BaseIDToken), args.Error(1)
}
