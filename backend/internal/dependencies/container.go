package dependencies

import (
	application_friend "palbum/internal/application/usecase/friend"
	application_user "palbum/internal/application/usecase/user"
	"palbum/internal/infrastructure/persistence"
	persistence_auth "palbum/internal/infrastructure/persistence/auth"
	persistence_friend "palbum/internal/infrastructure/persistence/friend"
	persistence_user "palbum/internal/infrastructure/persistence/user"
	"palbum/internal/presentation"
	"palbum/internal/presentation/middleware"
	"palbum/internal/presentation/utils"
	"palbum/internal/route"

	"palbum/internal/infrastructure/security"
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
		log.NewLogger,
		presentation.NewUserHandler,
		presentation.NewFriendHandler,
		utils.NewErrorHandler,
		persistence.NewBaseRepository,
		persistence_user.NewUserRepositoryImpl,
		persistence_auth.NewUserAuthRepositoryImpl,
		application_user.NewSignUpUsecase,
		application_user.NewUserLoginUsecase,
		application_user.NewUserLogoutUsecase,
		application_user.NewRefreshUsecase,
		application_friend.NewFriendCodeUsecase,
		application_friend.NewFriendRequestUsecase,
		application_friend.NewRequestStatusUsecase,
		application_friend.NewFriendListUsecase,
		persistence_friend.NewFriendshipRepositoryImpl,
		persistence_friend.NewFriendCodeRepositoryImpl,
		persistence_friend.NewFriendRequestRepositoryImpl,
		security.NewJWTService,
		security.NewBcryptHasher,
		middleware.NewAuthMiddleware,
		middleware.NewRequestMiddleware,
		route.NewHandler,
	}

	for _, dependency := range dependencies {
		err := container.Provide(dependency)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return container, nil
}
