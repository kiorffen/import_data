package main

import (
	"fmt"
	"log"
	"import_data/config"
	"import_data/database"
	"import_data/handlers"
	"import_data/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	if err := database.Initialize(cfg); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer database.Close()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware(cfg.Server.AuthToken))
	{
		api.POST("/articles", handlers.CreateArticle)
	}

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Server starting on port %d", cfg.Server.Port)
	if err := r.Run(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}