package models

import "time"

// Province represent province table
type Province struct {
	ID        int64     `json:"id"            db:"id"`
	Name      string    `json:"name"          db:"name"`
	CreatedAt time.Time `json:"created_at"    db:"created_at"`
	UpdatedAt time.Time `json:"updated_at"    db:"updated_at"`
}
