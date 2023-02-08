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

// func validateAttendance(pinEmploye string, state string) error {
// 	db := common.GetConnection()
// 	var validateAttendance models.Attendances
// 	if err := db.Model(&validateAttendance).Where("pin_employe_fk = ? AND DATE(created_at) = CURDATE()", pinEmploye).Find(&validateAttendance).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, "Register")
// 	}

// 	if validateAttendance.ID == 0 && state != "arrival" {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Debe registrar el estado 'arrival' primero")
// 	}

// 	if state == "arrival" && validateAttendance.Arrival != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado el estado 'arrival'")
// 	}

// 	if state == "breakIn" && validateAttendance.BreakIn != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado el estado 'breakIn'")
// 	}

// 	if state == "breakOut" && validateAttendance.BreakOut != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado el estado 'breakOut'")
// 	}

//		if state == "departure" && validateAttendance.Departure != nil {
//			return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado el estado 'departure'")
//		}
//		return nil
//	}
// func ValidateAttendance(c echo.Context) error {
// 	id := c.Param("pin")

// 	if id == "" {
// 		return echo.NewHTTPError(http.StatusBadRequest, errors.New("El id es necesario"))
// 	}

// 	db := common.GetConnection()
// 	var validateAttendance models.Attendances
// 	if err := db.Model(&validateAttendance).Where("pin_employe_fk = ? AND DATE(created_at) = CURDATE()", id).Find(&validateAttendance).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, err.Error())
// 	}

// 	if validateAttendance.BreakIn == nil && validateAttendance.Arrival == nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Tienes que marcar tu llegada primero")
// 	} else if validateAttendance.BreakIn != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado el estado 'breakIn")
// 	}

// 	return nil
// }

func SaveRegisterAttendance(c echo.Context) error {
	db := common.GetConnection()
	var attendance entity.Attendance
	err := c.Bind(&attendance)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var validateAttendance models.Attendances
	if err := db.Model(&validateAttendance).Where("pin_employe_fk = ? AND DATE(created_at) = CURDATE()", attendance.PinEmployeFK).Find(&validateAttendance).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Empleado no se encuentra registrado")
	}
	timeNow := time.Now()

	if attendance.State == "arrival" {
		if validateAttendance.ID == 0 {
			modelsAttendance := models.Attendances{
				PinEmployeFK: attendance.PinEmployeFK,
				Photo:        attendance.Photo,
				Arrival:      &timeNow,
			}

			err = db.Save(&modelsAttendance).Error
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return c.JSON(http.StatusOK, map[string]string{
				"message": "Registro creado exitosamente",
			})
		}
	}

	block := validateAttendance.Arrival == nil

	switch attendance.State {
	case "arrival":
		if validateAttendance.Arrival != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado el estado 'arrival'")
		}
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

	err = db.Save(&validateAttendance).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Registro actualizado exitosamente",
	})
}

func GetAllAttendance(c echo.Context) error {
	db := common.GetConnection()
	attendance := []models.GetAllAttendances{}

	db.Table("attendances").Select("*").Joins("INNER JOIN employes e on e.pin_employe = attendances.pin_employe_fk").Find(&attendance)

	return c.JSON(http.StatusOK, attendance)
}

//VALIDATIONS

func ValidateHorary(c echo.Context) error {
	db := common.GetConnection()

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

	db.Raw("select arrival from attendances a where pin_employe_fk = ? and date_format(arrival, '%d-%m-%Y') = date_format(?, '%d-%m-%Y')",
		validateHorary.PinEmployeFK, validateHorary.Date).Scan(&arrival)

	return c.JSON(http.StatusOK, arrival)

}

func ValidateEmploye(c echo.Context) error {
	id := c.Param("pin")
	db := common.GetConnection()

	var employe models.Employe
	if err := db.Table("employes").Where("pin_employe = ?", id).Scan(&employe).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if employe.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Empleado no se encuentra registrado")
	}
	return c.JSON(http.StatusOK, employe)

}
