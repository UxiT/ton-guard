package main

import (
	"decard/internal/interfaces/http/handlers"
	"log"
	"net/http"

	"decard/config"
	"decard/internal/domain/service"
	"decard/internal/infrastructure/database"
	orm_repositoty "decard/internal/infrastructure/orm/repositoty"
	"decard/internal/interfaces/http/routes"
)

func main() {
	// Load configuration
	cfg := config.NewConfig()

	// Initialize database
	db := database.NewDatabase(cfg)
	defer db.Close()

	// Initialize repositories
	ProfileRepository := orm_repositoty.NewProfileRepository(db.DB)

	// Initialize services
	authService := service.NewAuthService(ProfileRepository)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(*authService)

	// Setup router
	router := routes.NewRouter(authHandler)

	// Start server
	log.Printf("Server starting on %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
