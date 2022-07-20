package openhab

import (
	"github.com/gin-gonic/gin"
)

func RegisterEndpoints(r *gin.Engine, creator OpenhabHandlerCreator) {
	routes := r.Group("/openhab")

	routes.POST("/command", updateOpenHab(creator))
}

func updateOpenHab(creator OpenhabHandlerCreator) func(*gin.Context) {
	return func(c *gin.Context) {
		creator.CreateHandler(c).UpdateOpenHab(c)
	}
}
