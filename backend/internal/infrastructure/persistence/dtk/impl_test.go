package dtk_test

import (
	"testing"
	"time"

	domain "palbum/internal/domain/dtk"
	"palbum/internal/domain/user"
	"palbum/internal/infrastructure/persistence"
	infra "palbum/internal/infrastructure/persistence/dtk"
	"palbum/internal/test"

	"github.com/go-jose/go-jose/v4/testutils/assert"
	"github.com/google/uuid"
)

func setupDTK(t *testing.T) domain.DTKRepository {
	t.Helper()

	db := test.NewTestDB(t, &infra.DTK{})
	baseRepo := persistence.NewBaseRepository(db)

	return infra.NewDTKRepositoryImpl(baseRepo)
}

func TestDTKRepositoryImpl_Save(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		repo := setupDTK(t)
		ctx := t.Context()
		dtk, _ := domain.NewDTK(user.UUID(uuid.New()), domain.Today())

		err := repo.Save(ctx, dtk)
		assert.NoError(t, err)
	})
}

func TestDTKRepositoryImpl_FindByUserUUIDAndDate(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		repo := setupDTK(t)
		ctx := t.Context()
		userUUID := user.UUID(uuid.New())
		dtk, _ := domain.NewDTK(userUUID, domain.Today())
		_ = repo.Save(ctx, dtk)

		found, err := repo.FindByUserUUIDAndDate(ctx, userUUID, domain.Today())
		assert.NoError(t, err)
		assert.Equal(t, found.Key(), dtk.Key())
	})
}

func TestDTKRepositoryImpl_FindByUserUUIDsAndSince(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		repo := setupDTK(t)
		ctx := t.Context()

		userUUID1 := user.UUID(uuid.New())
		userUUID2 := user.UUID(uuid.New())

		dtk1, _ := domain.NewDTK(userUUID1, domain.Today())
		dtk2, _ := domain.NewDTK(userUUID2, domain.Today())
		_ = repo.Save(ctx, dtk1)
		_ = repo.Save(ctx, dtk2)

		since := time.Now().AddDate(0, 0, -1)
		found, err := repo.FindByUserUUIDsAndSince(ctx, []user.UUID{userUUID1, userUUID2}, since)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(found))
	})
}
