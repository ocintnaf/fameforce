package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres connection config
type Config struct {
	Host     string `mapstructure:"host"`
	Port     uint   `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

// GetDSN returns a data source name string
func GetDSN(cfg Config) string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.User,
		cfg.Password,
	)
}

func GetDatabaseURL(cfg Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
}

// NewConnection establishes a new database connection with the given config.
func NewConnection(cfg Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(GetDSN(cfg)), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("[database.NewConnection] Error connecting to postgres instance: %w", err)
	}

	return db, nil
}

// CloseConnection closes the database connection.
func CloseConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
