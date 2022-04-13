package model

type Books struct {
	No       int    `json:"no"`
	Book     string `json:"book"`
	Author   string `json:"author"`
	CreateBy string `json:"create_by"`
}

type BooksList []Books

type Pagination struct {
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	Sort       string `json:"sort"`
	Search     string `json:"search"`
	TotalRows  int    `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
}
