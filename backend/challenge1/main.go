package main

import (
	"github.com/Clarilab/gocloaksession"
	"github.com/go-resty/resty/v2"
	atreugo "github.com/savsgio/atreugo/v11"
)

func main() {
	gocloakSession, err := gocloaksession.NewSession(clientID, clientSecret, keycloakRealm, keycloakURL)
	if err != nil {
		zerologger.Error().Msg("could not initialize gocloaksession")

		return
	}

	restyClient := resty.New()
	restyClient.OnBeforeRequest(gocloakSession.AddAuthTokenToRequest)

	config := atreugo.Config{
		Addr: "0.0.0.0:8000",
	}

	server := atreugo.New(config)

	server.GET("/search/{name}", func(ctx *atreugo.RequestCtx) error {
		// Your code goes here
	})

	server.POST("/search", func(ctx *atreugo.RequestCtx) error {
		// Your code goes here
	})

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
