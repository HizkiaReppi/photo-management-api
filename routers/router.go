package routers

import (
	"github.com/gin-gonic/gin"
	"rest-api/controllers"
	"rest-api/databases"
	"rest-api/middlewares"
)

// RouteInit menginisialisasi dan mengonfigurasi router Gin.
func RouteInit() *gin.Engine {
	route := gin.Default()

	// Get the database instance.
	db := database.GetDB()

	// Initialize controllers.
	userController := controllers.NewUserController(db)
	photoController := controllers.NewPhotoController(db)

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

	// Serve static images from the "static/images" directory.
	route.Static("/images", "./static/images")

	// Photo routes with authentication middleware.
	photoRoute := api.Group("/photo").Use(middlewares.AuthMiddleware(db))
	{
		photoRoute.GET("/", photoController.Get)
		photoRoute.POST("/", photoController.Create)
		photoRoute.PUT("/", photoController.Update)
		photoRoute.DELETE("/", photoController.Delete)
	}

	return route
}
