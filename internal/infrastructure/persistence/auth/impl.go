package auth

import (
	"context"
	"palbum/internal/domain/auth"
	"palbum/internal/domain/commons"
	"palbum/internal/infrastructure/persistence"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserAuthRepositoryImpl struct {
	*persistence.Database
}

func NewUserAuthRepositoryImpl(baseRepo *persistence.Database) auth.AuthRepository {
	return &UserAuthRepositoryImpl{baseRepo}
}

func (r *UserAuthRepositoryImpl) Save(ctx context.Context, userAuth *auth.Auth) error {
	db := r.GetDBFromContext(ctx)

	err := gorm.G[AuthEntity](db, clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_uuid"}},
		DoUpdates: clause.AssignmentColumns([]string{"jti"}),
	}).Create(ctx, ToEntity(userAuth))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *UserAuthRepositoryImpl) Delete(ctx context.Context, userAuth *auth.Auth) error {
	db := r.GetDBFromContext(ctx)

	count, err := gorm.G[AuthEntity](db).Where("user_uuid = ? AND jti = ?", userAuth.UserUUID, userAuth.Jti).Delete(ctx)
	if count == 0 {
		return commons.ErrNotFound
	}

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *UserAuthRepositoryImpl) IsValidToken(ctx context.Context, userAuth *auth.Auth) (bool, error) {
	db := r.GetDBFromContext(ctx)

	_, err := gorm.G[AuthEntity](db).Where("user_uuid = ? AND jti = ?", userAuth.UserUUID, userAuth.Jti).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, errors.WithStack(err)
	}

	return true, nil
}
