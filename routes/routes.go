package routes

import (
	"github.com/Edwinfpirajan/Distrifabrica.git/common"
	"github.com/Edwinfpirajan/Distrifabrica.git/controller"
	"github.com/gorilla/mux"
)

func SetRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api").Subrouter()
	router.HandleFunc("/login", controller.Login).Methods("POST")

	subRoute.HandleFunc("/all", common.AuthMiddleware(controller.GetAll)).Methods("GET")
	subRoute.HandleFunc("/save", common.AuthMiddleware(controller.Save)).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", common.AuthMiddleware(controller.Delete)).Methods("POST")
	subRoute.HandleFunc("/find/{id}", common.AuthMiddleware(controller.Get)).Methods("GET")

	subRoute.HandleFunc("/register", common.AuthMiddleware(controller.SaveRegisterttendance)).Methods("POST")
	subRoute.HandleFunc("/attendance", common.AuthMiddleware(controller.GetAllAttendance)).Methods("GET")
	subRoute.HandleFunc("/attendance/validate", common.AuthMiddleware(controller.ValidateHorary)).Methods("POST")

	subRoute.HandleFunc("/schedule/all", common.AuthMiddleware(controller.GetAllSchedule)).Methods("GET")
	subRoute.HandleFunc("/schedule/save", common.AuthMiddleware(controller.SaveSchedule)).Methods("POST")
	subRoute.HandleFunc("/schedule/delete/{id}", common.AuthMiddleware(controller.DeleteSchedule)).Methods("DELETE")
}
