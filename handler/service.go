package handler

import (
	"context"
	"errors"

	"github.com/mochi-yu/dena-autumn-server/entity"
)

type PostLetterService interface {
	PostLetter(ctx context.Context, letterContent *entity.Strokes) (string, error)
}

type GetLetterService interface {
	GetLetter(ctx context.Context, letterID string) (*entity.Strokes, error)
}

// テスト用
type PostLetterServiceMock struct{}

func (ps *PostLetterServiceMock) PostLetter(ctx context.Context, letterContent *entity.Strokes) (string, error) {
	if letterContent == nil || len(*letterContent) == 0 {
		return "", errors.New("PostLetterServiceMock error")
	}
	return "this_is_uuid plus " + (*letterContent)[0].CompositeOperation, nil
}

// テスト用
type GetLetterServiceMock struct{}

func (ps *GetLetterServiceMock) GetLetter(ctx context.Context, letterID string) (*entity.Strokes, error) {
	if letterID == "" {
		return nil, errors.New("GetLetterServiceMock error")
	}
	strokesMock := entity.MakeStrokes()
	return &strokesMock, nil
}
