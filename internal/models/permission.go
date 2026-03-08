package models

type Permission struct {
	ID          uint `gorm:"primaryKey"`
	AccessLevel string
	Priority    int
}
