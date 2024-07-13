package models

import (
	"time"

	"frendler/common/constants"
)

type Set struct {
	Friend         Friend
	Addresses      []Address
	EmailAddresses []Email
	PhoneNumbers   []PhoneNumber
	Notes          []Note
	URLs           []URL
	Tags           []Tag
}

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
	ExternalID string             `json:"external_id"`
	ProfileURL string             `json:"profile_url"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	Params     string             `json:"params"`
	Token      string             `json:"token"`
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
	ID            int64      `json:"id"`
	OwnerID       int64      `json:"owner_id"`
	GivenName     string     `json:"given_name"`
	FamilyName    string     `json:"family_name"`
	DisplayName   string     `json:"display_name"`
	Birthdate     *time.Time `json:"birthdate,omitempty"`
	Organizations string     `json:"organizations,omitempty"`
	PhoneNumber   string     `json:"phone_number,omitempty"`
	AvatarURL     string     `json:"avatar_url,omitempty"`
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
