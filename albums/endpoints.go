package albums

import (
	v1 "github.com/dkosasih/meeting-light-proxy/albums/v1"
	v2 "github.com/dkosasih/meeting-light-proxy/albums/v2"
	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
}

type AlbumHandlerInterface interface {
	GetAlbums(*gin.Context)
	GetAlbumByID(*gin.Context)
	PostAlbums(*gin.Context)
}

func RegisterEndpoints(r *gin.Engine) {
	routes := r.Group("/albums")

	routes.GET("/", getAlbums)
	routes.GET("/:id", getAlbumByID)
	routes.POST("/", postAlbums)
}

func getAlbums(c *gin.Context) {
	createHandler(c).GetAlbums(c)
}

func getAlbumByID(c *gin.Context) {
	createHandler(c).GetAlbumByID(c)
}

func postAlbums(c *gin.Context) {
	createHandler(c).PostAlbums(c)
}

func createHandler(c *gin.Context) AlbumHandlerInterface {
	version, exists := c.Get("version")
	var ahi AlbumHandlerInterface = &v2.AlbumHandler{} // latest by default

	if exists {
		if version == "1" {
			return &v1.AlbumHandler{}
		}
	}

	return ahi
}
