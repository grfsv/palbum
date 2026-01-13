package dependencies

import (
	application_user "palbum/internal/application/usecase/user"
	"palbum/internal/infrastructure/persistence"
	persistence_auth "palbum/internal/infrastructure/persistence/auth"
	persistence_user "palbum/internal/infrastructure/persistence/user"
	"palbum/internal/presentation"
	"palbum/internal/presentation/middleware"
	"palbum/internal/route"

	"palbum/internal/infrastructure/security"
	"palbum/internal/utils/log"

	"github.com/cockroachdb/errors"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func InitContainer() (*dig.Container, error) {
	container := dig.New()

	err := container.Provide(NewConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = container.Provide(
		func(cfg *persistence.DBConfig) *gorm.DB {
			db, err := persistence.InitDB(cfg)
			if err != nil {
				panic(err)
			}

			return db
		})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = container.Provide(persistence.NewTransactionManager)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dependencies := []any{
		log.NewLogger,
		presentation.NewUserHandler,
		presentation.NewErrorHandler,
		persistence.NewBaseRepository,
		persistence_user.NewUserRepositoryImpl,
		persistence_auth.NewUserAuthRepositoryImpl,
		application_user.NewSignUpUsecase,
		application_user.NewUserLoginUsecase,
		application_user.NewUserLogoutUsecase,
		application_user.NewRefreshUsecase,
		security.NewJWTService,
		security.NewBcryptHasher,
		middleware.NewAuthMiddleware,
		middleware.NewRequestMiddleware,
	}

	for _, dependency := range dependencies {
		err := container.Provide(dependency)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	err = container.Provide(route.NewHandler)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return container, nil
}
