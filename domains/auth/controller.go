package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(authService IAuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginRequest LoginRequest
		if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		response, err := authService.LoginWithPassword(loginRequest.Email, loginRequest.Password)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"data":    response,
		})
	}
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func RegisterHandler(authService IAuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var registerRequest RegisterRequest
		if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		response, err := authService.RegisterWithPassword(registerRequest.Email, registerRequest.Password)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Register successful",
			"data":    response,
		})
	}
}
