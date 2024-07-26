package db

import (
	"time"

	"frendler/processor/models"
)

func (d *DBsql) CheckToken(id string, sessionToken string) *models.Token {
	var token models.Token
	query := `SELECT id, user_id, token, is_active FROM tokens WHERE user_id = ? AND token = ?`
	err := d.DB.QueryRow(query, id, sessionToken).Scan(&token.ID, &token.UserID, &token.Token, &token.IsActive)
	if err != nil {
		return nil
	}

	//todo  Проверить, что токен не протух
	return &token
}

func (d *DBsql) CreateToken(userID int64, token string) (int64, error) {
	query := `INSERT INTO tokens (user_id, token, created_at) VALUES (?, ?, ?)`
	result, err := d.DB.Exec(query, userID, token, time.Now())
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
