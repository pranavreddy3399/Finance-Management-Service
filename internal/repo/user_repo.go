package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"fms/internal/model"

	"github.com/jmoiron/sqlx"
)

// ErrNotFound is returned when a record does not exist.
var ErrNotFound = errors.New("record not found")

// UserRepoInt defines all persistence methods for users.
type UserRepoInt interface {
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
}

type UserRepo struct {
	db *sqlx.DB
}

// NewUserRepoSQL returns a SQL implementation of UserRepoInt.
func NewUserRepo(db *sqlx.DB) UserRepoInt {
	return &UserRepo{db: db}
}

func (u *UserRepo) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	var user model.User
	err := u.db.GetContext(ctx, &user,
		"SELECT id, name, email, phno, status, created_at, updated_at FROM user1 WHERE id = ?",
		userID,
	)
	if err != nil {
		fmt.Printf("GetUserByID | Error fetching user with ID %s: %v\n", userID, err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) CreateUser(ctx context.Context, user *model.User) error {
	const q = `
        INSERT INTO user1 
        (id, name, email, phno, status, created_at, updated_at)
        VALUES 
        (:id, :name, :email, :phno, :status, :created_at, :updated_at)
    `

	_, err := u.db.NamedExecContext(ctx, q, user)
	if err != nil {
		return fmt.Errorf("CreateUser: %w", err)
	}
	return nil
}
