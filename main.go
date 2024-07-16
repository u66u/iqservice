package main

import (
	"errors"
	"iq/db"
	"iq/handlers"
	"iq/routes"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		log.Fatal("JWT_SECRET_KEY is not set in the environment")
	}
  

	e := echo.New()

	e.Use(handlers.LogRequest)
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	db.InitDB()

	e.GET("/", routes.Home)
	e.POST("/users/create", handlers.HandleCreateUser)
	e.PUT("/users/:id", handlers.HandleUpdateUser)
	e.POST("/login", handlers.HandleLogin)

	r := e.Group("/api")
	r.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(jwtSecretKey),
		TokenLookup: "cookie:token",
	}))

	r.GET("/protected", handlers.HandleProtected)
	r.GET("/test", func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
		if !ok {
			return errors.New("JWT token missing or invalid")
		}
		claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
		if !ok {
			return errors.New("failed to cast claims as jwt.MapClaims")
		}
		return c.JSON(http.StatusOK, claims)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
