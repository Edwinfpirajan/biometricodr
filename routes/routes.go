package routes

import (
	"github.com/Edwinfpirajan/Distrifabrica.git/common"
	"github.com/Edwinfpirajan/Distrifabrica.git/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func EchoRoutes(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://distriramirez.com.co/"},
		// AllowMethods: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		// AllowHeaders: []string{"*"},
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	group := e.Group("/api")

	// LOGIN

	e.POST("/login", controller.Login)
	e.POST("/logout", controller.Logout, common.OnlyAdmin)

	// USER ROUTES

	//ATTENDANCE REGISTER

	group.POST("/attendance/register", controller.SaveRegisterAttendance)
	group.GET("/attendance", controller.GetAllAttendance)
	group.GET("/attendance/validate/:pin", controller.ValidateEmploye)

	// ADMIN ROUTES

	//EMPLOYE MANAGE

	group.GET("/all", controller.GetAll)
	group.POST("/save", controller.SaveEmploye, common.OnlyAdmin)
	group.DELETE("/delete/:id", controller.DeleteEmploye, common.OnlyAdmin)
	group.GET("/find/:id", controller.GetEmploye, common.OnlyAdmin)

	// SCHEDULE MANAGE

	group.GET("/schedule/all", controller.GetAllSchedule, common.OnlyAdmin)
	group.POST("/schedule/save", controller.SaveSchedule, common.OnlyAdmin)
	group.DELETE("/schedule/delete/:id", controller.DeleteSchedule, common.OnlyAdmin)
}
