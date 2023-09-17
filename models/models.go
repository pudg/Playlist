package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

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
