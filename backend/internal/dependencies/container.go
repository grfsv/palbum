package dependencies

import (
	application_dtk "palbum/internal/application/usecase/dtk"
	application_friend "palbum/internal/application/usecase/friend"
	application_user "palbum/internal/application/usecase/user"
	"palbum/internal/infrastructure/persistence"
	persistence_auth "palbum/internal/infrastructure/persistence/auth"
	persistence_dtk "palbum/internal/infrastructure/persistence/dtk"
	persistence_friend "palbum/internal/infrastructure/persistence/friend"
	persistence_user "palbum/internal/infrastructure/persistence/user"
	"palbum/internal/infrastructure/security"
	"palbum/internal/presentation"
	"palbum/internal/presentation/middleware"
	"palbum/internal/presentation/utils"
	"palbum/internal/route"
	"palbum/internal/utils/log"

	"github.com/cockroachdb/errors"
	"go.uber.org/dig"
)

func InitContainer() (*dig.Container, error) {
	container := dig.New()

	dependencies := []any{
		NewConfig,
		persistence.InitDB,
		persistence.NewTransactionManager,
		persistence.NewBaseRepository,
		log.NewLogger,
		utils.NewErrorHandler,
		security.NewJWTService,
		security.NewBcryptHasher,
		middleware.NewAuthMiddleware,
		middleware.NewRequestMiddleware,
		route.NewHandler,

		presentation.NewUserHandler,
		application_user.NewSignUpUsecase,
		application_user.NewUserLoginUsecase,
		application_user.NewUserLogoutUsecase,
		application_user.NewRefreshUsecase,
		persistence_user.NewUserRepositoryImpl,
		persistence_auth.NewUserAuthRepositoryImpl,

		presentation.NewFriendHandler,
		application_friend.NewFriendCodeUsecase,
		application_friend.NewFriendRequestUsecase,
		application_friend.NewRequestStatusUsecase,
		application_friend.NewFriendListUsecase,
		persistence_friend.NewFriendshipRepositoryImpl,
		persistence_friend.NewFriendCodeRepositoryImpl,
		persistence_friend.NewFriendRequestRepositoryImpl,

		presentation.NewDTKHandler,
		application_dtk.NewGetMyDTKUsecase,
		application_dtk.NewGetFriendDTKsUsecase,
		persistence_dtk.NewDTKRepositoryImpl,
	}

	for _, dependency := range dependencies {
		err := container.Provide(dependency)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return container, nil
}
