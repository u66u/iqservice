package handlers

import (
	"iq/db"
	"iq/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)
  
  func HandleCreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	newUser, err := db.CreateUser(user)
	if err != nil {
	  return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newUser)
  }
  
  func HandleUpdateUser(c echo.Context) error {
    user := new(models.User)
    if err := c.Bind(user); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    
    idParam := c.Param("id")
    userID, err := uuid.Parse(idParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid user ID")
    }

    updatedUser, err := db.UpdateUser(*user, userID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, updatedUser)
}
