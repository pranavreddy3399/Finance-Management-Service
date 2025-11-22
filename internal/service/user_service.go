package service

import (
	"encoding/json"
	userhandler "fms/internal/handler"
	"fms/internal/request"
	"fmt"
	"net/http"
)

type Service struct {
	// Add necessary fields here, e.g., repositories, configurations
	UserHandler userhandler.UserHandlerRepo
}

func NewService(userHandler *userhandler.UserHandlerRepo) *Service {
	return &Service{
		// Initialize fields here
		UserHandler: *userHandler,
	}
}

func (s *Service) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("GetUserById called")

	// 1. Read `id` query param
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "id query param is required", http.StatusBadRequest)
		return
	}

	// 2. Call repo/service to fetch user
	user, err := s.UserHandler.GetUserById(ctx, userID)
	if err != nil {
		http.Error(w, "failed to fetch user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. If not found return 404
	if user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	// 4. Return user as JSON
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "failed to encode user", http.StatusInternalServerError)
		return
	}
}

func (s *Service) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("CreateUser called")

	// 1. Parse request body
	var userReq request.UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	id, err := s.UserHandler.CreateUser(ctx, &userReq)
	if err != nil {
		http.Error(w, "failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Return user as JSON
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(id); err != nil {
		http.Error(w, "failed to encode user", http.StatusInternalServerError)
		return
	}
}
