package models

import "time"

type Image struct {
	ID            int64     `json:"id"                  db:"id"`
	ReferenceType int       `json:"reference_type"      db:"reference_type"`
	ReferenceID   int64     `json:"reference_id"        db:"reference_id"`
	Name          string    `json:"name"                db:"name"`
	Description   string    `json:"description"         db:"description"`
	CreatedAt     time.Time `json:"created_at"          db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"          db:"updated_at" `
}
