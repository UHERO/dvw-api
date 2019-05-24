package data

import "fmt"

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetDimensionAll(dim string, mod string) (dimList []Dimension, err error) {
	var xtraCols string
	if dim == "indicators" {
		xtraCols = ", t.unit, t.decimal"
	}
	//language=MySQL
	var query = fmt.Sprintf(`
          select t.module, t.handle, t.nameP, t.nameW, t.nameT, t.header, p.handle as parent, t.level, t.order%s
          from %s t left join %s p on p.id = t.parent_id
          where t.module = ?`, xtraCols, dim, dim)
	rows, err := Db.Query(query, mod)
	if err != nil {
		return
	}
	for rows.Next() {
	}
	return
}
