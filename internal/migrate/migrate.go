package migrate

import (
	"database/sql"
	"fmt"
	"io/fs"

	_ "github.com/jackc/pgx/v4/stdlib" //nolint:revive
	"github.com/pressly/goose/v3"
)

const driver = "pgx"

func Run(dataSourceName string, fsys fs.FS) error {
	db, err := sql.Open(driver, dataSourceName)
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}

	defer db.Close()

	goose.SetBaseFS(fsys)
	goose.SetLogger(&noopLogger{})

	err = goose.SetDialect(driver)
	if err != nil {
		return fmt.Errorf("goose.SetDialect: %w", err)
	}

	err = goose.Up(db, ".")
	if err != nil {
		return fmt.Errorf("goose.Up: %w", err)
	}

	return nil
}

func Reset(dataSourceName string, fsys fs.FS) error {
	db, err := sql.Open(driver, dataSourceName)
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}

	defer db.Close()

	goose.SetBaseFS(fsys)
	goose.SetLogger(&noopLogger{})

	err = goose.SetDialect(driver)
	if err != nil {
		return fmt.Errorf("goose.SetDialect: %w", err)
	}

	err = goose.Reset(db, ".")
	if err != nil {
		return fmt.Errorf("goose.Reset: %w", err)
	}

	return nil
}
