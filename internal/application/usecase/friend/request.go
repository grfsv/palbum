package friend

import (
	"context"
	"errors"
	"palbum/internal/domain/commons"
	"palbum/internal/domain/friend"
	"palbum/internal/domain/user"
)

type FriendRequestRequest struct {
	UserUUID user.UUID
	Code     friend.Code
}

type FriendRequestUsecase struct {
	codeRepo          friend.FriendCodeRepository
	friendRequestRepo friend.FriendRequestRepository
}

func NewFriendRequestUsecase(
	codeRepo friend.FriendCodeRepository,
	friendRequestRepo friend.FriendRequestRepository,
) *FriendRequestUsecase {
	return &FriendRequestUsecase{
		codeRepo:          codeRepo,
		friendRequestRepo: friendRequestRepo,
	}
}

func (u *FriendRequestUsecase) Execute(ctx context.Context, input FriendRequestRequest) error {
	friendCode, err := u.codeRepo.FindByCode(ctx, input.Code)
	if err != nil {
		if errors.Is(err, commons.ErrNotFound) {
			return commons.NewNotFoundError("user not found")
		}

		return err
	}

	if friendCode.UserUUID() == input.UserUUID {
		return commons.NewBadRequestError("cannot request yourself")
	}

	newRequest := friend.NewFriendRequest(input.UserUUID, friendCode.UserUUID())

	err = u.friendRequestRepo.Create(ctx, newRequest)
	if err != nil {
		if errors.Is(err, commons.ErrDuplicate) {
			return commons.NewConflictError("request already sent")
		}

		return err
	}

	return nil
}
