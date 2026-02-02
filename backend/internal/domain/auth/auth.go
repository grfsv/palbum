package auth

import (
	"palbum/internal/domain/user"

	"github.com/google/uuid"
)

type Jti uuid.UUID

type Auth struct {
	userUUID user.UUID
	jti      Jti
}

func NewAuth(userUUID user.UUID) *Auth {
	jti := Jti(uuid.New())

	return &Auth{
		userUUID: userUUID,
		jti:      jti,
	}
}

func NewAuthWithJti(userUUID user.UUID, jti Jti) *Auth {
	return &Auth{
		userUUID: userUUID,
		jti:      jti,
	}
}

func NewJtiFromString(jtiStr string) (Jti, error) {
	parsedJti, err := uuid.Parse(jtiStr)
	if err != nil {
		return Jti(uuid.Nil), err
	}

	return Jti(parsedJti), nil
}

func (a *Auth) Refresh() {
	jti := Jti(uuid.New())
	a.jti = jti
}

func (a *Auth) UserUUID() user.UUID {
	return a.userUUID
}

func (a *Auth) Jti() Jti {
	return a.jti
}
