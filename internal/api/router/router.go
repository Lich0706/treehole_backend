package router

import (
	"TreeHole/treehole_backend/internal/api/handler"
	"TreeHole/treehole_backend/internal/api/handler/auth"
	"TreeHole/treehole_backend/middleware/jwt"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{
			"Content-Type",
			"access-control-allow-origin",
			"access-control-allow-headers",
			"Authorization"},
	}))
	g := r.Group("/api")
	g.GET("/ping", handler.Ping)
	apiv1 := g.Group("/v1")
	{
		g := apiv1.Group("/user")
		g.POST("/create", handler.CreateUser)
		g.POST("/auth", auth.GetAuth)
	}
	{
		g := apiv1.Group("/post")
		g.Use(jwt.JWT())
		g.GET("/list", handler.ListPosts)
	} // /post
}
