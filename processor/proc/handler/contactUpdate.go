package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

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
	code, err := GetGoogleCode(socialProfile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to get authorization code from database"})
	}

	token, err := config.GoogleOauth.Exchange(c.Request().Context(), code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Failed to exchange token: %s", err.Error())})
	}

	client := config.GoogleOauth.Client(c.Request().Context(), token)

	contacts, err := getGoogleContacts(client)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Failed to get contacts: %s", err.Error())})
	}

	// Сохранение контактов в базу данных
	IDs, err := h.DB.CreateFriends(userId, contacts, constants.PlatformGoogle)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Failed to save contacts: %s", err.Error())})
	}

	return c.JSON(http.StatusOK, IDs)
}

type paramsCode struct {
	Code string `json:"code"`
}

func GetGoogleCode(socialProfile *models.SocialProfile) (string, error) {
	var params paramsCode

	err := json.Unmarshal([]byte(socialProfile.Params), &params)
	if err != nil {
		return "", err
	}

	return params.Code, nil

}

// Вспомогательная функция для запроса контактов пользователя
func getGoogleContacts(client *http.Client) ([]*models.Friend, error) {

	resp, err := client.Get("https://www.google.com/m8/feeds/contacts/default/full?alt=json&max-results=1000")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var contacts []googleModels.User
	if err := json.NewDecoder(resp.Body).Decode(&contacts); err != nil {
		return nil, err
	}

	// todo преобразовать contacts в []models.Friend
	return []*models.Friend{}, nil

}
