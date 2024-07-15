package handler

import (
	"log"
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
	cookieId, err := c.Cookie("userId")
	if err != nil {
		log.Printf("Server failed: %v", err)
	}
	stringId := cookieId.Value
	id, err := strconv.Atoi(stringId)
	if err != nil {
		log.Printf("Failed to formatting param: %v", err)
	}
	contactCount, conflictCount, err := h.DB.GetContactStats(id)
	if err != nil {
		log.Printf("Failed to get contact stats: %v", err)
	}

	return c.JSON(http.StatusOK, echo.Map{"conflicts": conflictCount, "new_contacts": contactCount})
}
