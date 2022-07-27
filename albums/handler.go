package albums

import "github.com/gin-gonic/gin"

type AlbumHandlerInterface interface {
	GetAlbums(*gin.Context)
	GetAlbumByID(*gin.Context)
	PostAlbums(*gin.Context)
}
