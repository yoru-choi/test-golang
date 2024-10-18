package router

import (
	"project/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin router with routes
func SetupRouter(userController *controller.UserController) *gin.Engine {
	router := gin.Default()

	usersGroup := router.Group("/users")
	{
		usersGroup.GET("/:id", userController.GetUserHandler)
		usersGroup.POST("", userController.CreateUserHandler)
		usersGroup.PUT("/:id", userController.UpdateUserHandler)
		usersGroup.DELETE("/:id", userController.DeleteUserHandler)
	}
	return router
}
