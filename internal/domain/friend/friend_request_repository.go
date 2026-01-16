package friend

import "context"

type FriendRequestRepository interface {
	Create(ctx context.Context, request *FriendRequest) error
	Update(ctx context.Context, request *FriendRequest) error
	FindByUUID(ctx context.Context, requestUUID RequestUUID) (*FriendRequest, error)
}
