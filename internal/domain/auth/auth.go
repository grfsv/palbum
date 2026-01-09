package auth

import (
	"github.com/google/uuid"
)

type Auth struct {
	UserUUID uuid.UUID `gorm:"type:char(36);primaryKey; foreignKey"`
	Jti      string    `gorm:"type:char(36);unique"`
}

func NewAuth(userUUID uuid.UUID) *Auth {
	jti := uuid.New()

	return &Auth{
		UserUUID: userUUID,
		Jti:      jti.String(),
	}
}

func (a *Auth) Refresh() {
	jti := uuid.New()
	a.Jti = jti.String()
}
