package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Edwinfpirajan/Distrifabrica.git/common"
	"github.com/Edwinfpirajan/Distrifabrica.git/entity"
	"github.com/Edwinfpirajan/Distrifabrica.git/models"
)

func SaveRegisterttendance(w http.ResponseWriter, r *http.Request) {
	db := common.GetConnection()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var attendance entity.Attendance
	err = json.Unmarshal(body, &attendance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var employe models.Employe
	if err := db.Where("pin_employe = ?", attendance.PinEmployeFK).First(&employe).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var validateAttendance models.Attendances
	if err := db.Model(&validateAttendance).Where("pin_employe_fk = ? AND DATE(created_at) = CURDATE()", attendance.PinEmployeFK).Find(&validateAttendance).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if validateAttendance.ID == 0 {

		modelsAttendance := models.Attendances{
			PinEmployeFK: attendance.PinEmployeFK,
			Photo:        attendance.Photo,
		}

		err = db.Save(&modelsAttendance).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Registro creado exitosamente"))
		return
	}

	timeNow := time.Now()

	switch attendance.State {

	case "breakIn":
		validateAttendance.BreakIn = &timeNow
		break
	case "breakOut":
		validateAttendance.BreakOut = &timeNow
		break
	case "departure":
		validateAttendance.Departure = &timeNow
		break
	}

	err = db.Save(&validateAttendance).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Registro creado exitosamente"))
}

func GetAllAttendance(w http.ResponseWriter, r *http.Request) {
	attendance := []models.Attendances{}
	db := common.GetConnection()

	db.Find(&attendance)
	json, _ := json.Marshal(attendance)
	common.SendResponse(w, http.StatusOK, json)
	// fmt.Println(attendance)
}

// func GetAllAttendance(w http.ResponseWriter, r *http.Request) {
// 	attendance := []models.Attendances{}
// 	db := common.GetConnection()

// 	db.Find(&attendance)
// 	timeNow := time.Now()

// 	for i := range attendance {

// 		if attendance[i].Arrival == nil {
// 			attendance[i].Arrival = &timeNow
// 		}
// 		if attendance[i].BreakIn == nil {
// 			attendance[i].BreakIn = &timeNow
// 		}
// 		if attendance[i].BreakOut == nil {
// 			attendance[i].BreakOut = &timeNow
// 		}
// 		if attendance[i].Departure == nil {
// 			attendance[i].Departure = &timeNow

// 			arrivalTime, _ := time.Parse("03:04:05 PM", attendance[i].Arrival.Format("03:04:05 PM"))
// 			breakInTime, _ := time.Parse("03:04:05 PM", attendance[i].BreakIn.Format("03:04:05 PM"))
// 			breakOutTime, _ := time.Parse("03:04:05 PM", attendance[i].BreakOut.Format("03:04:05 PM"))
// 			departureTime, _ := time.Parse("03:04:05 PM", attendance[i].Departure.Format("03:04:05 PM"))

// 			attendance[i].Arrival = &arrivalTime
// 			attendance[i].BreakIn = &breakInTime
// 			attendance[i].BreakOut = &breakOutTime
// 			attendance[i].Departure = &departureTime
// 		}

// 		json, _ := json.Marshal(attendance)
// 		common.SendResponse(w, http.StatusOK, json)
// 	}

// }
