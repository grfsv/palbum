package friend

import (
	"database/sql/driver"
	"palbum/internal/domain/commons"
	"palbum/internal/domain/user"
	"time"

	"github.com/google/uuid"
)

type RequestUUID uuid.UUID
type Status string

const (
	Pending  Status = "pending"
	Accepted Status = "accepted"
	Rejected Status = "rejected"
)

type FriendRequest struct {
	requestUUID RequestUUID
	fromUserID  user.UUID
	toUserID    user.UUID
	status      Status
	requestAt   time.Time
}

func NewFriendRequest(fromUUID user.UUID, toUUID user.UUID) *FriendRequest {
	requestUUID := RequestUUID(uuid.New())

	return &FriendRequest{
		requestUUID: requestUUID,
		fromUserID:  fromUUID,
		toUserID:    toUUID,
		status:      Pending,
		requestAt:   time.Now(),
	}
}

func ReconstructFriendRequest(
	requestUUID uuid.UUID,
	fromUUID uuid.UUID,
	toUUID uuid.UUID,
	status string,
	requestAt time.Time,
) *FriendRequest {
	return &FriendRequest{
		requestUUID: RequestUUID(requestUUID),
		fromUserID:  user.UUID(fromUUID),
		toUserID:    user.UUID(toUUID),
		status:      Status(status),
		requestAt:   requestAt,
	}
}

func ParseStatus(status string) (Status, error) {
	switch Status(status) {
	case Pending, Accepted, Rejected:
		return Status(status), nil
	default:
		return "", commons.NewBadRequestError("invalid status value")
	}
}

func (fr *FriendRequest) Accept() (*Friendship, error) {
	if fr.status != Pending {
		return nil, commons.NewConflictError("this request is not pending")
	}

	fr.status = Accepted

	return NewFriendship(fr), nil
}

func (fr *FriendRequest) Reject() error {
	if fr.status != Pending {
		return commons.NewConflictError("this request is not pending")
	}

	fr.status = Rejected

	return nil
}

func NewUUIDFromString(id string) (RequestUUID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return RequestUUID(uuid.Nil), err
	}

	return RequestUUID(uid), nil
}

func (fr *FriendRequest) RequestUUID() RequestUUID {
	return fr.requestUUID
}

func (fr *FriendRequest) FromUserID() user.UUID {
	return fr.fromUserID
}

func (fr *FriendRequest) ToUserID() user.UUID {
	return fr.toUserID
}

func (fr *FriendRequest) Status() Status {
	return fr.status
}

func (fr *FriendRequest) RequestAt() time.Time {
	return fr.requestAt
}

func (u RequestUUID) Value() (driver.Value, error) {
	return uuid.UUID(u).String(), nil
}
