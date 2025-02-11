package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/skip2/go-qrcode"
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
		log.Printf("Server could not get the userid: %v", err)
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

func (h *HandlerImpl) GetQRCode(c echo.Context) error {
	cookieId, err := c.Cookie("userId")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to formatting param: " + err.Error()})
	}

	id, err := strconv.Atoi(cookieId.Value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "\"Failed to formatting param: " + err.Error()})
	}

	user, err := h.DB.GetUserById(id)
	if err != nil {
		log.Printf("Failed to formatting param: %v", err)
	}

	vCard := generateVCard(user.FirstName+""+user.LastName, user.Gender, user.PhoneNumber, user.Email)

	png, err := qrcode.Encode(vCard, qrcode.Medium, 256)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to generate QR code")
	}

	return c.Blob(http.StatusOK, "image/png", png)
}

func generateVCard(name, organization, phone, email string) string {
	vCard := fmt.Sprintf(
		"BEGIN:VCARD\nVERSION:3.0\nFN:%s\nORG:%s\nTEL;TYPE=WORK,VOICE:%s\nEMAIL:%s\nEND:VCARD",
		name, organization, phone, email,
	)
	return vCard
}
