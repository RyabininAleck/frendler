package models

import "time"

type Conflict struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	OldFriendID int64     `json:"left_friend_id"`
	NewFriendID int64     `json:"right_friend_id"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
