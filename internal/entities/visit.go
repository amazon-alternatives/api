package entities

import (
	"time"
)

type Visit struct {
	ID        int       `json:"id"`
	Asin      string    `json:"asin"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
}

type PostVisitRequest struct {
	Asin    string `json:"asin" validate:"required"`
	Country string `json:"country" validate:"required"`
}
