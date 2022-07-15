package albums

import (
	"github.com/gin-gonic/gin"
)

type albumHandler struct {
}

func RegisterEndpoints(r *gin.Engine) {
	ah := &albumHandler{}

	routes := r.Group("/albums")

	routes.GET("/", ah.GetAlbums)
	routes.GET("/:id", ah.GetAlbumByID)
	routes.POST("/", ah.PostAlbums)
}
