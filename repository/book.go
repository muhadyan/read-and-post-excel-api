package repository

import (
	"excel-read/db"
	"excel-read/model"
	"log"
	"math"
)

func GetAllBooks(pagination *model.Pagination) (*model.BooksList, error) {
	db := db.DbManager()
	var books model.BooksList
	book := model.BooksList{}
	offset := (pagination.Page - 1) * pagination.Limit

	queryPagination := db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	queryResult := queryPagination.Model(&model.BooksList{})

	if pagination.Search != "" {
		queryResult = queryResult.Where(book).Find(&books, "book LIKE ? OR author LIKE ?", pagination.Search, pagination.Search)
	} else {
		queryResult = queryResult.Where(book).Find(&books)
	}

	if queryResult.Error != nil {
		err := queryResult.Error
		log.Println("QueryResult QueryResultError", err)
		return nil, err
	}
	return &books, nil
}

func GetTotalRowsAndPages(pagination *model.Pagination) *model.Pagination {
	db := db.DbManager()
	var totalRows int

	if pagination.Search != "" {
		db.
			Model(&model.BooksList{}).
			Where("book LIKE ? OR author LIKE ?", pagination.Search, pagination.Search).
			Count(&totalRows)
	} else {
		db.Model(&model.BooksList{}).Count(&totalRows)
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

	rowsAndPages := model.Pagination{
		TotalRows:  totalRows,
		TotalPages: totalPages,
	}

	return &rowsAndPages
}
