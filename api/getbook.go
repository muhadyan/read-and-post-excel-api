package api

import (
	"excel-read/db"
	"excel-read/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetBooks(c echo.Context) error {
	db := db.DbManager()
	books := []model.Books{}
	db.Find(&books)
	return c.JSON(http.StatusOK, books)
}
