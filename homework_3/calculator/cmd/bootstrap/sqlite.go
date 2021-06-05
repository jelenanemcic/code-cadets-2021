package bootstrap

import (
	"database/sql"

	"github.com/jelenanemcic/code-cadets-2021/homework_3/calculator/cmd/config"
	_ "github.com/mattn/go-sqlite3"
)

func Sqlite() *sql.DB {
	db, err := sql.Open("sqlite3", config.Cfg.SqliteDatabase)
	if err != nil {
		panic(err)
	}

	return db
}
