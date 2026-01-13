package auth

import (
	"palbum/internal/domain/auth"
	"palbum/internal/domain/user"

	"github.com/google/uuid"
)

type AuthEntity struct {
	UUID uuid.UUID `gorm:"type:char(36);primaryKey"`
	Jti  uuid.UUID `gorm:"type:char(36);unique"`
}

func ToEntity(auth *auth.Auth) *AuthEntity {
	return &AuthEntity{
		UUID: uuid.UUID(auth.UserUUID()),
		Jti:  uuid.UUID(auth.Jti()),
	}
}

func (e AuthEntity) ToDomain() (*auth.Auth, error) {
	userUUID, err := user.NewUUIDFromString(e.UUID.String())
	if err != nil {
		return nil, err
	}

	jti, err := auth.NewJtiFromString(e.Jti.String())
	if err != nil {
		return nil, err
	}

	return auth.NewAuthWithJti(userUUID, jti), nil
}
