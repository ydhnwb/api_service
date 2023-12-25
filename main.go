package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/api_service/controller/auth"
)

var (
	authController auth.AuthController = auth.NewAuthController()
)

func main() {
	setupGinServer()

}

func setupGinServer() {
	r := gin.Default()
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API is running..",
		})
	})

	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", authController.Login)
	}

	r.Run()
}
