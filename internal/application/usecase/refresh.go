package usecase

import (
	"context"
	"remind_map/internal/application/service"
	"remind_map/internal/domain/auth"
	"remind_map/internal/domain/commons"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
type RefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
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

	u.authRepo.IsValidToken(ctx, auth)
	isValid, err := u.authRepo.IsValidToken(ctx, auth)
	if err != nil {
		return res, err
	}
	if !isValid {
		return res, commons.NewUnAuthorizedError()
	}

	u.txRepo.WithInTx(ctx, func(ctx context.Context) error {

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

	return res, nil
}
