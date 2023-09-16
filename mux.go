package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mochi-yu/dena-autumn-server/config"
	"github.com/mochi-yu/dena-autumn-server/handler"
	"github.com/mochi-yu/dena-autumn-server/repository"
	"github.com/mochi-yu/dena-autumn-server/service"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	r := chi.NewRouter()
	r.HandleFunc("/", handler.JsonTest)

	db, cleanup, err := repository.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	repo := repository.Repository{
		DB: db,
	}

	pl := &handler.PostLetter{
		Service: &service.PostLetter{Repo: &repo},
	}
	gl := &handler.GetLetter{
		Service: &service.GetLetter{Repo: &repo},
	}

	r.Route("/letter", func(r chi.Router) {
		r.Post("/", pl.ServeHTTP)          // 作成した手紙の投稿
		r.Get("/{letterID}", gl.ServeHTTP) // 手紙の読み込み
	})
	return r, func() {}, nil
}
