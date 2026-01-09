package usecase

import (
	"context"
	"remind_map/internal/application/service"
	"remind_map/internal/domain/auth"
	"remind_map/internal/domain/commons"
	"remind_map/internal/domain/user"

	"github.com/cockroachdb/errors"
)

type UserLoginRequest struct {
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserLoginUsecase struct {
	userRepo     user.UserRepository
	authRepo     auth.AuthRepository
	tokenService service.TokenService
	txRepo       commons.TxManagerRepository
	passHasher   user.PasswordHasher
}

func NewUserLoginUsecase(
	userRepo user.UserRepository,
	authRepo auth.AuthRepository,
	txRepo commons.TxManagerRepository,
	tokenService service.TokenService,
	passHasher user.PasswordHasher,
) *UserLoginUsecase {
	return &UserLoginUsecase{
		userRepo:     userRepo,
		authRepo:     authRepo,
		txRepo:       txRepo,
		tokenService: tokenService,
		passHasher:   passHasher,
	}
}

func (u *UserLoginUsecase) Execute(ctx context.Context, input UserLoginRequest) (UserLoginResponse, error) {
	var res UserLoginResponse

	user, err := u.userRepo.FindByMail(ctx, input.Mail)
	if err != nil {
		if errors.Is(err, commons.ErrNotFound) {
			return res, commons.NewUnAuthorizedError()
		}

		return res, errors.WithStack(err)
	}

	err = u.passHasher.Compare(user.Password, input.Password)
	if err != nil {
		return res, commons.NewUnAuthorizedError()
	}

	auth := auth.NewAuth(user.UUID)

	u.txRepo.WithInTx(ctx, func(ctx context.Context) error {

		err = u.authRepo.Save(ctx, auth)
		if err != nil {
			return errors.WithStack(err)
		}

		refreshToken, err := u.tokenService.GenerateRefreshToken(auth)
		if err != nil {
			return err
		}

		accessToken, err := u.tokenService.GenerateAccessToken(auth)
		if err != nil {
			return errors.WithStack(err)
		}

		res = UserLoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
		return nil

	})

	return res, nil
}
