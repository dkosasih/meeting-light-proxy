package albums

import (
	v1 "github.com/dkosasih/meeting-light-proxy/albums/v1"
	v2 "github.com/dkosasih/meeting-light-proxy/albums/v2"
	"github.com/dkosasih/meeting-light-proxy/utils"
	"github.com/gin-gonic/gin"
)

type AlbumHandlerFactory struct {
}

type AlbumHandlerCreator interface {
	CreateHandler(*gin.Context) AlbumHandlerInterface
}

func (hf *AlbumHandlerFactory) CreateHandler(c *gin.Context) AlbumHandlerInterface {
	version := utils.GetVersionString(c)

	var ohi AlbumHandlerInterface

	switch version {
	case "1":
		// TODO: is there a better way to refactor this?
		c.Writer.Header().Add("x-version", "1")
		ohi = &v1.AlbumHandler{}
	default:
		c.Writer.Header().Add("x-version", "2")
		ohi = &v2.AlbumHandler{}
	}

	return ohi
}
