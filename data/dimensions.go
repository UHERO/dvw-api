package data

import "fmt"

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetDimensionAll(dim string, mod string) (dimList []Dimension, err error) {
	xtraCols := ", null, null"  // No unit or decimals for all tables other than indicators
	if dim == "indicators" {
		xtraCols = ", t.unit, t.decimal"
	}
	//language=MySQL
	var query = fmt.Sprintf(
	 	 `select t.module, t.handle, t.nameP, t.nameW, t.nameT, t.header, p.handle as parent, t.level, t.order %s
          from %s t left join %s p on p.id = t.parent_id
          where t.module = ?`, xtraCols, dim, dim)
	results, err := Db.Query(query, mod)
	if err != nil {
		return
	}
	for results.Next() {
		scanDim := Dimension{}
		err = results.Scan(&scanDim.Module, &scanDim.Handle, &scanDim.NameP, &scanDim.NameW, &scanDim.NameT,
						   &scanDim.Header, &scanDim.Parent, &scanDim.Level, &scanDim.Order, &scanDim.Unit, &scanDim.Decimal)
		if err != nil {
			return
		}
		dimList = append(dimList, scanDim)
	}
	return
}
