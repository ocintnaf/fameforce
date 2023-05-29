package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
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

// NewConnection establishes a new database connection with the given config
func NewConnection(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", GetDSN(cfg))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
