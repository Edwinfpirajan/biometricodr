package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Edwinfpirajan/Distrifabrica.git/common"
)

type Login struct {
	Username string
	Password string
}

type ErrResponse struct {
	Message string
}

func Logeo(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	common.HandleErr(err)

	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	common.HandleErr(err)
	login := common.Login(formattedBody.Username, formattedBody.Password)

	//Preparación de la respuestas

	if login["Message"] == "Bienvenido" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		// Manejo de error
		resp := ErrResponse{Message: "Nombre de usuario o contraseña incorrectos"}
		json.NewEncoder(w).Encode(resp)
	}

}

// func StartApi() {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/login", login).Methods("POST")
// 	fmt.Println("Login funcionando")

// }
