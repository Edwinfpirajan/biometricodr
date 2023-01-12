package routes

import (
	"github.com/Edwinfpirajan/Distrifabrica.git/api"
	"github.com/Edwinfpirajan/Distrifabrica.git/controller"
	"github.com/gorilla/mux"
)

func SetRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api").Subrouter()

	// EMPLOYE CONTROLLER
	subRoute.HandleFunc("/all", controller.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controller.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controller.Delete).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controller.Get).Methods("GET")

	// ATTENDANCE CONTROLLER
	subRoute.HandleFunc("/register", controller.SaveRegisterttendance).Methods("POST")
	subRoute.HandleFunc("/attendance", controller.GetAllAttendance).Methods("GET")
	subRoute.HandleFunc("/employes/pin/{pin}", controller.GetAllAttendance).Methods("GET")

	// SCHEDULE CONTROLLER

	subRoute.HandleFunc("/schedule", controller.GetAllSchedule).Methods("GET")
	subRoute.HandleFunc("/schedule/save", controller.SaveSchedule).Methods("POST")
	subRoute.HandleFunc("/schedule/delete/{id}", controller.DeleteSchedule).Methods("POST")

	//LOGIN CONTROLLER
	subRoute.HandleFunc("/login", api.Logeo).Methods("POST")
}
