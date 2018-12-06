package middleware

import (
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ggalihpp/shopit/auth"

	"github.com/labstack/echo"
)

var keyFunc = func(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET")), nil
}

// UserContext - Used for passing decoded token from middleware to handlers
type UserContext struct {
	echo.Context
	User *auth.JwtCustomClaim
}

// GetUserName - Return a logged in username
func (c *UserContext) GetUserName() string {

	return c.User.Username
}

// GetEmail - Return a logged in email
func (c *UserContext) GetEmail() string {

	return c.User.Email
}
