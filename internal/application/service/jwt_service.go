package service

import (
	"remind_map/internal/domain/auth"
)

type TokenService interface {
	GenerateAccessToken(auth *auth.Auth) (string, error)
	GenerateRefreshToken(auth *auth.Auth) (string, error)
	ConfirmAccessToken(tokenString string) (string, error)
	ConfirmRefreshToken(tokenString string) (*auth.Auth, error)
}
