package main

import (
	"crypto/tls"
	"net/http"

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

func NewHttpClient() *http.Client {
	httpClient := &http.Client{}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient.Transport = tr

	return httpClient
}

func registerServices(router *gin.Engine) {
	var openHabFactory openhab.OpenhabHandlerCreator = openhab.NewOpenhabHandlerFactory(NewHttpClient())
	var albumFactory albums.AlbumHandlerCreator = albums.NewAlbumHandlerFactory()

	albums.RegisterEndpoints(router, albumFactory)
	openhab.RegisterEndpoints(router, openHabFactory)
}
