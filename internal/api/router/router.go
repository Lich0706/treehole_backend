package router

import (
	"TreeHole/treehole_backend/internal/api/handler"
	"TreeHole/treehole_backend/internal/api/handler/auth"
	"TreeHole/treehole_backend/middleware/jwt"

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
	r.GET("/auth", auth.GetAuth)
	g := r.Group("/api")
	g.GET("/ping", handler.Ping)
	apiv1 := g.Group("/v1")
	apiv1.Use(jwt.JWT())
	{
		g := apiv1.Group("/post")
		g.GET("/list", handler.ListPosts)
	} // /message
}
