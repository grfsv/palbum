package friend

import (
	"context"
	"palbum/internal/domain/commons"
	"palbum/internal/domain/friend"
	"palbum/internal/infrastructure/persistence"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type FriendRequestRepositoryImpl struct {
	*persistence.Database
}

func NewFriendRequestRepositoryImpl(baseRepo *persistence.Database) friend.FriendRequestRepository {
	return &FriendRequestRepositoryImpl{baseRepo}
}

func (r *FriendRequestRepositoryImpl) Create(ctx context.Context, request *friend.FriendRequest) error {
	err := gorm.G[FriendRequest](r.GetDBFromContext(ctx)).Create(ctx, ToFriendRequestEntity(request))
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return commons.ErrDuplicate
		}

		return errors.WithStack(err)
	}

	return nil
}
func (r *FriendRequestRepositoryImpl) Update(ctx context.Context, request *friend.FriendRequest) error {
	rowsAffected, err := gorm.G[FriendRequest](r.GetDBFromContext(ctx)).
		Where("request_uuid = ?", request.RequestUUID()).Updates(ctx, *ToFriendRequestEntity(request))
	if err != nil {
		return errors.WithStack(err)
	}

	if rowsAffected == 0 {
		return commons.ErrNotFound
	}

	return nil
}

func (r *FriendRequestRepositoryImpl) FindByUUID(
	ctx context.Context,
	requestUUID friend.RequestUUID,
) (*friend.FriendRequest, error) {
	row, err := gorm.G[FriendRequest](r.GetDBFromContext(ctx)).Where("request_uuid = ?", requestUUID).First(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return row.ToDomain(), nil
}
