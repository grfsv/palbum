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
func (r *FriendRequestRepositoryImpl) ChangeStatus(ctx context.Context, request *friend.FriendRequest) error {
	rowsAffected, err := gorm.G[FriendRequest](r.GetDBFromContext(ctx)).
		Where("request_uuid = ?", request.RequestUUID()).
		Update(ctx, "status", request.Status())
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, commons.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return row.ToDomain(), nil
}

func (r *FriendRequestRepositoryImpl) FindByUserUUID(
	ctx context.Context,
	userUUID user.UUID,
	category friend.RequestCategory,
) ([]*friend.FriendRequest, error) {
	var filter string

	switch category {
	case friend.CategorySent:
		filter = "sender_uuid = ?"
	case friend.CategoryReceived:
		filter = "receiver_uuid = ?"
	}

	rows, err := gorm.G[FriendRequest](r.GetDBFromContext(ctx)).
		Where(filter, userUUID).Where("status = ?", friend.Pending).Find(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(rows) == 0 {
		return nil, commons.ErrNotFound
	}

	requests := make([]*friend.FriendRequest, len(rows))
	for i, row := range rows {
		requests[i] = row.ToDomain()
	}

	return requests, nil
}
