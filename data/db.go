package data

import "database/sql"

var Db *sql.DB

///////////////////////////////////////////////////////////////////////////////////////////////////
func CreateDatabase(connString string) (newDb *sql.DB, err error) {
	newDb, err = sql.Open("mysql", connString)
	if err != nil {
		return
	}
	err = newDb.Ping()
	if err != nil {
		return
	}
	Db = newDb
	return
}
