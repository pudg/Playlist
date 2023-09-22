package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pudg/Playlist/database"
	"pudg/Playlist/middleware"
	"pudg/Playlist/models"
	"pudg/Playlist/routes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllPlaylistsRoute(t *testing.T) {
	router := routes.InitRouter()
	middleware.InitMiddleware(router)
	database.ConnectDatabase()
	router.GET("/playlists", routes.GetAllPlaylists)
	req, _ := http.NewRequest("GET", "/playlists", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
	assert.Contains(t, w.Body.String(), "data")
	assert.Contains(t, w.Body.String(), "links")
}

func TestPlaylistDetailsRoute(t *testing.T) {
	router := routes.InitRouter()
	middleware.InitMiddleware(router)
	database.ConnectDatabase()
	router.GET("/playlists/:id", routes.GetPlaylistDetails)
	req, _ := http.NewRequest("GET", "/playlists/2", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
	assert.Contains(t, w.Body.String(), "data")
	assert.Contains(t, w.Body.String(), "links")
}

func TestEmptyPlaylistCreateRoute(t *testing.T) {
	router := routes.InitRouter()
	middleware.InitMiddleware(router)
	database.ConnectDatabase()
	router.POST("/playlists", routes.CreatePlaylist)
	empty_playlist := models.Playlist{}
	json_val, _ := json.Marshal(empty_playlist)
	req, _ := http.NewRequest("POST", "/playlists", bytes.NewBuffer(json_val))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestNonEmptyPlaylistCreateRoute(t *testing.T) {
	router := routes.InitRouter()
	middleware.InitMiddleware(router)
	database.ConnectDatabase()
	router.POST("/playlists", routes.CreatePlaylist)
	playlist := models.Playlist{
		Name:    "Orlando",
		Creator: "Not Nelson",
		Links:   []string{"Twitch", "YouTube"},
		Public:  func() *bool { b := true; return &b }(),
	}
	json_val, _ := json.Marshal(playlist)
	req, _ := http.NewRequest("POST", "/playlists", bytes.NewBuffer(json_val))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdatePlaylistRoute(t *testing.T) {
	router := routes.InitRouter()
	middleware.InitMiddleware(router)
	database.ConnectDatabase()
	router.PUT("/playlists/:id", routes.UpdatePlaylist)
	updated_playlist := models.Playlist{
		Name:    "Cool Sites",
		Creator: "Nelson",
		Links: []string{
			"@HackerNews",
			"@NetflixEngineering",
			"@DiscordEngineering",
		},
		Public: func() *bool { b := false; return &b }(),
	}
	json_val, _ := json.Marshal(updated_playlist)
	req, _ := http.NewRequest("PUT", "/playlists/9", bytes.NewBuffer(json_val))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeletePlaylistRoute(t *testing.T) {
	router := routes.InitRouter()
	middleware.InitMiddleware(router)
	database.ConnectDatabase()
	router.DELETE("/playlists/:id", routes.DeletePlaylist)
	req, _ := http.NewRequest("DELETE", "/playlists/11", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
