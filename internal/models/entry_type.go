package models

type EntryType struct {
	ID         uint `gorm:"primaryKey"`
	EntryTitle string
	Priority   int
}
