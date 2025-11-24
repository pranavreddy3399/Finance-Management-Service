package main

import (
	"fms/internal/db"
	"fms/internal/handler"
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
	expRepo := repo.NewExpenseRepo(dbConn)
	userHandler := handler.NewUserHandler(userRepo)
	expHandler := handler.NewExpenseHandler(&expRepo, &userRepo)
	service := service.NewService(&userHandler, &expHandler)
	server := server.NewServer(service)
	err = server.RegisterServerAndRoutes()
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Database connection established:", dbConn)
}
