package user

import (
	"go-api/internal/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *UserService
}

func CreateUserController(s *UserService) *UserController {
	return &UserController{userService: s}
}

func (ctrl *UserController) FindUserById(c *gin.Context) {
	userID := c.Param("id")
	userData, err := ctrl.userService.FindUserById(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	utils.OKResponse(c, userData)
}
