package usecase

import (
	"context"
	"remind_map/internal/application/service"
	"remind_map/internal/domain/auth"
	"remind_map/internal/domain/commons"
	"remind_map/internal/domain/user"

	"github.com/cockroachdb/errors"
)

type UserLogoutRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type UserLogoutUsecase struct {
	userRepo     user.UserRepository
	authRepo     auth.AuthRepository
	tokenService service.TokenService
}

func NewUserLogoutUsecase(
	userRepo user.UserRepository,
	authRepo auth.AuthRepository,
	tokenService service.TokenService,
) *UserLogoutUsecase {
	return &UserLogoutUsecase{
		userRepo:     userRepo,
		authRepo:     authRepo,
		tokenService: tokenService,
	}
}

func (u *UserLogoutUsecase) Execute(ctx context.Context, input UserLogoutRequest) error {
	auth, err := u.tokenService.ConfirmRefreshToken(input.RefreshToken)
	if err != nil {
		return err
	}

	// トークン削除
	err = u.authRepo.Delete(ctx, auth)
	if err != nil {
		if errors.Is(err, commons.ErrNotFound) {
			return commons.NewUnAuthorizedError()
		} else {
			return err
		}
	}

	return nil
}
