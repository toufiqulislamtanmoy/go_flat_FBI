package models

import "time"

type User struct {
	ID         uint `gorm:"primaryKey"`
	Username   string
	Fullname   string
	Password   string `gorm:"column:password"`
	Email      string `gorm:"column:email_address"`
	RoleID     uint
	Role       Role `gorm:"foreignKey:RoleID"`
	IsDisabled bool
	CreatedAt  time.Time
}
