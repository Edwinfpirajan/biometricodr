package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Edwinfpirajan/Distrifabrica.git/common"
	"github.com/Edwinfpirajan/Distrifabrica.git/entity"
	"github.com/Edwinfpirajan/Distrifabrica.git/models"
	"github.com/labstack/echo/v4"
)

func SaveRegisterAttendance(c echo.Context) error {
	var attendance entity.Attendance
	err := c.Bind(&attendance)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var validateAttendance models.Attendances
	if err := common.DB.Model(&validateAttendance).Where("pin_employe_fk = ? AND DATE(created_at) = CURDATE()", attendance.PinEmployeFK).Scan(&validateAttendance).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Empleado no se encuentra registrado")
	}
	location, _ := time.LoadLocation("America/Bogota")
	timeNow := time.Now().In(location)

	if attendance.State == "arrival" {
		if validateAttendance.ID == 0 && validateAttendance.Arrival == nil {
			modelsAttendance := models.Attendances{
				PinEmployeFK: attendance.PinEmployeFK,
				Photo:        attendance.Photo,
				Arrival:      &timeNow,
			}

			err = common.DB.Save(&modelsAttendance).Error
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return c.JSON(http.StatusOK, map[string]string{
				"message": "Registro creado exitosamente",
			})
		}
	}

	block := validateAttendance.Arrival == nil
	blockBreakInit := validateAttendance.BreakInit == nil
	blockBreakInitTwo := validateAttendance.BreakInit == nil
	blockBreakIn := validateAttendance.BreakIn == nil

	switch attendance.State {
	case "arrival":
		if validateAttendance.Arrival != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado el estado entrada")
		}
		break
	case "breakInit":
		if block {
			return echo.NewHTTPError(http.StatusBadRequest, "Debe registrar la llegada primero")
		}
		if validateAttendance.BreakInit != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado la salida a pausa")
		}
		validateAttendance.BreakInit = &timeNow
		break
	case "breakEnd":
		if block {
			return echo.NewHTTPError(http.StatusBadRequest, "Debe registrar la llegada primero")
		}
		if blockBreakInit {
			return echo.NewHTTPError(http.StatusBadRequest, "Debe registrar la salida a pausa primero")
		}
		if validateAttendance.BreakEnd != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado la entrada a la pausa")
		}
		validateAttendance.BreakEnd = &timeNow
		break
	case "breakInitTwo":
		if block {
			return echo.NewHTTPError(http.StatusBadRequest, "Debe registrar la llegada primero")
		}
		if validateAttendance.BreakInitTwo != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado la salida a pausa")
		}
		validateAttendance.BreakInitTwo = &timeNow
		break
	case "breakEndTwo":
		if block {
			return echo.NewHTTPError(http.StatusBadRequest, "Debe registrar la llegada primero")
		}
		if blockBreakInitTwo {
			return echo.NewHTTPError(http.StatusBadRequest, "Debe registrar la salida a pausa primero")
		}
		if validateAttendance.BreakEndTwo != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado la entrada a la pausa")
		}
		validateAttendance.BreakEndTwo = &timeNow
		break
	case "breakIn":
		if block {
			return echo.NewHTTPError(http.StatusBadRequest, "Debe registrar la llegada primero")
		}
		if validateAttendance.BreakIn != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado la salida a almuerzo")
		}
		validateAttendance.BreakIn = &timeNow
		break
	case "breakOut":
		if block {
			return echo.NewHTTPError(http.StatusBadRequest, "Debe registrar la llegada primero")
		}
		if blockBreakIn {
			return echo.NewHTTPError(http.StatusBadRequest, "Debe registrar la salida a almuerzo primero")
		}
		if validateAttendance.BreakOut != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado la entrada de almuerzo")
		}
		validateAttendance.BreakOut = &timeNow
		break
	case "departure":
		if block {
			return echo.NewHTTPError(http.StatusBadRequest, "Debe registrar la llegada primero")
		}
		if validateAttendance.Departure != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado el estado salida")
		}
		validateAttendance.Departure = &timeNow
		break
	}

	err = common.DB.Save(&validateAttendance).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Registro actualizado exitosamente",
	})
}

func GetAllAttendance(c echo.Context) error {

	attendance := []models.GetAllAttendances{}

	common.DB.Table("attendances a").Select("e.first_name, e.last_name, a.* ").Joins("INNER JOIN employes e on e.pin_employe = a.pin_employe_fk").Find(&attendance)

	return c.JSON(http.StatusOK, attendance)
}

//VALIDATIONS

func ValidateHorary(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var validateHorary entity.ValidateHorary
	err = json.Unmarshal(body, &validateHorary)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var arrival time.Time

	common.DB.Raw("select arrival from attendances a where pin_employe_fk = ? and date_format(arrival, '%d-%m-%Y') = date_format(?, '%d-%m-%Y')",
		validateHorary.PinEmployeFK, validateHorary.Date).Scan(&arrival)

	return c.JSON(http.StatusOK, arrival)
}

func ValidateEmploye(c echo.Context) error {
	id := c.Param("pin")

	var employe models.Employe
	if err := common.DB.Table("employes").Where("pin_employe = ?", id).Scan(&employe).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if employe.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Empleado no se encuentra registrado")
	}
	return c.JSON(http.StatusOK, employe)

}
