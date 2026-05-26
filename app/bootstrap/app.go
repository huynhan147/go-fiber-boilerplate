package bootstrap

import (
	"myapp/app/container"
	"myapp/app/http/handlers"
	"myapp/app/repositories"
	"myapp/app/services"
	"myapp/pkg/cache"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func BuildContainer(
	db *gorm.DB,
	rdb *redis.Client,
	cfg *viper.Viper,
) *container.Container {

	// cache
	cacheStore := cache.New(rdb)

	// repositories
	userRepo := repositories.NewUserRepository(db)

	// services
	userService := services.NewUserService(
		userRepo,
		cacheStore,
	)

	authService := services.NewAuthService(
		userRepo,
		cacheStore,
		cfg,
	)

	// repositories
	transactionRepo := repositories.NewTransactionRepository(db)

	// services
	transactionService := services.NewTransactionService(
		transactionRepo,
	)

	// handlers
	userHandler := handlers.NewUserHandler(userService)

	authHandler := handlers.NewAuthHandler(authService)

	transactionHandler := handlers.NewTransactionHandler(transactionService)

	return &container.Container{
		Config: cfg,

		Repositories: &repositories.Repositories{
			User:        userRepo,
			Transaction: transactionRepo,
		},

		Services: &services.Services{
			User:        userService,
			Auth:        authService,
			Transaction: transactionService,
		},

		Handlers: &handlers.Handlers{
			User:        userHandler,
			Auth:        authHandler,
			Transaction: transactionHandler,
		},
	}
}
