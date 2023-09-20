package main

import (
	"pudg/Playlist/routes"

	"pudg/Playlist/database"

	"pudg/Playlist/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	database.ConnectDatabase()
	middleware.InitMiddleware(engine)
	routes.InitRoutes(engine)
	engine.Run(":8000")
}
