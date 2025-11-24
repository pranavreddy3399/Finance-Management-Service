package handler

import (
	"context"
	"fms/internal/model"
	"fms/internal/request"
	"fmt"
)

func (eh *ExpenseHandler) validateUser(ctx context.Context, userID string) error {
	if userID == "" {
		return fmt.Errorf("user ID cannot be empty")
	}
	_, err := eh.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

func toPB(expReq *request.ExpenseCreateRequest) *model.ExpenseEntity {
	return &model.ExpenseEntity{
		UserID:   expReq.UserID,
		GroupID:  expReq.GroupID,
		Amount:   expReq.Amount,
		Category: expReq.Category,
		Note:     expReq.Note,
		Title:    expReq.Title,
	}
}
