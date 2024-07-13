package db

import (
	"database/sql"
	"fmt"

	"frendler/common/constants"
	"frendler/processor/models"
)

// GetSetting возвращает настройки пользователя по id
func (d *DBsql) GetSetting(userID int) (*models.Setting, error) {
	var setting models.Setting

	query := `
		SELECT id, theme, language,auto_update,created_at,updated_at
		FROM settings
		WHERE user_id = ?
	`

	err := d.DB.QueryRow(query, userID).Scan(&setting.ID, &setting.Theme, &setting.Language, &setting.AutoUpdate, &setting.CreatedAt, &setting.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &setting, fmt.Errorf("setting with id %d not found", userID)
		}
		return &setting, fmt.Errorf("error fetching setting: %v", err)
	}

	return &setting, nil
}

// GetSetting возвращает параметры пользователя по id
func (d *DBsql) GetUserById(userID int) (*models.User, error) {
	var user models.User
	query := `
		SELECT id, username, email, password, created_at, updated_at, first_name, last_name, role, status, avatar_url, phone_number, gender, birthdate
		FROM users
		WHERE id = ?
	`

	err := d.DB.QueryRow(query, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.Status,
		&user.AvatarURL,
		&user.PhoneNumber,
		&user.Gender,
		&user.Birthdate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, fmt.Errorf("profile for user with id %d not found", userID)
		}
		return &user, fmt.Errorf("error fetching profile: %v", err)
	}

	return &user, nil

}

// GetUserByExternalId возвращает параметры пользователя по внешнему id. Ищет externalUserID в SocialProfile
func (d *DBsql) GetSocialProfileByExternalId(externalUserID string, service constants.Platform) (*models.SocialProfile, error) {
	var socialProfile models.SocialProfile
	query := `
		SELECT id, user_id, platform, profile_url, external_id
		FROM social_profiles
		WHERE external_id = ? and platform = ?
	`

	err := d.DB.QueryRow(query, externalUserID, service).Scan(
		&socialProfile.ID,
		&socialProfile.UserID,
		&socialProfile.Platform,
		&socialProfile.ProfileURL,
		&socialProfile.ExternalID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &socialProfile, fmt.Errorf("error fetching profile: %v", err)
	}

	return &socialProfile, nil

}

// GetUserByExternalId возвращает параметры пользователя по внешнему id. Ищет externalUserID в SocialProfile
func (d *DBsql) GetSocialProfileByUserId(UserID int, service constants.Platform) (*models.SocialProfile, error) {
	var socialProfile models.SocialProfile
	query := `
		SELECT id, user_id, params, token, platform, profile_url, external_id
		FROM social_profiles
		WHERE user_id = ? and platform = ?
	`

	err := d.DB.QueryRow(query, UserID, service).Scan(
		&socialProfile.ID,
		&socialProfile.UserID,
		&socialProfile.Params,
		&socialProfile.Token,
		&socialProfile.Platform,
		&socialProfile.ProfileURL,
		&socialProfile.ExternalID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return &socialProfile, fmt.Errorf("error fetching profile: %v", err)
	}

	return &socialProfile, nil

}

// GetContactStats возвращает количество общее количество контактов и количество неразрешенных конфликтов
func (d *DBsql) GetContactStats(userID int) (int64, int64, error) {
	var contactCount, conflictCount int64

	// Query to get the number of contacts
	contactQuery := `SELECT COUNT(*) FROM friends WHERE ownerID = ?`
	err := d.DB.QueryRow(contactQuery, userID).Scan(&contactCount)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get contact count: %v", err)
	}

	// Query to get the number of conflicts
	conflictQuery := `SELECT COUNT(*) FROM conflicts WHERE user_id = ?`
	err = d.DB.QueryRow(conflictQuery, userID).Scan(&conflictCount)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get conflict count: %v", err)
	}

	return contactCount, conflictCount, nil
}
