package db

import (
	"fmt"
	"time"

	"frendler/processor/models"
)

func (d *DBsql) UpdateUserLogin(id int64) error {
	query := `
		UPDATE users
		SET updated_at = ?
		WHERE id = ?
	`
	_, err := d.DB.Exec(query, time.Now(), id)
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}
	return nil
}

func (d *DBsql) UpdateUser(user *models.User) error {
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

func (d *DBsql) UpdateUserSocialProfile(profile *models.SocialProfile) error {
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
