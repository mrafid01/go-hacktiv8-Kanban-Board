package entity

import "time"

type Category struct {
	ID        int `gorm:"primarykey"`
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
