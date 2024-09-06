package routes

import (
	"gin_api_rest/controllers"
	"gin_api_rest/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	userService := services.NewUserService() // Supondo que você tenha um construtor para UserService
	userController := controllers.NewUserController(userService)

	users := router.Group("/users")
	{
		users.GET("", userController.GetUsers)
		users.GET("/:id", userController.GetUserByID)
		users.POST("", userController.CreateUser)
		users.PUT("", userController.UpdateUser)
		users.DELETE("/:id", userController.DeleteUser)
	}
}
