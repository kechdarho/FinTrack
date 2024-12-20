package models

import "time"

type User struct {
	UserID    uint          `json:"user_id"`
	Password  string        `json:"password"`
	Role      string        `json:"role"`
	CreatedAT time.Time     `json:"created_at"`
	DeletedAT time.Duration `json:"deleted_at"`
}
