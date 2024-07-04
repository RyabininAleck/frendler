package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"frendler/common/constants"
	"frendler/processor/models"
)

func (h *HandlerImpl) CreateUserByEmail(c echo.Context) error {
	return c.String(http.StatusOK, "CreateUserByEmail")
}

/*
body

	{
	    "username": "exampleUser",
	    "email": "example@example.com",
	    "password": "examplePassword",
	    "role": "user",
	    "status": "active",
	    "phone_number": "1234567890"
	}
*/
func (h *HandlerImpl) CreateUserByNumber(c echo.Context) error {
	user := models.User{}

	err := c.Bind(user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	//todo проверить номер на действительность

	userExists := h.userExistsByNumber(user.PhoneNumber)
	if userExists == true {
		return c.String(http.StatusBadRequest, "User already exists")
	}

	err = h.DB.CreateUserSetting(
		models.User{
			Username:    user.Username,
			Email:       user.Email,
			Password:    user.Password,
			Role:        user.Role,
			Status:      user.Status,
			PhoneNumber: user.PhoneNumber,
		},
		models.Setting{
			Theme:      constants.ThemeLight,
			Language:   "ru",
			AutoUpdate: false,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())

	}

	//todo добавить настройки
	return c.String(http.StatusOK, "CreateUserByNumber")
}

func (h *HandlerImpl) CreateUserByVk(c echo.Context) error {
	return c.String(http.StatusOK, "CreateUserByVk!")
}
