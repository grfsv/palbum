package friend

import "context"

type FriendshipRepository interface {
	Save(ctx context.Context, friendship *Friendship) error
	Delete(ctx context.Context, friendship *Friendship) error
}
