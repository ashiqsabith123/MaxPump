package interfaces

import (
	"MAXPUMP1/pkg/domain/entity"
	"MAXPUMP1/pkg/model"
)

type UserUsecase interface {
	ExecuteSignup(user entity.User) (*entity.User, error)
	ExecuteSignupWithOtp(user model.Signup) (string, error)
	ExecuteSignupOtpValidation(key string, otp string) error
}
