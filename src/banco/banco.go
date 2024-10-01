package banco

import (
	"database/sql"
	"lavajato/src/config"

	_ "github.com/go-sql-driver/mysql"
)

func Conection() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
