package models

import "time"

type Book struct {
	ID            int        `json:"id" form:"id" query:"id"`
	Title         string     `json:"title" form:"title" query:"title"`
	Author        string     `json:"author" form:"author" query:"author"`
	PublishedDate string     `json:"published_date" form:"published_date" query:"published_date"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}
