package data

import (
	"fmt"
	"log"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetDimensionAll(dim string, mod string) (dimList []PortalDimension, err error) {
	xtraCols := ", null, null"  // No unit or decimals for all tables other than indicators
	if dim == "indicators" {
		xtraCols = ", t.unit, t.decimal"
	}
	//language=MySQL
	var query = fmt.Sprintf(
	 	 `select t.module, t.handle, t.nameP, t.nameW, t.nameT, t.header, p.handle as parent, t.level, t.order %s
          from %s t left join %s p on p.id = t.parent_id
          where t.module = ? `, xtraCols, dim, dim)
	results, err := Db.Query(query, mod)
	if err != nil {
		log.Printf("Database error: %s", err.Error())
		return
	}
	for results.Next() {
		scanDim := Dimension{}
		err = results.Scan(&scanDim.Module, &scanDim.Handle, &scanDim.NameP, &scanDim.NameW, &scanDim.NameT,
						   &scanDim.Header, &scanDim.Parent, &scanDim.Level, &scanDim.Order, &scanDim.Unit, &scanDim.Decimal)
		if err != nil {
			return
		}
		pDim := PortalDimension{
			Module: scanDim.Module,
			Handle: scanDim.Handle,
			NameW: scanDim.NameW,
			Header: scanDim.Header,
			Level: scanDim.Level,
			Order: scanDim.Order,
		}
		if scanDim.NameP.Valid {
			pDim.NameP = scanDim.NameP.String
		}
		if scanDim.NameT.Valid {
			pDim.NameT = scanDim.NameT.String
		}
		if scanDim.Parent.Valid {
			pDim.Parent = scanDim.Parent.String
		}
		if scanDim.Unit.Valid {
			pDim.Unit = scanDim.Unit.String
		}
		if scanDim.Decimal.Valid {
			pDim.Decimal = scanDim.Decimal.String
		}
		dimList = append(dimList, pDim)
	}
	return
}
