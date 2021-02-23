package main

import (
	"encoding/json"
	"errors"

	"github.com/Clarilab/codechallenge/backend/challenge2/database"
	"github.com/Clarilab/codechallenge/backend/challenge2/middlewares"
	"github.com/Clarilab/codechallenge/backend/challenge2/models"
	"github.com/rs/zerolog/log"
	"github.com/savsgio/atreugo/v11"
)

func main() {
	repo := database.NewRepository()
	api := &API{
		repo: repo,
	}

	config := atreugo.Config{
		Addr: "0.0.0.0:8000",
	}

	server := atreugo.New(config)
	server.SetLogOutput(log.Logger)
	server.UseBefore(middlewares.LogRequest)

	server.GET("/search/{name}", api.SearchForName)

	server.POST("/search", api.Search)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msgf("can not start service")
		panic(err)
	}
}

// API implements the rest api handlers
type API struct {
	repo database.Repository
}

// SearchForName handles the the name search route
func (api *API) SearchForName(ctx *atreugo.RequestCtx) error {
	var err error
	name := ctx.UserValue("name")
	nameString, ok := name.(string)
	searchResponse := &models.SearchResponse{}
	if !ok {
		err := errors.New("Param conversion to string doesn't work")
		log.Error().Err(err).Msg("something went wrong")
		return err
	}
	searchResponse.Companies = api.repo.GetByName(nameString)
	searchResponse.TotalHits = len(searchResponse.Companies)
	if err := json.NewEncoder(ctx).Encode(searchResponse); err != nil {
		return err
	}
	return err
}

// Search handles the search via search request route
func (api *API) Search(ctx *atreugo.RequestCtx) error {
	var err error
	request := models.SearchRequest{}
	if err := json.Unmarshal(ctx.Request.Body(), &request); err != nil {
		return err
	}
	searchResponse := &models.SearchResponse{}
	searchResponse.Companies = api.repo.Get(request)
	if err := json.NewEncoder(ctx).Encode(searchResponse); err != nil {
		return err
	}
	return err
}
