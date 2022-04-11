package api

import (
	"excel-read/service"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func GetBooks(c echo.Context) error {
	bookLists, rowPageList, err := service.GeneratePaginationFromRequest(c)
	if err != nil {
		log.Println("GeneratePaginationFromRequest PaginationRequestError", err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Total Rows":  rowPageList.TotalRows,
		"Total Pages": rowPageList.TotalPages,
		"Data":        bookLists,
	})
}

func GetPdf(c echo.Context) error {
	err := service.PdfBooksList(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}