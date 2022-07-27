package albums

import (
	v1 "github.com/dkosasih/meeting-light-proxy/albums/v1"
	v2 "github.com/dkosasih/meeting-light-proxy/albums/v2"
	"github.com/dkosasih/meeting-light-proxy/utils"
	"github.com/gin-gonic/gin"
)

type albumHandlerFactory struct {
}

type AlbumHandlerCreator interface {
	CreateHandler(*gin.Context) AlbumHandlerInterface
}

func NewAlbumHandlerFactory() *albumHandlerFactory {
	return &albumHandlerFactory{}
}

func (hf *albumHandlerFactory) CreateHandler(c *gin.Context) AlbumHandlerInterface {
	version := utils.GetVersionString(c)

	var ohi AlbumHandlerInterface

	switch version {
	case "1":
		// TODO: is there a better way to refactor this?
		c.Writer.Header().Add("x-version", "1")
		ohi = v1.NewAlbumHandler()
	default:
		c.Writer.Header().Add("x-version", "2")
		ohi = v2.NewAlbumHandler()
	}

	return ohi
}
