package dtk

import (
	"time"

	"palbum/internal/domain/dtk"
	"palbum/internal/domain/user"

	"github.com/google/uuid"
)

type DTK struct {
	UserUUID uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	Date     time.Time `gorm:"type:date;not null;primaryKey"`
	Key      []byte    `gorm:"type:binary(16);not null"`
}

func (e *DTK) ToDomain() *dtk.DTK {
	var key dtk.Key
	copy(key[:], e.Key)

	return dtk.ReconstructDTK(
		user.UUID(e.UserUUID),
		e.Date,
		key,
	)
}

func ToDTKEntity(d *dtk.DTK) *DTK {
	return &DTK{
		UserUUID: uuid.UUID(d.UserUUID()),
		Date:     d.Date(),
		Key:      d.Key().Bytes(),
	}
}
