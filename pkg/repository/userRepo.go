package repository

import (
	"MAXPUMP1/pkg/domain/entity"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetByID(id int) (*entity.User, error) {
	var user entity.User
	result := ur.db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepository) GetByEmail(email string) (*entity.User, error) {

	var user entity.User
	result := ur.db.First(&user, email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil

}

func (ur *UserRepository) GetByPhone(Phone string) (*entity.User, error) {

	var user entity.User
	result := ur.db.First(&user, Phone)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil

}

func (ur *UserRepository) Create(user *entity.User) error {
	return ur.db.Create(user).Error
}
