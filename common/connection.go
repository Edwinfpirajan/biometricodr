package common

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetConnection() *gorm.DB {
	// dsn := "root:@/distridb"
	pass := os.Getenv("MYSQLPASSWORD")
	host := os.Getenv("MYSQLHOST")
	port := os.Getenv("MYSQLPORT")
	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/railway?charset=utf8mb4&parseTime=True&loc=Local", pass, host, port)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Print("Error")
	}

	return db
}
