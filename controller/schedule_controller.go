package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Edwinfpirajan/Distrifabrica.git/common"
	"github.com/Edwinfpirajan/Distrifabrica.git/entity"
	"github.com/Edwinfpirajan/Distrifabrica.git/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllSchedule(writer http.ResponseWriter, request *http.Request) {
	schedule := []models.Horary{}
	db := common.GetConnection()
	db.Find(&schedule)
	json, _ := json.Marshal(schedule)
	common.SendResponse(writer, http.StatusOK, json)
	// fmt.Println(schedule)
}

func GetScheduleById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var schedule models.Horary
	db := common.GetConnection()

	db.First(&schedule, id)

	if schedule.Id_sch == 0 {
		common.SendError(w, http.StatusNotFound, fmt.Errorf(""))
		return
	}

	json, _ := json.Marshal(schedule)
	common.SendResponse(w, http.StatusOK, json)
}

//****************************************PARSER TIME

func saveSchedule(horary entity.Horary) (models.Horary, error) {

	schedule := models.Horary{}

	schedule.Id_sch = horary.Id_sch

	schedule.Arrival = horary.Arrival

	schedule.Departure = horary.Departure

	db := common.GetConnection()
	db.Save(&schedule)
	return schedule, nil
}

func SaveSchedule(w http.ResponseWriter, r *http.Request) {
	requestBody := entity.Horary{}
	error := json.NewDecoder(r.Body).Decode(&requestBody)
	if error != nil {
		common.SendError(w, http.StatusBadRequest, error.Error())
		return
	}

	schedule, err := saveSchedule(requestBody)
	if err != nil {
		common.SendError(w, http.StatusBadRequest, error)
		return
	}
	json, _ := json.Marshal(schedule)
	common.SendResponse(w, http.StatusOK, json)
}

func DeleteSchedule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	db := common.GetConnection()

	var schedule models.Horary
	if err := db.First(&schedule, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := db.Delete(&schedule).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
