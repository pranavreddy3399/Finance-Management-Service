package repo

import (
	"context"
	"database/sql"
	"errors"

	"fms/internal/model"
	"github.com/jmoiron/sqlx"
)

// ErrNotFound is returned when a record does not exist.
var ErrNotFound = errors.New("record not found")

// UserRepoInt defines all persistence methods for users.
type UserRepoInt interface {
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
}

type UserRepo struct {
	db *sqlx.DB
}

// NewUserRepoSQL returns a SQL implementation of UserRepoInt.
func NewUserRepoSQL(db *sqlx.DB) UserRepoInt {
	return &UserRepo{db: db}
}

func (u *UserRepo) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	var user model.User
	err := u.db.GetContext(ctx, &user,
		"SELECT id, name, email, phno, status, created_at, updated_at FROM users WHERE id = ?",
		userID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}
