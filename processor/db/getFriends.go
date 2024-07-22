package db

import "frendler/processor/models"

func (d *DBsql) GetFriendSets(userId int) ([]*models.Set, error) {
	friends, err := d.GetFriends(userId)
	if err != nil {
		return nil, err
	}

	var contactSets []*models.Set

	for _, friend := range friends {
		addresses, err := d.GetAddresses(friend.ID)
		if err != nil {
			return nil, err
		}

		emails, err := d.GetEmails(friend.ID)
		if err != nil {
			return nil, err
		}

		phoneNumbers, err := d.GetPhoneNumbers(friend.ID)
		if err != nil {
			return nil, err
		}

		notes, err := d.GetNotes(friend.ID)
		if err != nil {
			return nil, err
		}

		urls, err := d.GetURLs(friend.ID)
		if err != nil {
			return nil, err
		}

		tags, err := d.GetTags(friend.ID)
		if err != nil {
			return nil, err
		}

		contactSet := &models.Set{
			Friend:         friend,
			Addresses:      addresses,
			EmailAddresses: emails,
			PhoneNumbers:   phoneNumbers,
			Notes:          notes,
			URLs:           urls,
			Tags:           tags,
		}

		contactSets = append(contactSets, contactSet)
	}

	return contactSets, nil
}

func (d *DBsql) GetFriends(userId int) ([]models.Friend, error) {
	query := `
        SELECT id, ownerID, given_name, family_name, display_name, birthdate, organizations, platform, avatar_url
        FROM friends
        WHERE ownerID = ?
    `
	rows, err := d.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friends []models.Friend
	for rows.Next() {
		var friend models.Friend
		if err := rows.Scan(&friend.ID, &friend.OwnerID, &friend.GivenName, &friend.FamilyName, &friend.DisplayName, &friend.Birthdate, &friend.Organizations, &friend.Platform, &friend.AvatarURL); err != nil {
			return nil, err
		}
		friends = append(friends, friend)
	}

	return friends, nil
}

func (d *DBsql) GetAddresses(friendID int64) ([]models.Address, error) {
	query := `
        SELECT id, addresses, is_primary, address_type, country, country_code
        FROM friend_addresses
        WHERE friend_id = ?
    `
	rows, err := d.DB.Query(query, friendID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []models.Address
	for rows.Next() {
		var address models.Address
		if err := rows.Scan(&address.ID, &address.Address, &address.IsPrimary, &address.AddressType, &address.Country, &address.CountryCode); err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}

func (d *DBsql) GetEmails(friendID int64) ([]models.Email, error) {
	query := `
        SELECT id, email, email_type
        FROM friend_emails
        WHERE friend_id = ?
    `
	rows, err := d.DB.Query(query, friendID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emails []models.Email
	for rows.Next() {
		var email models.Email
		if err := rows.Scan(&email.ID, &email.Email, &email.EmailType); err != nil {
			return nil, err
		}
		emails = append(emails, email)
	}

	return emails, nil
}

func (d *DBsql) GetPhoneNumbers(friendID int64) ([]models.PhoneNumber, error) {
	query := `
        SELECT id, phone_number, is_primary, number_type
        FROM friend_phone_numbers
        WHERE friend_id = ?
    `
	rows, err := d.DB.Query(query, friendID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var phoneNumbers []models.PhoneNumber
	for rows.Next() {
		var phoneNumber models.PhoneNumber
		if err := rows.Scan(&phoneNumber.ID, &phoneNumber.PhoneNumber, &phoneNumber.IsPrimary, &phoneNumber.NumberType); err != nil {
			return nil, err
		}
		phoneNumbers = append(phoneNumbers, phoneNumber)
	}

	return phoneNumbers, nil
}

func (d *DBsql) GetNotes(friendID int64) ([]models.Note, error) {
	query := `
        SELECT note_id, title, category, content, event_time
        FROM friend_notes
        WHERE friend_id = ?
    `
	rows, err := d.DB.Query(query, friendID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.NoteID, &note.Title, &note.Category, &note.Content, &note.EventTime); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (d *DBsql) GetURLs(friendID int64) ([]models.URL, error) {
	query := `
        SELECT id, url, url_description, url_type
        FROM friend_urls
        WHERE friend_id = ?
    `
	rows, err := d.DB.Query(query, friendID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []models.URL
	for rows.Next() {
		var url models.URL
		if err := rows.Scan(&url.ID, &url.URL, &url.URLDescription, &url.URLType); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func (d *DBsql) GetTags(friendID int64) ([]models.Tag, error) {
	query := `
        SELECT id, tag, platform
        FROM friend_tags
        WHERE friend_id = ?
    `
	rows, err := d.DB.Query(query, friendID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []models.Tag
	for rows.Next() {
		var tag models.Tag
		if err := rows.Scan(&tag.ID, &tag.Tag, &tag.Platform); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
