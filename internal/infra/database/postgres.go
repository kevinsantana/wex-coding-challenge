package database

import (
	"context"
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/kevinsantana/wex-coding-challenge/internal/config"
	"github.com/kevinsantana/wex-coding-challenge/internal/share"
)

type Database struct {
	db *sql.DB
}

func Connect(ctx context.Context, cfg *config.Config) *sql.DB {
	var err error

	db, err := sql.Open("nrpostgres", cfg.Database.Host)

	if err != nil {
		log.Panicf("connecting postgres: %+v", err)
	}

	db.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime)
	db.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Database.MaxIdleConns)

	if err := db.Ping(); err != nil {
		log.WithError(err).Panic("conn error with database")
	}

	log.Info("database connected")

	return db
}

func (m Database) Check(ctx context.Context) error {
	if err := m.db.PingContext(ctx); err != nil {
		log.WithError(err).Error("ping database")

		return share.ErrDatabase
	}

	return nil
}

func (m Database) Close() error {
	if err := m.db.Close(); err != nil {
		log.WithError(err).Error("close database")

		return share.ErrDatabase
	}

	return nil
}

func New(db *sql.DB) *Database { //nolint
	return &Database{
		db: db,
	}
}

func InitDb(ctx context.Context, cfg *config.Config) *Database {
	db := Connect(ctx, cfg)
	RunMigrations(ctx, db)

	return New(db)
}
