package controller

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/Edwinfpirajan/Distrifabrica.git/common"
	"github.com/Edwinfpirajan/Distrifabrica.git/entity"
	"github.com/Edwinfpirajan/Distrifabrica.git/models"
	"github.com/labstack/echo/v4"
)

func GetAll(c echo.Context) error {
	employeesWithSchedule := []models.EmployeeWithSchedule{}
	db := common.GetConnection()
	db.Table("employes").Select("*").Joins("left join horaries h on h.id_sch = employes.schedule_id").Scan(&employeesWithSchedule)
	return c.JSON(http.StatusOK, employeesWithSchedule)
}

func GetEmploye(c echo.Context) error {
	id := c.Param("id")

	employeeWithSchedule := models.EmployeeWithSchedule{}
	db := common.GetConnection()

	db.Table("employes").Select("*").Joins("left join horaries h on h.id_sch = employes.schedule_id").Where("employes.id = ?", id).First(&employeeWithSchedule)

	if employeeWithSchedule.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Employee not found")
	}

	return c.JSON(http.StatusOK, employeeWithSchedule)
}

func SaveEmploye(c echo.Context) error {
	employe := entity.Employe{}
	db := common.GetConnection()
	err := c.Bind(&employe)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	employeFromDb := models.Employe{}

	db.Table("employes").Where("employes.pin_employe = ?", employe.PinEmploye).Scan(&employeFromDb)

	if employeFromDb.ID == 0 {

		num1 := make([]byte, 1)
		rand.Read(num1)
		num1[0] = num1[0]%10 + 48

		num2 := make([]byte, 1)
		rand.Read(num2)
		num2[0] = num2[0]%10 + 48

		letter1 := make([]byte, 1)
		rand.Read(letter1)
		letter1[0] = letter1[0]%26 + 65

		letter2 := make([]byte, 1)
		rand.Read(letter2)
		letter2[0] = letter2[0]%26 + 65

		pin := fmt.Sprintf("%c%c%d%d", letter1[0], letter2[0], num1[0]-48, num2[0]-48)

		employeFromDb.PinEmploye = pin
		employeFromDb.FirstName = employe.FirstName
		employeFromDb.LastName = employe.LastName
		employeFromDb.Company = employe.LastName
		employeFromDb.Position = employe.LastName
		employeFromDb.ScheduleId = employe.ScheduleId

		db.Save(&employeFromDb)

	} else {
		employeFromDb.FirstName = employe.FirstName
		employeFromDb.LastName = employe.LastName
		employeFromDb.Company = employe.LastName
		employeFromDb.Position = employe.LastName
		employeFromDb.ScheduleId = employe.ScheduleId

		db.Save(&employeFromDb)
	}

	return c.JSON(http.StatusCreated, employe)
}

// employe := models.Employe{}
// db := common.GetConnection()
// error := json.NewDecoder(request.Body).Decode(&employe)

// if error != nil {
// 	common.SendError(writer, http.StatusBadRequest, error)
// 	return
// }

// // Generate two random numbers
// num1 := make([]byte, 1)
// rand.Read(num1)
// num1[0] = num1[0]%10 + 48

// num2 := make([]byte, 1)
// rand.Read(num2)
// num2[0] = num2[0]%10 + 48

// // Generate two random uppercase letters
// letter1 := make([]byte, 1)
// rand.Read(letter1)
// letter1[0] = letter1[0]%26 + 65

// letter2 := make([]byte, 1)
// rand.Read(letter2)
// letter2[0] = letter2[0]%26 + 65

// // Concatenate the numbers and letters to create the pin
// pin := fmt.Sprintf("%c%c%d%d", letter1[0], letter2[0], num1[0]-48, num2[0]-48)

// employe.PinEmploye = pin

// if error != nil {
// 	common.SendError(writer, http.StatusInternalServerError, error)
// 	return
// }

// error = db.Save(&employe).Error
// if error != nil {
// 	common.SendError(writer, http.StatusInternalServerError, error)
// 	return
// }
// json, _ := json.Marshal(employe)
// common.SendResponse(writer, http.StatusCreated, json)

func DeleteEmploye(c echo.Context) error {
	id := c.Param("id")

	employee := models.Employe{}
	db := common.GetConnection()
	db.Find(&employee, id)

	if employee.ID > 0 {
		db.Delete(employee)
		return c.JSON(http.StatusOK, employee)
	} else {
		return echo.NewHTTPError(http.StatusNotFound, "Employee not found")
	}
}

// employe := models.Employe{}
// db := common.GetConnection()
// id := mux.Vars(request)["id"]
// db.Find(&employe, id)

// if employe.ID > 0 {
// 	db.Delete(employe)
// 	common.SendResponse(writer, http.StatusOK, []byte(`{}`))
// } else {
// 	common.SendError(writer, http.StatusNotFound, fmt.Errorf(""))
// }
