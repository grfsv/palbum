package user

import (
	"database/sql/driver"
	"palbum/internal/domain/commons"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type UUID uuid.UUID
type Name string
type Mail string
type Password string

type User struct {
	uuid     UUID
	name     Name
	mail     Mail
	password Password
}

const (
	NameMinLength     = 2
	NameMaxLength     = 30
	PasswordMinLength = 8
	PasswordMaxLength = 256
)

func NewUser(name Name, mail Mail, password Password) *User {
	uuid := UUID(uuid.New())

	return &User{
		uuid:     uuid,
		name:     name,
		mail:     mail,
		password: password,
	}
}

func NewUserWithUUID(uuid UUID, name Name, mail Mail, password Password) *User {
	return &User{
		uuid:     uuid,
		name:     name,
		mail:     mail,
		password: password,
	}
}

func NewUUIDFromString(uuidStr string) (UUID, error) {
	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		return UUID(uuid.Nil), err
	}

	return UUID(parsedUUID), nil
}

func NewName(name string) (Name, error) {
	trimmedName := strings.TrimSpace(name)

	err := validateName(trimmedName)
	if err != nil {
		return "", err
	}

	return Name(trimmedName), nil
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

func NewMail(mail string) (Mail, error) {
	trimmedMail := strings.TrimSpace(mail)

	err := validateMail(trimmedMail)
	if err != nil {
		return "", err
	}

	return Mail(trimmedMail), nil
}

func validateMail(mail string) error {
	ok := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(mail)
	if !ok {
		return commons.NewValidationError("invalid format email")
	}

	return nil
}

func NewPassword(password string, passHasher PasswordHasher) (Password, error) {
	trimmedPassword := strings.TrimSpace(password)

	err := validatePassword(trimmedPassword)
	if err != nil {
		return "", err
	}

	hashed, err := passHasher.Hash(trimmedPassword)
	if err != nil {
		return "", err
	}

	return Password(hashed), nil
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

func (u *User) UUID() UUID {
	return u.uuid
}

func (u *User) Name() Name {
	return u.name
}

func (u *User) Mail() Mail {
	return u.mail
}

func (u *User) Password() Password {
	return u.password
}

func (u UUID) Value() (driver.Value, error) {
	return uuid.UUID(u).String(), nil
}
