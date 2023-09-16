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
	mux.Route("/letter", func(r chi.Router) {
		r.Post("/", handler.PostLetterHandler)
		r.Get("/{letterId}", handler.GetLetterHandler) // GET /articles
	})
	return mux, func() {}, nil
}
