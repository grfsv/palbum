package commons

import (
	"context"

	"github.com/cockroachdb/errors"
)

type TxManagerRepository interface {
	WithInTx(ctx context.Context, f func(ctx context.Context) error) error
}

type txContextKey string

const TxKey txContextKey = "gorm_tx_session"

var ErrNotFound = errors.New("record not found")
