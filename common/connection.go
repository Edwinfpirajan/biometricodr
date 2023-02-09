package common

import (
	"fmt"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetConnection() *gorm.DB {
	dsn := "MYSQL_URL"
	// dsn := os.Getenv("distriramirezcom_boreal:Developer2023++@tcp(107.6.54.56:3306)/distriramirezcom_boreal?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Status:", http.StatusServiceUnavailable)
	}

	return db
}
