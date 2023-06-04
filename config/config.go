package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/ocintnaf/fameforce/pkg/database"
	"github.com/ocintnaf/fameforce/pkg/gcloud"
	"github.com/spf13/viper"
)

type Config struct {
	Database database.Config `mapstructure:"database"`
	GCloud   gcloud.Config   `mapstructure:"gcp"`
}

func Init() (*Config, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Load .env file (for local development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(cwd + "/config/config.yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config

	err = viper.Unmarshal(&config)

	return &config, err
}
