package friend

import (
	"context"
	"palbum/internal/domain/commons"
	"palbum/internal/domain/friend"
	"palbum/internal/infrastructure/persistence"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type FriendshipRepositoryImpl struct {
	*persistence.Database
}

func NewFriendshipRepositoryImpl(baseRepo *persistence.Database) friend.FriendshipRepository {
	return &FriendshipRepositoryImpl{baseRepo}
}

func (r *FriendshipRepositoryImpl) Save(ctx context.Context, friendship *friend.Friendship) error {
	entity := ToFriendshipEntity(friendship)

	err := gorm.G[Friendship](r.GetDBFromContext(ctx)).CreateInBatches(ctx, entity, len(*entity))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *FriendshipRepositoryImpl) Delete(ctx context.Context, friendship *friend.Friendship) error {
	row, err := gorm.G[Friendship](r.GetDBFromContext(ctx)).
		Where("friend_uuid = ?", friendship.FriendUUID()).Delete(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	if row == 0 {
		return commons.ErrNotFound
	}

	return nil
}
