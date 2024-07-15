package main

import (
	"iq/db"
	"iq/handlers"
	"iq/routes"

	"github.com/labstack/echo/v4"
)

func main() {
  e := echo.New()
  e.GET("/", routes.Home)
  db.InitDB()
  e.POST("/users", handlers.HandleCreateUser)
  e.PUT("/users/:id", handlers.HandleUpdateUser)
  e.Logger.Fatal(e.Start(":8080"))
}
