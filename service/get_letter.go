package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mochi-yu/dena-autumn-server/entity"
)

type GetLetter struct {
	Repo LetterGetter
}

func (g *GetLetter) GetLetter(ctx context.Context, letterID string) (*entity.Strokes, error) {
	letj, err := g.Repo.GetLetter(ctx, letterID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("GetLetter: such letterID does not exist")
		}
		return nil, fmt.Errorf("GetLetter: failed to get: %w", err)
	}
	fmt.Println(string(letj))

	var letter entity.Strokes
	if err := json.Unmarshal(letj, &letter); err != nil {
		return nil, fmt.Errorf("GetLetter: Unmarshal error: %w", err)
	}
	return &letter, nil
}
