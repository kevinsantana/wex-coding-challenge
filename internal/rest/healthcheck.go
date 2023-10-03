package rest

import (
	"net/http"


	"github.com/gofiber/fiber/v2"
	"github.com/kevinsantana/wex-coding-challenge/internal/infra/database"
	log "github.com/sirupsen/logrus"
)

type HealthWebHandler struct {
	database *database.Database
}

func (h HealthWebHandler) Readiness(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusOK)
}

func (h HealthWebHandler) Liveness(ctx *fiber.Ctx) error {
	if err := (*database.Database)(h.database).Check(ctx.Context()); err != nil {
		log.WithError(err).Errorf("database ping failed")

		return ctx.SendStatus(http.StatusServiceUnavailable)
	}

	return ctx.SendStatus(http.StatusOK)
}

func NewHealthWebHandler(database *database.Database) HealthWebHandler {
	return HealthWebHandler{database: database}
}

func InitializeHealthWeb(database *database.Database) HealthWebHandler {
	healthWebHandler := NewHealthWebHandler(database)
	return healthWebHandler
}