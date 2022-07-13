package main

import (
	handlers "github.com/dkosasih/meeting-light-proxy/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)
	router.POST("/albums", handlers.PostAlbums)
	router.POST("/openhab/command", handlers.UpdateOpenHab)

	router.RunTLS("localhost:443", "./static/devcerts/loclhost.crt", "./static/devcerts/localhost.key")
}
