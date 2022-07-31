package main

import (
	"github.com/dkosasih/meeting-light-proxy/interfaces"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router    *gin.Engine
	Endpoints []interfaces.EndpointRegistrator
}
