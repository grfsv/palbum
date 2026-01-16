package friend

import (
	"context"
	"palbum/internal/domain/commons"
	"palbum/internal/domain/friend"
	"palbum/internal/domain/user"
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

func (r *FriendshipRepositoryImpl) FindByUserUUID(
	ctx context.Context,
	userUUID user.UUID,
) ([]*friend.Friendship, error) {
	rows, err := gorm.G[Friendship](r.GetDBFromContext(ctx)).Where("user_uuid1 = ?", userUUID).Find(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	friendships := make([]*friend.Friendship, len(rows))
	for i, row := range rows {
		friendships[i] = row.ToDomain()
	}

	return friendships, nil
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
