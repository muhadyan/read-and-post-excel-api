package route

import (
	"excel-read/api"
	"excel-read/auth"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.POST("/signup", api.HandleSignUp)
	e.POST("/login", api.HandleLogin)

	book := e.Group("/book", auth.IsAuthenticated)
	{
		book.POST("/inputbooks", api.InputBooks)
		book.GET("/getbooks", api.GetBooks)
		book.GET("/getpdf", api.GetPdf)
	}

	return e
}
