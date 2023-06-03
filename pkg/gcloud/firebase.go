package gcloud

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
)

func NewFirebaseApp() (*firebase.App, error) {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("[gcloud.NewFirebaseApp] Error initializing Firebase app: %w", err)
	}

	return app, nil
}
