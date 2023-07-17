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
	g := r.Group("/api")
	g.GET("/ping", handler.Ping)
	g.GET("/auth", auth.GetAuth)
	apiv1 := g.Group("/v1")
	{
		g := apiv1.Group("/user")
		g.POST("/create", handler.CreateUser)
	}
	{
		g := apiv1.Group("/post")
		g.Use(jwt.JWT())
		g.GET("/list", handler.ListPosts)
	} // /post
}
