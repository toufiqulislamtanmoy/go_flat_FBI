package models

import "time"

type Plan struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Date      time.Time
	UserID    uint
	User      User `gorm:"foreignKey:UserID"`
	Status    string
	CreatedAt time.Time
}
