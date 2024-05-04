package models

import "database/sql"

type CategoryResponse struct {
	ID           int    `json : "ID"`
	CategoryName string `json : "category_name"`
}

type CategoryRequest struct {
	ID           int64  `json : "ID"`
	CategoryName sql.NullString `json : "category_name"`
}
