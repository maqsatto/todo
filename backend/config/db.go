package config

import (
	"backend/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД: ", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		log.Fatal("Ошибка миграции: ", err)
	}

	DB = db
}
