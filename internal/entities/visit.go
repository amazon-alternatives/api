package entities

import (
	"time"
)

// Visit entity describes a page view of a user
type Visit struct {
	ID        int       `json:"id"`
	Asin      string    `json:"asin"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
}

// PostVisitRequest describes the validation rules of a new visit
type PostVisitRequest struct {
	Asin    string `json:"asin" validate:"required"`
	Country string `json:"country" validate:"required"`
}
