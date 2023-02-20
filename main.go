package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Edwinfpirajan/Distrifabrica.git/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	routes.EchoRoutes(e)

	// e.Logger.Fatal(e.Start(":8080"))

	PORT, _ := strconv.Atoi(os.Getenv("PORT"))

	if PORT == 0 {
		PORT = 8080
	}
	HOST := os.Getenv("SERVER_HOST")

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", HOST, PORT)))
}

//PORT="8080" HOST="0.0.0.0" go run .
