package controller

import (
	"context"
	"net/http"
	"project/models"
	"project/service"

	"github.com/gin-gonic/gin"
)

// UserController handles user-related requests
type UserController struct {
	userService service.UserService
}

// NewUserController creates a new UserController
func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) GetUserHandler(ctx *gin.Context) {
	userID := ctx.Param("id")
	user, err := c.userService.GetUserByID(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) CreateUserHandler(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err := c.userService.CreateUser(context.Background(), &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) UpdateUserHandler(ctx *gin.Context) {
	userID := ctx.Param("id")
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err := c.userService.UpdateUser(context.Background(), userID, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) DeleteUserHandler(ctx *gin.Context) {
	userID := ctx.Param("id")
	err := c.userService.DeleteUser(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
