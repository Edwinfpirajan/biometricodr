package common

import (
	"log"

	"github.com/Edwinfpirajan/Distrifabrica.git/models"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, error := gorm.Open("mysql", "root:@/distridb?charset=utf8&parseTime=True&loc=Local")

	if error != nil {
		log.Fatal(error)
	}

	return db
}

func Migrate() {
	db := GetConnection()

	defer db.Close()
	log.Println("Iniciando...")

	db.AutoMigrate(
		&models.Employe{},
		&models.Attendance{},
		&models.User{},
		&models.Account{})
}
