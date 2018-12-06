package middleware

import (
	"github.com/ggalihpp/shopit/auth"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
)

// CheckJWT - Will check the origin of the requester valid or not
func CheckJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("auth")
		if err != nil {
			return echo.NewHTTPError(401, err.Error())
		}

		claims := new(auth.JwtCustomClaim)

		token, err := jwt.ParseWithClaims(cookie.Value, claims, keyFunc)
		if err != nil {
			return echo.NewHTTPError(401, err.Error())
		}

		user := token.Claims.(*auth.JwtCustomClaim)

		cc := &UserContext{
			c,
			user,
		}

		return next(cc)
	}
}
