package service

import (
	"excel-read/db"
	"excel-read/model"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/echo"
)

func ReadFile(c echo.Context) (multipart.File, error) {
	file, err := c.FormFile("excel")
	if err != nil {
		log.Println("FormFile WrongQueryParam", err)
		return nil, c.String(http.StatusBadRequest, "Parameter Not Found")
	}

	src, err := file.Open()
	if err != nil {
		log.Println("Open CannotOpenFile", err)
		return nil, c.String(http.StatusBadRequest, err.Error())
	}
	defer src.Close()
	return src, nil
}

func InputAndValidate(c echo.Context, xlsx *excelize.File) error {
	db := db.DbManager()
	res := model.BooksList{}
	getCell := func (cell string) string {
		return xlsx.GetCellValue("Sheet1", cell)
	}

	if getCell("B1") != "Books" || getCell("C1") != "Author" {
		return c.String(http.StatusBadRequest, "Wrong Column Name Input")
	} else {
		for i := 2; i > 1; i++ {
			number := getCell(fmt.Sprintf("A%d", i))
			bookname := getCell(fmt.Sprintf("B%d", i))
			author := getCell(fmt.Sprintf("C%d", i))

			if number == "" && bookname == "" && author == "" {
				break
			}

			if number == "" || bookname == "" || author == "" {
				return c.String(http.StatusBadRequest, "Data cannot empty")
			}

			no, err := strconv.Atoi(number)
			if err != nil {
				log.Println("Atoi CannotConvertToNum", err)
				return c.String(http.StatusBadRequest, "Column 'No' must be a number")
			}

			if _, err := strconv.Atoi(author); err == nil {
				return c.String(http.StatusBadRequest, "Column 'Author' must be a name")
			}

			res = append(res, model.Books{
				No:     no,
				Book:   bookname,
				Author: author,
			})
		}
		for _, b := range res {
			book := model.Books{
				No:     b.No,
				Book:   b.Book,
				Author: b.Author}

			db.Create(book)
		}

		return c.String(http.StatusOK, "Success")
	}
}
