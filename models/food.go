package models

import "time"

// Food represent food table
type Food struct {
	ID         int64     `json:"id"            db:"id"`
	UserID     int64     `json:"user_id"       db:"user_id"`
	ProvinceID int64     `json:"province_id"   db:"province_id"`
	Title      string    `json:"title"         db:"title"`
	CreatedAt  time.Time `json:"created_at"    db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"    db:"updated_at"`
}

type FoodRequest struct {
	ID            int64    `json:"id"`
	UserID        int64    `json:"user_id"`
	ProvinceID    int64    `json:"province_id"`
	Title         string   `json:"title"`
	Benefits      []string `json:"benefits"`
	Disadvantages []string `json:"disadvantages"`
	Images        []string `json:"images"`
}

type FoodResponse struct {
	Food
	ProvinceName  string         `json:"province_name"`
	Benefits      []Benefit      `json:"benefits"`
	Disadvantages []Disadvantage `json:"disadvantages"`
	Images        []Image        `json:"images"`
}

type Benefit struct {
	ID            int64  `json:"id"`
	ReferenceType int    `json:"reference_type"`
	ReferenceID   int64  `json:"reference_id"`
	Description   string `json:"description"`
}

type Disadvantage struct {
	ID            int64  `json:"id"`
	ReferenceType int    `json:"reference_type"`
	ReferenceID   int64  `json:"reference_id"`
	Description   string `json:"description"`
}
