package controller

import (
	"net/http"
	"time"

	"github.com/Edwinfpirajan/Distrifabrica.git/common"
	"github.com/Edwinfpirajan/Distrifabrica.git/models"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	db := common.GetConnection()

	var user models.User
	// error := json.NewDecoder(c.Request().Body).Decode(&user)

	if err := db.Table("admin_user").Where("email = ? and pass = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, _ := token.SignedString([]byte("secret"))

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}
