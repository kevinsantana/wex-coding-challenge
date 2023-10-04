package middlewares

import (
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinsantana/wex-coding-challenge/internal/rest/handlers"
	log "github.com/sirupsen/logrus"
)

func Recover() fiber.Handler {
	return func(ctx *fiber.Ctx) (errHandler error) {
		defer func() {
			if err := recover(); err != nil {
				errHandler = err.(error)

				ctx.Response().SetStatusCode(500)

				log.
					WithField("request", ctx.UserContext()).
					WithField("stack", string(debug.Stack())).
					WithError(err.(error)).
					Error("Panic recovered")
			}
		}()

		if errHandler != nil {
			return errHandler
		}

		return ctx.Next()
	}
}

func ErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {

		code, res, isLog := handlers.GetResponseError(err)
		if isLog {
			log.
				WithFields(log.Fields{
					"path":          string(ctx.Request().URI().Path()),
					"status_code":   code,
					"client_ip":     ctx.IP(),
					"method":        string(ctx.Context().Method()),
					"userAgent":     string(ctx.Request().Header.UserAgent()),
					"response_body": res,
					"request_body":  string(ctx.Request().Body()),
				}).
				WithError(err).
				Error("Error not map")
		}

		return ctx.Status(code).JSON(res)
	}
}

func RouteNotFound() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).
			JSON(handlers.ResponseError{
				Code:    "API|ROUTE_NOT_FOUND",
				Message: "This route not found",
			})
	}
}
