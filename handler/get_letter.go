package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mochi-yu/dena-autumn-server/service"
)

type GetLetter struct {
	Service GetLetterService
}

func (gl *GetLetter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	letterID := chi.URLParam(r, "letterID")

	// TODO: validate

	strokes, err := gl.Service.GetLetter(ctx, letterID)
	if err == service.ErrUUIDNotExist {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	RespondJSON(ctx, w, strokes, http.StatusOK)
}
