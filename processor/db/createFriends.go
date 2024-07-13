package db

import (
	"fmt"
	"log"

	"frendler/common/constants"
	"frendler/processor/models"
)

func (d *DBsql) CreateFriendSets(userId int, contacts []*models.Set, platform constants.Platform) ([]int64, error) {
	socialProfileIDs := []int64{}
	for _, contact := range contacts {
		friendId, err := d.CreateFriendSet(userId, contact, platform)
		if err != nil {
			log.Printf("error CreateFriendSet for id: %v platform:%v: %v", userId, platform, err)
			continue
		}
		socialProfileIDs = append(socialProfileIDs, friendId)
	}
	return socialProfileIDs, nil
}

func (d *DBsql) CreateFriendSet(userId int, contact *models.Set, platform constants.Platform) (int64, error) {
	// todo надо попробовать чтобы в каждом из списков было несколько значений
	friendID, _ := d.CreateFriend(userId, platform, contact.Friend)

	addressIDs, _ := d.CreateAddresses(friendID, contact.Addresses)
	emailIDs, _ := d.CreateEmails(friendID, contact.EmailAddresses)
	phoneIDs, err := d.CreatePhoneNumbers(friendID, contact.PhoneNumbers)
	noteIDs, _ := d.CreateNotes(friendID, contact.Notes)
	urlIDs, _ := d.CreateURLs(friendID, contact.URLs)
	tagIDs, _ := d.CreateTags(friendID, contact.Tags)

	if err != nil {
		log.Printf("error CreateFriendSet for id: %v platform:%v: %v", userId, platform, err)
	}
	log.Printf("For friendID:%d was created:\nAddresses: %v\nEmails: %v\nPhone Numbers: %v\nNotes: %v\nURLs: %v\nTags: %v", friendID, addressIDs, emailIDs, phoneIDs, noteIDs, urlIDs, tagIDs)

	return friendID, nil
}

func (d *DBsql) CreateFriend(userId int, platform constants.Platform, Friend models.Friend) (int64, error) {
	query := `
		INSERT INTO friends (ownerID, given_name, family_name, display_name, birthdate, organizations, platform, avatar_url )
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	res, err := d.DB.Exec(query, userId, Friend.GivenName, Friend.FamilyName, Friend.DisplayName, Friend.Birthdate, Friend.Organizations, platform, Friend.AvatarURL)
	if err != nil {
		return 0, fmt.Errorf("error creating social profile: %v", err)
	}

	friendID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert ID for user: %v", err)
	}

	return friendID, nil
}

// --
//func (d *DBsql) CreateAddresses(friendID int64, addresses []models.Address) (int64, error) {
//	// Ваш код для создания адресов
//	return nil
//}
//
//func (d *DBsql) CreateEmails(friendID int64, emails []models.Email) (int64, error) {
//	// Ваш код для создания электронных адресов
//	return nil
//}
//
//func (d *DBsql) CreatePhoneNumbers(friendID int64, phoneNumbers []models.PhoneNumber) (int64, error) {
//	// Ваш код для создания номеров телефонов
//	return nil
//}
//
//func (d *DBsql) CreateNotes(friendID int64, notes []models.Note) (int64, error) {
//	// Ваш код для создания заметок
//	return nil
//}
//
//func (d *DBsql) CreateURLs(friendID int64, urls []models.URL) (int64, error) {
//	// Ваш код для создания URL
//	return nil
//}
//
//func (d *DBsql) CreateTags(friendID int64, tags []models.Tag) (int64, error) {
//	// Ваш код для создания тегов
//	return nil
//}

// CreateAddresses - метод для создания адресов
func (d *DBsql) CreateAddresses(friendID int64, addresses []models.Address) ([]int64, error) {
	var ids []int64
	for _, address := range addresses {
		query := `
			INSERT INTO friend_addresses (friend_id, addresses, is_primary, address_type, country, country_code)
			VALUES (?, ?, ?, ?, ?, ?)
		`
		res, err := d.DB.Exec(query, friendID, address.Address, address.IsPrimary, address.AddressType, address.Country, address.CountryCode)
		if err != nil {
			return nil, fmt.Errorf("error creating address: %v", err)
		}

		addressID, err := res.LastInsertId()
		if err != nil {
			return nil, fmt.Errorf("error getting last insert ID for address: %v", err)
		}

		ids = append(ids, addressID)
	}
	return ids, nil
}

// CreateEmails - метод для создания электронных адресов
func (d *DBsql) CreateEmails(friendID int64, emails []models.Email) ([]int64, error) {
	var ids []int64
	for _, email := range emails {
		query := `
			INSERT INTO friend_emails (friend_id, email, email_type)
			VALUES (?, ?, ?)
		`
		res, err := d.DB.Exec(query, friendID, email.Email, email.EmailType)
		if err != nil {
			return nil, fmt.Errorf("error creating email: %v", err)
		}

		emailID, err := res.LastInsertId()
		if err != nil {
			return nil, fmt.Errorf("error getting last insert ID for email: %v", err)
		}

		ids = append(ids, emailID)
	}
	return ids, nil
}

// CreatePhoneNumbers - метод для создания номеров телефонов
func (d *DBsql) CreatePhoneNumbers(friendID int64, phoneNumbers []models.PhoneNumber) ([]int64, error) {
	var ids []int64
	for _, phoneNumber := range phoneNumbers {
		query := `
			INSERT INTO friend_phone_numbers (friend_id, phone_number, is_primary, number_type)
			VALUES (?, ?, ?, ?)
		`
		res, err := d.DB.Exec(query, friendID, phoneNumber.PhoneNumber, phoneNumber.IsPrimary, phoneNumber.NumberType)
		if err != nil {
			return nil, fmt.Errorf("error creating phone number: %v", err)
		}

		phoneID, err := res.LastInsertId()
		if err != nil {
			return nil, fmt.Errorf("error getting last insert ID for phone number: %v", err)
		}

		ids = append(ids, phoneID)
	}
	return ids, nil
}

// CreateNotes - метод для создания заметок
func (d *DBsql) CreateNotes(friendID int64, notes []models.Note) ([]int64, error) {
	var ids []int64
	for _, note := range notes {
		query := `
			INSERT INTO friend_notes (friend_id, title, category, content, event_time)
			VALUES (?, ?, ?, ?, ?)
		`
		res, err := d.DB.Exec(query, friendID, note.Title, note.Category, note.Content, note.EventTime)
		if err != nil {
			return nil, fmt.Errorf("error creating note: %v", err)
		}

		noteID, err := res.LastInsertId()
		if err != nil {
			return nil, fmt.Errorf("error getting last insert ID for note: %v", err)
		}

		ids = append(ids, noteID)
	}
	return ids, nil
}

// CreateURLs - метод для создания URL
func (d *DBsql) CreateURLs(friendID int64, urls []models.URL) ([]int64, error) {
	var ids []int64
	for _, url := range urls {
		query := `
			INSERT INTO friend_urls (friend_id, url, url_description, url_type)
			VALUES (?, ?, ?, ?)
		`
		res, err := d.DB.Exec(query, friendID, url.URL, url.URLDescription, url.URLType)
		if err != nil {
			return nil, fmt.Errorf("error creating URL: %v", err)
		}

		urlID, err := res.LastInsertId()
		if err != nil {
			return nil, fmt.Errorf("error getting last insert ID for URL: %v", err)
		}

		ids = append(ids, urlID)
	}
	return ids, nil
}

// CreateTags - метод для создания тегов
func (d *DBsql) CreateTags(friendID int64, tags []models.Tag) ([]int64, error) {
	var ids []int64
	for _, tag := range tags {
		query := `
			INSERT INTO friend_tags (friend_id, tag, platform)
			VALUES (?, ?, ?)
		`
		res, err := d.DB.Exec(query, friendID, tag.Tag, tag.Platform)
		if err != nil {
			return nil, fmt.Errorf("error creating tag: %v", err)
		}

		tagID, err := res.LastInsertId()
		if err != nil {
			return nil, fmt.Errorf("error getting last insert ID for tag: %v", err)
		}

		ids = append(ids, tagID)
	}
	return ids, nil
}
