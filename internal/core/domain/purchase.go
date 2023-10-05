package domain

import (
	"github.com/shopspring/decimal"
)

type Purchase struct {
	Description string           `json:"description" binding:"required,max=50" example:"This is a purchase description"`
	CreatedTime string           `json:"createdTime" example:"2021-01-01 15:04:05"`
	Amount      *decimal.Decimal `json:"amount" binding:"required,decimal,positive,max_precision=64,max_scale=8" example:"0.02"`
	PurchaseId  int              `json:"-"`
}
