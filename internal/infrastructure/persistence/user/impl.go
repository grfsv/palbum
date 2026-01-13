package user

import (
	"context"
	"palbum/internal/domain/commons"
	"palbum/internal/domain/user"
	"palbum/internal/infrastructure/persistence"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	*persistence.Database
}

func NewUserRepositoryImpl(baseRepo *persistence.Database) user.UserRepository {
	return &UserRepositoryImpl{baseRepo}
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, uuid user.UUID) (*user.User, error) {
	db := r.GetDBFromContext(ctx)

	record, err := gorm.G[UserEntity](db).Where("uuid = ?", uuid).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, commons.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return record.ToDomain()
}

func (r *UserRepositoryImpl) FindByMail(ctx context.Context, mail user.Mail) (*user.User, error) {
	db := r.GetDBFromContext(ctx)

	record, err := gorm.G[UserEntity](db).Where("mail = ?", mail).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, commons.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return record.ToDomain()
}

func (r *UserRepositoryImpl) ExistByMail(ctx context.Context, mail user.Mail) (bool, error) {
	db := r.GetDBFromContext(ctx)

	_, err := gorm.G[UserEntity](db).Where("mail = ?", mail).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, errors.WithStack(err)
	}

	return true, nil
}

func (r *UserRepositoryImpl) Create(ctx context.Context, u *user.User) error {
	db := r.GetDBFromContext(ctx)

	err := gorm.G[UserEntity](db).Create(ctx, ToEntity(u))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
