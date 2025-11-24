package repo

import (
	"context"
	"fms/internal/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ExpenseRepoInt interface {
	AddExpense(ctx context.Context, expense *model.ExpenseEntity) (string, error)
	AddExpenseSplit(ctx context.Context, expenseSplit *model.ExpenseSplitEntity) (string, error)
}

type ExpenseRepo struct {
	db *sqlx.DB
}

func NewExpenseRepo(db *sqlx.DB) ExpenseRepoInt {
	return &ExpenseRepo{
		db: db,
	}
}

func (er *ExpenseRepo) AddExpense(ctx context.Context, expense *model.ExpenseEntity) (string, error) {
	const q = `
        INSERT INTO expense (
            expense_id, user_id, group_id, amount, category,
            created_at, note, status, title, updated_at
        ) VALUES (
            :expense_id, :user_id, :group_id, :amount, :category,
            :created_at, :note, :status, :title, :updated_at
        )
    `

	_, err := er.db.NamedExecContext(ctx, q, expense)
	if err != nil {
		return "", fmt.Errorf("AddExpense: %w", err)
	}

	return expense.ExpenseID, nil
}

func (er *ExpenseRepo) AddExpenseSplit(ctx context.Context, expenseSplit *model.ExpenseSplitEntity) (string, error) {
	const q = `
        INSERT INTO expense_split (
            expense_split_id, expense_id, group_id, friend_id, amount,
            status, created_by, created_at, updated_at, paid_at
        ) VALUES (
            :expense_split_id, :expense_id, :group_id, :friend_id, :amount,
            :status, :created_by, :created_at, :updated_at, :paid_at
        )
    `

	if expenseSplit == nil {
		return "", fmt.Errorf("AddExpenseSplit: nil expenseSplit")
	}

	_, err := er.db.NamedExecContext(ctx, q, expenseSplit)
	if err != nil {
		return "", fmt.Errorf("AddExpenseSplit: %w", err)
	}

	return expenseSplit.ExpenseSplitID, nil
}
