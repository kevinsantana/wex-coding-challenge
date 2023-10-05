package modules

import (
	"context"

	"github.com/kevinsantana/wex-coding-challenge/internal/core"
	"github.com/kevinsantana/wex-coding-challenge/internal/core/domain"
	"github.com/kevinsantana/wex-coding-challenge/internal/infra/database"
	"github.com/kevinsantana/wex-coding-challenge/internal/share"
	log "github.com/sirupsen/logrus"
)

type Purchase interface {
	Create(ctx context.Context, purchase domain.Purchase) (savedPurchase domain.Purchase, err error)
}

type PurchaseCore struct {
	purchasePort core.PurchasePort
	database     *database.Database
}

func NewPurchaseCore(purchasePort core.PurchasePort, database *database.Database) PurchaseCore {
	return PurchaseCore{
		purchasePort: purchasePort,
		database:     database,
	}
}

func (c PurchaseCore) Create(ctx context.Context, purchase domain.Purchase) (savedPurchase domain.Purchase, err error) {
	log.Debug("starting saving purchase transcation to database")

	txnDB, _ := c.purchasePort.Start()
	savedPurchase, err = txnDB.SavePurchase(ctx, purchase)

	if err != nil {
		txnDB.Rollback() //nolint

		if err == share.ErrDuplicate {
			log.WithError(err).Error("duplicate, purchase transaction already exists")
		}

		// TODO: return already created purchase transaction

		log.WithError(err).Error("error create order into db")

		return domain.Purchase{}, share.ErrOrderNotCreated
	}

	txnDB.End() //nolint

	log.WithFields(map[string]interface{}{
		"purchaseId": savedPurchase.PurchaseId,
		"purchase":   savedPurchase,
	})

	log.Info("purchase transaction created")
	log.Debug("end save purchase transaction")

	return savedPurchase, nil
}
