package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Playlist struct {
	gorm.Model
	Name    string   `json:"name"`
	Creator string   `json:"creator"`
	Links   []string `json:"links"`
	Public  bool     `json:"public"`
}

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to database...")
	}

	db.AutoMigrate(&Playlist{})

	DB = db
}
