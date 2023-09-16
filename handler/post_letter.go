package handler

import (
	"net/http"

	"github.com/mochi-yu/dena-autumn-server/entity"
)

type PostLetter struct {
	Service PostLetterService
}

func (pl *PostLetter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// TODO: validate

	letterID, err := pl.Service.PostLetterService(ctx, &entity.Test{Data: "TODO: test"})
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	rsp := struct {
		LetterID string `json:"letter_id"`
	}{LetterID: letterID}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
