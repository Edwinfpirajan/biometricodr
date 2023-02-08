package main

import (
	"fmt"
	"os"

	"github.com/Edwinfpirajan/Distrifabrica.git/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	routes.EchoRoutes(e)

	PORT := os.Getenv("PORT")
	HOST := os.Getenv("SERVER_HOST")

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", HOST, PORT)))

}
