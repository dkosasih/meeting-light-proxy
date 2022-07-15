package openhab

import "github.com/gin-gonic/gin"

type openhabHandler struct {
}

func RegisterEndpoints(r *gin.Engine) {
	oh := &openhabHandler{}

	routes := r.Group("/openhab")

	routes.POST("/command", oh.UpdateOpenHab)
}
