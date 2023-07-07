package di

import (
	"MaxPump1/pkg/usecase"
	"MaxPump1/pkg/repository"
	"MaxPump1/pkg/db"

	"github.com/google/wire"
)
 
func InitializeApi(){
	wire.Build(
		db.ConnectDB,
		repository.NewUserRepository,
		usecase.NewUser,

	)
}