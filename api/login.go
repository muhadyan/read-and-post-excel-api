package api

import (
	"excel-read/service"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func HandleLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := service.CheckLogin(c, username, password)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	if !res {
		return echo.ErrUnauthorized
	}

	// generate token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func HandleSignUp(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	err := service.SignUp(c, username, password)
	if err != nil {
		return err
	}

	return nil
}
