package db

import (
	"database/sql"
	"github.com/rs/zerolog/log"
)

func GetConn() (*sql.DB, error) {
	var db *sql.DB
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/prototype")
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	return db, nil
}
