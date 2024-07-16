package handlers

import (
	"iq/db"
	"iq/domain"
	"iq/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func HandleCreateUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	hashedPassword, err := domain.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error processing password"})
	}
	user.Password = hashedPassword

	newUser, err := db.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	newUser.Password = ""
	return c.JSON(http.StatusCreated, newUser)
}

func HandleUpdateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	idParam := c.Param("id")
	userID, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// If a new password is provided, hash it
	if user.Password != "" {
		hashedPassword, err := domain.HashPassword(user.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error processing password"})
		}
		user.Password = hashedPassword
	}

	updatedUser, err := db.UpdateUser(*user, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	updatedUser.Password = ""
	return c.JSON(http.StatusOK, updatedUser)
}
