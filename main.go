package main

import (
	"pudg/Playlist/routes"

	"pudg/Playlist/database"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine = routes.RegisterRoutes(engine)
	database.ConnectDatabase()
	engine.Run()
}
