package openhab

import (
	v1 "github.com/dkosasih/meeting-light-proxy/openhab/v1"
	"github.com/gin-gonic/gin"
)

type OpenhabHandler struct {
}

type OpenhabHandlerInterface interface {
	UpdateOpenHab(*gin.Context)
}

func RegisterEndpoints(r *gin.Engine) {
	routes := r.Group("/openhab")

	routes.POST("/command", updateOpenHab)
}

func updateOpenHab(c *gin.Context) {
	createHandler(c).UpdateOpenHab(c)
}

func createHandler(c *gin.Context) OpenhabHandlerInterface {
	version, exists := c.Get("version")
	var ohi OpenhabHandlerInterface = &v1.OpenhabHandler{} // default
	if exists {
		if version == "1" {
			return &v1.OpenhabHandler{}
		}
	}

	return ohi
}
