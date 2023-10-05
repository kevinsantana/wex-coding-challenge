package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinsantana/wex-coding-challenge/internal/core/domain"
	"github.com/kevinsantana/wex-coding-challenge/internal/core/modules"
	"github.com/kevinsantana/wex-coding-challenge/internal/share"
	log "github.com/sirupsen/logrus"
)

type PurchaseHandler struct {
	core modules.Purchase
}

func (h PurchaseHandler) CreatePurchase(ctx *fiber.Ctx) error {
	var request domain.Purchase

	if err := ctx.BodyParser(&request); err != nil {
		log.WithError(err).
			WithField("request.body", string(ctx.Body())).
			Warn("Invalid request")

		return err
	}

	if errs := share.ValidateStruct(&request); errs != nil {
		log.WithField("validate.purchase", errs).
			Warn("Invalid pruchase transaction")

		return share.ErrValidation
	}

	purchase := domain.Purchase{
		Description: request.Description,
		Amount:      request.Amount,
	}

	p, err := h.core.Create(ctx.Context(), purchase)

	if err != nil {
		log.WithError(err).
			Error("Error creating purchase transaction")

		return err
	}

	return ctx.Status(http.StatusOK).JSON(p)
}

func NewPurchaseWebHandler(core modules.Purchase) PurchaseHandler {
	return PurchaseHandler{core: core}
}

func InitializePurchaseWebHanlder(core modules.Purchase) PurchaseHandler {
	purchaseWebHandler := NewPurchaseWebHandler(core)
	return purchaseWebHandler
}
