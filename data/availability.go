package data

import "log"

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetDataAvailability(module string, indicators HandleList, groups HandleList, markets HandleList, destinations HandleList, categories HandleList) (result *map[string]HandleList , err error) {
	//language=MySQL
	query :=
		`select * from data_toc toc		 where toc.module = ? `

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
	query += "order by 1,2,3,4,5" + "\n" // extra "\n" only to make GoLand shut up about an error :(

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
