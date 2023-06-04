package gcloud

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
)

func NewFirebaseApp(cfg Config) (*firebase.App, error) {
	config := &firebase.Config{
		ProjectID: cfg.ProjectID,
	}

	app, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("[gcloud.NewFirebaseApp] Error initializing Firebase app: %w", err)
	}

	return app, nil
}
