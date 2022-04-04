package api

import (
	"excel-read/service"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func CheckLogin(c echo.Context) error {
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

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, err := service.HashPassword(password)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, hash)
}
