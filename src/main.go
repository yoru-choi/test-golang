package main

import (
	"context"
	"fmt"
	"log"
	"runtime"

	"test-golang/src/config" // 수정된 경로
	"test-golang/src/controller"
	"test-golang/src/repository"
	"test-golang/src/router"
	"test-golang/src/service"
)
func printGoASCIIArt() {
	goVersion := runtime.Version() 
	asciiArt := `
	╔═╗╔═╗  ╦  ╔═╗╔╗╔╔═╗
	║ ╦║ ║  ║  ╠═╣║║║║ ╦  
	╚═╝╚═╝  ╩═╝╩ ╩╝╚╝╚═╝  %s
  `
	  fmt.Printf(asciiArt, goVersion)
}
func main() {

	// Print ASCII art
	printGoASCIIArt()
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize MongoDB client
	client, err := config.NewMongoClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO()) 

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


