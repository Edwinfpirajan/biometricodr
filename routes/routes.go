package routes

import (
	"github.com/Edwinfpirajan/Distrifabrica.git/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/* func SetRoutes(router *mux.Router) {
	common.EnableCORS(router)
	subRoute := router.PathPrefix("/api").Subrouter()
	router.HandleFunc("/login", controller.Login).Methods("POST")

	subRoute.HandleFunc("/all", common.AuthMiddleware(controller.GetAll)).Methods("GET")
	subRoute.HandleFunc("/save", common.AuthMiddleware(controller.Save)).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", common.AuthMiddleware(controller.Delete)).Methods("POST")
	subRoute.HandleFunc("/find/{id}", common.AuthMiddleware(controller.Get)).Methods("GET")

	subRoute.HandleFunc("/register", common.AuthMiddleware(controller.SaveRegisterttendance)).Methods("POST")
	subRoute.HandleFunc("/attendance", common.AuthMiddleware(controller.GetAllAttendance)).Methods("GET")
	subRoute.HandleFunc("/attendance/validate", common.AuthMiddleware(controller.ValidateHorary)).Methods("POST")

	// subRoute.HandleFunc("/schedule/all", common.AuthMiddleware(controller.GetAllSchedule)).Methods("GET")
	subRoute.HandleFunc("/schedule/save", common.AuthMiddleware(controller.SaveSchedule)).Methods("POST")
	subRoute.HandleFunc("/schedule/delete/{id}", common.AuthMiddleware(controller.DeleteSchedule)).Methods("DELETE")
} */

func EchoRoutes(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	group := e.Group("/api")

	e.POST("/login", controller.Login)

	group.GET("/all", controller.GetAll /* common.OnlyAdmin */)
	group.POST("/save", controller.SaveEmploye /* common.OnlyAdmin */)
	group.DELETE("/delete/:id", controller.DeleteEmploye /* common.OnlyAdmin */)
	group.GET("/find/:id", controller.GetEmploye /* common.OnlyAdmin */)

	// e.POST("/register", common.AuthMiddleware(controller.SaveRegisterttendance))
	// e.GET("/attendance", common.AuthMiddleware(controller.GetAllAttendance))
	// e.POST("/attendance/validate", common.AuthMiddleware(controller.ValidateHorary))

	group.GET("/schedule/all", controller.GetAllSchedule /* common.OnlyAdmin */)
	group.POST("/schedule/save", controller.SaveSchedule /* common.OnlyAdmin */)
	group.DELETE("/schedule/delete/:id", controller.DeleteSchedule)
}
