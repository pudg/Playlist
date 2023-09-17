package routes

import (
	"net/http"
	"pudg/Playlist/database"

	"github.com/gin-gonic/gin"
)

func CreatePlaylist(c *gin.Context) {
	var input database.CreatePlaylistInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playlist := database.Playlist{
		Name:    input.Name,
		Creator: input.Creator,
		Links:   input.Links,
		Public:  input.Public,
	}

	database.DB.Create(&playlist)

	c.JSON(http.StatusOK, gin.H{"data": playlist})
}

func GetAllPlaylists(c *gin.Context) {
	var playlists []database.Playlist

	database.DB.Find(&playlists)

	c.JSON(http.StatusOK, gin.H{
		"data": playlists,
	})
}

func GetPlaylistDetails(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "TODO: Implement GET playlist details",
	})
}

func UpdatePlaylist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "TODO: Implement UPDATE playlist",
	})
}

func DeletePlaylist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "TODO: Implement DELETE playlist",
	})
}

func RegisterRoutes(engine *gin.Engine) *gin.Engine {
	engine.POST("/playlists", CreatePlaylist)
	engine.GET("/playlists", GetAllPlaylists)
	engine.GET("/playlists/:id", GetPlaylistDetails)
	engine.PUT("/playlists/:id", UpdatePlaylist)
	engine.DELETE("/playlists/:id", DeletePlaylist)
	return engine
}
