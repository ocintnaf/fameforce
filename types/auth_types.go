package types

import "context"

type BaseIDToken struct {
	UID string
}

type IDTokenVerifier interface {
	VerifyIDToken(ctx context.Context, idToken string) (*BaseIDToken, error)
}
