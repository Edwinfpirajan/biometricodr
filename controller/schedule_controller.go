package controller

import (
	"net/http"

	"github.com/Edwinfpirajan/Distrifabrica.git/common"
	"github.com/Edwinfpirajan/Distrifabrica.git/entity"
	"github.com/Edwinfpirajan/Distrifabrica.git/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func GetAllSchedule(c echo.Context) error {
	schedule := []models.Horary{}
	db := common.GetConnection()
	defer common.CloseDB(&db)
	db.Find(&schedule)
	return c.JSON(http.StatusOK, schedule)
}

//****************************************PARSER TIME

func saveSchedule(horary entity.Horary) (models.Horary, error) {

	schedule := models.Horary{}

	schedule.Id_sch = horary.Id_sch

	schedule.Arrival = horary.Arrival

	schedule.Departure = horary.Departure

	db := common.GetConnection()
	defer common.CloseDB(&db)
	err := db.Save(&schedule).Error
	if err != nil {
		return models.Horary{}, err
	}
	return schedule, nil
}

func SaveSchedule(c echo.Context) error {
	requestBody := entity.Horary{}
	if err := c.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	schedule, err := saveSchedule(requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, schedule)
}

func DeleteSchedule(c echo.Context) error {
	id := entity.Horary{}

	if err := c.Bind(&id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db := common.GetConnection()
	defer common.CloseDB(&db)

	var schedule models.Horary
	if err := db.First(&schedule, id.Id_sch).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Schedule not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := db.Delete(&schedule).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Schedule deleted")
}
