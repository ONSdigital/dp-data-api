package api

import (
	"context"
	"net/http"
)

// API provides a struct to wrap the api around
type API struct {
	Mux *http.ServeMux
}

// Setup function sets up the api and returns an api
func Setup(ctx context.Context, m *http.ServeMux) *API {
	api := &API{
		Mux: m,
	}

	// TODO: remove hello world example handler route
	m.HandleFunc("GET /hello", HelloHandler(ctx))
	return api
}
