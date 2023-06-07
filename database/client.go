package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"test/models"
)

var Db *gorm.DB
var err error

func Connect() {
	dsn := "root:0he0dFqkGG738S8Eo9gc@tcp(containers-us-west-95.railway.app:6719)/railway?charset=utf8mb4&parseTime=True&loc=Local"
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
