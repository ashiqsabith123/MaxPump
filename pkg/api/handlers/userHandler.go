package handlers

import (
	"MaxPump/delivery/middleware"
	"MaxPump/delivery/models"
	"MaxPump/domain/entity"
	"MaxPump/repository/infrastructure"
	usecase "MaxPump/usecase/user"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserUsecase *usecase.UserUsecase
}

func NewUserHandler(UserUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase}
}

func (uh *UserHandler) UserSignup(c *gin.Context) {

	var userInput models.Users

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user entity.User
	copier.Copy(&user, &userInput)
	newUser, err := uh.UserUsecase.ExecuteSignup(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func (uh *UserHandler) SignupWithOtp(c *gin.Context) {
	var user models.Signup
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	key, err := uh.UserUsecase.ExecuteSignupWithOtp()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Otp send succesfuly to": user.Phone, "Key": key})
	}
}

func (uh *UserHandler) SignupOtpValidation(c *gin.Context) {
	key := c.PostForm("key")
	otp := c.PostForm("otp")
	err := uh.UserUsecase.ExecuteSignupOtpValidation(key, otp)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"massage": "user signup succesfull"})
	}

}

func UserLogin(c *gin.Context) {
	ok := middleware.ValidateCookie(c)

	if !ok {
		c.Status(http.StatusNotFound) //404 (Not Found)
	} else {
		c.Status(http.StatusOK) // Set HTTP status to 200 OK
	}
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginPost(c *gin.Context) {
	// Parse and decode the request body into a `Credentials` struct
	creds := &Credentials{}
	if err := c.ShouldBindJSON(creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	enteredPassword := creds.Password

	var storedPassword string
	result := infrastructure.DB.Raw("SELECT password FROM users WHERE email = ?", creds.Email).Scan(&storedPassword)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	fmt.Println("Stored Password:", storedPassword)
	fmt.Println("Entered Password (Hashed):", enteredPassword)

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(enteredPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	// Generate the token for the authenticated user
	token, err := middleware.GenToken(creds.Email, c)
	if err != nil {
		// If there is an error in token generation, return a 500 error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation error", "token": token})
		return
	}

	// If we reach this point, the user is considered authorized
	// Send a success response
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func UserLogout(c *gin.Context) {

	err := middleware.DeleteCookie(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user cookie deletion failed"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
	}

}