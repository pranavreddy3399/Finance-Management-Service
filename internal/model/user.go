package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phno      string    `json:"phno"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserToFriendsEntity struct {
	UserID   string `json:"user_id"`
	FriendID string `json:"friend_id"`
	Status   string `json:"status"`
}
