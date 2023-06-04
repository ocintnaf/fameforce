package logger

import (
	"strings"

	"go.uber.org/zap"
)

func NewLogger(env string) (*zap.Logger, error) {
	if strings.ToLower(env) == "development" {
		return zap.NewDevelopment()
	}

	return zap.NewProduction()
}
