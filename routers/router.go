package routers

import (
	"github.com/gin-gonic/gin"
	"rest-api/controllers"
	"rest-api/databases"
)

// RouteInit menginisialisasi dan mengonfigurasi router Gin.
func RouteInit() *gin.Engine {
	route := gin.Default()

	// Get the database instance.
	db := database.GetDB()

	// Initialize controllers.
	userController := controllers.NewUserController(db)

	// Create API groups.
	api := route.Group("/api/v1")

	// User routes.
	userRoute := api.Group("/users")
	{
		userRoute.POST("/register", userController.Register)
		userRoute.POST("/login", userController.Login)
		userRoute.PUT("/:userId", userController.Update)
		userRoute.DELETE("/:userId", userController.Delete)
	}

	return route
}
