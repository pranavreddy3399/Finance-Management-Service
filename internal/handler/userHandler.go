package userhandler

import (
	"context"
	"fms/internal/model"
	"fms/internal/repo"
	"fms/internal/request"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserHandlerRepo interface {
	GetUserById(ctx context.Context, userId string) (*model.User, error)
	CreateUser(ctx context.Context, userReq *request.UserCreateRequest) (string, error)
}

type UserHandler struct {
	UserRepo repo.UserRepoInt
}

func NewUserHandler(userRepo repo.UserRepoInt) UserHandlerRepo {
	return &UserHandler{
		UserRepo: userRepo,
	}
}

func (uh *UserHandler) GetUserById(ctx context.Context, userId string) (*model.User, error) {
	return uh.UserRepo.GetUserByID(ctx, userId)
}

func (uh *UserHandler) CreateUser(ctx context.Context, userReq *request.UserCreateRequest) (string, error) {
	userID := uuid.New().String()
	user := &model.User{
		ID:        userID,
		Name:      userReq.Name,
		Email:     userReq.Email,
		Phno:      userReq.Phno,
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := uh.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}
	fmt.Printf("User created with ID: %s\n", userID)
	return userID, nil
}
