package main

import (
	"fms/internal/db"
	userhandler "fms/internal/handler"
	"fms/internal/repo"
	"fms/internal/server"
	"fms/internal/service"
	"fmt"
)

func main() {
	fmt.Println("Hello, Finance Management Service!")

	dbConn, err := db.InitMySQL()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}

	defer dbConn.Close()
	userRepo := repo.NewUserRepo(dbConn)
	userHandler := userhandler.NewUserHandler(userRepo)
	service := service.NewService(&userHandler)
	server := server.NewServer(service)
	err = server.RegisterServerAndRoutes()
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Database connection established:", dbConn)
}
