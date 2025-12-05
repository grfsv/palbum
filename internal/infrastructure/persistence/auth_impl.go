package persistence

import (
	"context"
	"remind_map/internal/domain/auth"
	"remind_map/internal/domain/commons"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserAuthRepositoryImpl struct {
	*Database
}

func NewUserAuthRepositoryImpl(baseRepo *Database) auth.AuthRepository {
	return &UserAuthRepositoryImpl{baseRepo}
}

func (r *UserAuthRepositoryImpl) Save(ctx context.Context, userAuth *auth.Auth) error {
	db := r.GetDBFromContext(ctx)

	err := gorm.G[auth.Auth](db, clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_uuid"}},
		DoUpdates: clause.AssignmentColumns([]string{"jti"}),
	}).Create(ctx, userAuth)
	if err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(err)
}

func (r *UserAuthRepositoryImpl) Delete(ctx context.Context, userAuth *auth.Auth) error {
	db := r.GetDBFromContext(ctx)

	count, err := gorm.G[auth.Auth](db).Where("user_uuid = ? AND jti = ?", userAuth.UserUUID, userAuth.Jti).Delete(ctx)
	if count == 0 {
		return commons.ErrNotFound
	}

	return errors.WithStack(err)
}

func (r *UserAuthRepositoryImpl) IsValidToken(ctx context.Context, userAuth *auth.Auth) (bool, error) {
	db := r.GetDBFromContext(ctx)

	_, err := gorm.G[auth.Auth](db).Where("user_uuid = ? AND jti = ?", userAuth.UserUUID, userAuth.Jti).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, errors.WithStack(err)
	}

	return true, nil
}
