package models

import "time"

type User struct {
	UserID    int
	Password  string
	Role      string
	CreatedAT time.Time
	DeletedAT time.Duration
}
