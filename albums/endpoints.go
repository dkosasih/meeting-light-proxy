package albums

import (
	"github.com/gin-gonic/gin"
)

type Endpoints struct {
	r       *gin.Engine
	creator AlbumHandlerCreator
}

func NewEndpoints(r *gin.Engine, creator AlbumHandlerCreator) *Endpoints {
	return &Endpoints{r: r, creator: creator}
}

func (e *Endpoints) Register() {
	routes := e.r.Group("/albums")

	routes.GET("/", e.getAlbums())
	routes.GET("/:id", e.getAlbumByID())
	routes.POST("/", e.postAlbums())
}

func (e *Endpoints) getAlbums() func(*gin.Context) {
	return func(c *gin.Context) {
		e.creator.CreateHandler(c).GetAlbums(c)
	}
}

func (e *Endpoints) getAlbumByID() func(c *gin.Context) {
	return func(c *gin.Context) {
		e.creator.CreateHandler(c).GetAlbumByID(c)
	}
}

func (e *Endpoints) postAlbums() func(c *gin.Context) {
	return func(c *gin.Context) {
		e.creator.CreateHandler(c).PostAlbums(c)
	}
}
