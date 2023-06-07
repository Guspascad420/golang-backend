package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"test/models"
)

var Db *gorm.DB
var err error

func Connect() {
	port := os.Getenv("PORT")
	dsn := "root:rahasia@tcp(0.0.0.0:" + port + ")/projects?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to DB")
	} else {
		log.Println("Connected to Database!")
	}
}

func Migrate() {
	Db.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}
