package service

import (
	"encoding/json"
	"fms/internal/request"
	"fmt"
	"net/http"
)

func (s *Service) CreateGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("CreateGroup called")

	// 1. Parse request body
	var groupReq request.GroupCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&groupReq); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	groupID, err := s.GroupHandler.CreateGroup(ctx, &groupReq)
	if err != nil {
		http.Error(w, "failed to create group: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Return group ID as JSON
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(groupID); err != nil {
		http.Error(w, "failed to encode group ID", http.StatusInternalServerError)
		return
	}
}

func (s *Service) AddMembersToGroup(w http.ResponseWriter, r *http.Request) {
	// Implementation for adding members to a group
	ctx := r.Context()
	fmt.Println("AddMembersToGroup called")

	// 1. Parse request body
	var memberReq request.GroupMemberAddRequest
	if err := json.NewDecoder(r.Body).Decode(&memberReq); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	err := s.GroupHandler.AddGroupMember(ctx, &memberReq)
	if err != nil {
		http.Error(w, "failed to add member to group: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Member added to group successfully"))
}
