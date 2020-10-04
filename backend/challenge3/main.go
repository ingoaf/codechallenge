package main

import (
	"fmt"

	"github.com/Clarilab/gocloaksession"
	"github.com/go-resty/resty/v2"
	"github.com/savsgio/atreugo/v11"
)

const clientID = ""
const clientSecret = ""
const keycloakRealm = ""
const keycloakURL = ""

func main() {
	gocloakSession, err := gocloaksession.NewSession(clientID, clientSecret, keycloakRealm, keycloakURL)
	if err != nil {
		fmt.Println("could not initialize gocloaksession")
		return
	}

	restyClient := resty.New()
	restyClient.OnBeforeRequest(gocloakSession.AddAuthTokenToRequest)

	api := &API{}

	config := atreugo.Config{
		Addr: "0.0.0.0:8000",
	}

	server := atreugo.New(config)

	server.GET("/search/{name}", api.SearchForName)

	server.POST("/search", api.Search)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// API implements the rest api handlers
type API struct {
}

func (api *API) SearchForName(ctx *atreugo.RequestCtx) error {
	// Your code goes here
	return nil
}

func (api *API) Search(ctx *atreugo.RequestCtx) error {
	// Your code goes here
	return nil
}
