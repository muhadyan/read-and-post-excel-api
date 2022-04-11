package pdf

import (
	"bytes"
	"excel-read/model"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/labstack/echo"
)

func ParseTemplate(c echo.Context, data model.BooksList) (*bytes.Buffer, error) {
	t, err := template.ParseFiles("./templates/books.html")
	if err != nil {
		return nil, err
	}

	tmpl := template.Must(t, err)

	buf := new(bytes.Buffer)

	err = tmpl.ExecuteTemplate(buf, "books.html", data)
	if err != nil {
		log.Println("GeneratePdf ExecuteTemplate")
		return nil, c.String(http.StatusBadRequest, "Cannot Execute Template")
	}

	return buf, nil
}

func GeneratePdf(c echo.Context, buf *bytes.Buffer) (*wkhtmltopdf.PDFGenerator, error) {
	t := time.Now().Unix()
	// write whole the body

	if _, err := os.Stat("cloneTemplate/"); os.IsNotExist(err) {
		errDir := os.Mkdir("cloneTemplate/", 0777)
		if errDir != nil {
			log.Fatal(errDir)
		}
	}

	err1 := ioutil.WriteFile("cloneTemplate/"+strconv.FormatInt(int64(t), 10)+".html", []byte(buf.String()), 0644)
	if err1 != nil {
		return nil, err1
	}

	f, err := os.Open("cloneTemplate/"+strconv.FormatInt(int64(t), 10)+".html")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		log.Println("GeneratePdf OsOpen")
		return nil, err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Println("GeneratePdf PDFGenerator")
		return nil, err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		log.Println("GeneratePdf PdfCreate")
		return nil, c.String(http.StatusConflict, err.Error())
	}

	err = pdfg.WriteFile("./downloads/books.pdf")
	if err != nil {
		log.Println("GeneratePdf PdfWriteFile")
		return nil, c.String(http.StatusConflict, err.Error())
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	defer os.RemoveAll(dir + "/cloneTemplate")

	return pdfg, nil
}