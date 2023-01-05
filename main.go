package main

import (
	"log"
	"net/http"

	"github.com/Edwinfpirajan/Distrifabrica.git/common"
	"github.com/Edwinfpirajan/Distrifabrica.git/routes"
	"github.com/gorilla/mux"
)

func main() {
	// common.Migrate()
	// controller.Query(common.GetConnection(), 3)
	router := mux.NewRouter()
	routes.SetRoutes(router)
	common.EnableCORS(router)
	// common.Login("", "")

	server := http.Server{
		Addr:    ":3001",
		Handler: router,
	}

	log.Println("Servidor ejecutandose: ", server.Addr)
	log.Println(server.ListenAndServe())

}
