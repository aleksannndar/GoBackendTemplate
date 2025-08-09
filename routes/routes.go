package routes

import (
	"GoBackendTemplate/dependency"
	"GoBackendTemplate/domains/auth"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, container *dependency.Container) {
	authGroup := r.Group("auth")
	authGroup.POST("/login", auth.LoginHandler(container.AuthService))
	authGroup.POST("/register", auth.RegisterHandler(container.AuthService))
}
