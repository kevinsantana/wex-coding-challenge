package database

import (
	"context"
	"database/sql"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

var (
	FileMigrations string = "./db/migration"
)

func RunMigrations(ctx context.Context, db *sql.DB) {
	PostgresMigrate(ctx, db)
}

func PostgresMigrate(ctx context.Context, db *sql.DB) {
	if os.Getenv("DATABASE_MIGRATE") == "true" {
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			log.WithError(err).Panic("creating migration driver instance")
		}

		m, err := migrate.NewWithDatabaseInstance("file://"+FileMigrations, "purchase", driver)
		if err != nil {
			log.WithError(err).Panic("creating migrator instance")
		}

		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.WithError(err).Panic("migrating database error")
		}
	}
}