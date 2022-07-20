package albums

import "github.com/gin-gonic/gin"

type AlbumHandler struct {
}

type AlbumHandlerInterface interface {
	GetAlbums(*gin.Context)
	GetAlbumByID(*gin.Context)
	PostAlbums(*gin.Context)
}
