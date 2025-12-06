package model

import "time"

type GroupEntity struct {
	GroupID     string    `json:"group_id" db:"group_id"`
	GroupName   string    `json:"group_name" db:"group_name"`
	Description string    `json:"description" db:"description"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type GroupMemberEntity struct {
	GroupMemberID string    `json:"group_member_id" db:"group_member_id"`
	GroupID       string    `json:"group_id" db:"group_id"`
	Role          string    `json:"role" db:"role"`
	Status        string    `json:"status" db:"status"`
	JoinedAt      time.Time `json:"joined_at" db:"joined_at"`
}
