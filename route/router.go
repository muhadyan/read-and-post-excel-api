package route

import (
	"excel-read/api"
	"excel-read/service"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/generate-hash/:password", api.GenerateHashPassword)
	e.POST("/login", api.CheckLogin)

	e.POST("/inputbooks", api.InputBooks, service.IsAuthenticated)
	e.GET("/getbooks", api.GetBooks)

	return e
}
