package models

import "time"

type MealEntry struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Amount      float64
	Date        time.Time
	Details     string
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
	PlanID      uint
	Plan        Plan `gorm:"foreignKey:PlanID"`
	EntryTypeID uint
	EntryType   EntryType `gorm:"foreignKey:EntryTypeID"`
	CreatedAt   time.Time
}
