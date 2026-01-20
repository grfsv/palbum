package friend

import (
	"database/sql/driver"
	"palbum/internal/domain/user"
	"time"

	"github.com/google/uuid"
)

type FriendshipUUID uuid.UUID

type Friendship struct {
	friendshipUUID FriendshipUUID
	userUUID       user.UUID
	friendUUID     user.UUID
	acceptedAt     time.Time
}

func NewFriendship(request *FriendRequest) *Friendship {
	friendshipUUID := FriendshipUUID(uuid.New())

	return &Friendship{
		friendshipUUID: friendshipUUID,
		userUUID:       request.ToUserID(),
		friendUUID:     request.FromUserID(),
		acceptedAt:     time.Now(),
	}
}

func ReconstructFriendship(
	friendshipUUID uuid.UUID,
	userUUID uuid.UUID,
	friendUUID uuid.UUID,
	acceptedAt time.Time,
) *Friendship {
	return &Friendship{
		friendshipUUID: FriendshipUUID(friendshipUUID),
		userUUID:       user.UUID(userUUID),
		friendUUID:     user.UUID(friendUUID),
		acceptedAt:     acceptedAt,
	}
}

func (f *Friendship) FriendshipUUID() FriendshipUUID {
	return f.friendshipUUID
}

func (f *Friendship) FriendUUID() user.UUID {
	return f.friendUUID
}

func (f *Friendship) UserUUID() user.UUID {
	return f.userUUID
}

func (f *Friendship) AcceptedAt() time.Time {
	return f.acceptedAt
}

func (u FriendshipUUID) Value() (driver.Value, error) {
	return uuid.UUID(u).String(), nil
}
