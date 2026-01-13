package user

import (
	"context"
)

type UserRepository interface {
	FindByID(ctx context.Context, uuid UUID) (*User, error)
	FindByMail(ctx context.Context, mail Mail) (*User, error)
	ExistByMail(ctx context.Context, mail Mail) (bool, error)
	Create(ctx context.Context, user *User) error
}
