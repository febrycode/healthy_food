package models

import "time"

// Food represent food table
type Food struct {
	ID         int64     `json:"id"            db:"id"`
	ProvinceID int64     `json:"province_id"   db:"province_id"`
	Title      string    `json:"title"         db:"title"`
	CreatedAt  time.Time `json:"created_at"    db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"    db:"updated_at"`
}

type FoodRequest struct {
	ID           int64    `json:"id"`
	ProvinceID   int64    `json:"province_id"`
	Title        string   `json:"title"`
	Benefit      []string `json:"benefit"`
	Disadvantage []string `json:"disadvantage"`
}
