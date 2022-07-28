//go:build wireinject
// +build wireinject

package main

import (
	"crypto/tls"
	"net/http"

	"github.com/dkosasih/meeting-light-proxy/albums"
	"github.com/dkosasih/meeting-light-proxy/interfaces"
	"github.com/dkosasih/meeting-light-proxy/middlewares"
	"github.com/dkosasih/meeting-light-proxy/openhab"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func NewHttpClient() *http.Client {
	httpClient := &http.Client{}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient.Transport = tr

	return httpClient
}

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.ApiVersion())

	return router
}

func ProvideEndpoints(
	albumEp *albums.Endpoints,
	openHabEp *openhab.Endpoints) []interfaces.EndpointRegistrator {
	return []interfaces.EndpointRegistrator{albumEp, openHabEp}
}

func InitialiseProviders() *App {
	panic(
		wire.Build(
			NewRouter,
			NewHttpClient,
			openhab.Provider,
			albums.Provider,
			ProvideEndpoints,
			wire.Struct(new(App), "*"),
		),
	)
}
