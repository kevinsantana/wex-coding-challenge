package database

import (
	"context"
	"fmt"
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/kevinsantana/wex-coding-challenge/internal/config"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
)

var (
	FileMigrations string = "./db/migration"
)

func RunMigrations(ctx context.Context, db *sql.DB, cfg *config.Config) {
	PostgresMigrate(ctx, db, cfg)
}

func PostgresMigrate(ctx context.Context, db *sql.DB, cfg *config.Config) {
	if cfg.Database.Migrate {
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
		log.Info("database migrated")
	}
}
