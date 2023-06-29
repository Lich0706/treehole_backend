package router

import (
	"TreeHole/treehole_backend/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// RegisterRoutes
	Register(r)

	return r
}

func Register(r *gin.Engine) {
	g := r.Group("/api")
	g.GET("/ping", handler.Ping)
	{
		g := g.Group("/message")
		g.GET("/list", handler.ListMessage)
	} // /message
}
