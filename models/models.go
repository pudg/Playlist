package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Playlist struct {
	gorm.Model
	Name    string         `json:"name"`
	Creator string         `json:"creator"`
	Links   pq.StringArray `json:"links" gorm:"type:text[]"`
	Public  *bool          `json:"public" gorm:"default:false"`
}

type CreatePlaylistInput struct {
	Name    string         `json:"name" binding:"required"`
	Creator string         `json:"creator" binding:"required"`
	Links   pq.StringArray `json:"links" binding:"required"`
	Public  *bool          `json:"public" binding:"required"`
}
