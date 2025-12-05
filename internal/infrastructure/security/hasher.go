package security

import (
	"remind_map/internal/domain/user"

	"github.com/cockroachdb/errors"
	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct {
	Cost int
}

func NewBcryptHasher() user.PasswordHasher {
	return &BcryptHasher{Cost: bcrypt.DefaultCost}
}

func (h *BcryptHasher) Hash(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), h.Cost)
	if err != nil {
		return "", errors.WithStack(err)
	}
	
	return string(hashedBytes), nil
}

func (h *BcryptHasher) Compare(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
