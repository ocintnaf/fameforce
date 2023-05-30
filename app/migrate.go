//go:build migrate

package app

import (
	"errors"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/ocintnaf/fameforce/config"
	"github.com/ocintnaf/fameforce/pkg/database"

	// migrate tools
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	defaultAttempts = 5
	defaultDelay    = time.Second * 5
)

func init() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal("Migrate: failed to load config. Error: ", err)
	}

	var (
		attempts = defaultAttempts
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://migrations", database.GetDatabaseURL(cfg.Database))
		if err == nil {
			break
		}

		log.Printf("Migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(defaultDelay)
		attempts--
	}

	if err != nil {
		log.Fatal("Migrate: failed to connect to postgres. Error: ", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("Migrate: failed to run migrations. Error: ", err)
	}

	defer m.Close()

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: no change")
		return
	}

	log.Println("Migrate: migrations ran successfully")
}
