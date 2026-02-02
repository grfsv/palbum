package friend

import (
	"context"
	"palbum/internal/domain/user"
)

type FriendRequestRepository interface {
	Create(ctx context.Context, request *FriendRequest) error
	ChangeStatus(ctx context.Context, request *FriendRequest) error
	FindByUUID(ctx context.Context, requestUUID RequestUUID) (*FriendRequest, error)
	FindByUserUUID(ctx context.Context, userUUID user.UUID, category RequestCategory) ([]*FriendRequest, error)
}
