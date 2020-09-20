package entities

import (
	"time"
)

type Product struct {
	ID         int       `json:"id"`
	CategoryID int       `json:"category_id"`
	Asin       string    `json:"asin"`
	Gtin       string    `json:"gtin"`
	Title      string    `json:"title"`
	Url        string    `json:"url"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

type PostProductRequest struct {
	Category string  `json:"category" validate:"required"`
	Asin     string  `json:"asin" validate:"required"`
	Gtin     *string `json:"gtin"`
	Title    string  `json:"title" validate:"required"`
	Url      string  `json:"url" validate:"required"`
}
