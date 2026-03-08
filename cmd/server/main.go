package main

import (
	"github.com/toufiqulislamtanmoy/go_flat_FBI/config"
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/models"
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/router"
)

func main() {
	db := config.ConnectDatabase()

	db.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Permission{},
		&models.Plan{},
		&models.PlanMember{},
		&models.EntryType{},
		&models.MealEntry{},
	)
	r := router.SetupRouter(db)
	r.Run(":5000")
}
