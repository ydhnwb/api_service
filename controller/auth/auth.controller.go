package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/api_service/common"
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
	lang := c.DefaultQuery("lang", "en")
	var loginRequest dto.LoginRequest
	err := c.ShouldBind(&loginRequest)
	if err != nil {
		res := common.BuildErrorResponse(err, true, lang)
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// dummy check
	if loginRequest.Email != "johndoe@mail.com" || loginRequest.Password != "yudhanewbie" {
		res := common.BuildErrorResponse(errors.New("auth_wrong_credentials"), true, lang)
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    1,
		"name":  "John Doe",
		"email": "johndoe@mail.com",
	})
}
