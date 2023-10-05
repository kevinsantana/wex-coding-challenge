package config

import (
	"context"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Host            string        `envconfig:"HOST" default:"0.0.0.0"`
	Port            string        `envconfig:"PORT" default:"3060"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"10s"`
}

type Database struct {
	ConnMaxLifetime time.Duration `envconfig:"DATABASE_CONNMAXLIFETIME" default:"4m"`
	MaxOpenConns    int           `envconfig:"DATABASE_MAXOPENCONNS" default:"25"`
	MaxIdleConns    int           `envconfig:"DATABASE_MAXIDLECONNS" default:"25"`
	Host            string        `envconfig:"DATABASE_HOST" default:"postgres://postgres:secret@172.17.0.1:5432/purchase?sslmode=disable"`
}

type Config struct {
	ApplicationVersion string
	Server             Server
	Database           Database
}

func InitConfig(ctx context.Context) *Config {
	conf := &Config{}
	_ = godotenv.Load()

	err := envconfig.Process("", conf)
	if err != nil {
		log.WithError(err).Panic("Error loading .env file")
	}

	log.WithField("Config", conf).
		Info("Success on loading .env file")

	return conf

}
