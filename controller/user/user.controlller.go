package user_controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Register(c *gin.Context)
}

type userControllerDependencies struct {
}

func NewUserController() UserController {
	return &userControllerDependencies{}
}

func (userCtl *userControllerDependencies) Register(c *gin.Context) {

}
