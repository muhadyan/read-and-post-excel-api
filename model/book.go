package model

type Books struct {
	No     int    `json:"no"`
	Book   string `json:"book"`
	Author string `json:"author"`
}

type BooksList []Books

type Pagination struct {
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	Sort       string `json:"sort"`
	TotalRows  int    `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
}
