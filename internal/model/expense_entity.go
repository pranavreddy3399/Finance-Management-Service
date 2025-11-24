package model

import (
	"fms/internal/model/enums"
	"time"
)

type ExpenseEntity struct {
	ExpenseID string                `json:"expense_id" db:"expense_id"`
	UserID    string                `json:"user_id" db:"user_id"`
	GroupID   string                `json:"group_id" db:"group_id"`
	Amount    float64               `json:"amount" db:"amount"`
	Category  enums.ExpenseCategory `json:"category" db:"category"`
	CreatedAt time.Time             `json:"created_at" db:"created_at"`
	Note      string                `json:"note" db:"note"`
	Status    string                `json:"status" db:"status"`
	Title     string                `json:"title" db:"title"`
	UpdatedAt time.Time             `json:"updated_at" db:"updated_at"`
}

// we can add txnId like those for more extension
type ExpenseSplitEntity struct {
	ExpenseSplitID string    `json:"expense_split_id" db:"expense_split_id"`
	ExpenseID      string    `json:"expense_id" db:"expense_id"`
	GroupID        string    `json:"group_id" db:"group_id"`
	FriendID       string    `json:"friend_id" db:"friend_id"`
	Amount         float64   `json:"amount" db:"amount"`
	Status         string    `json:"status" db:"status"`
	CreatedBy      string    `json:"created_by" db:"created_by"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	PaidAt         time.Time `json:"paid_at" db:"paid_at"`
}
