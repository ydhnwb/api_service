package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/api_service/controller/auth"
	"github.com/ydhnwb/api_service/controller/backoffice"
	"github.com/ydhnwb/api_service/controller/health"
	"github.com/ydhnwb/api_service/database"
	authrepo "github.com/ydhnwb/api_service/repository/auth_repo"
	authservice "github.com/ydhnwb/api_service/service/auth_service"
)

var (
	db                   *sql.DB                         = database.SetupDatabaseConnection()
	authRepo             authrepo.AuthRepository         = authrepo.NewAuthRepository(db)
	authService          authservice.AuthService         = authservice.NewAuthService(authRepo)
	authController       auth.AuthController             = auth.NewAuthController(authService)
	backofficeController backoffice.BackofficeController = backoffice.NewBackofficeController()
)

func main() {
	setupGinServer()
}

func setupGinServer() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{

		v1.GET("/health", health.HealthCheck)
		v1.POST("/login", authController.Login)

		// Backoffice endpoint
		backoffice := v1.Group("/backoffice")
		{
			backoffice.POST("/login", backofficeController.Login)
		}
	}

	r.Run(":8080")
}
