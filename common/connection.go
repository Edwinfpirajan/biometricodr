package common

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	dsn := "root:@/distridb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Migrate() {
	db := GetConnection()

	log.Println("Iniciando...")

	db.AutoMigrate(
	// &models.Employe{},
	// &models.Attendances{},
	// &models.User{},
	// &models.Account{},
	// &models.Horary{},
	)
}
