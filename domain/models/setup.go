package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	connection_string, founded := os.LookupEnv("DB_STRING")

	if !founded {
		panic("DB_STRING environment variable not found!")
	}

	database, err := gorm.Open(postgres.Open(connection_string), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Book{})
	if err != nil {
		return
	}

	DB = database
}
