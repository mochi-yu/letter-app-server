package service

import (
	"context"
)

type LetterPoster interface {
	PostLetter(ctx context.Context, letterContent []byte) (string, error)
}

type LetterGetter interface {
	GetLetter(ctx context.Context, letterID string) ([]byte, error)
}
