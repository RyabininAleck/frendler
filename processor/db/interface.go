package db

import "frendler/processor/models"

type DB interface {
	//Stop()
	CheckToken(hash string) (bool, error) // проверяет есть ли такой хэш(логин пароль)

	CreateUser(user models.User) error
	UpdateUser(user models.User) error

	CreateSocialProfile(user models.SocialProfile) error // что если пользователь изменить ссылку на себя ?
	UpdateUserSocialProfile(user models.SocialProfile) error

	AddFriends(friends []models.Friend) error
	AddTags(tags []models.FriendTag) error

	//
}
