package server

import (
	"fms/internal/service"
	"fmt"

	"net/http"
)

type Server struct {
	Service *service.Service
}

func NewServer(service *service.Service) *Server {
	return &Server{
		Service: service,
	}
}

func (ser *Server) RegisterServerAndRoutes() error {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	//user routes
	http.HandleFunc("/user/", ser.Service.GetUserById)
	http.HandleFunc("/create-user/", ser.Service.CreateUser)

	//expense routes
	http.HandleFunc("/add-expense/", ser.Service.AddExpense)

	fmt.Println("connecting to port 9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("server not conneted to 9090")
		return err
	}
	return nil
}
