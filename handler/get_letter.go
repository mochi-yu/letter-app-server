package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type GetLetter struct {
	Service GetLetterService
}

func (gl *GetLetter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	letterID := chi.URLParam(r, "letterID")

	// TODO: validate

	// TODO: 構造体で帰ってきて欲しい
	obj, err := gl.Service.GetLetter(ctx, letterID)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
	}
	RespondJSON(ctx, w, &obj, http.StatusOK)
}
