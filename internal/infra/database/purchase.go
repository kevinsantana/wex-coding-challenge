package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kevinsantana/wex-coding-challenge/internal/core"
	"github.com/kevinsantana/wex-coding-challenge/internal/core/domain"
	"github.com/kevinsantana/wex-coding-challenge/internal/share"
	"github.com/kevinsantana/wex-coding-challenge/pkg/utils"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type DBPurchase struct {
	database *Database
}

type TxOrderAdapter struct {
	tx *sql.Tx
}

func (db DBPurchase) Start() (core.TxPurchasePort, error) { //nolint
	tx, err := db.database.db.Begin()
	if err != nil {
		log.WithError(err).Error("error begin purchase database")

		return nil, share.ErrDatabase
	}

	return TxOrderAdapter{
		tx: tx,
	}, nil
}

func (t TxOrderAdapter) Rollback() error {
	if err := t.tx.Rollback(); err != nil {
		log.WithError(err).Error("error rollback purchase database")

		return share.ErrDatabase
	}

	return nil
}

func (t TxOrderAdapter) End() error {
	if err := t.tx.Commit(); err != nil {
		log.WithError(err).Error("error commit purchase database")

		return share.ErrDatabase
	}

	return nil
}

func (t TxOrderAdapter) SavePurchase(ctx context.Context, purchase domain.Purchase) (domain.Purchase, error) {
	sqlStmt := `
	INSERT INTO purchase_transaction (
		description,
		transaction_date,
		purchase_amount,
	) VALUES ($1, $2, $3)
	`
	purchase.CreatedTime = utils.GetCurrentDate()
	_, err := t.tx.ExecContext(ctx, sqlStmt,
		purchase.Description,
		purchase.CreatedTime,
		purchase.Amount)

	var pqerr *pq.Error

	if err != nil {
		log.WithError(err).Error("error save order")

		if errors.As(err, &pqerr) && err.(*pq.Error).Code == "23505" {
			return domain.Purchase{}, share.ErrDuplicate
		}

		return domain.Purchase{}, share.ErrDatabase
	}

	return purchase, nil
}

func NewDBPurchase(database *Database) DBPurchase {
	return DBPurchase{database: database}
}
