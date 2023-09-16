package repository

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mochi-yu/dena-autumn-server/config"
)

const (
	ErrCodeMySQLDuplicateEntry = 1062 // 重複時のエラーコード
)

type Repository struct {
	DB *sql.DB
}

func New(ctx context.Context, cfg *config.Config) (*sql.DB, func(), error) {
	parameters := []string{
		"tls=true",
		"charset=utf8mb4",
		"collation=utf8mb4_general_ci",
		"interpolateParams=true",
		// "loc=Asia%2FTokyo",
		// "parseTime=true",
	}

	// dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ")/" + cfg.DBName + "?" + strings.Join(parameters, "&") // 以前はこれで接続成功していた
	dsn := cfg.DBUser + ":" + cfg.DBPassword +
		"@tcp(" + cfg.DBHost + ":" + strconv.Itoa(cfg.DBPort) + ")/" + cfg.DBName + "?" + strings.Join(parameters, "&")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, func() {}, err
	}
	return db, func() { db.Close() }, nil
}
