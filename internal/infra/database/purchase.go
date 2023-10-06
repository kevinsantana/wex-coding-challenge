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
	sqlStmt := `INSERT INTO purchase_transaction (description, transaction_date, purchase_amount) VALUES ($1, $2, $3) RETURNING id;`

	var lastId int64
	var args []interface{}

	purchase.CreatedTime = utils.GetCurrentDate()

	args = append(args, purchase.Description)
	args = append(args, purchase.CreatedTime)
	args = append(args, purchase.Amount)

	rows, err := t.tx.QueryContext(ctx, sqlStmt, args...)

	var pqerr *pq.Error

	if err != nil {
		log.WithError(err).Error("error to save purchase transaction")

		if errors.As(err, &pqerr) && err.(*pq.Error).Code == "23505" {
			return domain.Purchase{}, share.ErrDuplicate
		}

		return domain.Purchase{}, share.ErrDatabase
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&lastId)
		
		if err != nil {
            log.WithError(err).Error("error to retrieve id")
        }

	}

	purchase.PurchaseId = lastId

	return purchase, nil
}

func NewDBPurchase(database *Database) DBPurchase {
	return DBPurchase{database: database}
}
