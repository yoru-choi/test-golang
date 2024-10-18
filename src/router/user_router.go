package router

import (
	"test-golang/src/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin router and sets up routes
func SetupRouter(userController *controller.UserController) *gin.Engine {
	r := gin.Default()
	usersGroup := r.Group("/users") // /users 경로로 그룹화

	{
		usersGroup.GET("", userController.GetUsersHandler)
		usersGroup.GET("/:id", userController.GetUserHandler)
		usersGroup.POST("", userController.CreateUserHandler)         // 빈 문자열로 POST를 처리
		usersGroup.PUT("/:id", userController.UpdateUserHandler)
		usersGroup.DELETE("/:id", userController.DeleteUserHandler)
	}

	return r // *gin.Engine 반환
}
