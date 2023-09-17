package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//ID
//Name
//Creator
//Links
//Public

func CreatePlaylist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "TODO: Implement POST playlist",
	})
}

func GetAllPlaylists(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "TODO: Implement GET all playlist",
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
