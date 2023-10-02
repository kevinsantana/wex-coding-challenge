package server

import (
	"net/http"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"

	"github.com/kevinsantana/wex-coding-challenge/internal/rest"
	"github.com/kevinsantana/wex-coding-challenge/internal/rest/middlewares"
)

type Routes []Route

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc fiber.Handler
	Public      bool
	Scopes      []string
}

var healthCheck = Routes{
	{
		Name:        "Healthcheck",
		Method:      http.MethodGet,
		Pattern:     "/healthcheck",
		HandlerFunc: rest.Liveness,
		Public:      true,
	},
	{
		Name:        "Readiness",
		Method:      http.MethodGet,
		Pattern:     "/readiness",
		HandlerFunc: rest.Readiness,
		Public:      true,
	},
}

func Router() *fiber.App {
	r := fiber.New(fiber.Config{
		Prefork:               false,
		CaseSensitive:         false,
		StrictRouting:         false,
		ServerHeader:          "*",
		AppName:               "WEX TAG and Gateways Product",
		Immutable:             true,
		DisableStartupMessage: true,
		ErrorHandler:          middlewares.ErrorHandler(),
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
	})

	api := r.Group("/")
	for _, route := range healthCheck {
		api.Add(route.Method, route.Pattern, route.HandlerFunc)
	}

	v1 := api.Group("/api/v1")

	var routes []Route

	for _, route := range routes {
		v1.Add(route.Method, route.Pattern, route.HandlerFunc)
	}

	r.Use(middlewares.RouteNotFound())

	return r
}
