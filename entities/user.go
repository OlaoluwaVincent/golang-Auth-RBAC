package entities

import "time"

type User struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	Name         string    `gorm:"size:255" json:"name"`
	Email        string    `gorm:"size:255;unique_index" json:"email"`
	PasswordHash string    `gorm:"size:255" json:"-"`
	Role         string    `gorm:"size:50" json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
