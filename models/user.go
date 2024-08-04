package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        int    `gorm:"primaryKey"`
	Email     string `gorm:"unique;not null"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
