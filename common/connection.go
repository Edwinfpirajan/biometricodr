package common

import (
	"fmt"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetConnection() *gorm.DB {
	dsn := "root:@/distridb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Status:", http.StatusServiceUnavailable)
	}

	return db
}

// func Migrate() {
// 	db := GetConnection()

// 	log.Println("Iniciando...")

// 	db.AutoMigrate(
// 	// &models.Employe{},
// 	// &models.Attendances{},
// 	// &models.User{},
// 	// &models.Account{},
// 	// &models.Horary{},
// 	)
// }
