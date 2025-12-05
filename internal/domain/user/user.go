package user

import (
	"regexp"
	"remind_map/internal/domain/commons"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type User struct {
	UUID     uuid.UUID `gorm:"type:char(36);primaryKey"`
	Name     string    `gorm:"type:varchar(255);not null"`
	Mail     string    `gorm:"unique;not null;uniqueIndex"`
	Password string    `gorm:"type:varchar(255);not null"`
}

const (
	NameMinLength     = 2
	NameMaxLength     = 30
	PasswordMinLength = 8
	PasswordMaxLength = 256
)

func NewUser(name string, mail string, password string, passHasher PasswordHasher) (*User, error) {
	var errorList []error

	trimmedName := strings.TrimSpace(name)
	trimmedMail := strings.TrimSpace(mail)
	trimmedPassword := strings.TrimSpace(password)

	customErr := validateName(trimmedName)
	if customErr != nil {
		errorList = append(errorList, customErr)
	}

	customErr = validateMail(trimmedMail)
	if customErr != nil {
		errorList = append(errorList, customErr)
	}

	customErr = validatePassword(trimmedPassword)
	if customErr != nil {
		errorList = append(errorList, customErr)
	}

	switch len(errorList) {
	case 0:
		hashed, err := passHasher.Hash(trimmedPassword)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		uuid, err := uuid.NewV7()
		if err != nil {
			return nil, errors.WithStack(err)
		}

		return &User{
			UUID:     uuid,
			Name:     trimmedName,
			Mail:     trimmedMail,
			Password: hashed,
		}, nil
	case 1:
		return nil, errorList[0]
	default:
		err := commons.NewMultipleValidationError(errorList)

		return nil, err
	}
}

func validateName(name string) error {
	if len(name) < NameMinLength {
		return commons.NewValidationError("too short name")
	}

	if len(name) > NameMaxLength {
		return commons.NewValidationError("too long name")
	}

	return nil
}

func validateMail(mail string) error {
	ok := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(mail)
	if !ok {
		return commons.NewValidationError("invalid format email")
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < PasswordMinLength {
		return commons.NewValidationError("too short password")
	}

	if len(password) >= PasswordMaxLength {
		return commons.NewValidationError("too long password")
	}

	return nil
}
