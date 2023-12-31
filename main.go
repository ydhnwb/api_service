package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/api_service/controller/auth"
	"github.com/ydhnwb/api_service/controller/backoffice"
	"github.com/ydhnwb/api_service/controller/health"
)

var (
	authController       auth.AuthController             = auth.NewAuthController()
	backofficeController backoffice.BackofficeController = backoffice.NewBackofficeController()
)

func main() {
	setupGinServer()

}

func setupGinServer() {
	r := gin.Default()

	r.GET("/api/health", health.HealthCheck)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", authController.Login)
		// Backoffice endpoint
		bo := v1.Group("/backoffice")
		{
			bo.POST("/login", backofficeController.Login)
		}
	}

	r.Run()
}
