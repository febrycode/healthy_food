package models

import "time"

// FoodDetail represent food_detail table
type FoodDetail struct {
	ID            int64     `json:"id"`
	ReferenceType int       `json:"reference_type"`
	ReferenceID   int64     `json:"reference_id"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
