package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

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

func (h *HandlerImpl) HandleGoogleCallback(c echo.Context) error {
	state := c.QueryParam("state")
	if state != config.OauthStateString {
		log.Println("Invalid OAuth state")
		return c.Redirect(http.StatusTemporaryRedirect, "/error")
	}

	code := c.QueryParam("code")
	token, err := config.GoogleOauth.Exchange(c.Request().Context(), code)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to exchange token: %s", err.Error()))
	}

	client := config.GoogleOauth.Client(c.Request().Context(), token)
	userInfo, err := getUserInfo(client)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to get user info: %s", err.Error()))
	}

	user, err := h.DB.GetUserByExternalId(userInfo.ID, constants.PlatformGoogle)
	if err != nil {
		log.Printf("Database error: %v\n", err)
		return c.Redirect(http.StatusTemporaryRedirect, "/error")
	}

	var userId int64

	if user != nil {
		err := h.DB.UpdateUserLogin(user.UserID)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to updete user info: %s", err.Error()))
		}
	} else {
		userId, _, err = h.DB.CreateUserAndSetting(
			&models.User{
				Username:    userInfo.Email,
				Email:       userInfo.Email,
				Password:    "",
				Role:        constants.RoleUser,
				Status:      constants.StatusActive,
				AvatarURL:   userInfo.Picture,
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

		jsonData, _ := json.Marshal(c.QueryParams())

		_, err = h.DB.CreateSocialProfile(&models.SocialProfile{
			UserID:     userId,
			Platform:   constants.PlatformGoogle,
			ExternalID: userInfo.ID,
			ProfileURL: userInfo.Email,
			Params:     string(jsonData),
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
