package friend_test

import (
	"palbum/internal/domain/commons"
	domain "palbum/internal/domain/friend"
	"palbum/internal/domain/user"
	"palbum/internal/infrastructure/persistence"
	infra "palbum/internal/infrastructure/persistence/friend"
	"palbum/internal/test"
	"testing"
	"time"

	"github.com/go-jose/go-jose/v4/testutils/assert"
	"github.com/google/uuid"
)

func setupFriendship(t *testing.T) domain.FriendshipRepository {
	t.Helper()
	db := test.NewTestDB(t, &infra.Friendship{})
	baseRepo := persistence.NewBaseRepository(db)

	return infra.NewFriendshipRepositoryImpl(baseRepo)
}

func TestFriendshipRepositoryImpl_Save(t *testing.T) {
	t.Parallel()
	t.Run("Success: can save data", func(t *testing.T) {
		t.Parallel()
		repo := setupFriendship(t)

		friendUUID := uuid.New()
		userUUID1 := uuid.New()
		userUUID2 := uuid.New()

		request := domain.ReconstructFriendship(
			friendUUID,
			userUUID1,
			userUUID2,
			time.Now(),
		)

		err := repo.Create(t.Context(), request)
		if err != nil {
			assert.Error(t, err, "should not be error")
		}
	})
	t.Run("Error: can not save duplicate data", func(t *testing.T) {
		t.Parallel()
		repo := setupFriendship(t)

		friendUUID := uuid.New()
		userUUID1 := uuid.New()
		userUUID2 := uuid.New()

		request := domain.ReconstructFriendship(
			friendUUID,
			userUUID1,
			userUUID2,
			time.Now(),
		)
		_ = repo.Create(t.Context(), request)

		err := repo.Create(t.Context(), request)
		assert.ErrorIs(t, err, commons.ErrDuplicate)
	})
}

func TestFriendshipRepositoryImpl_FindByUserUUID(t *testing.T) {
	t.Parallel()
	// フレンド関係にあるユーザーがそれぞれ1件ずつ取得できる
	t.Run("Success: can get data by user1", func(t *testing.T) {
		t.Parallel()

		repo := setupFriendship(t)

		friendUUID := uuid.New()
		userUUID1 := uuid.New()
		userUUID2 := uuid.New()

		request := domain.ReconstructFriendship(
			friendUUID,
			userUUID1,
			userUUID2,
			time.Now(),
		)

		_ = repo.Create(t.Context(), request)

		friendships, err := repo.FindByUserUUID(t.Context(), user.UUID(userUUID1))
		assert.NoError(t, err, "should not be error")
		assert.Equal(t, len(friendships), 1)
	})
	t.Run("Success: can get data by user2", func(t *testing.T) {
		t.Parallel()

		repo := setupFriendship(t)

		friendUUID := uuid.New()
		userUUID1 := uuid.New()
		userUUID2 := uuid.New()

		request := domain.ReconstructFriendship(
			friendUUID,
			userUUID1,
			userUUID2,
			time.Now(),
		)

		_ = repo.Create(t.Context(), request)
		friendships, err := repo.FindByUserUUID(t.Context(), user.UUID(userUUID2))
		assert.NoError(t, err, "should not be error")
		assert.Equal(t, len(friendships), 1)
	})
}

func TestFriendshipRepositoryImpl_Delete(t *testing.T) {
	t.Parallel()
}
