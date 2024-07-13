package googleModels

import "frendler/common/constants"

type LoginRequest struct {
	RedirectURL string `json:"redirect_url"`
}

type User struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type PhoneNumber struct {
	ID          int    `json:"id"`
	FriendID    int    `json:"friend_id"`
	PhoneNumber string `json:"phone_number"`
	IsPrimary   bool   `json:"is_primary"`
	NumberType  string `json:"number_type"`
}

type Tags struct {
	ID          int                `json:"id"`
	FriendID    int                `json:"friend_id"`
	Value       string             `json:"value"`
	ContentType string             `json:"contentType"`
	Platform    constants.Platform `json:"platform"`
}
