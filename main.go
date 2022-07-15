package main

import (
	"github.com/dkosasih/meeting-light-proxy/albums"
	"github.com/dkosasih/meeting-light-proxy/openhab"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	albums.RegisterEndpoints(router)
	openhab.RegisterEndpoints(router)

	router.RunTLS(":443", "./static/devcerts/loclhost.crt", "./static/devcerts/localhost.key")
}
