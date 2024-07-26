package db

import (
	"frendler/common/constants"
	"frendler/processor/models"
)

type DB interface {
	GetUserById(userID int) (*models.User, error)
	GetSocialProfileByExternalId(externalUserID string, service constants.Platform) (*models.SocialProfile, error)
	GetSocialProfileByUserId(UserID int, service constants.Platform) (*models.SocialProfile, error)
	GetFriendSets(userId int) ([]*models.Set, error)

	CreateSocialProfile(profile *models.SocialProfile) (int64, error)
	CreateUserAndSetting(user *models.User, set *models.Setting) (int64, int64, error)

	CreateFriendSet(userId int, contact *models.Set, platform constants.Platform) (int64, error)
	CreateFriendSets(userId int, contact []*models.Set, platform constants.Platform) ([]int64, error)
	CreateConflicts(userId int, conflicts []*models.Conflict) ([]int64, error)

	UpdateUserLogin(id int64) error

	CreateToken(id int64, token string) (int64, error)
	CheckToken(id string, token string) *models.Token

	UpdateUser(user *models.User) error
	//UpdateUserSocialProfile(user *models.SocialProfile) error// что если пользователь изменить ссылку на себя ?

	GetSetting(userID int) (*models.Setting, error)

	GetContactStats(userID int) (int64, int64, error)
}
