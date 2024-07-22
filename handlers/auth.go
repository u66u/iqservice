package handlers

import (
	"iq/db"
	"iq/domain"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func HandleLogin(c echo.Context) error {
    var loginRequest struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.Bind(&loginRequest); err != nil {
        c.Logger().Errorf("Error binding request: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }

    user, err := db.SelectUser(loginRequest.Email, loginRequest.Password)
    if err != nil {
        c.Logger().Errorf("Invalid credentials: %v", err)
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
    }

    token, err := domain.GenerateToken(user)
    if err != nil {
        c.Logger().Errorf("Token generation failure: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not generate token"})
    }

    cookie := new(http.Cookie)
    cookie.Name = "token"
    cookie.Value = token
    cookie.Expires = time.Now().Add(24 * time.Hour)
    cookie.Path = "/"
    cookie.HttpOnly = true
    cookie.Secure = false // Set to true if using HTTPS
    c.SetCookie(cookie)

    return c.JSON(http.StatusOK, map[string]string{"message": "Successfully logged in"})
}


func HandleProtected(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
