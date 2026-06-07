package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	userService := CreateUserService()
	ctrl := CreateUserController(userService)
	userGroup := router.Group("/users")
	{
		userGroup.GET("/:id", ctrl.FindUserById)
	}
}
