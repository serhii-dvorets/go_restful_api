package database

import (
	"fmt"
	"log"
	"os"
	"restful_api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectDB() *gorm.DB {
	
	host := os.Getenv("HOST")
	dbport := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")
	
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, password, dbport)
	
	var db *gorm.DB
	var err error

	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database!")
	}

	db.AutoMigrate(&models.User{}, &models.Message{})

	return db
}
