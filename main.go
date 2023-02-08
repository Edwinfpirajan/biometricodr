package main

import (
	"github.com/Edwinfpirajan/Distrifabrica.git/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	routes.EchoRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))

}
