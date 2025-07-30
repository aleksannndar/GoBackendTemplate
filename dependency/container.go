package dependency

import (
	"GoBackendTemplate/database"

	"gorm.io/gorm"
)

type Container struct {
	DB *gorm.DB
}

func BuildContainer() *Container {
	db := database.Init()

	return &Container{
		DB: db,
	}
}
