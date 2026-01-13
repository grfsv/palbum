package user

import (
	"context"
	"palbum/internal/application/service"
	"palbum/internal/domain/auth"
	"palbum/internal/domain/commons"
	"palbum/internal/domain/user"

	"github.com/cockroachdb/errors"
)

type UserLoginRequest struct {
	Mail     string `binding:"required" json:"mail"`
	Password string `binding:"required" json:"password"`
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

	mail, err := user.NewMail(input.Mail)
	if err != nil {
		return res, err
	}

	user, err := u.userRepo.FindByMail(ctx, mail)
	if err != nil {
		if errors.Is(err, commons.ErrNotFound) {
			return res, commons.NewUnAuthorizedError()
		}

		return res, err
	}

	err = u.passHasher.Compare(user.Password(), input.Password)
	if err != nil {
		return res, commons.NewUnAuthorizedError()
	}

	auth := auth.NewAuth(user.UUID())

	err = u.txRepo.WithInTx(ctx, func(ctx context.Context) error {
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

		res = UserLoginResponse{
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
