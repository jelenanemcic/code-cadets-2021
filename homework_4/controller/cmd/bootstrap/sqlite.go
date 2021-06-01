package bootstrap

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/jelenanemcic/code-cadets-2021/homework_4/controller/cmd/config"
)

func Sqlite() *sql.DB {
	db, err := sql.Open("sqlite3", config.Cfg.SqliteDatabase)
	if err != nil {
		panic(err)
	}

	return db
}
