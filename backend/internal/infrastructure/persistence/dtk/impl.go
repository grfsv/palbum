package dtk

import (
	"context"
	"time"

	"palbum/internal/domain/commons"
	"palbum/internal/domain/dtk"
	"palbum/internal/domain/user"
	"palbum/internal/infrastructure/persistence"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type DTKRepositoryImpl struct {
	*persistence.Database
}

func NewDTKRepositoryImpl(baseRepo *persistence.Database) dtk.DTKRepository {
	return &DTKRepositoryImpl{baseRepo}
}

func (r *DTKRepositoryImpl) FindByUserUUIDAndDate(
	ctx context.Context,
	userUUID user.UUID,
	date time.Time,
) (*dtk.DTK, error) {
	entity, err := gorm.G[DTK](r.GetDBFromContext(ctx)).
		Where("user_uuid = ? AND date = ?", uuid.UUID(userUUID), date).
		First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, commons.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return entity.ToDomain(), nil
}

func (r *DTKRepositoryImpl) FindByUserUUIDsAndSince(
	ctx context.Context,
	userUUIDs []user.UUID,
	since time.Time,
) ([]*dtk.DTK, error) {
	if len(userUUIDs) == 0 {
		return []*dtk.DTK{}, nil
	}

	uuids := make([]uuid.UUID, len(userUUIDs))
	for i, u := range userUUIDs {
		uuids[i] = uuid.UUID(u)
	}

	rows, err := gorm.G[DTK](r.GetDBFromContext(ctx)).
		Where("user_uuid IN ? AND date >= ?", uuids, since).
		Find(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dtks := make([]*dtk.DTK, len(rows))
	for i, row := range rows {
		dtks[i] = row.ToDomain()
	}

	return dtks, nil
}

func (r *DTKRepositoryImpl) Save(ctx context.Context, d *dtk.DTK) error {
	entity := ToDTKEntity(d)

	err := gorm.G[DTK](r.GetDBFromContext(ctx)).Create(ctx, entity)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return commons.ErrDuplicate
		}

		return errors.WithStack(err)
	}

	return nil
}
