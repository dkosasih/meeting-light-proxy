package main

func main() {
	app := InitialiseProviders()

	for _, s := range app.Endpoints {
		s.Register()
	}
	app.Router.RunTLS(":443", "./static/devcerts/loclhost.crt", "./static/devcerts/localhost.key")
}
