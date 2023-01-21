package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Edwinfpirajan/Distrifabrica.git/common"
	"github.com/Edwinfpirajan/Distrifabrica.git/models"
	jwt "github.com/golang-jwt/jwt/v4"
)

func Login(w http.ResponseWriter, r *http.Request) {
	db := common.GetConnection()

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	if err := db.Table("admin_user").Where("email = ? and pass = ?", user.Email, user.Password).First(&user).Error; err != nil {
		http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, _ := token.SignedString([]byte("secret"))

	json.NewEncoder(w).Encode(models.JwtToken{Token: tokenString})
}
