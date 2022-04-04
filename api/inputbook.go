package api

import (
	"excel-read/service"
	"log"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/echo"
)

func InputBooks(c echo.Context) error {
	src, err := service.ReadFile(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	xlsx, err := excelize.OpenReader(src)
	if err != nil {
		log.Println("OpenReader WrongFileTypeInput", err)
		return c.String(http.StatusBadRequest, "Wrong File Type Input")
	}

	err = service.InputAndValidate(c, xlsx)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return nil
}
