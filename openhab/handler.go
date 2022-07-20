package openhab

import "github.com/gin-gonic/gin"

type OpenhabHandler struct {
}

type OpenhabHandlerInterface interface {
	UpdateOpenHab(*gin.Context)
}
