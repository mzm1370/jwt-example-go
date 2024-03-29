package main

import (
	"jwt-example-go/config"
	"jwt-example-go/db"
	"jwt-example-go/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.NewConfig()

	_, err := db.Initialize(cfg)

	if err != nil {
		log.Fatalf("Failed top initialize database: %v", err)

	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}()

	r := gin.Default()

	handlers.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
