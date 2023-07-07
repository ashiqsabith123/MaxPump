package routes

import (
	"MAXPUMP1/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, userHandler *handlers.UserHandler) *gin.Engine {
	// Register the routes and the handlers
	r.POST("/signup", userHandler.UserSignup)
	r.POST("/login", userHandler.SignupWithOtp)
	r.POST("/loginpost", userHandler.SignupOtpValidation)
	r.POST("/logout", userHandler.UserLogout)

	return r // Return the gin.Engine instance
}
