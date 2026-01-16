package friend

import (
	"context"
	"palbum/internal/domain/friend"
	"palbum/internal/domain/user"
	"time"

	"github.com/google/uuid"
)

type RequestListInput struct {
	UserUUID user.UUID
	Category friend.RequestCategory
}

type RequestListOutput struct {
	Requests []Request `json:"requests"`
}

type Request struct {
	RequestUUID string `json:"requestUuid"`
	UserName    string `json:"userName"`
	RequestAt   string `json:"requestAt"`
}

type RequestListUsecase struct {
	requestRepo friend.FriendRequestRepository
	userRepo    user.UserRepository
}

func NewRequestListUsecase(
	requestRepo friend.FriendRequestRepository,
	userRepo user.UserRepository,
) *RequestListUsecase {
	return &RequestListUsecase{
		requestRepo: requestRepo,
		userRepo:    userRepo,
	}
}

func (u *RequestListUsecase) Execute(ctx context.Context, input RequestListInput) (RequestListOutput, error) {
	var out RequestListOutput

	requests, err := u.requestRepo.FindByUserUUID(ctx, input.UserUUID, input.Category)
	if err != nil {
		return out, err
	}

	userUUIDs := make([]user.UUID, 0, len(requests))
	for _, request := range requests {
		switch input.Category {
		case friend.CategorySent:
			userUUIDs = append(userUUIDs, request.ToUserID())
		case friend.CategoryReceived:
			userUUIDs = append(userUUIDs, request.FromUserID())
		}
	}

	users, err := u.userRepo.ListInUUID(ctx, userUUIDs)
	if err != nil {
		return out, err
	}

	out.Requests = make([]Request, len(requests))
	for i, request := range requests {
		out.Requests[i] = Request{
			RequestUUID: uuid.UUID(request.RequestUUID()).String(),
			UserName:    string(users[i].Name()),
			RequestAt:   request.RequestAt().Format(time.RFC3339),
		}
	}

	return out, nil
}
