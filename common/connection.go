package common

import (
	"fmt"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetConnection() *gorm.DB {
	// dsn := "root:@/distridb?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:1zkzAauj8BwLeBjxyZiC@tcp(containers-us-west-148.railway.app:6027)/railway"
	// dsn := os.Getenv("MYSQL_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Status:", http.StatusServiceUnavailable)
	}

	return db
}
