package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"

	"frendler/common/constants"
	"frendler/processor/config"
	"frendler/processor/models"
	googleModels "frendler/processor/models/google"
)

func (h *HandlerImpl) GoogleContactUpdate(c echo.Context) error {
	cookieId, err := c.Cookie("userId")
	userId, err := strconv.Atoi(cookieId.Value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	//todo получить из базы данных code := sql:socialProfile(user_id).params.(json)[code]

	//todo получить токен token, err := config.GoogleOauth.Exchange(c.Request().Context(), code)
	//	if err != nil {
	//		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to exchange token: %s", err.Error()))
	//	}

	//todo получить контакты пользователя

	//todo занести контакты пользователя в базу данных, в таблицу friends
	// Получение авторизационного кода из базы данных
	socialProfile, err := h.DB.GetSocialProfileByUserId(userId, constants.PlatformGoogle)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	token, err := GetGoogleToken(socialProfile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to get authorization code from database"})
	}

	client := config.GoogleOauth.Client(c.Request().Context(), token)

	//todo функция getGoogleContacts должна возвращать комплекты
	// 1 комплект это контакт, его emails, numbers,Addresses, notes, tags( из Memberships),URLs

	contacts, err := getGoogleContacts(userId, client)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Failed to get contacts: %s", err.Error())})
	}

	//todo тут надо достать контакты из базы, обновить старые, добавить новые

	// Сохранение контактов в базу данных
	IDs, err := h.DB.CreateFriendSets(userId, contacts, constants.PlatformGoogle)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Failed to save contacts: %s", err.Error())})
	}

	return c.JSON(http.StatusOK, IDs)
}

func GetGoogleToken(socialProfile *models.SocialProfile) (*oauth2.Token, error) {
	var token oauth2.Token

	err := json.Unmarshal([]byte(socialProfile.Token), &token)
	if err != nil {
		return nil, err
	}

	return &token, nil

}

func getGoogleContacts(userId int, client *http.Client) ([]*models.Set, error) {
	resp, err := client.Get("https://people.googleapis.com/v1/people/me/connections?personFields=addresses,ageRanges,biographies,birthdays,calendarUrls,clientData,coverPhotos,emailAddresses,events,externalIds,genders,imClients,interests,locales,locations,memberships,metadata,miscKeywords,names,nicknames,occupations,organizations,phoneNumbers,photos,relations,sipAddresses,skills,urls,userDefined")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get contacts: %s", resp.Status)
	}

	var contactsResponse googleModels.GoogleContactsResponse
	if err := json.NewDecoder(resp.Body).Decode(&contactsResponse); err != nil {
		return nil, err
	}

	var friendSets []*models.Set

	for _, contact := range contactsResponse.Connections {
		var friendSet models.Set
		friendSet.Friend = models.Friend{
			OwnerID:       int64(userId),
			GivenName:     contact.Names[0].GivenName,
			FamilyName:    contact.Names[0].FamilyName,
			DisplayName:   contact.Names[0].DisplayName,
			Birthdate:     contact.GetBirthdateData(),
			Organizations: contact.GetOrganization(),
			PhoneNumber:   contact.GetPhoneNumber(),
			AvatarURL:     contact.GetPhoto(),
		}

		for _, address := range contact.Addresses {
			friendSet.Addresses = append(friendSet.Addresses, models.Address{
				Address:     address.FormattedValue,
				IsPrimary:   false,
				AddressType: address.Type,
				Country:     address.Country,
				CountryCode: address.CountryCode,
			})
		}

		for _, email := range contact.EmailAddresses {
			friendSet.EmailAddresses = append(friendSet.EmailAddresses, models.Email{
				Email:     email.Value,
				EmailType: email.Type,
			})
		}

		for _, phone := range contact.PhoneNumbers {
			friendSet.PhoneNumbers = append(friendSet.PhoneNumbers, models.PhoneNumber{
				PhoneNumber: phone.CanonicalForm,
				IsPrimary:   false,
				NumberType:  phone.Type,
			})
		}

		for _, biograph := range contact.Biographies {
			friendSet.Notes = append(friendSet.Notes, models.Note{
				Title:     biograph.Value,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Category:  biograph.ContentType,
			})
		}

		for _, event := range contact.Events {
			eventTime, err := event.Date.DateToTimeStamp() //todo добавить логирование ошибки
			if err != nil {
				continue
			}

			friendSet.Notes = append(friendSet.Notes, models.Note{
				Title:     event.Type,
				Content:   eventTime.String(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Category:  "event", //todo сделать категорию отдельным типом и сделать для нее варианты в constants
			})
		}

		for _, url := range contact.URLs {
			friendSet.URLs = append(friendSet.URLs, models.URL{
				URL:            url.Value,
				URLDescription: url.FormattedType, // todo внимательно проверить что тут все на своих местах
				URLType:        url.Type,
			})
		}

		for _, group := range contact.Memberships {
			friendSet.Tags = append(friendSet.Tags, models.Tag{
				Tag:      group.ContactGroupMembership.ContactGroupId,
				Platform: constants.PlatformGoogle,
			})
		}

		friendSets = append(friendSets, &friendSet)
	}

	return friendSets, nil

}
