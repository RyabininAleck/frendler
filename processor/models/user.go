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

func RemoveDuplicates(newSets, OldSets []*Set) []*Set {

	oldContactsPhoneMap := make(map[string]string)
	for _, oldContact := range OldSets {
		for _, phone := range oldContact.PhoneNumbers {
			oldContactsPhoneMap[phone.PhoneNumber] = oldContact.Friend.DisplayName
		}
	}

	type s struct {
		displayName string
		index       int
	}
	newContactsPhoneMap := make(map[string]s)
	for i, newContact := range newSets {
		for _, phone := range newContact.PhoneNumbers {
			newContactsPhoneMap[phone.PhoneNumber] = s{newContact.Friend.DisplayName, i}
		}
	}

	ApprovedSetIDs := make(map[int]struct{})

	//todo если у контакта несколько нометор и есть дубль этого контакта, то оно всегда проходит. Это надо пофиксить
	for newContactsPhone, newContactStruct := range newContactsPhoneMap {
		oldContactName, exist := oldContactsPhoneMap[newContactsPhone]
		if !exist || oldContactName != newContactStruct.displayName {
			ApprovedSetIDs[newContactStruct.index] = struct{}{}
		}
	}

	return nil
}
