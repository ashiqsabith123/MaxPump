package interfaces

import "MAXPUMP1/pkg/domain/entity"

type UserInterface interface {
	GetByID(id int) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	GetByPhone(Phone string) (*entity.User, error)
	Create(user *entity.User) error
}
