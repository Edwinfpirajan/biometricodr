package common

import (
	"fmt"
	"os"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB = getConnection()

func getConnection() *gorm.DB {
	// dsn := "root:@/distridb"
	pass := os.Getenv("MYSQLPASSWORD")
	host := os.Getenv("MYSQLHOST")
	port := os.Getenv("MYSQLPORT")
	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/railway?charset=utf8mb4&parseTime=True&loc=Local", pass, host, port)
	// dsn := "biometrico:_u7825Son@tcp(serverpruebas.tk:3306)/biometricov?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err.Error())
	}

	log.Info("Connection Successfully to Mysql")

	return db
}
