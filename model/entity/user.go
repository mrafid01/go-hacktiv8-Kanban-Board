package entity

import "time"

type User struct {
	ID        int `gorm:"primarykey"`
	FullName  string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
