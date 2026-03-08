package models

type Role struct {
	ID        uint   `gorm:"primaryKey"`
	RoleTitle string `gorm:"unique;not null"`
	Priority  int
}
