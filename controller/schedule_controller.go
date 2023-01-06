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
// 	json, _ := json.Marshal(schedule)
// 	common.SendResponse(w, http.StatusOK, json)

// 	fmt.Println(schedule)

// }

func SaveSchedule(w http.ResponseWriter, r *http.Request) {
	// Almacenar el cuerpo de la solicitud en una variable
	requestBody := map[string]string{}
	error := json.NewDecoder(r.Body).Decode(&requestBody)
	if error != nil {
		common.SendError(w, http.StatusBadRequest)
		return
	}
	// Obtener la hora de llegada del cuerpo de la solicitud
	arrival, ok := requestBody["arrival"]
	if !ok {
		common.SendError(w, http.StatusBadRequest)
		return
	}
	// Parsear la hora de llegada en formato HH:MM:SS
	t, err := time.Parse("15:04", arrival)
	if err != nil {
		common.SendError(w, http.StatusBadRequest)
		return
	}
	// Establecer el día y el año en el valor predeterminado "0001-01-01"
	t = time.Date(1, 1, 1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
	// Agregar la hora a la fecha
	t = t.Add(time.Duration(t.Hour()) * time.Hour)
	t = t.Add(time.Duration(t.Minute()) * time.Minute)
	t = t.Add(time.Duration(t.Second()) * time.Second)
	t = t.Add(time.Duration(t.Nanosecond()) * time.Nanosecond)
	// Crear una instancia de la estructura Horary con la hora parseada
	schedule := models.Horary{Arrival: t}
	// Obtener la hora de salida del cuerpo de la solicitud
	departure, ok := requestBody["departure"]
	if !ok {
		common.SendError(w, http.StatusBadRequest)
		return
	}
	// Parsear la hora de salida en formato HH:MM:SS
	t, err = time.Parse("15:04", departure)
	if err != nil {
		common.SendError(w, http.StatusBadRequest)
		return
	}
	// Establecer el día y el año en el valor predeterminado "0001-01-01"
	t = time.Date(1, 1, 1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
	// Agregar la hora a la fecha
	t = t.Add(time.Duration(t.Hour()) * time.Hour)
	t = t.Add(time.Duration(t.Minute()) * time.Minute)
	t = t.Add(time.Duration(t.Second()) * time.Second)
	t = t.Add(time.Duration(t.Nanosecond()) * time.Nanosecond)
	// Asignar la hora parseada al campo Departure de la estructura Horary
	schedule.Departure = t
	// Obtener una conexión a la base de datos
	db := common.GetConnection()
	// Guardar el horario en la base de datos
	db.Save(&schedule)
	// Convertir la estructura Horary a formato JSON
	json, _ := json.Marshal(schedule)
	// Enviar la respuesta al cliente
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
