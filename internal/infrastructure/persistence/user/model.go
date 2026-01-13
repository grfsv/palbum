package user

import (
	"palbum/internal/domain/user"

	"github.com/google/uuid"
)

type UserEntity struct {
	UUID     uuid.UUID `gorm:"type:char(36);primaryKey"`
	Name     string    `gorm:"type:varchar(255);not null"`
	Mail     string    `gorm:"unique;not null"`
	Password string    `gorm:"type:varchar(255);not null"`
}

func ToEntity(user *user.User) *UserEntity {
	return &UserEntity{
		UUID:     uuid.UUID(user.UUID()),
		Name:     string(user.Name()),
		Mail:     string(user.Mail()),
		Password: string(user.Password()),
	}
}

func (e UserEntity) ToDomain() (*user.User, error) {
	uuid, err := user.NewUUIDFromString(e.UUID.String())
	if err != nil {
		return nil, err
	}

	name := user.Name(e.Name)
	mail := user.Mail(e.Mail)
	password := user.Password(e.Password)

	return user.NewUserWithUUID(uuid, name, mail, password), nil
}
