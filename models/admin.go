package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"user" `
	Password string `json:"pass" `
	// Accounts []Account
}

type Account struct {
	gorm.Model
	Type   string
	Name   string
	UserID uint
}

type ResponseAccount struct {
	ID   uint
	Name string
}

type ResponseUser struct {
	ID       uint
	Username string
	Password string
	Accounts []ResponseAccount
}

//Creación de usuario administrador

// func createAdmin(){
// 	users := [1]User{
// 		Username: "Admin"
// 		// Password: "1234"
// 	}

// 	for i = 0; i <len(users); i++{
// 		generatedPassword := common.HashAndSalt([]byte(users[i].Username))
// 	}
// }
