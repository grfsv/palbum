package dtk

import (
	"crypto/rand"
	"time"

	"palbum/internal/domain/commons"
	"palbum/internal/domain/user"
)

const KeyLength = 16

type Key [KeyLength]byte

type DTK struct {
	userUUID user.UUID
	date     time.Time
	key      Key
}

func NewDTK(userUUID user.UUID, date time.Time) (*DTK, error) {
	var key Key

	_, err := rand.Read(key[:])
	if err != nil {
		return nil, commons.NewInternalError("failed to generate DTK", err)
	}

	return &DTK{
		userUUID: userUUID,
		date:     normalizeDate(date),
		key:      key,
	}, nil
}

func ReconstructDTK(userUUID user.UUID, date time.Time, key Key) *DTK {
	return &DTK{
		userUUID: userUUID,
		date:     normalizeDate(date),
		key:      key,
	}
}

func (d *DTK) UserUUID() user.UUID {
	return d.userUUID
}

func (d *DTK) Date() time.Time {
	return d.date
}

func (d *DTK) Key() Key {
	return d.key
}

func (k Key) Bytes() []byte {
	return k[:]
}

func normalizeDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

func Today() time.Time {
	return normalizeDate(time.Now())
}
