package storage

import (
	"embed"
	"log"

	"github.com/pressly/goose/v3"

	"frendler/processor/db"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func Migrations(db *db.DBsql) {

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite"); err != nil {
		log.Fatalf("Failed to set goose dialect: %v", err)
	}

	if err := goose.Up(db.DB, "migrations"); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Success migrations")
}
