package persistence

import (
	"context"
	"remind_map/internal/domain/commons"
	"remind_map/internal/domain/user"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	*Database
}

func NewUserRepositoryImpl(baseRepo *Database) user.UserRepository {
	return &UserRepositoryImpl{baseRepo}
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, uuid uuid.UUID) (*user.User, error) {
	db := r.GetDBFromContext(ctx)

	record, err := gorm.G[user.User](db).Where("uuid = ?", uuid.String()).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, commons.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &record, nil
}

func (r *UserRepositoryImpl) FindByMail(ctx context.Context, mail string) (*user.User, error) {
	db := r.GetDBFromContext(ctx)

	record, err := gorm.G[user.User](db).Where("mail = ?", mail).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, commons.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &record, nil
}

func (r *UserRepositoryImpl) ExistByMail(ctx context.Context, mail string) (bool, error) {
	db := r.GetDBFromContext(ctx)

	_, err := gorm.G[user.User](db).Where("mail = ?", mail).First(ctx)
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

	err := gorm.G[user.User](db).Create(ctx, u)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
