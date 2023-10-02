package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Readiness(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusOK)
}

func Liveness(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusOK)
}
