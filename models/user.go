package models

import "time"

// User represent the user table
type User struct {
	ID        int64     `json:"id" form:"id"`
	Email     string    `json:"email" form:"email"`
	Name      string    `json:"name" form:"name"`
	AvatarURL string    `json:"avatar_url" form:"avatar_url"`
	Address   string    `json:"address" form:"address"`
	Bio       string    `json:"bio" form:"bio"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}
