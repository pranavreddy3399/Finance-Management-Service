package request

import (
	"fms/internal/model/enums"
)

type UserCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phno  string `json:"phno"`
}

type ExpenseCreateRequest struct {
	UserID   string                `json:"user_id"`
	GroupID  string                `json:"group_id"`
	Amount   float64               `json:"amount"`
	Note     string                `json:"note"`
	Title    string                `json:"title"`
	FriendID string                `json:"friend_id"`
	Category enums.ExpenseCategory `json:"category"`
}

//Group Requests

type GroupCreateRequest struct {
	GroupName   string `json:"group_name"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
}

type GroupMemberAddRequest struct {
	GroupID       string `json:"group_id"`
	GroupMemberID string `json:"group_member_id"`
}
