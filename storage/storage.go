package storage

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sqlx.DB
}

func New() *Storage {
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE")))
	if err != nil {
		slog.Error("storage: cannot connect to DB", slog.String("error", err.Error()))
		os.Exit(1)
	}

	return &Storage{db: db}
}
