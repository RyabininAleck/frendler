package models

import "time"

type Token struct {
	ID        int       `json:"ID"`
	UserID    int       `json:"UserID"`
	Token     string    `json:"Token"`
	IsActive  bool      `json:"IsActive"`
	CreatedAt time.Time `json:"CreatedAt"`
	ExpiresAt time.Time `json:"ExpiresAt"`
}
