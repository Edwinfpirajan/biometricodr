package main

import (
	"github.com/Edwinfpirajan/Distrifabrica.git/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	// common.Migrate()
	// controller.Query(common.GetConnection(), 3)
	// router := mux.NewRouter()
	// routes.SetRoutes(router)
	// common.Login("", "")

	e := echo.New()
	routes.EchoRoutes(e)

	e.Logger.Fatal(e.Start(":3001"))

	// server := http.Server{
	// 	Addr:    ":3001",
	// 	Handler: router,
	// }

	// log.Println("Servidor ejecutandose: ", server.Addr)
	// log.Println(server.ListenAndServe())

}
