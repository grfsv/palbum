package auth

import (
	"palbum/internal/domain/auth"
	"palbum/internal/domain/user"

	"github.com/google/uuid"
)

type Auth struct {
	UUID uuid.UUID `gorm:"type:char(36);primaryKey"`
	Jti  uuid.UUID `gorm:"type:char(36);unique"`
}

func ToEntity(auth *auth.Auth) *Auth {
	return &Auth{
		UUID: uuid.UUID(auth.UserUUID()),
		Jti:  uuid.UUID(auth.Jti()),
	}
}

func (e Auth) ToDomain() (*auth.Auth, error) {
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
