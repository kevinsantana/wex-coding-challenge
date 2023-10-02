package config

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Host               string `envconfig:"HOST" default:"127.0.0.1"`
	Port               string `envconfig:"PORT" default:"3060"`
	ApplicationVersion string
}

var Configuration Config


func InitConfig(ctx context.Context) {
	_ = godotenv.Load()

	if err := envconfig.Process("", &Configuration); err != nil {
		log.WithError(err).
		Error("Error loading .env file")
	}

	log.WithField("Config", Configuration).
	Info("Success on loading .env file")

}