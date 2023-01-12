package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Edwinfpirajan/Distrifabrica.git/common"
	"github.com/Edwinfpirajan/Distrifabrica.git/models"
)

func GetAllSchedule(writer http.ResponseWriter, request *http.Request) {
	schedule := []models.Horary{}
	db := common.GetConnection()
	db.Find(&schedule)
	json, _ := json.Marshal(schedule)
	common.SendResponse(writer, http.StatusOK, json)
}

// func SaveSchedule(w http.ResponseWriter, r *http.Request) {

// 	schedule := models.Horary{}

// 	db := common.GetConnection()

// 	error := json.NewDecoder(r.Body).Decode(&schedule)

// 	if error != nil {
// 		common.SendError(w, http.StatusBadRequest)
// 		return
// 	}

// 	db.Create(&schedule)
// 	fmt.Println("consulta: ", schedule)
// 	json, _ := json.Marshal(schedule)
// 	common.SendResponse(w, http.StatusOK, json)

// }

func SaveSchedule(w http.ResponseWriter, r *http.Request) {
	requestBody := map[string]string{}
	error := json.NewDecoder(r.Body).Decode(&requestBody)
	if error != nil {
		common.SendError(w, http.StatusBadRequest)
		return
	}

	arrival, ok := requestBody["arrival"]
	if !ok {
		common.SendError(w, http.StatusBadRequest)
		return
	}

	t, err := time.Parse("15:04", arrival)
	if err != nil {
		common.SendError(w, http.StatusBadRequest)
		return
	}

	now := time.Now()
	schedule := models.Horary{}
	schedule.Arrival = time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	departure, ok := requestBody["departure"]
	if !ok {
		common.SendError(w, http.StatusBadRequest)
		return
	}

	t, err = time.Parse("15:04", departure)
	if err != nil {
		common.SendError(w, http.StatusBadRequest)
		return
	}

	schedule.Departure = time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	db := common.GetConnection()
	db.Save(&schedule)
	json, _ := json.Marshal(schedule)
	common.SendResponse(w, http.StatusOK, json)
}

func UpdateSchedule(w http.ResponseWriter, r *http.Request) {
	schedule := models.Horary{}
	db := common.GetConnection()
	error := json.NewDecoder(r.Body).Decode(&schedule)
	if error != nil {
		common.SendError(w, http.StatusBadRequest)
		return
	}

	db.Save(&schedule)
	json, _ := json.Marshal(schedule)
	common.SendResponse(w, http.StatusOK, json)

}

func DeleteSchedule(w http.ResponseWriter, r *http.Request) {
	schedule := models.Horary{}
	db := common.GetConnection()
	error := json.NewDecoder(r.Body).Decode(&schedule)
	if error != nil {
		common.SendError(w, http.StatusBadRequest)
		return
	}

	db.Delete(&schedule)
	json, _ := json.Marshal(schedule)
	common.SendResponse(w, http.StatusOK, json)
}
