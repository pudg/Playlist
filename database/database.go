package database

import (
	"gorm.io/datatypes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Playlist struct {
	gorm.Model
	Name    string         `json:"name"`
	Creator string         `json:"creator"`
	Links   datatypes.JSON `json:"link"`
	Public  *bool          `json:"public" gorm:"default:true"`
}

type CreatePlaylistInput struct {
	Name    string         `json:"name" binding:"required"`
	Creator string         `json:"creator" binding:"required"`
	Links   datatypes.JSON `json:"links" binding:"required"`
	Public  *bool          `json:"public" binding:"required"`
}

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to database...")
	}

	db.AutoMigrate(&Playlist{})

	DB = db
}
