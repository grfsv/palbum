package user

import (
	"context"
	"palbum/internal/application/service"
	"palbum/internal/domain/auth"
	"palbum/internal/domain/commons"
)

type RefreshRequest struct {
	RefreshToken string `binding:"required" json:"refreshToken"`
}
type RefreshResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshUsecase struct {
	tokenService service.TokenService
	authRepo     auth.AuthRepository
	txRepo       commons.TxManagerRepository
}

func NewRefreshUsecase(
	tokenService service.TokenService,
	authRepo auth.AuthRepository,
	txRepo commons.TxManagerRepository,
) *RefreshUsecase {
	return &RefreshUsecase{
		tokenService: tokenService,
		authRepo:     authRepo,
		txRepo:       txRepo,
	}
}

func (u *RefreshUsecase) Execute(ctx context.Context, input RefreshRequest) (RefreshResponse, error) {
	var res RefreshResponse

	auth, err := u.tokenService.ConfirmRefreshToken(input.RefreshToken)
	if err != nil {
		return res, err
	}

	isValid, err := u.authRepo.IsValidToken(ctx, auth)
	if err != nil {
		return res, err
	}

	if !isValid {
		return res, commons.NewUnAuthorizedError()
	}

	err = u.txRepo.WithInTx(ctx, func(ctx context.Context) error {
		auth.Refresh()

		err = u.authRepo.Save(ctx, auth)
		if err != nil {
			return err
		}

		refreshToken, err := u.tokenService.GenerateRefreshToken(auth)
		if err != nil {
			return err
		}

		accessToken, err := u.tokenService.GenerateAccessToken(auth)
		if err != nil {
			return err
		}

		res = RefreshResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		return nil
	})
	if err != nil {
		return res, err
	}

	return res, nil
}
