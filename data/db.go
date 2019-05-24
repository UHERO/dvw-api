package data

import (
	"database/sql"
	"fmt"
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

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetDimensionAll(dim string, mod string) (dimList []Dimension, err error) {
	//language=MySQL
	var query = fmt.Sprintf(`select * from %s where module = ?`, dim)
	rows, err := Db.Query(query, mod)
	if err != nil {
		return
	}
	for rows.Next() {
	}
	return
}
