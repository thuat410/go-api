package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(router *gin.RouterGroup, dbPool *pgxpool.Pool) {
	userService := CreateUserService()
	ctrl := CreateUserController(userService)
	userGroup := router.Group("/users")
	{
		userGroup.GET("/:id", ctrl.FindUserById)
	}
}
