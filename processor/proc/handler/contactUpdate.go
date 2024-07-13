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
	IDs, err := h.DB.CreateFriends(userId, contacts, constants.PlatformGoogle)
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

// Вспомогательная функция для запроса контактов пользователя
func getGoogleContacts(userId int, client *http.Client) ([]*models.Friend, error) {
	resp, err := client.Get("https://people.googleapis.com/v1/people/me/connections?personFields=addresses,ageRanges,biographies,birthdays,calendarUrls,clientData,coverPhotos,emailAddresses,events,externalIds,genders,imClients,interests,locales,locations,memberships,metadata,miscKeywords,names,nicknames,occupations,organizations,phoneNumbers,photos,relations,sipAddresses,skills,urls,userDefined")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get contacts: %s", resp.Status)
	}

	var contactsResponse GoogleContactsResponse
	if err := json.NewDecoder(resp.Body).Decode(&contactsResponse); err != nil {
		return nil, err
	}

	var friends []*models.Friend
	//var addresses []*models.Address
	//var emails []*models.Email
	//var phoneNumbers []*models.PhoneNumber
	//var urls []*models.URL
	//var note []*models.Note

	for _, contact := range contactsResponse.Connections {
		friend := models.Friend{
			OwnerID:       int64(userId),
			GivenName:     contact.Names[0].GivenName,
			FamilyName:    contact.Names[0].FamilyName,
			DisplayName:   contact.Names[0].DisplayName,
			Birthdate:     contact.getBirthdateData(),
			Organizations: contact.getOrganization(),
			PhoneNumber:   contact.getPhoneNumber(),
			AvatarURL:     contact.getPhoto(),
		}
		friends = append(friends, &friend)
		//todo вот тут не проставляется friend_id потому что сам friend еще не в базе и не известен его Id недо не забыть подставить Id по время проброса в базу данных
		//		for _, address := range contact.Addresses {
		//
		//		}
		//
		//		address := &models.Address{
		//			Address:     [],
		//			IsPrimary:   false,
		//			AddressType: "",
		//			Country:     "",
		//			CountryCode: "",
		//		}
		//		addresses = append(addresses,address)

	}

	// todo преобразовать contacts в []models.Friend
	return friends, nil

}

type Contact struct {
	Names []struct {
		DisplayName string `json:"displayName"`
		FamilyName  string `json:"familyName"`
		GivenName   string `json:"givenName"`
	} `json:"names"` //одно
	Photos []struct {
		URL string `json:"url"`
	} `json:"photos"` // одно
	Birthdays []struct {
		Date Date `json:"date"`
	} `json:"birthdays"` //одно
	Organizations []struct {
		Name string `json:"name"`
	} `json:"organizations"` //одно

	Addresses []struct {
		FormattedValue string `json:"formattedValue"`
		Type           string `json:"type"`
		Country        string `json:"country"`
		CountryCode    string `json:"countryCode"`
	} `json:"addresses"` //несколько
	EmailAddresses []struct {
		Value string `json:"value"`
		Type  string `json:"type"`
	} `json:"emailAddresses"` //несколько
	PhoneNumbers []struct {
		CanonicalForm string `json:"canonicalForm"`
		Type          string `json:"type"`
	} `json:"phoneNumbers"` //несколько
	Biographies []struct {
		Value       string `json:"value"`
		ContentType string `json:"contentType"`
	} `json:"biographies"` //несколько
	URLs []struct {
		Value         string `json:"value"`
		Type          string `json:"type"`
		FormattedType string `json:"formattedType"`
	} `json:"urls"` //несколько
	Memberships []struct {
		ContactGroupMembership struct {
			ContactGroupId string `json:"contactGroupId"`
		} `json:"contactGroupMembership"`
	} `json:"memberships"` //несколько
	Events []struct {
		Date struct {
			Year  int `json:"year"`
			Month int `json:"month"`
			Day   int `json:"day"`
		} `json:"date"`
		Type          string `json:"type"`
		FormattedType string `json:"formattedType"`
	} `json:"events"` //несколько
}

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}
type GoogleContactsResponse struct {
	Connections []Contact `json:"connections"`
	TotalPeople int       `json:"totalPeople"`
	TotalItems  int       `json:"totalItems"`
}

func (c *Contact) getBirthdateData() *time.Time {
	if len(c.Birthdays) == 0 {
		return nil
	}

	date, _ := dateToTimeStamp(c.Birthdays[0].Date)
	return date
}

func dateToTimeStamp(date Date) (*time.Time, error) {
	// Проверяем корректность значений месяца и дня
	if date.Month < 1 || date.Month > 12 {
		return nil, fmt.Errorf("invalid month: %d", date.Month)
	}
	if date.Day < 1 || date.Day > 31 {
		return nil, fmt.Errorf("invalid day: %d", date.Day)
	}

	// Пытаемся создать объект time.Time
	t := time.Date(date.Year, time.Month(date.Month), date.Day, 0, 0, 0, 0, time.UTC)

	// Проверяем, корректна ли созданная дата
	if t.Month() != time.Month(date.Month) || t.Day() != date.Day {
		return nil, fmt.Errorf("invalid date: %v", date)
	}

	return &t, nil
}

func (c *Contact) getOrganization() string {
	if len(c.Organizations) == 0 {
		return ""
	}
	return c.Organizations[0].Name
}

func (c *Contact) getPhoneNumber() string {
	if len(c.PhoneNumbers) == 0 {
		return ""
	}
	return c.PhoneNumbers[0].CanonicalForm
}
func (c *Contact) getPhoto() string {
	if len(c.Photos) == 0 {
		return ""
	}
	return c.Photos[0].URL
}

//type complect struct {
//	Friend models.Friend
//	Addresses []models.Address
//	EmailAddresses []models.Email
//	PhoneNumbers []models.PhoneNumber
//	Notes []models.Note
//	URLs []models.URL
//	Tags []models.Tags
//
//}
