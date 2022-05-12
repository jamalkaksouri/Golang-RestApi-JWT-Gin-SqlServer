package database

import (
	"jwt-authentication-golang/models"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// defining an instance of the database
var Instance *gorm.DB
var dbError error

func Connet(connectionString string) {
	Instance, dbError = gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB!")
	}
	log.Println("Connected to Database!")
}
func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}
