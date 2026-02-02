package dtk

import (
	"context"
	"time"

	"palbum/internal/domain/user"
)

type DTKRepository interface {
	FindByUserUUIDAndDate(ctx context.Context, userUUID user.UUID, date time.Time) (*DTK, error)
	FindByUserUUIDsAndSince(ctx context.Context, userUUIDs []user.UUID, since time.Time) ([]*DTK, error)
	Save(ctx context.Context, dtk *DTK) error
}
