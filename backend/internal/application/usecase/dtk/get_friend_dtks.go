package dtk

import (
	"context"
	"encoding/hex"
	"time"

	"palbum/internal/domain/dtk"
	"palbum/internal/domain/friend"
	"palbum/internal/domain/user"

	"github.com/google/uuid"
)

type GetFriendDTKsInput struct {
	UserUUID user.UUID
	Since    *time.Time
}

type FriendDTKItem struct {
	UserUUID string `json:"userUuid"`
	Key      string `json:"key"`
	Date     string `json:"date"`
}

type GetFriendDTKsOutput struct {
	DTKs []FriendDTKItem `json:"dtks"`
}

type GetFriendDTKsUsecase struct {
	dtkRepo        dtk.DTKRepository
	friendshipRepo friend.FriendshipRepository
}

func NewGetFriendDTKsUsecase(
	dtkRepo dtk.DTKRepository,
	friendshipRepo friend.FriendshipRepository,
) *GetFriendDTKsUsecase {
	return &GetFriendDTKsUsecase{
		dtkRepo:        dtkRepo,
		friendshipRepo: friendshipRepo,
	}
}

func (u *GetFriendDTKsUsecase) Execute(ctx context.Context, input GetFriendDTKsInput) (GetFriendDTKsOutput, error) {
	var out GetFriendDTKsOutput

	out.DTKs = []FriendDTKItem{}

	friendships, err := u.friendshipRepo.FindByUserUUID(ctx, input.UserUUID)
	if err != nil {
		return out, err
	}

	if len(friendships) == 0 {
		return out, nil
	}

	friendUUIDs := make([]user.UUID, len(friendships))
	for i, f := range friendships {
		friendUUIDs[i] = f.FriendUUID()
	}

	since := dtk.Today()
	if input.Since != nil {
		since = *input.Since
	}

	dtks, err := u.dtkRepo.FindByUserUUIDsAndSince(ctx, friendUUIDs, since)
	if err != nil {
		return out, err
	}

	for _, d := range dtks {
		out.DTKs = append(out.DTKs, FriendDTKItem{
			UserUUID: uuid.UUID(d.UserUUID()).String(),
			Key:      hex.EncodeToString(d.Key().Bytes()),
			Date:     d.Date().Format("2006-01-02"),
		})
	}

	return out, nil
}
