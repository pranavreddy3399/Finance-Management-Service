package handler

import (
	"context"
	"fms/internal/model"
	"fms/internal/repo"
	"fms/internal/request"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ExpenseHandlerRepo interface {
	// Define necessary methods here
	AddExpense(ctx context.Context, expReq *request.ExpenseCreateRequest) (string, error)
}

type ExpenseHandler struct {
	ExpenseRepo repo.ExpenseRepoInt
	UserRepo    repo.UserRepoInt
}

func NewExpenseHandler(expenseRepo *repo.ExpenseRepoInt, userRepo *repo.UserRepoInt) ExpenseHandlerRepo {
	return &ExpenseHandler{
		ExpenseRepo: *expenseRepo,
		UserRepo:    *userRepo,
	}
}

func (eh *ExpenseHandler) AddExpense(ctx context.Context, expReq *request.ExpenseCreateRequest) (string, error) {
	fmt.Printf("request body is : %v", expReq)
	err := eh.validateUser(ctx, expReq.UserID)
	if err != nil {
		fmt.Printf("Add Expense | error validating userId error: %v", err)
		return "", err
	}

	if expReq.Amount <= 0 {
		return "", fmt.Errorf("amount must be greater than zero")
	}

	expenseEntity := toPB(expReq)
	expenseEntity.ExpenseID = uuid.New().String()
	expenseEntity.CreatedAt = time.Now()
	expenseEntity.UpdatedAt = time.Now()
	expenseEntity.Status = "ACTIVE"

	expenseId, err := eh.ExpenseRepo.AddExpense(ctx, expenseEntity)
	if err != nil {
		fmt.Printf("Add Expense | error adding expense error: %v", err)
		return "", err
	}

	if expReq.FriendID != "" {
		expenseSplitEntity := &model.ExpenseSplitEntity{
			ExpenseSplitID: uuid.New().String(),
			ExpenseID:      expenseEntity.ExpenseID,
			GroupID:        expReq.GroupID,
			FriendID:       expReq.FriendID,
			Amount:         expReq.Amount,
			Status:         "PENDING",
			CreatedBy:      expReq.UserID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		_, err = eh.ExpenseRepo.AddExpenseSplit(ctx, expenseSplitEntity)
		if err != nil {
			fmt.Printf("Add Expense | error adding expense split error: %v", err)
			return "", err
		}
	}

	return expenseId, nil
}
