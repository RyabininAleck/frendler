package db

import (
	"fmt"

	"frendler/processor/models"
)

func (d *DBsql) CheckToken(hash string) (bool, error) {
	return true, nil
}

func (d *DBsql) CreateUser(user models.User) error {
	query := `
		INSERT INTO users (username, email, password, created_at, updated_at, first_name, last_name, role, status, avatar_url, phone_number, gender, birthdate)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := d.DB.Exec(query, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.FirstName, user.LastName, user.Role, user.Status, user.AvatarURL, user.PhoneNumber, user.Gender, user.Birthdate)
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	return nil
}

func (d *DBsql) UpdateUser(user models.User) error {
	query := `
		UPDATE users
		SET username = ?, email = ?, password = ?, updated_at = ?, first_name = ?, last_name = ?, role = ?, status = ?, avatar_url = ?, phone_number = ?, gender = ?, birthdate = ?
		WHERE id = ?
	`
	_, err := d.DB.Exec(query, user.Username, user.Email, user.Password, user.UpdatedAt, user.FirstName, user.LastName, user.Role, user.Status, user.AvatarURL, user.PhoneNumber, user.Gender, user.Birthdate, user.ID)
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}
	return nil
}

func (d *DBsql) CreateSocialProfile(profile models.SocialProfile) error {
	query := `
		INSERT INTO social_profiles (user_id, platform, profile_url, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`
	_, err := d.DB.Exec(query, profile.UserID, profile.Platform, profile.ProfileURL, profile.CreatedAt, profile.UpdatedAt)
	if err != nil {
		return fmt.Errorf("error creating social profile: %v", err)
	}
	return nil
}

func (d *DBsql) UpdateUserSocialProfile(profile models.SocialProfile) error {
	query := `
		UPDATE social_profiles
		SET platform = ?, profile_url = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := d.DB.Exec(query, profile.Platform, profile.ProfileURL, profile.UpdatedAt, profile.ID)
	if err != nil {
		return fmt.Errorf("error updating social profile: %v", err)
	}
	return nil
}

func (d *DBsql) AddFriends(friends []models.Friend) error {
	tx, err := d.DB.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
		if err != nil {
			err = fmt.Errorf("error committing transaction: %v", err)
		}
	}()

	stmt, err := tx.Prepare(`
		INSERT INTO friends (ownerID, name, alternate_names, birthdate, phone_number, alternate_phone_numbers, avatar_url)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close()

	for _, friend := range friends {
		_, err = stmt.Exec(friend.OwnerID, friend.Name, friend.AlternateNames, friend.Birthdate, friend.PhoneNumber, friend.AlternatePhoneNumbers, friend.AvatarURL)
		if err != nil {
			return fmt.Errorf("error adding friend: %v", err)
		}
	}

	return nil
}

func (d *DBsql) AddTags(tags []models.FriendTag) error {
	tx, err := d.DB.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
		if err != nil {
			err = fmt.Errorf("error committing transaction: %v", err)
		}
	}()

	stmt, err := tx.Prepare(`
		INSERT INTO friend_tags (friend_id, tag, platform)
		VALUES (?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close()

	for _, tag := range tags {
		_, err = stmt.Exec(tag.FriendID, tag.Tag, tag.Platform)
		if err != nil {
			return fmt.Errorf("error adding tag: %v", err)
		}
	}

	return nil
}
