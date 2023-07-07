package di

import (
	"MAXPUMP1/pkg/api/handlers"
	"MAXPUMP1/pkg/db"
	"MAXPUMP1/pkg/repository"
	"MAXPUMP1/pkg/usecase"

	"github.com/google/wire"
)

func InitializeApi() *handlers.UserHandler {
	wire.Build(
		db.ConnectDB,
		repository.NewUserRepository,
		usecase.NewUser,
		handlers.NewUserHandler,
	)

	return &handlers.UserHandler{}
}
