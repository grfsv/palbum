package friend

import (
	"database/sql/driver"
	"palbum/internal/domain/user"
	"time"

	"github.com/google/uuid"
)

type FriendUUID uuid.UUID
type Friend map[user.UUID]struct{}

type Friendship struct {
	friendUUID FriendUUID
	friend     Friend
	acceptedAt time.Time
}

func NewFriendship(request *FriendRequest) *Friendship {
	friendUUID := FriendUUID(uuid.New())

	return &Friendship{
		friendUUID: friendUUID,
		friend: Friend{
			request.fromUserID: {},
			request.toUserID:   {},
		},
		acceptedAt: time.Now(),
	}
}

func ReconstructFriendship(
	friendUUID uuid.UUID,
	userUUID1 uuid.UUID,
	userUUID2 uuid.UUID,
	acceptedAt time.Time,
) *Friendship {
	return &Friendship{
		friendUUID: FriendUUID(friendUUID),
		friend: Friend{
			user.UUID(userUUID1): {},
			user.UUID(userUUID2): {},
		},
		acceptedAt: acceptedAt,
	}
}

func (f *Friendship) Friend() Friend {
	return f.friend
}

func (f *Friendship) FriendUUID() FriendUUID {
	return f.friendUUID
}

func (f *Friendship) AcceptedAt() time.Time {
	return f.acceptedAt
}

func (u FriendUUID) Value() (driver.Value, error) {
	return uuid.UUID(u).String(), nil
}
