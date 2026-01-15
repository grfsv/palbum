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

type FriendCodeRepositoryImpl struct {
	*persistence.Database
}

func NewFriendCodeRepositoryImpl(baseRepo *persistence.Database) friend.FriendCodeRepository {
	return &FriendCodeRepositoryImpl{baseRepo}
}

func (r *FriendCodeRepositoryImpl) Save(ctx context.Context, code *friend.FriendCode) error {
	err := gorm.G[FriendCode](r.GetDBFromContext(ctx)).Create(ctx, ToFriendCodeEntity(code))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *FriendCodeRepositoryImpl) FindByUUID(ctx context.Context, userUUID user.UUID) (*friend.FriendCode, error) {
	row, err := gorm.G[FriendCode](r.GetDBFromContext(ctx)).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, commons.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return row.ToDomain(), nil
}

func (r *FriendCodeRepositoryImpl) FindByCode(ctx context.Context, code friend.Code) (*friend.FriendCode, error) {
	row, err := gorm.G[FriendCode](r.GetDBFromContext(ctx)).Where("code = ?", code).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, commons.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return row.ToDomain(), nil
}
