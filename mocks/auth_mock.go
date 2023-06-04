package mocks

import (
	"context"

	"firebase.google.com/go/v4/auth"
	"github.com/stretchr/testify/mock"
)

type idTokenVerifierMock struct {
	mock.Mock
}

func NewIDTokenVerifierMock() *idTokenVerifierMock {
	return &idTokenVerifierMock{}
}

func (m *idTokenVerifierMock) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	args := m.Called(ctx, idToken)

	return args.Get(0).(*auth.Token), args.Error(1)
}
