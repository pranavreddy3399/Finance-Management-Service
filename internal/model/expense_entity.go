package model

import "time"

type ExpenseEntity struct {
	ExpenseID string    `json:"expense_id"`
	UserID    string    `json:"user_id"`
	GroupID   string    `json:"group_id"`
	Amount    float64   `json:"amount"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	Note      string    `json:"note"`
	Status    string    `json:"status"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ExpenseSplitEntity struct {
	ExpenseSplitID string    `json:"expense_split_id"`
	GroupID        string    `json:"group_id"`
	FriendID       string    `json:"friend_id"`
	Amount         float64   `json:"amount"`
	Status         string    `json:"status"`
	CreatedBy      string    `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
