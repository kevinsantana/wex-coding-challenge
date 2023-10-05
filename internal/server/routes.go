package server

import (
	"net/http"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"

	"github.com/kevinsantana/wex-coding-challenge/internal/rest"
	"github.com/kevinsantana/wex-coding-challenge/internal/rest/handlers"
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

func Router(health rest.HealthWebHandler, purchase handlers.PurchaseHandler) *fiber.App {
	var healthCheck = Routes{
		{
			Name:        "Healthcheck",
			Method:      http.MethodGet,
			Pattern:     "/healthcheck",
			HandlerFunc: health.Liveness,
			Public:      true,
		},
		{
			Name:        "Readiness",
			Method:      http.MethodGet,
			Pattern:     "/readiness",
			HandlerFunc: health.Readiness,
			Public:      true,
		},
	}

	var purchaseTransaction = Routes{
		{
			Name:        "Create Purchase Transaction",
			Method:      http.MethodPost,
			Pattern:     "/purchase",
			HandlerFunc: purchase.CreatePurchase,
			Public:      true,
		},
	}

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
	routes = append(routes, purchaseTransaction...)

	for _, route := range routes {
		v1.Add(route.Method, route.Pattern, route.HandlerFunc)
	}

	r.Use(middlewares.RouteNotFound())

	return r
}
