package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createPlaylist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "TODO: Implement POST playlists",
	})
}

func playlistsAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "TODO: Implement GET all playlists...",
	})
}

func playlistDetails(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "TODO: Implement GET playlist detail...",
	})
}

func updatePlaylist(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "TODO: Implement UPDATE playlist",
	})
}

func deletePlaylist(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"message": "TODO: Implement DELETE playlist",
	})
}

func RegisterRoutes(engine *gin.Engine) *gin.Engine {
	engine.POST("/playlists", createPlaylist)
	engine.GET("/playlists", playlistsAll)
	engine.GET("/playlists/:id", playlistDetails)
	engine.PUT("/playlists/:id", updatePlaylist)
	engine.DELETE("/playlists/:id", deletePlaylist)
	return engine
}
