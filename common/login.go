package common

import (
	"fmt"
	"time"

	"github.com/Edwinfpirajan/Distrifabrica.git/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, pass string) map[string]interface{} {
	db := GetConnection()
	user := &models.User{}
	// sql := "SELECT * FROM users WHERE username'" + user.Username + "'and password ='" + user.Password + "'"

	if db.Where("username = ?", username).First(&user).RecordNotFound() {
		return map[string]interface{}{"message": "Usuario no encontrado"}
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return map[string]interface{}{"mesagge": "Usuario o contraseña incorrectos"}
	}

	// buscar cuenta usuario

	accounts := []models.ResponseAccount{}

	db.Table("accounts").Select("id, name").Where("user_id = ?", user.ID).Scan(&accounts)

	// Configuración del response
	responseUser := &models.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Accounts: accounts,
	}

	defer db.Close()

	fmt.Println(responseUser)

	//Entrada token

	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute ^ 60).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	HandleErr(err)

	var response = map[string]interface{}{"Message": "Bienvenido"}
	response["jwt"] = token
	response["data"] = responseUser

	return response
}
