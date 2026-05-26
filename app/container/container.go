package container

import (
	"myapp/app/http/handlers"
	"myapp/app/repositories"
	"myapp/app/services"

	"github.com/spf13/viper"
)

type Container struct {
	Config       *viper.Viper
	Repositories *repositories.Repositories
	Services     *services.Services
	Handlers     *handlers.Handlers
}
