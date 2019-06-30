package models

import "time"

// Food represent food table
type Food struct {
	ID         int64     `json:"id"`
	ProvinceID int64     `json:"province_id"`
	Title      string    `json:"title"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
