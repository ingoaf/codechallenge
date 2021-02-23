package main

import (
	"encoding/json"

	"github.com/Clarilab/codechallenge/backend/challenge4/database"
	"github.com/Clarilab/codechallenge/backend/challenge4/models"
	"github.com/rs/zerolog/log"
	"github.com/savsgio/atreugo/v11"
)

const (
	Port = 6379
)

func main() {
	repo := database.NewRepository()
	cache, err := database.NewCache(Port)
	if err != nil {
		log.Fatal().Err(err)
	}
	api := &API{
		cache: cache,
		repo:  repo,
	}

	config := atreugo.Config{
		Addr: "0.0.0.0:8000",
	}

	server := atreugo.New(config)
	server.SetLogOutput(log.Logger)

	server.GET("/search/{name}", api.Search)

	server.POST("/search", api.Search)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msgf("can not start service")
		panic(err)
	}
}

// API implements the rest api handlers
type API struct {
	cache database.Cache
	repo  database.Repository
}

// Search handles the search via search request route
func (api *API) Search(ctx *atreugo.RequestCtx) error {
	var err error
	request := models.SearchRequest{}
	if err := json.Unmarshal(ctx.Request.Body(), &request); err != nil {
		return err
	}
	searchResponse := &models.SearchResponse{}
	exists, err := api.cache.Exists(&request)
	if exists && err == nil {
		searchResponse.Companies, err = api.cache.Get(&request)
		log.Info().Msg("response retrieved from cache")
	} else {
		companies, _ := api.repo.Get(&request)
		searchResponse.Companies = companies
		err = api.cache.Set(&request, searchResponse)
	}
	if err := json.NewEncoder(ctx).Encode(searchResponse); err != nil {
		return err
	}
	return err
}
