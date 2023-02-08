package main

import (
	"os"

	"github.com/Edwinfpirajan/Distrifabrica.git/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	routes.EchoRoutes(e)

	port := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(port))

}
