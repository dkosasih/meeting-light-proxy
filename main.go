package main

import (
	"github.com/dkosasih/meeting-light-proxy/albums"
	"github.com/dkosasih/meeting-light-proxy/middleware"
	"github.com/dkosasih/meeting-light-proxy/openhab"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(middleware.ApiVersion())

	registerServices(router)

	router.RunTLS(":443", "./static/devcerts/loclhost.crt", "./static/devcerts/localhost.key")
}

func registerServices(router *gin.Engine) {
	var openHabFactory openhab.OpenhabHandlerCreator = &openhab.OpenhabHandlerFactory{}
	var albumFactory albums.AlbumHandlerCreator = &albums.AlbumHandlerFactory{}

	albums.RegisterEndpoints(router, albumFactory)
	openhab.RegisterEndpoints(router, openHabFactory)
}
