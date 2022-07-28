package openhab

import (
	"github.com/gin-gonic/gin"
)

type Endpoints struct {
	r       *gin.Engine
	creator OpenhabHandlerCreator
}

func NewEndpoints(r *gin.Engine, creator OpenhabHandlerCreator) *Endpoints {
	return &Endpoints{r: r, creator: creator}
}

func (e *Endpoints) Register() {
	routes := e.r.Group("/openhab")

	routes.POST("/command", e.updateOpenHab())
}

func (e *Endpoints) updateOpenHab() func(*gin.Context) {
	return func(c *gin.Context) {
		e.creator.CreateHandler(c).UpdateOpenHab(c)
	}
}
