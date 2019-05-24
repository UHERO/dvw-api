package data

import "fmt"

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetDimensionAll(dim string, mod string) (dimList []Dimension, err error) {
	var xtraCols string
	if dim == "indicators" {
		xtraCols = ", t1.unit, t1.decimal"
	}
	//language=MySQL
	var query = fmt.Sprintf(`select t1.module, t1.handle, t1.nameP, t1.nameW, t1.nameT, t1.header, t2.handle as parent, t1.level, t1.order%s
							 from %s t1 left join %s t2 on t2.id = t1.parent_id
                             where module = ?`, xtraCols, dim, dim)
	rows, err := Db.Query(query, mod)
	if err != nil {
		return
	}
	for rows.Next() {
	}
	return
}
