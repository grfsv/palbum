package persistence

import (
	"context"
	"palbum/internal/domain/commons"

	"gorm.io/gorm"
)

type GormTransactionManager struct {
	db *gorm.DB
}

func NewTransactionManager(db *gorm.DB) commons.TxManagerRepository {
	return &GormTransactionManager{db: db}
}

func (tm *GormTransactionManager) WithInTx(ctx context.Context, function func(ctx context.Context) error) error {
	txDB := tm.db.Begin()
	if txDB.Error != nil {
		return txDB.Error
	}

	txCtx := context.WithValue(ctx, commons.TxKey, txDB)

	err := function(txCtx)
	if err != nil {
		txDB.Rollback()

		return err
	}

	return txDB.Commit().Error
}
