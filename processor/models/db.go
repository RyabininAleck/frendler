package models

import (
	"time"

	"frendler/common/constants"
)

type User struct {
	ID          int64            `json:"id"`
	Username    string           `json:"username"`
	Email       string           `json:"email"`
	Password    string           `json:"password"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	FirstName   string           `json:"first_name,omitempty"`
	LastName    string           `json:"last_name,omitempty"`
	Role        constants.Role   `json:"role"`   // CHECK(role IN ('admin', 'user', 'tester')) NOT NULL DEFAULT 'user',
	Status      constants.Status `json:"status"` // CHECK(status IN ('active', 'block', 'deleted'))
	AvatarURL   string           `json:"avatar_url,omitempty"`
	PhoneNumber string           `json:"phone_number,omitempty"` // todo номер должен приходить уже предобработанный
	Gender      string           `json:"gender,omitempty"`
	Birthdate   *time.Time       `json:"birthdate,omitempty"`
}

type SocialProfile struct {
	ID         int64              `json:"id"`
	UserID     int64              `json:"user_id"`
	Platform   constants.Platform `json:"platform"`
	ProfileURL string             `json:"profile_url"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

type Setting struct {
	ID         int64           `json:"id"`
	UserID     int64           `json:"user_id,omitempty"`
	Theme      constants.Theme `json:"theme"`
	Language   string          `json:"language"`
	AutoUpdate bool            `json:"auto_update"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}
type Friend struct {
	ID                    int64      `json:"id"`
	OwnerID               int64      `json:"owner_id"`
	Name                  string     `json:"name"`
	AlternateNames        []string   `json:"alternate_names,omitempty"`
	Birthdate             *time.Time `json:"birthdate,omitempty"`
	PhoneNumber           string     `json:"phone_number,omitempty"`
	AlternatePhoneNumbers []string   `json:"alternate_phone_numbers,omitempty"`
	AvatarURL             string     `json:"avatar_url,omitempty"`
}
type FriendNote struct {
	NoteID    int64     `json:"note_id"`
	FriendID  int64     `json:"friend_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Category  string    `json:"category,omitempty"`
}

type FriendTag struct {
	ID       int64              `json:"id"`
	FriendID int64              `json:"friend_id"`
	Tag      string             `json:"tag"`
	Platform constants.Platform `json:"platform,omitempty"`
}
