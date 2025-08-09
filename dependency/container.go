package dependency

import (
	"GoBackendTemplate/database"
	"GoBackendTemplate/domains/auth"
	"GoBackendTemplate/domains/authtoken"

	"gorm.io/gorm"
)

type Container struct {
	DB          *gorm.DB
	JwtService  authtoken.IJWTService
	AuthService auth.IAuthService
}

func BuildContainer() *Container {
	db := database.Init()
	jwtService := authtoken.NewJWTService()
	authService := auth.NewAuthService(db, jwtService)

	return &Container{
		DB:          db,
		JwtService:  jwtService,
		AuthService: authService,
	}
}
