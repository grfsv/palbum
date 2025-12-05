package persistence

import (
	"context"
	"remind_map/internal/domain/commons"

	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *Database {
	return &Database{db: db}
}

func (r *Database) GetDBFromContext(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(commons.TxKey).(*gorm.DB); ok {
		return tx
	}

	return r.db
}
