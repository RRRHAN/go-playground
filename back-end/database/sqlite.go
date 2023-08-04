package database

import (
	"embed"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed migrations
var migrations embed.FS

func NewDB() (db *sqlx.DB, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Error get WD - %w", err)
	}

	path := filepath.Join(wd, "data")
	_, err = os.Stat(path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("failed to check directory - %w", err)
		}
		err = os.Mkdir(path, 0700)
		if err != nil {
			return nil, fmt.Errorf("Error make directory - %w", err)
		}
	}

	db, err = sqlx.Open("sqlite3", "file:data/main.sqlite3")
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite connection - %w", err)
	}

	sourceInstance, err := httpfs.New(http.FS(migrations), "migrations")
	if err != nil {
		return nil, fmt.Errorf("invalid source instance, %w", err)
	}
	targetInstance, err := sqlite.WithInstance(db.DB, new(sqlite.Config))
	if err != nil {
		return nil, fmt.Errorf("invalid target sqlite instance, %w", err)
	}
	m, err := migrate.NewWithInstance(
		"httpfs", sourceInstance, "sqlite3", targetInstance)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize migrate instance, %w", err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return nil, err
	}
	err = sourceInstance.Close()
	if err != nil {
		return nil, err
	}

	return
}
