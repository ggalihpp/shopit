package user

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ggalihpp/shopit/auth"
	"github.com/labstack/echo"
)

// Handler -
type Handler struct {
	Repository *Repository
}

// SetRoutes -
func (h *Handler) SetRoutes(r *echo.Group) {
	r.POST("/", h.register)
	r.POST("/l", h.login)
	r.GET("/test", h.test)
}

func (h *Handler) register(c echo.Context) error {
	input := new(User)

	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	err := h.Repository.InsertUser(input)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, input)
}

func (h *Handler) login(c echo.Context) error {
	var (
		checkPW bool
		err     error
	)

	u := new(User)

	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	if h.Repository.IsUserExist(u) {
		u, checkPW, err = h.Repository.CheckUserPassword(u)
		if err != nil {
			return echo.NewHTTPError(500, err.Error())
		}

		if !checkPW {
			return echo.NewHTTPError(403, "Password incorrect sir.")
		}
	}

	// TODO: RETURN A TOKEN TO THE LOGGED IN USER
	claims := &auth.JwtCustomClaim{
		Username: u.Username,
		Email:    u.Email,
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "auth"
	cookie.Value = t
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (h *Handler) test(c echo.Context) error {
	cookie, err := c.Cookie("auth")
	if err != nil {
		return err
	}
	spew.Dump(c.Cookie("auth"))
	return c.JSON(200, echo.Map{
		"name":  cookie.Name,
		"value": cookie.Value,
	})
}
