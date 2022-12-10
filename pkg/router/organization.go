package router

import "github.com/gin-gonic/gin"

func createOrganizationRoutes(g *gin.RouterGroup) {
	org := g.Group("/organizations")
	org.GET("/", orgHandler)
}

func orgHandler(context *gin.Context) {

}
