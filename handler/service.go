package handler

import (
	"context"
	"errors"

	"github.com/mochi-yu/dena-autumn-server/entity"
)

type PostLetterService interface {
	PostLetterService(ctx context.Context, letterContent *entity.Test) (string, error)
}

type PostLetterServiceMock struct{}

func (ps *PostLetterServiceMock) PostLetterService(ctx context.Context, letterContent *entity.Test) (string, error) {
	if letterContent.Data == "" {
		return "", errors.New("PostLetterServiceMock error")
	}
	return "this_is_uuid", nil
}

type GetLetterService interface {
	GetLetter(ctx context.Context, letterID string) (*entity.Test, error)
}

type GetLetterServiceMock struct{}

func (ps *GetLetterServiceMock) GetLetter(ctx context.Context, letterID string) (*entity.Test, error) {
	if letterID == "" {
		return nil, errors.New("GetLetterServiceMock error")
	}
	return &entity.Test{Data: "this_is_test"}, nil
}
