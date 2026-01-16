package friend

import (
	"palbum/internal/domain/commons"
	"palbum/internal/domain/user"

	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

const (
	CodeLength = 21
)

type Code string

type FriendCode struct {
	userUUID user.UUID
	code     Code
}

func NewFriendCode(userUUID user.UUID) (*FriendCode, error) {
	uuid, err := gonanoid.New(CodeLength)
	if err != nil {
		return nil, commons.NewInternalError("failed to generate friend code", err)
	}

	return &FriendCode{
		userUUID: userUUID,
		code:     Code(uuid),
	}, nil
}

func ReconstructFriendCode(userUUID uuid.UUID, code string) *FriendCode {
	return &FriendCode{
		userUUID: user.UUID(userUUID),
		code:     Code(code),
	}
}

func NewCodeFromString(code string) (Code, error) {
	if len(code) != CodeLength {
		return "", commons.NewValidationError("invalid code length")
	}

	return Code(code), nil
}

func (c *FriendCode) UserUUID() user.UUID {
	return c.userUUID
}

func (c *FriendCode) Code() string {
	return string(c.code)
}
