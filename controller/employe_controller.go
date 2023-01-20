package controller

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Edwinfpirajan/Distrifabrica.git/common"
	"github.com/Edwinfpirajan/Distrifabrica.git/models"
	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	employeesWithSchedule := []models.EmployeeWithSchedule{}
	db := common.GetConnection()
	db.Table("employes").Select("*").Joins("left join horaries h on h.id_sch = employes.schedule_id").Scan(&employeesWithSchedule)
	json, _ := json.Marshal(employeesWithSchedule)

	// t, _ := time.Parse("15:04", employeesWithSchedule[0].Horary.Arrival)
	// t.Add()
	common.SendResponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	employeeWithSchedule := models.EmployeeWithSchedule{}
	db := common.GetConnection()
	db.Table("employes").Select("*").Joins("left join horaries h on h.id_sch = employes.schedule_id").Where("employes.id = ?", id).First(&employeeWithSchedule)

	if employeeWithSchedule.ID == 0 {
		common.SendError(writer, http.StatusNotFound, fmt.Errorf(""))
		return
	}

	json, _ := json.Marshal(employeeWithSchedule)
	common.SendResponse(writer, http.StatusOK, json)
}

func Save(writer http.ResponseWriter, request *http.Request) {
	employe := models.Employe{}
	db := common.GetConnection()
	error := json.NewDecoder(request.Body).Decode(&employe)

	if error != nil {
		common.SendError(writer, http.StatusBadRequest, error)
		return
	}

	// Generate two random numbers
	num1 := make([]byte, 1)
	rand.Read(num1)
	num1[0] = num1[0]%10 + 48

	num2 := make([]byte, 1)
	rand.Read(num2)
	num2[0] = num2[0]%10 + 48

	// Generate two random uppercase letters
	letter1 := make([]byte, 1)
	rand.Read(letter1)
	letter1[0] = letter1[0]%26 + 65

	letter2 := make([]byte, 1)
	rand.Read(letter2)
	letter2[0] = letter2[0]%26 + 65

	// Concatenate the numbers and letters to create the pin
	pin := fmt.Sprintf("%c%c%d%d", letter1[0], letter2[0], num1[0]-48, num2[0]-48)

	employe.PinEmploye = pin

	if error != nil {
		common.SendError(writer, http.StatusInternalServerError, error)
		return
	}

	error = db.Save(&employe).Error
	if error != nil {
		common.SendError(writer, http.StatusInternalServerError, error)
		return
	}
	json, _ := json.Marshal(employe)
	common.SendResponse(writer, http.StatusCreated, json)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	employe := models.Employe{}
	db := common.GetConnection()
	id := mux.Vars(request)["id"]
	db.Find(&employe, id)

	if employe.ID > 0 {
		db.Delete(employe)
		common.SendResponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		common.SendError(writer, http.StatusNotFound, fmt.Errorf(""))
	}
}
