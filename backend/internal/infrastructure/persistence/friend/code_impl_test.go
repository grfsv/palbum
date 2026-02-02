package friend_test

import (
	domain "palbum/internal/domain/friend"
	"palbum/internal/domain/user"
	"palbum/internal/infrastructure/persistence"
	infra "palbum/internal/infrastructure/persistence/friend"
	"palbum/internal/test"
	"testing"

	"github.com/go-jose/go-jose/v4/testutils/assert"
	"github.com/google/uuid"
)

func setupFriendCode(t *testing.T) domain.FriendCodeRepository {
	t.Helper()

	db := test.NewTestDB(t, &infra.FriendCode{})
	baseRepo := persistence.NewBaseRepository(db)

	return infra.NewFriendCodeRepositoryImpl(baseRepo)
}

func TestFriendCodeRepositoryImpl_Save(t *testing.T) {
	t.Parallel()

	t.Run("Success: ", func(t *testing.T) {
		t.Parallel()

		repo := setupFriendCode(t)
		ctx := t.Context()
		code, _ := domain.NewFriendCode(user.UUID(uuid.New()))

		err := repo.Save(ctx, code)
		assert.NoError(t, err)
	})
}

func TestFriendCodeRepositoryImpl_FindByUUID(t *testing.T) {
	t.Parallel()

	t.Run("Success: ", func(t *testing.T) {
		t.Parallel()

		repo := setupFriendCode(t)
		ctx := t.Context()
		code, _ := domain.NewFriendCode(user.UUID(uuid.New()))
		_ = repo.Save(ctx, code)

		get, err := repo.FindByUUID(ctx, code.UserUUID())
		assert.NoError(t, err)
		assert.Equal(t, get.Code(), code.Code())
	})
}

func TestFriendCodeRepositoryImpl_FindByCode(t *testing.T) {
	t.Parallel()

	t.Run("Success: ", func(t *testing.T) {
		t.Parallel()

		repo := setupFriendCode(t)
		ctx := t.Context()
		code, _ := domain.NewFriendCode(user.UUID(uuid.New()))
		_ = repo.Save(ctx, code)

		get, err := repo.FindByCode(ctx, code.Code())
		assert.NoError(t, err)
		assert.Equal(t, get.Code(), code.Code())
	})
}
