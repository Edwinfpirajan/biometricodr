package controller

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Edwinfpirajan/Distrifabrica.git/common"
	"github.com/Edwinfpirajan/Distrifabrica.git/models"
	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	employeesWithSchedule := []models.EmployeeWithSchedule{}
	db := common.GetConnection()
	// db.Raw("SELECT * FROM employes e INNER JOIN horaries h ON h.id = e.schedule_id").Find(&employeesWithSchedule)
	db.Table("employes").Select("*").Joins("left join horaries h on h.id_sch = employes.schedule_id").Scan(&employeesWithSchedule)
	fmt.Println("consulta: ", employeesWithSchedule)
	json, _ := json.Marshal(employeesWithSchedule)
	common.SendResponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	employeeWithSchedule := models.EmployeeWithSchedule{}
	db := common.GetConnection()
	db.Table("employes").Select("*").Joins("left join horaries h on h.id_sch = employes.schedule_id").Where("employes.id = ?", id).First(&employeeWithSchedule)

	if employeeWithSchedule.ID == 0 {
		common.SendError(writer, http.StatusNotFound)
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
		log.Fatal(error)
		common.SendError(writer, http.StatusBadRequest)
		return
	}

	randomBytes := make([]byte, 2)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal(error)
		common.SendError(writer, http.StatusBadRequest)
		return
	}

	employe.PinEmploye = hex.EncodeToString(randomBytes)
	// horary := models.Employe{Arrival: employe.Arrival, Departure: employe.Departure}
	// error = db.Save(&horary).Error
	if error != nil {
		log.Fatal(error)
		common.SendError(writer, http.StatusInternalServerError)
		return
	}

	// employe.ScheduleId = horary.ID
	error = db.Save(&employe).Error
	if error != nil {
		log.Fatal(error)
		common.SendError(writer, http.StatusInternalServerError)
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
		common.SendError(writer, http.StatusNotFound)
	}
}
