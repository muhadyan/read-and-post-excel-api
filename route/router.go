package route

import (
	"excel-read/api"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.POST("/inputbooks", api.InputBooks)
	e.GET("/getbooks", api.GetBooks)
	return e
}
