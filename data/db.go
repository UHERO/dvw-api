package data

import (
	"database/sql"
	"strings"
)

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

///////////////////////////////////////////////////////////////////////////////////////////////////
func makeQlist(length int) string {
	var list []string
	for i := 0; i < length; i++ {
		list = append(list, "?")
	}
	return strings.Join(list, ",")
}
