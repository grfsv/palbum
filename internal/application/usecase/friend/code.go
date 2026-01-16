package friend

import (
	"context"
	"errors"
	"palbum/internal/domain/commons"
	"palbum/internal/domain/friend"
	"palbum/internal/domain/user"
)

type FriendCodeInput struct {
	UserUUID user.UUID
}

type FriendCodeOutput struct {
	Code string `json:"code"`
}

type FriendCodeUsecase struct {
	codeRepo friend.FriendCodeRepository
}

func NewFriendCodeUsecase(
	codeRepo friend.FriendCodeRepository,
) *FriendCodeUsecase {
	return &FriendCodeUsecase{
		codeRepo: codeRepo,
	}
}

func (u *FriendCodeUsecase) Execute(ctx context.Context, input FriendCodeInput) (FriendCodeOutput, error) {
	var out FriendCodeOutput

	friendCode, err := u.codeRepo.FindByUUID(ctx, input.UserUUID)
	if err != nil {
		if errors.Is(err, commons.ErrNotFound) {
			newFriendCode, err := u.createCode(ctx, input.UserUUID)
			if err != nil {
				return out, err
			}

			out.Code = newFriendCode.Code()

			return out, nil
		}

		return out, err
	}

	out.Code = friendCode.Code()

	return out, nil
}

func (u *FriendCodeUsecase) createCode(ctx context.Context, userUUID user.UUID) (*friend.FriendCode, error) {
	code, err := friend.NewFriendCode(userUUID)
	if err != nil {
		return nil, err
	}

	err = u.codeRepo.Save(ctx, code)
	if err != nil {
		return nil, err
	}

	return code, nil
}
