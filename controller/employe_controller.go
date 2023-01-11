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

// func GetAll(writer http.ResponseWriter, request *http.Request) {
// 	employes := []models.Employe{}
// 	db := common.GetConnection()

// 	db.Find(&employes)
// 	json, _ := json.Marshal(employes)
// 	common.SendResponse(writer, http.StatusOK, json)
// }

// func GetAll(writer http.ResponseWriter, request *http.Request) {
// 	employes := []entity.Employe{}
// 	db := common.GetConnection()

// 	// db.Raw("SELECT * FROM employes e LEFT JOIN horaries h ON h.id = e.schedule_id").Scan(&employes)
// 	db.Model(&employes).Select("*").Joins("left join horaries h on h.id = employes.schedule_id").Scan(&employes)
// 	db.Find(&employes)
// 	json, _ := json.Marshal(employes)
// 	common.SendResponse(writer, http.StatusOK, json)
// 	// fmt.Println(employes)
// }

func GetAll(writer http.ResponseWriter, request *http.Request) {

	employeesWithSchedule := []models.EmployeeWithSchedule{}
	db := common.GetConnection()
	// db.Raw("SELECT * FROM employes e INNER JOIN horaries h ON h.id = e.schedule_id").Find(&employeesWithSchedule)
	db.Table("employes").Select("*").Joins("left join horaries h on h.id = employes.schedule_id").Scan(&employeesWithSchedule)
	fmt.Println("consulta: ", employeesWithSchedule)
	json, _ := json.Marshal(employeesWithSchedule)
	common.SendResponse(writer, http.StatusOK, json)

}

// func GetAll(writer http.ResponseWriter, request *http.Request) {
// 	employes := []models.Employe{}
// 	db := common.GetConnection()

// 	db.Find(&employes)
// 	json, _ := json.Marshal(employes)
// 	common.SendResponse(writer, http.StatusOK, json)
// }

// func Get(writer http.ResponseWriter, request *http.Request) {
// 	employe := models.Employe{}

// 	id := mux.Vars(request)["id"]

// 	db := common.GetConnection()

// 	db.Find(&employe, id)

// 	if employe.ID > 0 {
// 		json, _ := json.Marshal(employe)
// 		common.SendResponse(writer, http.StatusOK, json)
// 	} else {
// 		common.SendError(writer, http.StatusNotFound)
// 	}
// }

// func Get(writer http.ResponseWriter, request *http.Request) {
// 	employe := entity.Employe{}

// 	id := mux.Vars(request)["id"]

// 	db := common.GetConnection()

// 	db.Select("*").Joins("left join horaries h on h.id = employes.schedule_id").Find(&employe, id)

// 	db.Find(&employe, id)

// 	if employe.ID > 0 {
// 		json, _ := json.Marshal(employe)
// 		common.SendResponse(writer, http.StatusOK, json)
// 	} else {
// 		common.SendError(writer, http.StatusNotFound)
// 	}
// }

// func Save(writer http.ResponseWriter, request *http.Request) {
// 	employe := models.Employe{}

// 	db := common.GetConnection()

// 	error := json.NewDecoder(request.Body).Decode(&employe)

// 	if error != nil {
// 		log.Fatal(error)
// 		common.SendError(writer, http.StatusBadRequest)
// 		return
// 	}

// 	randomBytes := make([]byte, 2)
// 	_, err := rand.Read(randomBytes)
// 	if err != nil {
// 	}

// 	// Convierte los bytes a una cadena hexadecimal
// 	employe.PinEmploye = hex.EncodeToString(randomBytes)

// 	// Guarda el empleado en la base de datos
// 	error = db.Save(&employe).Error

// 	if error != nil {
// 		log.Fatal(error)
// 		common.SendError(writer, http.StatusInternalServerError)
// 		return
// 	}

// 	json, _ := json.Marshal(employe)

// 	common.SendResponse(writer, http.StatusCreated, json)
// }

// func Save(writer http.ResponseWriter, request *http.Request) {
// 	employe := entity.Employe{}

// 	db := common.GetConnection()

// 	error := json.NewDecoder(request.Body).Decode(&employe)

// 	if error != nil {
// 		log.Fatal(error)
// 		common.SendError(writer, http.StatusBadRequest)
// 		return
// 	}

// 	randomBytes := make([]byte, 2)
// 	_, err := rand.Read(randomBytes)
// 	if err != nil {
// 	}

// 	// Convierte los bytes a una cadena hexadecimal
// 	employe.PinEmploye = hex.EncodeToString(randomBytes)

// 	// Incluye los datos de la tabla Horary en el objeto employe
// 	horary := entity.Employe{Arrival: employe.Arrival, Departure: employe.Departure}
// 	error = db.Save(&horary).Error

// 	if error != nil {
// 		log.Fatal(error)
// 		common.SendError(writer, http.StatusInternalServerError)
// 		return
// 	}

// 	// Asigna el ID de la tabla Horary al objeto employe
// 	employe.ScheduleId = horary.ID

// 	// Guarda el empleado en la base de datos
// 	error = db.Save(&employe).Error

// 	if error != nil {
// 		log.Fatal(error)
// 		common.SendError(writer, http.StatusInternalServerError)
// 		return
// 	}

// 	json, _ := json.Marshal(employe)

// 	common.SendResponse(writer, http.StatusCreated, json)

// 	fmt.Println(employe)
// }

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
	fmt.Println(employe)
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
