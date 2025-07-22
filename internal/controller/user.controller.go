package controller

import (
	"github.com/3xhvy/go-backend/internal/service"
	"github.com/3xhvy/go-backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	// response.SuccessResponse(c, 20001, []string{"hvy", "tuongby"})
	response.ErrorResponse(c, 20003, "unkown error")
}
