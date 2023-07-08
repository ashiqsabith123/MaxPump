package usecase

import (
	"MAXPUMP1/pkg/domain/entity"
	"MAXPUMP1/pkg/model"
	repo "MAXPUMP1/pkg/repository/interfaces"
	use "MAXPUMP1/pkg/usecase/interfaces"
	"MAXPUMP1/pkg/utils"
	"fmt"

	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepo repo.UserInterface
}

func NewUser(userRepo repo.UserInterface) use.UserUsecaseInterface {
	return &UserUsecase{userRepo: userRepo}
}

func (us *UserUsecase) ExecuteSignup(user entity.User) (*entity.User, error) {
	fmt.Println("hello")
	fmt.Println("sa", user)

	email, err := us.userRepo.GetByEmail(user.Email) //call GetByEmail function to check the email is already exists or not
	if err != nil {
		return nil, errors.New("error with server")
	}

	if email != nil {
		return nil, errors.New("Email Already Exists")
	}

	phone, err := us.userRepo.GetByPhone(user.Phone) //call GetByPhone function to check the Phone number is already exists or not
	if err != nil {
		return nil, errors.New("error with serever")

	}
	if phone != nil {

		return nil, errors.New("Phone Number Already Exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {

		return nil, err
	}

	newUser := &entity.User{

		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  string(hashedPassword),
	}

	err1 := us.userRepo.Create(newUser)

	if err1 != nil {
		return nil, err1
	}

	return newUser, nil

}

func (uu *UserUsecase) ExecuteSignupWithOtp(user model.Signup) (string, error) {
	var otpKey entity.OtpKey
	email, err := uu.userRepo.GetByEmail(user.Email)
	if err != nil {
		return "", errors.New("error with server")
	}
	if email != nil {
		return "", errors.New("user with this email already exists")
	}
	phone, err := uu.userRepo.GetByPhone(user.Phone)
	if err != nil {
		return "", errors.New("error with server")
	}
	if phone != nil {
		return "", errors.New("user with this phone no already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)
	key, err := utils.SendOtp(user.Phone)
	if err != nil {
		return "", err
	} else {
		//err = uu.userRepo.CreateSignup(&user)
		otpKey.Key = key
		otpKey.Phone = user.Phone
		//err = uu.userRepo.CreateOtpKey(&otpKey)
		if err != nil {
			return "", err
		}
		return key, nil
	}
}

func (uu *UserUsecase) ExecuteSignupOtpValidation(key string, otp string) error {
	//result, err := uu.userRepo.GetByKey(key)
	// if err != nil {
	// 	return err
	// }
	// user, err := uu.userRepo.GetSignupByPhone(result.Phone)
	// if err != nil {
	// 	return err
	// }
	// err = utils.CheckOtp(result.Phone, otp)
	// if err != nil {
	// 	return err
	// } else {
	// 	newUser := &entity.User{
	// 		FirstName: user.FirstName,
	// 		LastName:  user.LastName,
	// 		Email:     user.Email,
	// 		Phone:     user.Phone,
	// 		Password:  user.Password,
	// 	}

	// 	err1 := uu.userRepo.Create(newUser)
	// 	if err1 != nil {
	// 		return err1
	// 	} else {
	// 		return nil
	// 	}
	// }

	return nil
}
