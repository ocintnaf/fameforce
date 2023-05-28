package config

import (
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/ocintnaf/fameforce/pkg/database"
	"github.com/spf13/viper"
)

type Config struct {
	Database database.Config `mapstructure:"database"`
}

func Init() (*Config, error) {
	// Load .env file (for local development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	viper.SetConfigType("yaml")
	viper.SetConfigFile("config/config.yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config

	err := viper.Unmarshal(&config)

	return &config, err
}
