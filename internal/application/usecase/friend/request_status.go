package friend

import (
	"context"
	"palbum/internal/domain/commons"
	"palbum/internal/domain/friend"
	"palbum/internal/domain/user"
)

type RequestStatusInput struct {
	UserUUID    user.UUID
	RequestUUID friend.RequestUUID
	Status      friend.Status
}

type RequestStatusUsecase struct {
	txRepo         commons.TxManagerRepository
	requestRepo    friend.FriendRequestRepository
	friendshipRepo friend.FriendshipRepository
}

func NewRequestStatusUsecase(
	txRepo commons.TxManagerRepository,
	requestRepo friend.FriendRequestRepository,
	friendshipRepo friend.FriendshipRepository,
) *RequestStatusUsecase {
	return &RequestStatusUsecase{
		txRepo:         txRepo,
		requestRepo:    requestRepo,
		friendshipRepo: friendshipRepo,
	}
}

func (u *RequestStatusUsecase) Execute(ctx context.Context, input RequestStatusInput) error {
	if input.Status == friend.Pending {
		return commons.NewBadRequestError("cannot change status to pending")
	}

	friendRequest, err := u.requestRepo.FindByUUID(ctx, input.RequestUUID)
	if err != nil {
		return err
	}

	if friendRequest.ToUserID() != input.UserUUID {
		return commons.NewForbiddenError("you are not allowed to change this request status")
	}

	var newFriendship *friend.Friendship

	newFriendship, err = u.changeRequestStatus(input.Status, friendRequest)
	if err != nil {
		return err
	}

	err = u.persistTransaction(ctx, friendRequest, newFriendship)
	if err != nil {
		return err
	}

	return nil
}

func (u *RequestStatusUsecase) changeRequestStatus(
	status friend.Status,
	friendRequest *friend.FriendRequest,
) (*friend.Friendship, error) {
	switch status {
	case friend.Rejected:
		return nil, friendRequest.Reject()
	case friend.Accepted:
		return friendRequest.Accept()
	case friend.Pending:
		return nil, commons.NewBadRequestError("cannot change status to pending")
	default:
		return nil, commons.NewBadRequestError("invalid status")
	}
}

func (u *RequestStatusUsecase) persistTransaction(
	ctx context.Context,
	req *friend.FriendRequest,
	newFriendship *friend.Friendship,
) error {
	return u.txRepo.WithInTx(ctx, func(txCtx context.Context) error {
		if newFriendship != nil {
			err := u.friendshipRepo.Save(txCtx, newFriendship)
			if err != nil {
				return err
			}
		}

		return u.requestRepo.Update(txCtx, req)
	})
}
