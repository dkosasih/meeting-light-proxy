// Package proxy API of the openhab and sample.
//
// OpenHAB proxy and sample api.
//
// Version: 1.0.0
// BasePath: /
// Schemes: https
// Consumes:
// - application/json
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
)

func main() {
	app := InitialiseProviders()

	for _, s := range app.Endpoints {
		s.Register()
	}

	app.Router.Handle("GET", "/swagger.yaml", gin.WrapH(http.FileServer(http.Dir("./"))))

	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	app.Router.Handle("GET", "/docs", gin.WrapH(sh))

	app.Router.RunTLS(":443", "./static/devcerts/loclhost.crt", "./static/devcerts/localhost.key")
}
