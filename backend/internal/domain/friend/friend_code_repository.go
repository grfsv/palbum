package friend

import (
	"context"
	"palbum/internal/domain/user"
)

type FriendCodeRepository interface {
	Save(ctx context.Context, code *FriendCode) error
	FindByUUID(ctx context.Context, userUUID user.UUID) (*FriendCode, error)
	FindByCode(ctx context.Context, code Code) (*FriendCode, error)
}
