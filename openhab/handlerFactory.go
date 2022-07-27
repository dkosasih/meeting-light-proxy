package openhab

import (
	"net/http"

	v1 "github.com/dkosasih/meeting-light-proxy/openhab/v1"
	"github.com/dkosasih/meeting-light-proxy/utils"
	"github.com/gin-gonic/gin"
)

type openhabHandlerFactory struct {
	HttpClient *http.Client
}

type OpenhabHandlerCreator interface {
	CreateHandler(*gin.Context) OpenhabHandlerInterface
}

func NewOpenhabHandlerFactory(client *http.Client) *openhabHandlerFactory {
	return &openhabHandlerFactory{client}
}

func (hf *openhabHandlerFactory) CreateHandler(c *gin.Context) OpenhabHandlerInterface {
	version := utils.GetVersionString(c)

	var ohi OpenhabHandlerInterface

	switch version {
	case "1":
		// TODO: is there a better way to refactor this?
		c.Writer.Header().Add("x-version", "1")
		ohi = v1.NewOpenHabHandler(hf.HttpClient)
	default:
		c.Writer.Header().Add("x-version", "2")
		ohi = v1.NewOpenHabHandler(hf.HttpClient)
	}

	return ohi
}
