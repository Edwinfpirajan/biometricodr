package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Edwinfpirajan/Distrifabrica.git/common"
	"github.com/Edwinfpirajan/Distrifabrica.git/models"
)

func SaveRegisterttendance(w http.ResponseWriter, r *http.Request) {
	db := common.GetConnection()
	defer db.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var attendance models.Attendance
	err = json.Unmarshal(body, &attendance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// PIN VALIDATE
	var employe models.Employe
	if err := db.Where("pin_employe = ?", attendance.PinEmploye).First(&employe).Error; err != nil {
		http.Error(w, "Pin del empleado no encontrado", http.StatusBadRequest)
		return
	}

	form := models.Attendance{
		PinEmploye: attendance.PinEmploye,
		State:      attendance.State,
		Photo:      attendance.Photo,
	}

	err = db.Create(&form).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Registro creado exitosamente"))
}

func GetAllAttendance(w http.ResponseWriter, r *http.Request) {
	attendance := []models.Attendance{}
	db := common.GetConnection()
	defer db.Close()

	db.Find(&attendance)
	json, _ := json.Marshal(attendance)
	common.SendResponse(w, http.StatusOK, json)
}
