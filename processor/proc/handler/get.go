package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *HandlerImpl) GetProfile(c echo.Context) error {
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	user, err := h.DB.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *HandlerImpl) GetSettings(c echo.Context) error {
	cookieId, err := c.Cookie("userId")
	id, err := strconv.Atoi(cookieId.Value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	setting, err := h.DB.GetSetting(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, setting)
}

func (h *HandlerImpl) GetContactStats(c echo.Context) error {
	//todo добавить обработку ошибок. при ошибке отвечать ошибка сервера
	cookieId, _ := c.Cookie("userId")
	stringId := cookieId.Value
	id, _ := strconv.Atoi(stringId)
	contactCount, conflictCount, _ := h.DB.GetContactStats(id)

	return c.JSON(http.StatusOK, echo.Map{"conflicts": conflictCount, "new_contacts": contactCount})
}
