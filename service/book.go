package service

import (
	"excel-read/db"
	"excel-read/model"
	"excel-read/pdf"
	"excel-read/repository"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func ReadFile(c echo.Context) (multipart.File, error) {
	file, err := c.FormFile("excel")
	if err != nil {
		log.Println("ReadFile FormFile", err)
		return nil, c.String(http.StatusBadRequest, "Parameter Not Found")
	}

	src, err := file.Open()
	if err != nil {
		log.Println("ReadFile FileOpen", err)
		return nil, c.String(http.StatusBadRequest, err.Error())
	}
	defer src.Close()
	return src, nil
}

func InputAndValidate(c echo.Context, xlsx *excelize.File, createby string) error {
	db := db.DbManager()
	res := model.BooksList{}
	getCell := func(cell string) string {
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
				log.Println("InputAndValidate NoAtoi", err)
				return c.String(http.StatusBadRequest, "Column 'No' must be a number")
			}

			if _, err := strconv.Atoi(author); err == nil {
				log.Println("InputAndValidate AuthorAtoi", err)
				return c.String(http.StatusBadRequest, "Column 'Author' must be a name")
			}

			res = append(res, model.Books{
				No:       no,
				Book:     bookname,
				Author:   author,
				CreateBy: createby,
			})
		}
		for _, b := range res {
			book := model.Books{
				No:       b.No,
				Book:     b.Book,
				Author:   b.Author,
				CreateBy: b.CreateBy,
			}

			db.Create(book)
		}

		return c.String(http.StatusOK, "Success")
	}
}

func GeneratePaginationFromRequest(c echo.Context, createby interface{}) (*model.BooksList, *model.Pagination, error) {
	// Initializing default
	//	var mode string
	limit := 10
	page := 1
	sort := "no"
	search := ""
	query := c.Request().URL.Query()
	var err error

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, err = strconv.Atoi(queryValue)
			if err != nil {
				log.Println("GeneratePaginationFromRequest LimitAtoi", err)
				return nil, nil, c.String(http.StatusBadRequest, "Limit must be a number")
			}
			if limit == 0 {
				limit = 10
			}
			break
		case "page":
			page, err = strconv.Atoi(queryValue)
			if err != nil {
				log.Println("GeneratePaginationFromRequest PageAtoi", err)
				return nil, nil, c.String(http.StatusBadRequest, "Page must be a number")
			}
			if page == 0 {
				page = 1
			}
			break
		case "sort":
			sort = queryValue
			break
		case "search":
			search = "%" + queryValue + "%"
			break
		}
	}

	pagination := model.Pagination{
		Limit:  limit,
		Page:   page,
		Sort:   sort,
		Search: search,
	}

	bookLists, err := repository.GetAllBooks(&pagination, createby)
	if err != nil {
		return nil, nil, c.String(http.StatusBadRequest, err.Error())
	}

	totRowsAndPages, err := repository.GetTotalRowsAndPages(&pagination, createby)
	if err != nil {
		return nil, nil, err
	}

	return bookLists, totRowsAndPages, nil
}

func PdfBooksList(c echo.Context, createby interface{}) error {
	db := db.DbManager()
	book := model.BooksList{}

	db.Where("create_by = ?", createby).Find(&book)

	buf, err := pdf.ParseTemplate(c, book)
	if err != nil {
		return err
	}

	pdfg, err := pdf.GeneratePdf(c, buf)
	if err != nil {
		return err
	}

	c.Response().Header().Set("Content-Disposition", "attachment; filename=books.pdf")
	c.Response().Header().Set("Content-Type", "application/pdf")
	c.Response().WriteHeader(http.StatusOK)
	c.Response().Write(pdfg.Bytes())
	return nil
}

func GetTokenData(c echo.Context, key string) interface{} {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	res := claims[key]
	return res
}
