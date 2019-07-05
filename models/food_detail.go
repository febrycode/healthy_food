package models

import "time"

// FoodDetail represent food_detail table
type FoodDetail struct {
	ID            int64     `json:"id"                  db:"id"`
	ReferenceType int       `json:"reference_type"      db:"reference_type"`
	ReferenceID   int64     `json:"reference_id"        db:"reference_id"`
	Description   string    `json:"description"         db:"description"`
	CreatedAt     time.Time `json:"created_at"          db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"          db:"updated_at"`
}
