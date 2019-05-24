package data

import "fmt"

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
