package database

import (
	"pudg/Playlist/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to database...")
	}

	db.AutoMigrate(&models.Playlist{})

	DB = db
}
