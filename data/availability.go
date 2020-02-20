package data

import "log"

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetDataAvailability(module string, indicators HandleList, groups HandleList, markets HandleList, destinations HandleList, categories HandleList) (result *map[string]HandleList , err error) {
	//language=MySQL
	query :=
		`select frequency from data_toc toc
		 left join indicators i on i.id = toc.indicator_id
		 left join groups g on g.id = toc.group_id
		 left join markets m on m.id = toc.market_id
		 left join destinations d on d.id = toc.destination_id
		 left join categories c on c.id = toc.category_id
		 where module = ? `
	var bindVals []interface{}
	bindVals = append(bindVals, module)

	query += "and i.handle in (" + makeQlist(len(indicators)) + ")\n"
	for _, ind := range indicators {
		bindVals = append(bindVals, ind)
	}
	if len(markets) > 0 {
		query += "and m.handle in (" + makeQlist(len(markets)) + ")\n"
		for _, mkt := range markets {
			bindVals = append(bindVals, mkt)
		}
	}
	if len(destinations) > 0 {
		query += "and d.handle in (" + makeQlist(len(destinations)) + ")\n"
		for _, dest := range destinations {
			bindVals = append(bindVals, dest)
		}
	}
	if len(groups) > 0 {
		query += "and g.handle in (" + makeQlist(len(groups)) + ")\n"
		for _, grp := range groups {
			bindVals = append(bindVals, grp)
		}
	}
	if len(categories) > 0 {
		query += "and c.handle in (" + makeQlist(len(categories)) + ")\n"
		for _, cat := range categories {
			bindVals = append(bindVals, cat)
		}
	}
	// query += "order by 1,2,3,4,5" + "\n" // extra "\n" only to make GoLand shut up about an error :(

	dbResults, err := Db.Query(query, bindVals...)
	if err != nil {
		log.Printf("Database error: %s", err.Error())
		return
	}
	currentSlug := ""
	var series Series
	for dbResults.Next() {
	}
	return
}
