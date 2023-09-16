package repository

import (
	"github.com/google/uuid"
)

// データの挿入処理
func (ms *MySQL) InsertLetterData(letterContent string) (letterId string, err error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	query := "INSERT INTO LetterList(LetterID, LetterContent) VALUES(?, ?, ?)"
	_, err = ms.db.Exec(query, uuid.String(), letterContent)
	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}

// データの取得処理
func (ms *MySQL) GetLetterData(letterId string) (string, error) {
	var letterContent string

	rows, err := ms.db.Query("SELECT LetterContent FROM LetterList WHERE LetterID = ?", letterId)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&letterContent); err != nil {
			return "", err
		}
	}

	return letterContent, nil
}
