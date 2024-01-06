package user_controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/api_service/common"
	"github.com/ydhnwb/api_service/dto"
)

type UserController interface {
	Register(c *gin.Context)
}

type userControllerDependencies struct {
}

func NewUserController() UserController {
	return &userControllerDependencies{}
}

func (userCtl *userControllerDependencies) Register(c *gin.Context) {
	lang := c.DefaultQuery("lang", "en")
	var registerRequest dto.UserRegister

	err := c.ShouldBind(&registerRequest)
	if err != nil {
		res := common.BuildErrorResponse(err, true, lang)
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	dob, dobErr := time.Parse("2006-01-02", registerRequest.DateOfBirth)
	if dobErr != nil {
		res := common.BuildErrorResponse(errors.New("invalid_date_format"), true, lang)
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	registerRequest.DateOfBirth = dob.String()

	c.JSON(200, gin.H{
		"message": registerRequest,
	})
}
