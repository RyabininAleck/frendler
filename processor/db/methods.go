package db

import (
	"fmt"

	"frendler/processor/models"
)

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
		INSERT INTO friends (ownerID, given_name, family_name, display_name, birthdate, avatar_url, organizations, platform)
		VALUES (?, ?, ?, ?, ?,?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close()

	for _, friend := range friends {
		//todo обработать friend.PhoneNumber при наличии
		_, err = stmt.Exec(friend.OwnerID, friend.GivenName, friend.FamilyName, friend.DisplayName, friend.Birthdate, friend.AvatarURL, friend.Organizations, friend.Platform)
		if err != nil {
			return fmt.Errorf("error adding friend: %v", err)
		}
	}

	return nil
}

func (d *DBsql) AddTags(tags []models.Tag) error {
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
