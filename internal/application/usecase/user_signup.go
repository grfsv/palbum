package usecase

import (
	"context"
	"remind_map/internal/application/service"
	"remind_map/internal/domain/auth"
	"remind_map/internal/domain/commons"
	"remind_map/internal/domain/user"
)

type SignUpUsecase struct {
	txRepo       commons.TxManagerRepository
	userRepo     user.UserRepository
	authRepo     auth.AuthRepository
	passHasher   user.PasswordHasher
	tokenService service.TokenService
}

func NewSignUpUsecase(
	txRepo commons.TxManagerRepository,
	userRepo user.UserRepository,
	authRepo auth.AuthRepository,
	passHasher user.PasswordHasher,
	tokenService service.TokenService,
) *SignUpUsecase {
	return &SignUpUsecase{
		txRepo:       txRepo,
		userRepo:     userRepo,
		authRepo:     authRepo,
		passHasher:   passHasher,
		tokenService: tokenService,
	}
}

type SignUpRequest struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type SignUpResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Name         string `json:"name"`
	Mail         string `json:"mail"`
}

func (u *SignUpUsecase) Execute(ctx context.Context, input SignUpRequest) (SignUpResponse, error) {
	var res SignUpResponse

	err := u.txRepo.WithInTx(ctx, func(ctx context.Context) error {
		exist, err := u.userRepo.ExistByMail(ctx, input.Mail)
		if err != nil {
			return err
		}
		if exist {
			return commons.NewConflictError("email already in use")
		}

		newUser, err := user.NewUser(input.Name, input.Mail, input.Password, u.passHasher)
		if err != nil {
			return err
		}
		err = u.userRepo.Create(ctx, newUser)
		if err != nil {
			return err
		}

		newAuth, err := auth.NewAuth(newUser.UUID)
		if err != nil {
			return err
		}
		err = u.authRepo.Save(ctx, newAuth)
		if err != nil {
			return err
		}

		refreshToken, err := u.tokenService.GenerateRefreshToken(newAuth)
		if err != nil {
			return err
		}
		accessToken, err := u.tokenService.GenerateAccessToken(newAuth)
		if err != nil {
			return err
		}

		res = SignUpResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			Name:         newUser.Name,
			Mail:         newUser.Mail,
		}

		return nil
	})

	return res, err
}
