package db

import (
	"frendler/common/constants"
	"frendler/processor/models"
)

type DB interface {
	GetUserById(userID int) (*models.User, error)
	GetUserByExternalId(externalUserID string, service constants.Platform) (*models.SocialProfile, error)

	CreateSocialProfile(profile *models.SocialProfile) (int64, error)
	CreateUserAndSetting(user *models.User, set *models.Setting) (int64, int64, error)

	UpdateUserLogin(id int64) error

	CreateToken(id int64, token string) (int64, error)
	CheckToken(id string, token string) (bool, error) // проверяет есть ли такой хэш(логин пароль)

	UpdateUser(user *models.User) error
	//
	//CreateSocialProfile(user *models.SocialProfile) error // что если пользователь изменить ссылку на себя ?
	//UpdateUserSocialProfile(user *models.SocialProfile) error
	//
	//AddFriends(friends []models.Friend) error
	//AddTags(tags []models.FriendTag) error
	//
	GetSetting(userID int) (*models.Setting, error)
}
