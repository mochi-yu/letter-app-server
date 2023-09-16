package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mochi-yu/dena-autumn-server/entity"
)

type PostLetter struct {
	Repo LetterPoster
}

func (a *PostLetter) PostLetter(ctx context.Context, letterContent *entity.Strokes) (string, error) {
	latj, err := json.Marshal(letterContent)
	if err != nil {
		return "", fmt.Errorf("PostLetter: Marshal error: %w", err)
	}

	letterID, err := a.Repo.PostLetter(ctx, latj)
	if err != nil {
		return "", fmt.Errorf("PostLetter: failed to register: %w", err)
	}
	return letterID, nil
}
