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
	var baseTime int64 = 9223372036854775807

	for i := 0; i < len(*letterContent); i++ {
		st := &((*letterContent)[i].Points)
		for j := 0; j < len(*st); j++ {
			baseTime = min(baseTime, (*st)[j].Time)
		}
	}
	for i := 0; i < len(*letterContent); i++ {
		st := &((*letterContent)[i].Points)
		for j := 0; j < len(*st); j++ {
			(*st)[j].Time -= baseTime
		}
	}

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
