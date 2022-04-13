package api

import (
	"excel-read/auth"
	"net/http"

	"github.com/labstack/echo"
)

func HandleLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := auth.CheckLogin(c, username, password)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	if !res {
		return echo.ErrUnauthorized
	}

	err = auth.GenerateToken(c, username)
	if err != nil {
		return err
	}

	return nil
}

func HandleSignUp(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	err := auth.SignUp(c, username, password)
	if err != nil {
		return err
	}

	return nil
}
