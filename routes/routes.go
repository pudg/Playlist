package routes

import (
	"net/http"
	"pudg/Playlist/database"
	"pudg/Playlist/models"

	"github.com/gin-gonic/gin"
)

func CreatePlaylist(c *gin.Context) {
	var input models.CreatePlaylistInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playlist := models.Playlist{
		Name:    input.Name,
		Creator: input.Creator,
		Links:   input.Links,
		Public:  input.Public,
	}

	database.DB.Create(&playlist)

	c.JSON(http.StatusOK, gin.H{"data": playlist})
}

func GetAllPlaylists(c *gin.Context) {
	var playlists []models.Playlist

	database.DB.Where("public = ?", true).Find(&playlists)

	c.JSON(http.StatusOK, gin.H{
		"data": playlists,
	})
}

func GetPlaylistDetails(c *gin.Context) {
	var playlist models.Playlist

	id := c.Param("id")
	database.DB.Find(&playlist, id)

	c.JSON(http.StatusOK, gin.H{"data": playlist})
}

func UpdatePlaylist(c *gin.Context) {
	var input models.CreatePlaylistInput
	var playlist models.Playlist

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	database.DB.Find(&playlist, id)

	playlist.Name = input.Name
	playlist.Creator = input.Creator
	playlist.Links = input.Links
	playlist.Public = input.Public

	database.DB.Save(&playlist)

	c.JSON(http.StatusOK, gin.H{"message": playlist})
}

func DeletePlaylist(c *gin.Context) {
	id := c.Param("id")

	database.DB.Unscoped().Delete(&models.Playlist{}, id)

	c.JSON(http.StatusNoContent, gin.H{})
}

func InitRoutes(engine *gin.Engine) {
	engine.POST("/playlists", CreatePlaylist)
	engine.GET("/playlists", GetAllPlaylists)
	engine.GET("/playlists/:id", GetPlaylistDetails)
	engine.PUT("/playlists/:id", UpdatePlaylist)
	engine.DELETE("/playlists/:id", DeletePlaylist)
}
