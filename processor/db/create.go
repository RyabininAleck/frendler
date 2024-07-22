package db

import (
	"fmt"
	"time"

	"frendler/processor/models"
)

func (d *DBsql) CreateUserAndSetting(user *models.User, set *models.Setting) (int64, int64, error) {
	query := `
		INSERT INTO users (username, email, password, first_name, last_name, role, status, avatar_url, phone_number, gender, birthdate)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	res, err := d.DB.Exec(query, user.Username, user.Email, user.Password, user.FirstName, user.LastName, user.Role, user.Status, user.AvatarURL, user.PhoneNumber, user.Gender, user.Birthdate)
	if err != nil {
		return 0, 0, fmt.Errorf("error creating user: %v", err)
	}

	userID, err := res.LastInsertId()
	if err != nil {
		return 0, 0, fmt.Errorf("error getting last insert ID for user: %v", err)
	}

	querySetting := `
		INSERT INTO settings (user_id, theme, language, auto_update, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	res, err = d.DB.Exec(querySetting, userID, set.Theme, set.Language, set.AutoUpdate, time.Now(), time.Now())
	if err != nil {
		return 0, 0, fmt.Errorf("error creating setting: %v", err)
	}

	settingID, err := res.LastInsertId()
	if err != nil {
		return 0, 0, fmt.Errorf("error getting last insert ID for setting: %v", err)
	}

	return userID, settingID, nil
}

func (d *DBsql) CreateSocialProfile(profile *models.SocialProfile) (int64, error) {
	query := `
		INSERT INTO social_profiles (user_id, external_id, platform, profile_url, created_at, updated_at, params, token)
		VALUES (?, ?, ?, ?, ?, ?,?, ? )
	`
	res, err := d.DB.Exec(query, profile.UserID, profile.ExternalID, profile.Platform, profile.ProfileURL, profile.CreatedAt, profile.UpdatedAt, profile.Params, profile.Token)
	if err != nil {
		return 0, fmt.Errorf("error creating social profile: %v", err)
	}

	socialProfileID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert ID for user: %v", err)
	}

	return socialProfileID, nil
}

// todo добавить создание друга, создание друзей
func (d *DBsql) CreateConflicts(userId int, conflicts []*models.Conflict) ([]int64, error) {
	var ids []int64
	query := `
		INSERT INTO conflicts (user_id, left_friend_id, right_friend_id, is_active)
		VALUES (?, ?, ?, ?)
	`
	for _, conflict := range conflicts {
		res, err := d.DB.Exec(query, userId, conflict.OldFriendID, conflict.NewFriendID, conflict.IsActive)
		if err != nil {
			return nil, fmt.Errorf("error creating conflict: %v", err)
		}

		conflictID, err := res.LastInsertId()
		if err != nil {
			return nil, fmt.Errorf("error getting last insert ID for conflict: %v", err)
		}

		ids = append(ids, conflictID)
	}

	return ids, nil
}
