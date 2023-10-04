package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
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
			log.WithError(err).Panic("error to create migration driver instance")
		}

		m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("%s%s", "file://", FileMigrations), "purchase", driver)
		if err != nil {
			log.WithError(err).Panic("error to create migrator instance")
		}

		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.WithError(err).Panic("migrating database error")
		}
	}
	log.Info("database migrated success")
}
