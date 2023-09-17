package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(engine *gin.Engine) {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
}
