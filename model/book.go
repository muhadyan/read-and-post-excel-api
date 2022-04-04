package model

type Books struct {
	No     int    `json:"no"`
	Book   string `json:"books"`
	Author string `json:"authors"`
}

type BooksList []Books