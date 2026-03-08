package models

type PlanMember struct {
	ID           uint `gorm:"primaryKey"`
	PlanID       uint
	Plan         Plan `gorm:"foreignKey:PlanID"`
	UserID       uint
	User         User `gorm:"foreignKey:UserID"`
	PermissionID uint
	Permission   Permission `gorm:"foreignKey:PermissionID"`
}
