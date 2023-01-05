package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func SaveSchedule(w http.ResponseWriter, r *http.Request) {

	schedule := models.Horary{}

	db := common.GetConnection()

	error := json.NewDecoder(r.Body).Decode(&schedule)

	if error != nil {
		common.SendError(w, http.StatusBadRequest)
		return
	}

	db.Create(&schedule)
	json, _ := json.Marshal(schedule)
	common.SendResponse(w, http.StatusOK, json)

	fmt.Println(schedule)

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
