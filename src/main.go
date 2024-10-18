package main

import (
	"config"
	"controller"
	"log"
	"repository"
	"router"
	"service"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize MongoDB client
	client, err := config.NewMongoClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(nil)

	// Initialize repository, service, controller
	db := client.Database(cfg.Database)
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Setup and run the Gin router
	r := router.SetupRouter(userController)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
