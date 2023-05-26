package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/ocintnaf/fameforce/pkg/database"
	"github.com/spf13/viper"
)

type Config struct {
	Database database.Config `mapstructure:"database"`
}

func Init() (*Config, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(cwd + "/config/config.yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config

	err = viper.Unmarshal(&config)

	return &config, err
}
