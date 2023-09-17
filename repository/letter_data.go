package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/lithammer/shortuuid/v4"
)

// データの挿入処理
func (r *Repository) PostLetter(ctx context.Context, letterContent []byte) (letterID string, err error) {
	uuid := shortuuid.New()
	if err != nil {
		return "", err
	}

	query := "INSERT INTO LetterList (LetterID, LetterContent) VALUES(?, ?)"
	_, err = r.DB.ExecContext(ctx, query, uuid, string(letterContent))
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == ErrCodeMySQLDuplicateEntry {
			return "", fmt.Errorf("PostLetter: duplicate UUID")
		}
		return "", err
	}

	return uuid, nil
}

// データの取得処理
func (r *Repository) GetLetter(ctx context.Context, letterID string) ([]byte, error) {

	query := "SELECT LetterContent FROM LetterList WHERE LetterID = ?"
	row := r.DB.QueryRowContext(ctx, query, letterID)

	var letterContent []byte
	if err := row.Scan(&letterContent); err != nil {
		return nil, err
	}

	return letterContent, nil
}
