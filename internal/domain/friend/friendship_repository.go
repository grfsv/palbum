package friend

import (
	"context"
	"palbum/internal/domain/user"
)

type FriendshipRepository interface {
	Save(ctx context.Context, friendship *Friendship) error
	FindByUserUUID(ctx context.Context, userUUID user.UUID) ([]*Friendship, error)
	Delete(ctx context.Context, friendship *Friendship) error
}
