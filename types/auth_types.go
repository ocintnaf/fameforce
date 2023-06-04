package types

import (
	"context"

	"firebase.google.com/go/v4/auth"
)

type BaseIDToken struct {
	UID string
}

type IDTokenVerifier interface {
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
}
