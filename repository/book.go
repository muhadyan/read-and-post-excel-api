package repository

import (
	"excel-read/db"
	"excel-read/model"
	"log"
	"math"
)

func GetAllBooks(pagination *model.Pagination, createby interface{}) (*model.BooksList, error) {
	db := db.DbManager()
	var books model.BooksList
	offset := (pagination.Page - 1) * pagination.Limit

	queryPagination := db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	queryResult := queryPagination.Model(&model.BooksList{}).Where("create_by = ?", createby)

	if pagination.Search != "" {
		queryResult = queryResult.Find(&books, "book LIKE ? OR author LIKE ?", pagination.Search, pagination.Search)
	} else {
		queryResult = queryResult.Find(&books)
	}

	if queryResult.Error != nil {
		err := queryResult.Error
		log.Println("GetAllBooks QueryResult", err)
		return nil, err
	}
	return &books, nil
}

func GetTotalRowsAndPages(pagination *model.Pagination, createby interface{}) (*model.Pagination, error) {
	db := db.DbManager()
	var totalRows int
	queryResult := db.Model(&model.BooksList{}).Where("create_by = ?", createby)

	if pagination.Search != "" {
		queryResult = queryResult.
			Where("book LIKE ? OR author LIKE ?", pagination.Search, pagination.Search).
			Count(&totalRows)
	} else {
		queryResult = queryResult.Count(&totalRows)
	}

	if queryResult.Error != nil {
		err := queryResult.Error
		log.Println("GetTotalRowsAndPages QueryResult", err)
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

	rowsAndPages := model.Pagination{
		TotalRows:  totalRows,
		TotalPages: totalPages,
	}

	return &rowsAndPages, nil
}
