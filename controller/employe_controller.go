package controller

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Edwinfpirajan/Distrifabrica.git/common"
	"github.com/Edwinfpirajan/Distrifabrica.git/models"
	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	employes := []models.Employe{}
	db := common.GetConnection()
	defer db.Close()

	db.Find(&employes)
	json, _ := json.Marshal(employes)
	common.SendResponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	employe := models.Employe{}

	id := mux.Vars(request)["id"]

	db := common.GetConnection()
	defer db.Close()

	db.Find(&employe, id)

	if employe.ID > 0 {
		json, _ := json.Marshal(employe)
		common.SendResponse(writer, http.StatusOK, json)
	} else {
		common.SendError(writer, http.StatusNotFound)
	}

}

// func Save(writer http.ResponseWriter, request *http.Request) {
// 	employe := models.Employe{}

// 	db := common.GetConnection()
// 	defer db.Close()

// 	error := json.NewDecoder(request.Body).Decode(&employe)

// 	if error != nil {
// 		log.Fatal(error)
// 		common.SendError(writer, http.StatusBadRequest)
// 		return
// 	}

// 	error = db.Save(&employe).Error

// 	if error != nil {
// 		log.Fatal(error)
// 		common.SendError(writer, http.StatusInternalServerError)
// 		return
// 	}

// 	json, _ := json.Marshal(employe)

// 	common.SendResponse(writer, http.StatusCreated, json)
// }

func Save(writer http.ResponseWriter, request *http.Request) {
	employe := models.Employe{}

	db := common.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&employe)

	if error != nil {
		log.Fatal(error)
		common.SendError(writer, http.StatusBadRequest)
		return
	}

	// Genera un slice de 8 bytes aleatorios
	// Genera un slice de 4 bytes aleatorios
	randomBytes := make([]byte, 2)
	_, err := rand.Read(randomBytes)
	if err != nil {
		// maneja el error
	}

	// Convierte los bytes a una cadena hexadecimal
	employe.PinEmploye = hex.EncodeToString(randomBytes)

	// Guarda el empleado en la base de datos
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
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&employe, id)

	if employe.ID > 0 {
		db.Delete(employe)
		common.SendResponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		common.SendError(writer, http.StatusNotFound)
	}
}
