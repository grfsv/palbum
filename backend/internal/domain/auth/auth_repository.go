package auth

import (
	"context"
)

type AuthRepository interface {
	Save(ctx context.Context, userAuth *Auth) error
	Delete(ctx context.Context, auth *Auth) error
	IsValidToken(ctx context.Context, userAuth *Auth) (bool, error)
}
