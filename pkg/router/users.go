package router

import "github.com/gin-gonic/gin"

func createUserRouter(g *gin.RouterGroup) {
	users := g.Group("/users")
	users.GET("/", userHandler)
}

func userHandler(context *gin.Context) {
}
