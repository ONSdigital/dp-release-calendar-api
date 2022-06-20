package api

import (
	"context"

	"github.com/gorilla/mux"
)

//API provides a struct to wrap the api around
type API struct {
	Router *mux.Router
}

//Setup function sets up the api and returns an api
func Setup(ctx context.Context, r *mux.Router, zc ZebedeeClient) *API {
	api := &API{
		Router: r,
	}

	r.HandleFunc("/v1/releases/legacy", LegacyHandler(ctx, zc)).Methods("GET")
	return api
}
