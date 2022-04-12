package api

import (
	"excel-read/service"
	"net/http"

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

	err = service.GenerateToken(c, username)
	if err != nil {
		return err
	}

	return nil
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
