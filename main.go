package main

import (
	"pudg/Playlist/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine = routes.RegisterRoutes(engine)
	engine.Run()
}
