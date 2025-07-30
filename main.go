package main

import (
	"GoBackendTemplate/dependency"
	"GoBackendTemplate/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	container := dependency.BuildContainer()

	routes.SetupRoutes(r, container)

	if err := r.Run(":8088"); err != nil {
		log.Fatalf("Failed to start GoBackendTemplate: %v", err)
	}
}
