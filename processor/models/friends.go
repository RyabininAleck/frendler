package models

import (
	"time"

	"frendler/common/constants"
)

type Friend struct {
	ID            int64      `json:"id"`
	OwnerID       int64      `json:"owner_id"`
	GivenName     string     `json:"given_name"`
	FamilyName    string     `json:"family_name"`
	DisplayName   string     `json:"display_name"`
	Birthdate     *time.Time `json:"birthdate,omitempty"`
	Organizations string     `json:"organizations,omitempty"`
	PhoneNumber   string     `json:"phone_number,omitempty"`
	AvatarURL     string     `json:"avatar_url,omitempty"`
	Platform      string     `json:"platform,omitempty"`
}

type PhoneNumber struct {
	ID          int    `json:"id"`
	FriendID    int    `json:"friend_id"`
	PhoneNumber string `json:"phone_number"`
	IsPrimary   bool   `json:"is_primary"`
	NumberType  string `json:"number_type"`
}

type Email struct {
	ID        int    `json:"id"`
	FriendID  int    `json:"friend_id"`
	Email     string `json:"email"`
	EmailType string `json:"email_type"`
}

type URL struct {
	ID             int    `json:"id"`
	FriendID       int    `json:"friend_id"`
	URL            string `json:"url"`
	URLDescription string `json:"url_description"`
	URLType        string `json:"url_type"`
}

type Address struct {
	ID          int    `json:"id"`
	FriendID    int    `json:"friend_id"`
	Address     string `json:"address"`
	IsPrimary   bool   `json:"is_primary"`
	AddressType string `json:"address_type"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

type Note struct {
	NoteID    int64     `json:"note_id"`
	FriendID  int64     `json:"friend_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	EventTime time.Time `json:"event_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Category  string    `json:"category,omitempty"`
}

type Tag struct {
	ID       int64              `json:"id"`
	FriendID int64              `json:"friend_id"`
	Tag      string             `json:"tag"`
	Platform constants.Platform `json:"platform,omitempty"`
}
