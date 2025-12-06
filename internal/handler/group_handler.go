package handler

import (
	"context"
	"fms/internal/model"
	"fms/internal/repo"
	"fms/internal/request"

	"github.com/google/uuid"
)

type GroupHandlerRepo interface {
	// Define necessary methods here
	CreateGroup(ctx context.Context, groupRequest *request.GroupCreateRequest) (string, error)
	AddGroupMember(ctx context.Context, memberRequest *request.GroupMemberAddRequest) error
}

type GroupHandler struct {
	// Define necessary repositories here
	GroupRepo repo.GroupRepoInt
	UserRepo  repo.UserRepoInt
}

func NewGroupHandler(groupRepo repo.GroupRepoInt) GroupHandlerRepo {
	return &GroupHandler{
		GroupRepo: groupRepo,
	}
}

func (gh *GroupHandler) CreateGroup(ctx context.Context, groupRequest *request.GroupCreateRequest) (string, error) { //basically we need to read userId from context from fontend
	groupId := uuid.New().String()
	groupId, err := gh.GroupRepo.CreateGroup(ctx, &model.GroupEntity{
		GroupID:     groupId,
		GroupName:   groupRequest.GroupName,
		Description: groupRequest.Description,
		CreatedBy:   groupRequest.CreatedBy,
	})

	if err != nil {
		return "", err
	}

	err = gh.GroupRepo.AddGroupMember(ctx, &model.GroupMemberEntity{
		GroupMemberID: groupRequest.CreatedBy,
		GroupID:       groupId,
		Role:          "ADMIN",
		Status:        "ACTIVE",
	})

	if err != nil {
		return "", err
	}

	return groupId, err
}

func (gh *GroupHandler) AddGroupMember(ctx context.Context, memberRequest *request.GroupMemberAddRequest) error {
	return gh.GroupRepo.AddGroupMember(ctx, &model.GroupMemberEntity{
		GroupMemberID: memberRequest.GroupMemberID,
		GroupID:       memberRequest.GroupID,
		Role:          "MEMBER",
		Status:        "ACTIVE",
	})
}
