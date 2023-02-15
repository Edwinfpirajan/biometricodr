package controller

import (
	"net/http"

	"github.com/Edwinfpirajan/Distrifabrica.git/common"
	"github.com/Edwinfpirajan/Distrifabrica.git/models"
	"github.com/labstack/echo/v4"
)

func GetAttendenceData(c echo.Context) error {
	db := common.GetConnection()
	attendances := models.GetAllAttendances{}

	db.Raw("SELECT COUNT(id) FROM attendances").Find(&attendances)

	return c.JSON(http.StatusOK, attendances)
}
