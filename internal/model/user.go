package model

import "time"

type User struct {
	ID        string    `db:"id"         json:"id"`
	Name      string    `db:"name"       json:"name"`
	Email     string    `db:"email"      json:"email"`
	Phno      string    `db:"phno"       json:"phno"`
	Status    string    `db:"status"     json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type UserToFriendsEntity struct {
	UserID   string `json:"user_id" db:"user_id"`
	FriendID string `json:"friend_id" db:"friend_id"`
	Status   string `json:"status" db:"status"`
}
