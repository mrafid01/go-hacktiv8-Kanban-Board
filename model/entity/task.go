package entity

import "time"

type Task struct {
	ID          int
	Title       string
	Description string
	Status      *bool
	UserID      int
	CategoryID  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User     `gorm:"foreignKey:UserID;Constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Category    Category `gorm:"foreignKey:CategoryID;Constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
