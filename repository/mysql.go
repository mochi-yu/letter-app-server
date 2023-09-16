package repository

import (
	"database/sql"

	"github.com/mochi-yu/dena-autumn-server/config"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL() (MySQL, error) {
	conf, err := config.New()
	if err != nil {
		return MySQL{}, err
	}

	db, err := sql.Open("mysql", conf.DNS)
	if err != nil {
		return MySQL{}, err
	}

	return MySQL{db: db}, nil
}
