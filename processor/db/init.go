package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"frendler/processor/config"
)

func Init(cfg config.DBConf) *DBsql {
	db, err := sql.Open("sqlite3", cfg.Path)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return &DBsql{DB: db}
}
