package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/api_service/dto"
)

type AuthController interface {
	Login(c *gin.Context)
}

type authControllerDependencies struct {
}

func NewAuthController() AuthController {
	return &authControllerDependencies{}
}

func (auth *authControllerDependencies) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest
	err := c.ShouldBind(&loginRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// dummy check
	if loginRequest.Email != "yudhanewbie@mail.com" || loginRequest.Password != "yudhanewbie" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Wrong credentials",
		})
		return
	}

	fmt.Println("Hello")
}
