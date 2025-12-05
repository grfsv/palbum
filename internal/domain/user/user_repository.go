package user

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindByID(ctx context.Context, uuid uuid.UUID) (*User, error)
	FindByMail(ctx context.Context, mail string) (*User, error)
	ExistByMail(ctx context.Context, mail string) (bool, error)
	Create(ctx context.Context, user *User) error
}
