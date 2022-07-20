package albums

import (
	"github.com/gin-gonic/gin"
)

func RegisterEndpoints(r *gin.Engine, creator AlbumHandlerCreator) {
	routes := r.Group("/albums")

	routes.GET("/", getAlbums(creator))
	routes.GET("/:id", getAlbumByID(creator))
	routes.POST("/", postAlbums(creator))
}

func getAlbums(creator AlbumHandlerCreator) func(*gin.Context) {
	return func(c *gin.Context) {
		creator.CreateHandler(c).GetAlbums(c)
	}
}

func getAlbumByID(creator AlbumHandlerCreator) func(c *gin.Context) {
	return func(c *gin.Context) {
		creator.CreateHandler(c).GetAlbumByID(c)
	}
}

func postAlbums(creator AlbumHandlerCreator) func(c *gin.Context) {
	return func(c *gin.Context) {
		creator.CreateHandler(c).PostAlbums(c)
	}
}
