package core

import (
	"context"

	"github.com/kevinsantana/wex-coding-challenge/internal/core/domain"
)

type PurchasePort interface {
	Start() (TxPurchasePort, error)
}

// transactional order operations
type TxPurchasePort interface {
	Rollback() error
	End() error
	SavePurchase(ctx context.Context, purchase domain.Purchase) (domain.Purchase, error)
}
