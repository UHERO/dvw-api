package data

import (
	"database/sql"
)

var Db *sql.DB

///////////////////////////////////////////////////////////////////////////////////////////////////
func CreateDatabase(connString string) (newDb *sql.DB, err error) {
	if newDb, err = sql.Open("mysql", connString); err != nil {
		return
	}
	if err = newDb.Ping(); err != nil {
		return
	}
	Db = newDb
	return
}
