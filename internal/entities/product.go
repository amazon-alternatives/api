package entities

import (
	"time"
)

// Product entity describes the different identificators of an ASIN number
type Product struct {
	Asin      string    `json:"asin"`
	Isbn      string    `json:"isbn"`
	Ean       string    `json:"ean"`
	Gtin      string    `json:"gtin"`
	Upc       string    `json:"upc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
