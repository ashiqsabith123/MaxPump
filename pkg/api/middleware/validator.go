package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// MyClaims represents custom claims for JWT
type MyClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// we define the expiration time of JWT, taking 2 hours
const TokenExpireDuration = time.Hour * 2

func GenToken(username string, c *gin.Context) (string, error) {
	//MySecretKEY := os.Getenv("MySecretKEY")
	MySecretKEY := "hello Hii"
	// Create our own statement
	claims := MyClaims{
		Username: username, // Custom field
		Role:     "user",   // Set the user's role
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // Expiration time
			Issuer:    "my-project",                               // Issuer
		},
	}
	// Creates a signed object using the specified signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(MySecretKEY))
	if err != nil {
		return "", err
	}

	// Calculate the number of seconds until the token expires
	expireSeconds := int(TokenExpireDuration.Seconds())

	// Set the cookie with the specified name and expiration
	c.SetCookie("Authorize", tokenString, expireSeconds, "", "", false, true)

	return tokenString, nil
}

func ValidateCookie(c *gin.Context) bool {
	cookie, err := c.Cookie("Authorize")
	if err != nil || cookie == "" {
		// Cookie not found or error occurred
		return false
	}
	// Cookie found
	return true
}

func DeleteCookie(c *gin.Context) error {
	c.SetCookie("Authorise", "", 0, "", "", false, true)
	fmt.Println("Cookie deleted")
	return nil
}
