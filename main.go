package main

import (
	"pudg/Playlist/database"
	"pudg/Playlist/middleware"
	"pudg/Playlist/routes"
)

func main() {
	router := routes.InitRouter()
	database.ConnectDatabase()
	middleware.InitMiddleware(router)
	routes.RegisterRoutes(router)
	router.Run(":8000")
}
