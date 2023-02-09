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
	dsn := os.Getenv("root:vjGaLrgKEoNnaJKjwlr7@tcp(containers-us-west-118.railway.app:7762)/railway?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Status:", http.StatusServiceUnavailable)
	}

	return db
}
