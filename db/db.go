package db

import (
	"database/sql"
	"main/utils/singleton"

	"github.com/mattn/go-sqlite3"
)

var (
	Dialect = "sqlite3"
	Driver  = Dialect + "_custom"
	Path    string
)

func Db() *sql.DB {
	return singleton.GetInstance(func() *sql.DB {
		sql.Register(Driver, &sqlite3.SQLiteDriver{
			ConnectHook: func(sc *sqlite3.SQLiteConn) error {
				var err error
				return err
			},
		})

		db, err := sql.Open(Driver, Path)
		if err != nil {
			return nil
		}

		return db
	})
}
