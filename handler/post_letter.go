package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mochi-yu/dena-autumn-server/entity"
)

type PostLetter struct {
	Service PostLetterService
}

func (pl *PostLetter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// strokes := entity.MakeStrokes()
	var strokes entity.Strokes
	if err := json.NewDecoder(r.Body).Decode(&strokes); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	// TODO: validate

	letterID, err := pl.Service.PostLetter(ctx, &strokes)
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
