package googleModels

import (
	"fmt"
	"time"
)

type LoginRequest struct {
	RedirectURL string `json:"redirect_url"`
}

type User struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type GoogleContactsResponse struct {
	Connections []Contact `json:"connections"`
	TotalPeople int       `json:"totalPeople"`
	TotalItems  int       `json:"totalItems"`
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

	Addresses      []Addresses `json:"addresses"` //несколько
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
		Date          Date   `json:"date"`
		Type          string `json:"type"`
		FormattedType string `json:"formattedType"`
	} `json:"events"` //несколько
}

func (c *Contact) GetBirthdateData() *time.Time {
	if len(c.Birthdays) == 0 {
		return nil
	}

	date, _ := c.Birthdays[0].Date.DateToTimeStamp()
	return date
}

func (c *Contact) GetOrganization() string {
	if len(c.Organizations) == 0 {
		return ""
	}
	return c.Organizations[0].Name
}

func (c *Contact) GetPhoneNumber() string {
	if len(c.PhoneNumbers) == 0 {
		return ""
	}
	return c.PhoneNumbers[0].CanonicalForm
}

func (c *Contact) GetPhoto() string {
	if len(c.Photos) == 0 {
		return ""
	}
	return c.Photos[0].URL
}

type Addresses struct {
	FormattedValue string `json:"formattedValue"`
	Type           string `json:"type"`
	Country        string `json:"country"`
	CountryCode    string `json:"countryCode"`
}

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

func (date *Date) DateToTimeStamp() (*time.Time, error) {

	if date.Month < 1 || date.Month > 12 {
		return nil, fmt.Errorf("invalid month: %d", date.Month)
	}
	if date.Day < 1 || date.Day > 31 {
		return nil, fmt.Errorf("invalid day: %d", date.Day)
	}

	t := time.Date(date.Year, time.Month(date.Month), date.Day, 0, 0, 0, 0, time.UTC)

	if t.Month() != time.Month(date.Month) || t.Day() != date.Day {
		return nil, fmt.Errorf("invalid date: %v", date)
	}

	return &t, nil
}
