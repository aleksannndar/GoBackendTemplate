package main

import (
	"GoBackendTemplate/database/migrate"
	"GoBackendTemplate/dependency"
	"GoBackendTemplate/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env: %v", err)
	}

	err = migrate.RunMigrations()
	if err != nil {
		log.Fatalf("Failed to run migration: %v", err)
	}

	r := gin.Default()
	container := dependency.BuildContainer()
	routes.SetupRoutes(r, container)

	if err := r.Run(":8088"); err != nil {
		log.Fatalf("Failed to start GoBackendTemplate: %v", err)
	}
}
