package common

import (
	"fmt"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetConnection() *gorm.DB {
	// dsn := "root:@/distridb?charset=utf8mb4&parseTime=True&loc=Local"
	host := os.Getenv("MYSQLHOST")
	pass := os.Getenv("MYSQLPASSWORD")
	port := os.Getenv("MYSQLPORT")
	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/railway", host, pass, port)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Status:", http.StatusServiceUnavailable)
	}

	return db
}
