package router

import (
	"github.com/gin-gonic/gin"
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/handler"
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/repository"
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/service"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// role
	roleRepo := repository.NewRoleRepository(db)
	roleService := service.NewRoleService(roleRepo)
	roleHandler := handler.NewRoleHandler(roleService)
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/signup", userHandler.CreateUser)
		userRoutes.GET("/", userHandler.GetAllUsers)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
		userRoutes.GET("/login", userHandler.Login)
	}
	roleRoutes := r.Group("/roles")
	{
		roleRoutes.GET("/", roleHandler.GetAllRole)
		roleRoutes.POST("/create", roleHandler.CreateRole)
		roleRoutes.GET("/:id", roleHandler.GetRoleByID)
		roleRoutes.PUT("/:id", roleHandler.UpdateRole)
	}

	return r
}
