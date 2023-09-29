package database

import (
	"log"

	"github.com/abdulmanafc2001/testing_ap/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	db := createDB()
	autoMigrate(db)
	return db
}

func createDB() *gorm.DB {
	DSN := "host=localhost user=root password=12345 dbname=cicd port=5432"
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
