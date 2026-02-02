package friend_test

import (
	"context"
	"palbum/internal/domain/commons"
	domain "palbum/internal/domain/friend"
	"palbum/internal/domain/user"
	"palbum/internal/infrastructure/persistence"
	infra "palbum/internal/infrastructure/persistence/friend"
	"palbum/internal/test"
	"testing"
	"time"

	"github.com/go-jose/go-jose/v4/testutils/assert"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
)

func setupRequest(t *testing.T) domain.FriendRequestRepository {
	t.Helper()

	testDB := test.NewTestDB(t, &infra.FriendRequest{})
	baseRepo := persistence.NewBaseRepository(testDB)

	return infra.NewFriendRequestRepositoryImpl(baseRepo)
}

func createRequest(t *testing.T) *domain.FriendRequest {
	t.Helper()
	
	fromUUID := user.UUID(uuid.New())
	toUUID := user.UUID(uuid.New())

	request := domain.NewFriendRequest(fromUUID, toUUID)

	return request
}

func TestFriendRequestRepositoryImpl_Create(t *testing.T) {
	t.Parallel()

	t.Run("Success: create request", func(t *testing.T) {
		t.Parallel()
		repo := setupRequest(t)

		ctx := context.Background()
		request := createRequest(t)

		err := repo.Create(ctx, request)
		assert.NoError(t, err, "should not be error")
	})
}

func TestFriendRequestRepositoryImpl_FindByUUID(t *testing.T) {
	t.Parallel()

	t.Run("Success: find by uuid", func(t *testing.T) {
		t.Parallel()

		repo := setupRequest(t)
		request := createRequest(t)
		_ = repo.Create(t.Context(), request)

		get, err := repo.FindByUUID(t.Context(), request.RequestUUID())
		assert.NoError(t, err, "should not be error")

		if diff := cmp.Diff(get, request, cmp.Options{
			cmp.AllowUnexported(domain.FriendRequest{}),
			cmpopts.EquateApproxTime(1 * time.Second),
		}); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("Error: not exist record", func(t *testing.T) {
		t.Parallel()

		repo := setupRequest(t)
		request := createRequest(t)

		get, err := repo.FindByUUID(t.Context(), request.RequestUUID())
		assert.ErrorIs(t, err, commons.ErrNotFound)
		assert.Equal(t, get, nil)
	})
}

func TestFriendRequestRepositoryImpl_FindByUserUUID(t *testing.T) {
	t.Parallel()

	t.Run("Success: find by user uuid", func(t *testing.T) {
		t.Parallel()
		repo := setupRequest(t)
		request := createRequest(t)
		_ = repo.Create(t.Context(), request)

		get, err := repo.FindByUserUUID(t.Context(), request.FromUserID(), domain.CategorySent)
		assert.NoError(t, err, "should not be error")
		assert.Equal(t, len(get), 1)
		assert.Equal(t, request.ToUserID(), get[0].ToUserID())

		get, err = repo.FindByUserUUID(t.Context(), request.ToUserID(), domain.CategoryReceived)
		assert.NoError(t, err, "should not be error")
		assert.Equal(t, len(get), 1)
		assert.Equal(t, request.ToUserID(), get[0].ToUserID())
	})

	t.Run("Error: find by user uuid", func(t *testing.T) {
		t.Parallel()
		repo := setupRequest(t)
		request := createRequest(t)
		_ = repo.Create(t.Context(), request)

		get, err := repo.FindByUserUUID(t.Context(), user.UUID(uuid.New()), domain.CategoryReceived)
		assert.ErrorIs(t, err, commons.ErrNotFound)
		assert.Equal(t, len(get), 0)
	})
}
func TestFriendRequestRepositoryImpl_ChangeStatus(t *testing.T) {
	t.Parallel()

	t.Run("Success: change status to accepted", func(t *testing.T) {
		t.Parallel()

		repo := setupRequest(t)
		request := createRequest(t)
		_ = repo.Create(t.Context(), request)

		ctx := t.Context()

		_, _ = request.Accept()

		err := repo.ChangeStatus(ctx, request)
		assert.NoError(t, err, "should not be error")
	})

	t.Run("Error: not exist record", func(t *testing.T) {
		t.Parallel()

		repo := setupRequest(t)
		request := createRequest(t)

		ctx := t.Context()

		_, _ = request.Accept()
		err := repo.ChangeStatus(ctx, request)
		assert.ErrorIs(t, err, commons.ErrNotFound)
	})
}
