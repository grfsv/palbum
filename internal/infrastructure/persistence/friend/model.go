package friend

import (
	"palbum/internal/domain/friend"
	"time"

	"github.com/google/uuid"
)

type FriendCode struct {
	UserUUID uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	Code     string    `gorm:"type:char(36);not null;uniqueIndex"`
}

type FriendRequest struct {
	RequestUUID  uuid.UUID `gorm:"type:char(36);primaryKey;index"`
	SenderUUID   uuid.UUID `gorm:"type:char(36);not null;uniqueIndex:idx_friend"`
	ReceiverUUID uuid.UUID `gorm:"type:char(36);not null;uniqueIndex:idx_friend"`
	Status       string    `gorm:"type:varchar(20);not null"`
	RequestAt    time.Time `gorm:"not null"`
}

type Friendship struct {
	FriendshipUUID uuid.UUID `gorm:"type:char(36);"`
	UserUUID       uuid.UUID `gorm:"type:char(36);not null;primaryKey;index"`
	FriendUUID     uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	AcceptedAt     time.Time `gorm:"not null"`
}

func (fe *FriendCode) ToDomain() *friend.FriendCode {
	return friend.ReconstructFriendCode(fe.UserUUID, fe.Code)
}

func (fre *FriendRequest) ToDomain() *friend.FriendRequest {
	return friend.ReconstructFriendRequest(
		fre.RequestUUID,
		fre.SenderUUID,
		fre.ReceiverUUID,
		fre.Status,
		fre.RequestAt,
	)
}

func (fe *Friendship) ToDomain() *friend.Friendship {
	return friend.ReconstructFriendship(
		fe.FriendshipUUID,
		fe.UserUUID,
		fe.FriendshipUUID,
		fe.AcceptedAt,
	)
}

func ToFriendCodeEntity(code *friend.FriendCode) *FriendCode {
	return &FriendCode{
		UserUUID: uuid.UUID(code.UserUUID()),
		Code:     string(code.Code()),
	}
}

func ToFriendRequestEntity(request *friend.FriendRequest) *FriendRequest {
	return &FriendRequest{
		RequestUUID:  uuid.UUID(request.RequestUUID()),
		SenderUUID:   uuid.UUID(request.FromUserID()),
		ReceiverUUID: uuid.UUID(request.ToUserID()),
		Status:       string(request.Status()),
		RequestAt:    request.RequestAt(),
	}
}

func ToFriendshipEntity(friendship *friend.Friendship) *[]Friendship {
	var entities []Friendship

	entities = append(entities, Friendship{
		FriendshipUUID: uuid.UUID(friendship.FriendshipUUID()),
		UserUUID:       uuid.UUID(friendship.UserUUID()),
		FriendUUID:     uuid.UUID(friendship.FriendUUID()),
		AcceptedAt:     friendship.AcceptedAt(),
	})

	entities = append(entities, Friendship{
		FriendshipUUID: uuid.UUID(friendship.FriendshipUUID()),
		UserUUID:       uuid.UUID(friendship.FriendUUID()),
		FriendUUID:     uuid.UUID(friendship.UserUUID()),
		AcceptedAt:     friendship.AcceptedAt(),
	})

	return &entities
}
