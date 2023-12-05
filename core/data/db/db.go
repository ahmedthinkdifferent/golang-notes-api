package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"noteapp/core/data/model"
)

var Database *gorm.DB

func Connect() {
	dsn := "root:ahmed123@tcp(127.0.0.1:3306)/notes?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	Database = db
	// Migrate the schema
	Database.AutoMigrate(&model.User{}, &model.Note{})
}
