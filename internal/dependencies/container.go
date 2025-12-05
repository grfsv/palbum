package dependencies

import (
	"remind_map/internal/application/usecase"
	"remind_map/internal/infrastructure/persistence"
	"remind_map/internal/presentation"
	"remind_map/internal/presentation/middleware"
	"remind_map/internal/route"

	"remind_map/internal/infrastructure/security"
	"remind_map/internal/utils/log"

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
		persistence.NewUserRepositoryImpl,
		persistence.NewUserAuthRepositoryImpl,
		usecase.NewSignUpUsecase,
		usecase.NewUserLoginUsecase,
		usecase.NewUserLogoutUsecase,
		usecase.NewRefreshUsecase,
		security.NewJWTService,
		security.NewBcryptHasher,
		middleware.NewAuthMiddleware,
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
