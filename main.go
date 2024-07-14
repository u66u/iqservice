package main

import (
	"iq/routes"

	"github.com/labstack/echo/v4"
)

func main() {
  e := echo.New()
  e.GET("/", routes.Home)
  e.Logger.Fatal(e.Start(":8080"))
}
