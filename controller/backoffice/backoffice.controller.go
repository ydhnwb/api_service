package backoffice

import "github.com/gin-gonic/gin"

type BackofficeController interface {
	Login(c *gin.Context)
}

type backofficeControllerDependencies struct {
}

func NewBackofficeController() BackofficeController {
	return &backofficeControllerDependencies{}
}

func (backoffice *backofficeControllerDependencies) Login(c *gin.Context) {}
