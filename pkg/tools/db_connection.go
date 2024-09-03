package tools

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func NewDBConnection() (*sql.DB, error) {
	dbConn, err := sql.Open("sqlite3", "./games.db")
	if err != nil {
		return nil, err
	}
	_, err = dbConn.Exec("CREATE TABLE IF NOT EXISTS `games` (id TEXT PRIMARY KEY, name TEXT NOT NULL, platform TEXT NOT NULL, price REAL NOT NULL);")
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}