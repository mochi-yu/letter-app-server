package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mochi-yu/dena-autumn-server/config"
	"github.com/mochi-yu/dena-autumn-server/handler"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/", handler.JsontestHandle)

	return mux, func() {}, nil
}
