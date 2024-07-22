package handler

import (
	"encoding/json"
	"fmt"
	"log"
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

	socialProfile, err := h.DB.GetSocialProfileByUserId(userId, constants.PlatformGoogle)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	token, err := GetGoogleToken(socialProfile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to get authorization code from database"})
	}

	client := config.GoogleOauth.Client(c.Request().Context(), token)

	newContacts, err := getGoogleContacts(userId, client)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Failed to get contacts: %s", err.Error())})
	}

	oldContacts, err := h.DB.GetFriendSets(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Failed to get contacts: %s", err.Error())})
	}

	newContacts = models.RemoveDuplicates(newContacts, oldContacts)

	IDs, err := h.DB.CreateFriendSets(userId, newContacts, constants.PlatformGoogle)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Failed to save contacts: %s", err.Error())})
	}

	conflicts := DetectFriendConflicts(userId, newContacts, oldContacts)
	ConflictIDs, err := h.DB.CreateConflicts(userId, conflicts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Failed to CreateConflicts: %s", err.Error())})
	}

	return c.JSON(http.StatusOK, echo.Map{"userID": userId, "newFriendIDs": IDs, "ConflictIDs": ConflictIDs})
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
			eventTime, err := event.Date.DateToTimeStamp()
			if err != nil {
				log.Fatalf("Error converting event date to timestamp: %v", err)
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
				URLDescription: url.Type,
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

func DetectFriendConflicts(userId int, newContacts []*models.Set, oldContacts []*models.Set) []*models.Conflict {
	var conflicts []*models.Conflict

	type IdAndName struct { //todo naming
		id   int64
		name string
	}

	oldContactsPhoneMap := make(map[string]IdAndName)
	for _, oldContact := range oldContacts {
		for _, phone := range oldContact.PhoneNumbers {
			oldContactsPhoneMap[phone.PhoneNumber] = IdAndName{id: oldContact.Friend.ID, name: oldContact.Friend.DisplayName}
		}
	}

	newContactsPhoneMap := make(map[string]IdAndName)
	for _, newContact := range newContacts {
		for _, phone := range newContact.PhoneNumbers {
			newContactsPhoneMap[phone.PhoneNumber] = IdAndName{id: newContact.Friend.ID, name: newContact.Friend.DisplayName}
		}
	}

	for newContactsPhone, newContact := range newContactsPhoneMap {
		oldContact, exist := oldContactsPhoneMap[newContactsPhone]
		if exist && oldContact.name != newContact.name {
			conflicts = append(conflicts, &models.Conflict{
				UserID:      int64(userId),
				OldFriendID: oldContact.id,
				NewFriendID: newContact.id,
				IsActive:    true,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			})
		}
	}

	return conflicts
}
