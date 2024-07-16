package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"

	"frendler/common/constants"
	"frendler/processor/config"
	"frendler/processor/models"
	googleModels "frendler/processor/models/google"
)

//func handleGoogleCallback(c echo.Context) error {
//	//todo делается проверка на успешный статус
//	//todo проверка что пользователь такой есть
//	//todo если пользователь есть в базе данных,
//	//	//todo то обновить его update time
//	//	//todo иначе,
//	//		//todo запросить данные пользователя
//	//		//todo создать пользователя
//
//	state := c.QueryParam("state")
//	if state != config.OauthStateString {
//		log.Println("invalid oauth state")
//		return c.Redirect(http.StatusTemporaryRedirect, "/")
//	}
//
//	return c.NoContent(http.StatusOK)
//}

type QueryParams struct {
	Authuser string `json:"authuser"`
	Code     string `json:"code"`
	Prompt   string `json:"prompt"`
	Scope    string `json:"scope"`
	State    string `json:"state"`
	Token    string `json:"token"`
}

func (h *HandlerImpl) HandleGoogleCallback(c echo.Context) error {
	state := c.QueryParam("state")
	if state != config.OauthStateString {
		log.Println("Invalid OAuth state")
		return c.Redirect(http.StatusTemporaryRedirect, "/error")
	}

	code := c.QueryParam("code")
	token, err := config.GoogleOauth.Exchange(c.Request().Context(), code)
	if err != nil {
		//todo сделать редирект на страинцу плагина
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to exchange token: %s", err.Error()))
	}

	client := config.GoogleOauth.Client(c.Request().Context(), token)
	userGoogleInfo, err := getUserInfo(client)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to get user info: %s", err.Error()))
	}

	user, err := h.DB.GetSocialProfileByExternalId(userGoogleInfo.ID, constants.PlatformGoogle)
	if err != nil {
		log.Printf("Database error: %v\n", err)
		return c.Redirect(http.StatusTemporaryRedirect, "/error")
	}

	var userId int64

	if user != nil {
		userId = user.UserID
		err := h.DB.UpdateUserLogin(user.UserID)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to updete user info: %s", err.Error()))
		}
	} else {
		userId, _, err = h.DB.CreateUserAndSetting(
			&models.User{
				Username:    userGoogleInfo.Email,
				Email:       userGoogleInfo.Email,
				Password:    "",
				Role:        constants.RoleUser,
				Status:      constants.StatusActive,
				AvatarURL:   userGoogleInfo.Picture,
				PhoneNumber: "",
			},
			&models.Setting{
				Theme:      constants.ThemeLight,
				Language:   "ru",
				AutoUpdate: false,
			},
		)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to create user: %s", err.Error()))
		}

		jsonData, err := getJsonQueryParams(c)
		if err != nil {
			log.Printf("Failed to get JSON params: %v", err)
		}

		stringTokenOauth2, err := getStringToken(token)
		if err != nil {
			log.Printf("Failed to get string token: %v", err)
		}

		_, err = h.DB.CreateSocialProfile(&models.SocialProfile{
			UserID:     userId,
			Platform:   constants.PlatformGoogle,
			ExternalID: userGoogleInfo.ID,
			ProfileURL: userGoogleInfo.Email,
			Params:     jsonData,
			Token:      stringTokenOauth2,
		})
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to create social profile: %s", err.Error()))
		}
	}

	sessionToken := uuid.New().String()

	_, err = h.DB.CreateToken(userId, sessionToken)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to create token: %s", err.Error()))
	}

	//todo вынести в конфиг путь до фронта
	return c.JSON(http.StatusOK, map[string]string{
		"redirectUrl":  "http://localhost:3000/main/?id=" + strconv.Itoa(int(userId)),
		"sessionToken": sessionToken,
		"userId":       strconv.FormatInt(userId, 10),
	})
}

func getStringToken(token *oauth2.Token) (string, error) {
	byteToken, err := json.Marshal(token)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal token: %s", err.Error())
	}
	return string(byteToken), nil
}

func getJsonQueryParams(c echo.Context) (string, error) {
	qp := QueryParams{
		Authuser: c.QueryParam("authuser"),
		Code:     c.QueryParam("code"),
		Prompt:   c.QueryParam("prompt"),
		Scope:    c.QueryParam("scope"),
		State:    c.QueryParam("state"),
		Token:    "",
	}

	jsonData, err := json.Marshal(qp)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func getUserInfo(client *http.Client) (*googleModels.User, error) {
	url := "https://www.googleapis.com/oauth2/v2/userinfo"
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request to Google API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	var userInfo googleModels.User
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return &userInfo, nil
}
