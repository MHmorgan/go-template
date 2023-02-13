package data

import (
	"database/sql"
	"embed"
	"github.com/mhmorgan/rodla/phys"
	_ "rsc.io/sqlite"
)

var db *sql.DB

//go:embed sql
var sqlFiles embed.FS

func Init(dbFile string) (err error) {
	db, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		return err
	}

	schema, err := sqlFiles.ReadFile("sql/schema.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}

func Close() error {
	return db.Close()
}

// FETCH OPTIONS

type FetchOption func(*fetchOptions)

type fetchOptions struct {
	location phys.Sized
}

func AtLocation(l phys.Sized) FetchOption {
	return func(opts *fetchOptions) {
		opts.location = l
	}
}
