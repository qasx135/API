package main

import (
	"log"
	"music-store/config"
	"music-store/db"
	//_ "music-store/docs"
	"music-store/handlers"
	"music-store/models"
	"music-store/routes"
)

func main() {
	cfg := config.ConfigNew()
	database := db.ConnectDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	if err := models.Migrate(database); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	handler := &handlers.Handler{DB: database}
	r := routes.SetupRoutes(handler)

	log.Printf("Server running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
