package friend

import (
	"context"
	"fmt"
	"palbum/internal/domain/friend"
	"palbum/internal/domain/user"

	"github.com/google/uuid"
)

type FriendListInput struct {
	UserUUID user.UUID
}

type FriendListOutput struct {
	Friends []Friend `json:"friends"`
}

type Friend struct {
	UserUUID string `json:"userUuid"`
	Username string `json:"username"`
}

type FriendListUsecase struct {
	friendshipRepo friend.FriendshipRepository
	userRepo       user.UserRepository
}

func NewFriendListUsecase(
	friendshipRepo friend.FriendshipRepository,
	userRepo user.UserRepository,
) *FriendListUsecase {
	return &FriendListUsecase{
		friendshipRepo: friendshipRepo,
		userRepo:       userRepo,
	}
}

func (u *FriendListUsecase) Execute(ctx context.Context, input FriendListInput) (FriendListOutput, error) {
	var out FriendListOutput

	friendships, err := u.friendshipRepo.FindByUserUUID(ctx, input.UserUUID)
	if err != nil {
		return out, err
	}

	userUUIDs := make([]user.UUID, 0, len(friendships))
	for _, friendship := range friendships {
		for userUUID := range friendship.Friend() {
			if userUUID == input.UserUUID {
				continue
			}

			userUUIDs = append(userUUIDs, userUUID)
		}
	}

	fmt.Println(userUUIDs)

	users, err := u.userRepo.ListInUUID(ctx, userUUIDs)
	if err != nil {
		return out, err
	}

	out.Friends = make([]Friend, len(users))
	for i, user := range users {
		out.Friends[i] = Friend{
			UserUUID: uuid.UUID(user.UUID()).String(),
			Username: string(user.Name()),
		}
	}

	return out, nil
}
