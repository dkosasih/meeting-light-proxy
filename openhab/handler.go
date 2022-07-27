package openhab

import "github.com/gin-gonic/gin"

type OpenhabHandlerInterface interface {
	UpdateOpenHab(*gin.Context)
}
