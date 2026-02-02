package dtk

import (
	"context"
	"encoding/hex"
	"errors"

	"palbum/internal/domain/commons"
	"palbum/internal/domain/dtk"
	"palbum/internal/domain/user"
)

type GetMyDTKInput struct {
	UserUUID user.UUID
}

type GetMyDTKOutput struct {
	Key  string `json:"key"`
	Date string `json:"date"`
}

type GetMyDTKUsecase struct {
	dtkRepo dtk.DTKRepository
}

func NewGetMyDTKUsecase(dtkRepo dtk.DTKRepository) *GetMyDTKUsecase {
	return &GetMyDTKUsecase{
		dtkRepo: dtkRepo,
	}
}

func (u *GetMyDTKUsecase) Execute(ctx context.Context, input GetMyDTKInput) (GetMyDTKOutput, error) {
	var out GetMyDTKOutput

	today := dtk.Today()

	existingDTK, err := u.dtkRepo.FindByUserUUIDAndDate(ctx, input.UserUUID, today)
	if err != nil {
		if errors.Is(err, commons.ErrNotFound) {
			newDTK, err := u.createDTK(ctx, input.UserUUID)
			if err != nil {
				return out, err
			}

			out.Key = hex.EncodeToString(newDTK.Key().Bytes())
			out.Date = newDTK.Date().Format("2006-01-02")

			return out, nil
		}

		return out, err
	}

	out.Key = hex.EncodeToString(existingDTK.Key().Bytes())
	out.Date = existingDTK.Date().Format("2006-01-02")

	return out, nil
}

func (u *GetMyDTKUsecase) createDTK(ctx context.Context, userUUID user.UUID) (*dtk.DTK, error) {
	today := dtk.Today()

	newDTK, err := dtk.NewDTK(userUUID, today)
	if err != nil {
		return nil, err
	}

	err = u.dtkRepo.Save(ctx, newDTK)
	if err != nil {
		return nil, err
	}

	return newDTK, nil
}
