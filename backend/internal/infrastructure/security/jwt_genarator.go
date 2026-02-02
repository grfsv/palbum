package security

import (
	"time"

	"palbum/internal/application/service"

	"palbum/internal/domain/auth"
	"palbum/internal/domain/commons"
	"palbum/internal/domain/user"

	"github.com/cockroachdb/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTGenerator struct {
	accessTokenSecret  []byte
	refreshTokenSecret []byte
}

type JWTConfig struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
}

func NewJWTService(cfg *JWTConfig) service.TokenService {
	return &JWTGenerator{
		accessTokenSecret:  []byte(cfg.AccessTokenSecret),
		refreshTokenSecret: []byte(cfg.RefreshTokenSecret),
	}
}

func (s *JWTGenerator) GenerateAccessToken(auth *auth.Auth) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   uuid.UUID(auth.UserUUID()).String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	})

	signedToken, err := token.SignedString(s.accessTokenSecret)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return signedToken, nil
}

func (s *JWTGenerator) GenerateRefreshToken(auth *auth.Auth) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   uuid.UUID(auth.UserUUID()).String(),
		ID:        uuid.UUID(auth.Jti()).String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
	})

	signedToken, err := claims.SignedString(s.refreshTokenSecret)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return signedToken, nil
}

func (s *JWTGenerator) ConfirmAccessToken(tokenString string) (string, error) {
	parser := jwt.NewParser(
		jwt.WithExpirationRequired(),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)

	token, err := parser.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		return s.accessTokenSecret, nil
	})
	if err != nil {
		return "", commons.NewUnAuthorizedError(err)
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return "", commons.NewUnAuthorizedError()
	}

	return claims.Subject, nil
}

func (s *JWTGenerator) ConfirmRefreshToken(tokenString string) (*auth.Auth, error) {
	parser := jwt.NewParser(
		jwt.WithExpirationRequired(),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)

	token, err := parser.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		return s.refreshTokenSecret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, commons.NewUnAuthorizedError()
		}

		return nil, errors.WithStack(err)
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, commons.NewUnAuthorizedError()
	}

	userUUIDStr, err := uuid.Parse(claims.Subject)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userUUID := user.UUID(userUUIDStr)

	jti, err := auth.NewJtiFromString(claims.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	newAuth := auth.NewAuthWithJti(
		userUUID,
		jti,
	)

	return newAuth, nil
}
