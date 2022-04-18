package api

import (
	"excel-read/service"
	"net/http"

	"github.com/labstack/echo"
)

func GetBooks(c echo.Context) error {
	createby := service.GetTokenData(c, "username")

	bookLists, rowPageList, err := service.GeneratePaginationFromRequest(c, createby)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Total Rows":  rowPageList.TotalRows,
		"Total Pages": rowPageList.TotalPages,
		"Data":        bookLists,
	})
}

func GetPdf(c echo.Context) error {
	createby := service.GetTokenData(c, "username")

	err := service.PdfBooksList(c, createby)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}